/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"golang.org/x/tools/go/packages"
)

var flHelp = pflag.BoolP("help", "h", false, "print help and exit")
var flDbg = pflag.BoolP("debug", "d", false, "enable debugging output")
var flDbgTime = pflag.BoolP("debug-time", "D", false, "enable debugging output with timestamps")
var flOut = pflag.StringP("output", "o", "make", "output format (mainly for debugging): one of make | json)")
var flTags = pflag.StringSlice("tag", nil, "build tags to pass to Go (see 'go help build', may be specified multiple times)")
var flRelPath = pflag.String("relative-to", ".", "emit by-path rules for packages relative to this path")
var flStateDir = pflag.String("state-dir", ".go2make", "directory in which to store state used by make")
var flIgnoreErrors = pflag.BoolP("ignore-errors", "e", false, "ignore package errors")

var lastDebugTime time.Time

func debug(items ...interface{}) {
	if *flDbg {
		x := []interface{}{}
		if *flDbgTime {
			elapsed := time.Since(lastDebugTime)
			if lastDebugTime.IsZero() {
				elapsed = 0
			}
			lastDebugTime = time.Now()
			x = append(x, fmt.Sprintf("DBG(+%v):", elapsed))
		} else {
			x = append(x, "DBG:")
		}
		x = append(x, items...)
		fmt.Fprintln(os.Stderr, x...)
	}

}

type emitter struct {
	tags         []string
	ignoreErrors bool
	relPath      string
	stateDir     string
}

func main() {
	pflag.Parse()

	if *flHelp {
		help(os.Stdout)
		os.Exit(0)
	}

	if *flDbgTime {
		*flDbg = true
	}

	switch *flOut {
	case "make":
	case "json":
	default:
		fmt.Fprintf(os.Stderr, "unknown output format %q\n", *flOut)
		pflag.Usage()
		os.Exit(1)
	}

	if *flRelPath == "" {
		fmt.Fprintf(os.Stderr, "error: --relative-to must be defined\n")
		os.Exit(1)
	}

	if *flStateDir == "" {
		fmt.Fprintf(os.Stderr, "error: --state-dir must be defined\n")
		os.Exit(1)
	}

	targets := pflag.Args()
	if len(targets) == 0 {
		targets = append(targets, ".")
	}
	debug("targets:", targets)

	// Gather flag values for easier testing.
	emit := emitter{
		tags:         *flTags,
		ignoreErrors: *flIgnoreErrors,
		relPath:      strings.TrimRight(absOrExit(*flRelPath), "/"),
		stateDir:     strings.TrimRight(*flStateDir, "/"),
	}
	debug("tags:", emit.tags)
	debug("relative-to:", emit.relPath)

	pkgs, err := emit.loadPackages(targets...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading packages: %v\n", err)
		os.Exit(1)
	}

	pkgMap := emit.visitPackages(pkgs)
	if pkgMap == nil {
		os.Exit(1)
	}

	switch *flOut {
	case "make":
		emit.emitMake(os.Stdout, pkgMap)
	case "json":
		emit.emitJSON(os.Stdout, pkgMap)
	}
}

func help(out io.Writer) {
	prog := filepath.Base(os.Args[0])
	fmt.Fprintf(out, "Usage: %s [FLAG...] <PKG...>\n", prog)
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "%s calculates all of the dependencies of a set of Go packages and\n", prog)
	fmt.Fprintf(out, "emits a Makfile (unless otherwise specified) which can be used to track dependencies.\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "Package specifications may be simple (e.g. 'example.com/txt/color') or\n")
	fmt.Fprintf(out, "recursive (e.g. 'example.com/txt/...'), and may be Go package names or\n")
	fmt.Fprintf(out, "relative file paths (e.g. './...')\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, " Example output:\n")
	fmt.Fprintf(out, "  .go2make/by-pkg/example.com/txt/color/_pkg: .go2make/by-pkg/example.com/txt/color/_files \\\n")
	fmt.Fprintf(out, "    color/color.go \\\n")
	fmt.Fprintf(out, "    .go2make/by-pkg/bytes/_pkg \\\n")
	fmt.Fprintf(out, "    .go2make/by-pkg/example.com/pretty/_pkg\n")
	fmt.Fprintf(out, "          @mkdir -p $(@D)\n")
	fmt.Fprintf(out, "          @touch $@\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "  .go2make/by-path/pkg/example.com/txt/color/_pkg: .go2make/by-pkg/example.com/txt/color/_pkg\n")
	fmt.Fprintf(out, "          @mkdir -p $(@D)\n")
	fmt.Fprintf(out, "          @touch $@\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "User Makefiles can include the generated output and trigger actions when the Go packages need\n")
	fmt.Fprintf(out, "to be rebuilt.  The 'by-pkg/.../_pkg' rules are defined by the Go package name (e.g.\n")
	fmt.Fprintf(out, "example.com/txt/color).  The 'by-path/.../_pkg' rules are defined by the relative path of the\n")
	fmt.Fprintf(out, "Go package when that path is below the value of the --relative-to flag.\n")
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, " Flags:\n")

	pflag.PrintDefaults()
}

func absOrExit(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	return abs
}

func (emit emitter) loadPackages(targets ...string) ([]*packages.Package, error) {
	cfg := packages.Config{
		Mode:       packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedModule,
		Tests:      false,
		BuildFlags: []string{"-tags", strings.Join(emit.tags, ",")},
	}
	return packages.Load(&cfg, targets...)
}

func (emit emitter) visitPackages(pkgs []*packages.Package) map[string]*packages.Package {
	pkgMap := map[string]*packages.Package{}
	errs := false
	for _, p := range pkgs {
		ok := emit.visitPackage(p, pkgMap)
		if !ok {
			errs = true
		}
	}
	if errs {
		return nil
	}
	return pkgMap
}

func (emit emitter) visitPackage(pkg *packages.Package, pkgMap map[string]*packages.Package) bool {
	debug("visiting package", pkg.PkgPath)
	if pkgMap[pkg.PkgPath] == pkg {
		debug("  ", pkg.PkgPath, "was already visited")
		return true
	}
	debug("  ", pkg.PkgPath, "is new")
	pkgMap[pkg.PkgPath] = pkg

	ok := true
	for _, e := range pkg.Errors {
		if emit.ignoreErrors {
			debug("    ignoring error:", e.Msg)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", e.Msg)
			ok = false
		}
	}

	return ok
}

func keys(m map[string]*packages.Package) []string {
	sl := make([]string, 0, len(m))
	for k := range m {
		sl = append(sl, k)
	}
	sort.Strings(sl)
	return sl
}

func maybeRelative(path, relativeTo string) (string, bool) {
	if path == relativeTo || strings.HasPrefix(path, relativeTo+"/") {
		return strings.TrimPrefix(path, relativeTo+"/"), true
	}
	return path, false
}

func (emit emitter) emitMake(out io.Writer, pkgMap map[string]*packages.Package) {
	for _, k := range keys(pkgMap) {
		pkg := pkgMap[k]

		codeDir := ""
		isRel := false
		if len(pkg.GoFiles) > 0 {
			codeDir, isRel = maybeRelative(filepath.Dir(pkg.GoFiles[0]), emit.relPath)
			// Emit a rule to represent changes to the directory contents.
			// This rule will be evaluated whenever the code-directory is
			// newer than the saved file-list, but the file-list will only get
			// touched (triggering downstream rebuilds) if the set of files
			// actually changes.
			fmt.Fprintf(out, "%s/by-pkg/%s/_files: %s\n", emit.stateDir, pkg.PkgPath, codeDir)
			fmt.Fprintf(out, "\t@mkdir -p $(@D)\n")
			fmt.Fprintf(out, "\t@ls $</*.go | LC_ALL=C sort > $@.tmp\n")
			fmt.Fprintf(out, "\t@if ! cmp -s $@.tmp $@; then \\\n")
			fmt.Fprintf(out, "\t    cat $@.tmp > $@; \\\n")
			fmt.Fprintf(out, "\tfi\n")
			fmt.Fprintf(out, "\t@rm -f $@.tmp\n")
			fmt.Fprintf(out, "\n")
		}

		// Emit a rule to represent the whole package.  This uses a file,
		// rather than the directory itself, to avoid nested dir creation
		// changing the directory's timestamp.
		fmt.Fprintf(out, "%s/by-pkg/%s/_pkg:", emit.stateDir, pkg.PkgPath)
		if len(pkg.GoFiles) > 0 {
			fmt.Fprintf(out, " %s/by-pkg/%s/_files", emit.stateDir, pkg.PkgPath)
		}
		for _, f := range pkg.GoFiles {
			rel, _ := maybeRelative(f, emit.relPath)
			fmt.Fprintf(out, " \\\n  %s", rel)
		}
		for _, imp := range keys(pkg.Imports) {
			if pkgMap[pkg.Imports[imp].PkgPath] != nil {
				fmt.Fprintf(out, " \\\n  %s/by-pkg/%s/_pkg", emit.stateDir, pkg.Imports[imp].PkgPath)
			}
		}
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "\t@mkdir -p $(@D)\n")
		fmt.Fprintf(out, "\t@touch $@\n")
		fmt.Fprintf(out, "\n")

		if isRel {
			// Emit a rule to represent the package, but by a relative path.  This
			// is useful when you know the path to something but maybe not which Go
			// package it is (e.g. you have a bunch of packages).  Like the by-pkg
			// equivalent, this uses a file, to avoid nested dir creation changing
			// the directory's timestamp.
			fmt.Fprintf(out, "%s/by-path/%s/_pkg: %s/by-pkg/%s/_pkg\n", emit.stateDir, codeDir, emit.stateDir, pkg.PkgPath)
			fmt.Fprintf(out, "\t@mkdir -p $(@D)\n")
			fmt.Fprintf(out, "\t@touch $@\n")
			fmt.Fprintf(out, "\n")
		}
	}
}

func (emit emitter) emitJSON(out io.Writer, pkgMap map[string]*packages.Package) {
	jb, err := json.Marshal(pkgMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON error: %v", err)
		os.Exit(1)
	}
	fmt.Fprintln(out, string(jb))
}
