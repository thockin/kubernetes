/*
Copyright 2022 The Kubernetes Authors.

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
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lithammer/dedent"
	"golang.org/x/tools/go/packages"
)

func initModule(t *testing.T, name string, files map[string]string) string {
	dir := t.TempDir()
	writeFile(t, dir, "go.mod", dedent.Dedent(`
		module example.com/mod
		go 1.18
	`))
	for path, content := range files {
		writeFile(t, dir, path, content)
	}
	return dir
}

func writeFile(t *testing.T, dir, path, content string) {
	path = filepath.Join(dir, path)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

// build a map of pkg -> filenames
func pkgFiles(pkgs []*packages.Package) map[string][]string {
	out := make(map[string][]string, len(pkgs))
	for _, p := range pkgs {
		skip := false
		for _, e := range p.Errors {
			// Some test cases try bad patterns on purpose.
			if e.Kind == packages.ListError {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		files := make([]string, 0, len(p.GoFiles))
		for _, f := range p.GoFiles {
			files = append(files, filepath.Base(f))
		}
		out[p.PkgPath] = files
	}
	return out
}

func TestLoadPackages(t *testing.T) {
	cases := []struct {
		name       string
		files      map[string]string
		tags       []string
		expectPkgs map[string][]string
	}{{
		name: "one_pkg_no_imports",
		files: map[string]string{
			"file.go": dedent.Dedent(`
				package p
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod": {"file.go"},
		},
	}, {
		name: "one_pkg_with_other_files",
		files: map[string]string{
			"README": "",
			"file.go": dedent.Dedent(`
				package p
				var V string
			`),
			"file_test.go": dedent.Dedent(`
				package p
				var T string
			`),
			"_ignore.go": dedent.Dedent(`
				package p
				var I string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod": {"file.go"},
		},
	}, {
		name: "one_pkg_with_imports",
		files: map[string]string{
			"file.go": dedent.Dedent(`
				package p
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod": {"file.go"},
		},
	}, {
		name: "multi_pkg_no_imports",
		files: map[string]string{
			"p1/file1.go": dedent.Dedent(`
				package p1
				var V string
			`),
			"p2/file2.go": dedent.Dedent(`
				package p2
				var V string
			`),
			"p3/file3.go": dedent.Dedent(`
				package p3
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1.go"},
			"example.com/mod/p2": {"file2.go"},
			"example.com/mod/p3": {"file3.go"},
		},
	}, {
		name: "multi_pkg_with_imports",
		files: map[string]string{
			"p1/file1.go": dedent.Dedent(`
				package p1
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
			"p2/file2.go": dedent.Dedent(`
				package p2
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
			"p3/file3.go": dedent.Dedent(`
				package p3
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1.go"},
			"example.com/mod/p2": {"file2.go"},
			"example.com/mod/p3": {"file3.go"},
		},
	}, {
		name: "multi_pkg_with_tags",
		tags: []string{"foo"},
		files: map[string]string{
			"p1/file1.go": dedent.Dedent(`
				package p1
				var V string
			`),
			"p2/file2.go": dedent.Dedent(`
				//go:build !foo
				// +build !foo
				package p2
				var V string
			`),
			"p3/file3.go": dedent.Dedent(`
				//go:build foo
				// +build foo
				package p3
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1.go"},
			"example.com/mod/p3": {"file3.go"},
		},
	}, {
		name: "multi_pkg_multi_file_with_tags",
		tags: []string{"foo", "bar"},
		files: map[string]string{
			"p1/file1a.go": dedent.Dedent(`
				package p1
				var Va string
			`),
			"p1/file1b.go": dedent.Dedent(`
				package p1
				var Vb string
			`),
			"p2/file2a.go": dedent.Dedent(`
				//go:build !foo
				// +build !foo
				package p2
				var Va string
			`),
			"p2/file2b.go": dedent.Dedent(`
				package p2
				var Vb string
			`),
			"p3/file3a.go": dedent.Dedent(`
				//go:build bar
				// +build bar
				package p3
				var Va string
			`),
			"p3/file3b.go": dedent.Dedent(`
				package p3
				var Vb string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1a.go", "file1b.go"},
			"example.com/mod/p2": {"file2b.go"},
			"example.com/mod/p3": {"file3a.go", "file3b.go"},
		},
	}}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := initModule(t, "example.com/mod", tc.files)

			emit := emitter{
				tags: tc.tags,
			}

			// pushd
			if err := os.Chdir(dir); err != nil {
				t.Fatal(err)
			}

			for _, pattern := range []string{"example.com/mod/...", "./..."} {
				pkgs, err := emit.loadPackages(pattern)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if want, got := tc.expectPkgs, pkgFiles(pkgs); !cmp.Equal(want, got) {
					t.Errorf("wrong result for pattern %q:\n\twant: %v\n\t got: %v", pattern, want, got)
				}
			}

			// popd
			if err := os.Chdir(wd); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestLoadPackagesMultiModule(t *testing.T) {
	cases := []struct {
		name          string
		files         map[string]string
		tags          []string
		expectPkgs    map[string][]string
		testSubByPath bool
	}{{
		name: "one_pkg",
		files: map[string]string{
			"file.go": dedent.Dedent(`
				package p
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod": {"file.go"},
		},
	}, {
		name: "multi_pkg",
		files: map[string]string{
			"p1/file1.go": dedent.Dedent(`
				package p1
				var V string
			`),
			"p2/file2.go": dedent.Dedent(`
				package p2
				var V string
			`),
			"p3/file3.go": dedent.Dedent(`
				package p3
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1.go"},
			"example.com/mod/p2": {"file2.go"},
			"example.com/mod/p3": {"file3.go"},
		},
	}, {
		name: "multi_module_no_workspace",
		files: map[string]string{
			"p1/file1.go": dedent.Dedent(`
				package p1
				var V string
			`),
			"m2/go.mod": dedent.Dedent(`
				module example.com/m2
				go 1.18
			`),
			"m2/file2.go": dedent.Dedent(`
				package m2
				var V string
			`),
			"m3/go.mod": dedent.Dedent(`
				module example.com/m3
				go 1.18
			`),
			"m3/file3.go": dedent.Dedent(`
				package m3
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1.go"},
		},
	}, {
		name: "multi_module_workspace",
		files: map[string]string{
			"go.work": dedent.Dedent(`
				go 1.18
				use (
					.
					./m2
				)
				replace (
					example.com/m2 v0.0.0 => ./m2
				)
			`),
			"p1/file1.go": dedent.Dedent(`
				package p1
				var V string
			`),
			"m2/go.mod": dedent.Dedent(`
				module example.com/m2
				go 1.18
			`),
			"m2/file2.go": dedent.Dedent(`
				package m2
				var V string
			`),
			"m3/go.mod": dedent.Dedent(`
				module example.com/m3
				go 1.18
			`),
			"m3/file3.go": dedent.Dedent(`
				package m3
				var V string
			`),
		},
		expectPkgs: map[string][]string{
			"example.com/mod/p1": {"file1.go"},
			"example.com/m2":     {"file2.go"},
		},
	}}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			emit := emitter{
				tags: tc.tags,
			}

			// pushd
			dir := initModule(t, "example.com/mod", tc.files)
			if err := os.Chdir(dir); err != nil {
				t.Fatal(err)
			}

			for _, pattern := range [][]string{{"example.com/mod/...", "example.com/m2/..."}, {"./...", "./m2/..."}, {"all"}} {
				pkgs, err := emit.loadPackages(pattern...)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if want, got := tc.expectPkgs, pkgFiles(pkgs); !cmp.Equal(want, got) {
					t.Errorf("wrong result for pattern(s) %q:\n\twant: %v\n\t got: %v", pattern, want, got)
				}
			}

			// popd
			if err := os.Chdir(wd); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestVisitPackage(t *testing.T) {
	pkgpath := "example.com/mod/pkg"

	cases := []struct {
		name       string
		pkg        packages.Package
		initMap    func(pkgMap map[string]*packages.Package, pkg *packages.Package) // optional
		expectErrs bool
	}{{
		name: "already_present",
		pkg: packages.Package{
			PkgPath: pkgpath,
		},
		initMap: func(pkgMap map[string]*packages.Package, pkg *packages.Package) {
			pkgMap[pkgpath] = pkg
		},
	}, {
		name: "success",
		pkg: packages.Package{
			PkgPath: pkgpath,
		},
	}, {
		name: "list_error",
		pkg: packages.Package{
			PkgPath: pkgpath,
			Errors:  []packages.Error{{Kind: packages.ListError}},
		},
		expectErrs: true,
	}, {
		name: "parse_error",
		pkg: packages.Package{
			PkgPath: pkgpath,
			Errors: []packages.Error{{
				Kind: packages.ParseError,
			}},
		},
		expectErrs: true,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pkgMap := map[string]*packages.Package{}
			if tc.initMap != nil {
				tc.initMap(pkgMap, &tc.pkg)
			}
			emit := emitter{}

			ok := emit.visitPackage(&tc.pkg, pkgMap)
			if ok && tc.expectErrs {
				t.Errorf("unexpected success")
			}
			if !ok && !tc.expectErrs {
				t.Errorf("unexpected failure")
			}
			if want, got := 1, len(pkgMap); want != got {
				t.Errorf("unexpected number of packages: want %d, got %d, %v", want, got, pkgMap)
			}
			if p, found := pkgMap[pkgpath]; !found {
				t.Errorf("package %q not found in map: %v", pkgpath, pkgMap)
			} else if p != &tc.pkg {
				t.Errorf("package %q in map is different pointer: %v", pkgpath, p)
			}
		})
	}
}

func TestEmitMake(t *testing.T) {
	cases := []struct {
		name   string
		files  map[string]string
		tags   []string
		expect string
	}{{
		name: "one_pkg_no_imports",
		files: map[string]string{
			"file.go": dedent.Dedent(`
				package p
				var V string
			`),
		},
		expect: dedent.Dedent(`
			.go2make/by-pkg/./m2/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/./m3/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/_files: ./
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/_pkg: .go2make/by-pkg/example.com/mod/_files \
			  ./file.go
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./_pkg: .go2make/by-pkg/example.com/mod/_pkg
				@mkdir -p $(@D)
				@touch $@
		`),
	}, {
		name: "one_pkg_with_other_files",
		files: map[string]string{
			"README": "",
			"file.go": dedent.Dedent(`
				package p
				var V string
			`),
			"file_test.go": dedent.Dedent(`
				package p
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
			"_ignore.go": dedent.Dedent(`
				package p
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
		},
		expect: dedent.Dedent(`
			.go2make/by-pkg/./m2/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/./m3/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/_files: ./
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/_pkg: .go2make/by-pkg/example.com/mod/_files \
			  ./file.go
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./_pkg: .go2make/by-pkg/example.com/mod/_pkg
				@mkdir -p $(@D)
				@touch $@
		`),
	}, {
		name: "one_pkg_with_imports",
		files: map[string]string{
			"file.go": dedent.Dedent(`
				package p
				import "io"
				import "os"
				func init() { io.WriteString(os.Stdout, "") }
			`),
		},
		expect: dedent.Dedent(`
			.go2make/by-pkg/./m2/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/./m3/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/_files: ./
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/_pkg: .go2make/by-pkg/example.com/mod/_files \
			  ./file.go
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./_pkg: .go2make/by-pkg/example.com/mod/_pkg
				@mkdir -p $(@D)
				@touch $@
		`),
	}, {
		name: "multi_pkg_no_imports",
		files: map[string]string{
			"p1/file1.go": dedent.Dedent(`
				package p1
				var V string
			`),
			"p2/file2.go": dedent.Dedent(`
				package p2
				import "example.com/mod/p1"
				var V = p1.V
			`),
			"p3/file3.go": dedent.Dedent(`
				package p3
				import "example.com/mod/p1"
				import "example.com/mod/p2"
				var V = p1.V + p2.V
			`),
		},
		expect: dedent.Dedent(`
			.go2make/by-pkg/./m2/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/./m3/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/p1/_files: ./p1/
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/p1/_pkg: .go2make/by-pkg/example.com/mod/p1/_files \
			  ./p1/file1.go
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./p1/_pkg: .go2make/by-pkg/example.com/mod/p1/_pkg
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/p2/_files: ./p2/
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/p2/_pkg: .go2make/by-pkg/example.com/mod/p2/_files \
			  ./p2/file2.go \
			  .go2make/by-pkg/example.com/mod/p1/_pkg
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./p2/_pkg: .go2make/by-pkg/example.com/mod/p2/_pkg
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/p3/_files: ./p3/
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/p3/_pkg: .go2make/by-pkg/example.com/mod/p3/_files \
			  ./p3/file3.go \
			  .go2make/by-pkg/example.com/mod/p1/_pkg \
			  .go2make/by-pkg/example.com/mod/p2/_pkg
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./p3/_pkg: .go2make/by-pkg/example.com/mod/p3/_pkg
				@mkdir -p $(@D)
				@touch $@
		`),
	}, {
		name: "multi_module_workspace",
		files: map[string]string{
			"go.work": dedent.Dedent(`
				go 1.18
				use (
					.
					./m2
				)
				replace (
					example.com/m2 v0.0.0 => ./m2
				)
			`),
			"p1/file1.go": dedent.Dedent(`
				package p1
				import "example.com/m2"
				var V = m2.V
			`),
			"m2/go.mod": dedent.Dedent(`
				module example.com/m2
				go 1.18
			`),
			"m2/file2.go": dedent.Dedent(`
				package m2
				var V string
			`),
			"m3/go.mod": dedent.Dedent(`
				module example.com/m3
				go 1.18
			`),
			"m3/file3.go": dedent.Dedent(`
				package m3
				var V string
			`),
		},
		expect: dedent.Dedent(`
			.go2make/by-pkg/./m3/.../_pkg:
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/m2/_files: ./m2/
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/m2/_pkg: .go2make/by-pkg/example.com/m2/_files \
			  ./m2/file2.go
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./m2/_pkg: .go2make/by-pkg/example.com/m2/_pkg
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-pkg/example.com/mod/p1/_files: ./p1/
				@mkdir -p $(@D)
				@ls $</*.go | LC_ALL=C sort > $@.tmp
				@if ! cmp -s $@.tmp $@; then \
				    cat $@.tmp > $@; \
				fi
				@rm -f $@.tmp

			.go2make/by-pkg/example.com/mod/p1/_pkg: .go2make/by-pkg/example.com/mod/p1/_files \
			  ./p1/file1.go \
			  .go2make/by-pkg/example.com/m2/_pkg
				@mkdir -p $(@D)
				@touch $@

			.go2make/by-path/./p1/_pkg: .go2make/by-pkg/example.com/mod/p1/_pkg
				@mkdir -p $(@D)
				@touch $@
		`),
	}}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := initModule(t, "example.com/mod", tc.files)

			emit := emitter{
				stateDir:     ".go2make",
				relPath:      dir,
				ignoreErrors: true, // easier output comparison
			}

			// pushd
			if err := os.Chdir(dir); err != nil {
				t.Fatal(err)
			}

			pkgs, err := emit.loadPackages("./...", "./m2/...", "./m3/...")
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			pkgMap := emit.visitPackages(pkgs)
			if pkgMap == nil {
				t.Errorf("unexpected error")
			}
			buf := bytes.Buffer{}
			emit.emitMake(&buf, pkgMap)
			if want, got := strings.Trim(tc.expect, "\n"), strings.Trim(buf.String(), "\n"); want != got {
				t.Errorf("wrong result:\n%s", cmp.Diff(want, got))
			}

			// popd
			if err := os.Chdir(wd); err != nil {
				t.Fatal(err)
			}
		})
	}
}
