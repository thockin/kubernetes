/*
Copyright 2015 The Kubernetes Authors.

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

package generator

import (
	"go/token"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"

	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
)

// NewImportTrackerForPackage creates a new import tracker which is aware
// of a generator's output package. The tracker will not add import lines
// when symbols or types are added from the same package, and LocalNameOf
// will return empty string for the output package.
//
// e.g.:
//
//	tracker := NewImportTrackerForPackage("bar.com/pkg/foo")
//	tracker.AddSymbol(types.Name{"bar.com/pkg/foo.MyType"})
//	tracker.AddSymbol(types.Name{"bar.com/pkg/baz.MyType"})
//	tracker.AddSymbol(types.Name{"bar.com/pkg/baz/baz.MyType"})
//
//	tracker.LocalNameOf("bar.com/pkg/foo") -> ""
//	tracker.LocalNameOf("bar.com/pkg/baz") -> "baz"
//	tracker.LocalNameOf("bar.com/pkg/baz/baz") -> "bazbaz"
//	tracker.ImportLines() -> {`baz "bar.com/pkg/baz"`, `bazbaz "bar.com/pkg/baz/baz"`}
func NewImportTrackerForPackage(local string, typesToAdd ...*types.Type) *namer.DefaultImportTracker {
	tracker := namer.NewDefaultImportTracker(types.Name{Package: local})
	tracker.IsInvalidType = func(*types.Type) bool { return false }
	tracker.LocalName = func(name types.Name) string { return golangTrackerLocalName(&tracker, name) }
	tracker.PrintImport = func(path, name string) string {
		if filepath.Base(path) == name {
			return `"` + path + `"`
		} else {
			return name + " \"" + path + "\""
		}
	}

	tracker.AddTypes(typesToAdd...)
	return &tracker
}

func NewImportTracker(typesToAdd ...*types.Type) *namer.DefaultImportTracker {
	return NewImportTrackerForPackage("", typesToAdd...)
}

func golangTrackerLocalName(tracker namer.ImportTracker, t types.Name) string {
	path := t.Package

	// Using backslashes in package names causes gengo to produce Go code which
	// will not compile with the gc compiler. See the comment on GoSeperator.
	if strings.ContainsRune(path, '\\') {
		klog.Warningf("Warning: backslash used in import path '%v', this is unsupported.\n", path)
	}

	dirs := strings.Split(path, namer.GoSeperator)

	// By default we will try to construct a local name with as few as 1
	// directories (e.g. "net/http" -> "http").
	min := 1
	// If it seems like "example.com/some/pkg" then we try for longer names.
	if strings.Contains(dirs[0], ".") {
		if len(dirs) >= 4 {
			// e.g. [ "example.com", "some", "long", "pkg" ]
			min = 3
		} else if len(dirs) == 3 {
			// e.g. [ "example.com", "short", "pkg" ]
			min = 2
		}
	}

	// Build up a potential local-name from directories, and see if it works.
	for n := len(dirs) - min; n >= 0; n-- {
		// follow kube convention of not having anything between directory names
		name := strings.Join(dirs[n:], "")
		name = strings.Replace(name, "_", "", -1)
		// These characters commonly appear in import paths for go
		// packages, but aren't legal go names. So we'll sanitize.
		name = strings.Replace(name, ".", "", -1)
		name = strings.Replace(name, "-", "", -1)
		if _, found := tracker.PathOf(name); found {
			// This name collides with some other package
			continue
		}

		// If the import name is a Go keyword, prefix with an underscore.
		if token.Lookup(name).IsKeyword() {
			name = "_" + name
		}
		return name
	}
	panic("can't find import for " + path)
}
