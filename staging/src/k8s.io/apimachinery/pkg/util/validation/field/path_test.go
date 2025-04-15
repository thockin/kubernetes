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

package field

import (
	"testing"
)

func TestPath(t *testing.T) {
	testCases := []struct {
		op       func(*Path) *Path
		expected string
	}{
		{
			func(p *Path) *Path { return p },
			"root",
		},
		{
			func(p *Path) *Path { return p.Child("first") },
			"root.first",
		},
		{
			func(p *Path) *Path { return p.Child("second") },
			"root.first.second",
		},
		{
			func(p *Path) *Path { return p.Index(0) },
			"root.first.second[0]",
		},
		{
			func(p *Path) *Path { return p.Child("third") },
			"root.first.second[0].third",
		},
		{
			func(p *Path) *Path { return p.Index(93) },
			"root.first.second[0].third[93]",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first.second[0].third",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first.second[0]",
		},
		{
			func(p *Path) *Path { return p.Key("key") },
			"root.first.second[0][key]",
		},
	}

	root := NewPath("root")
	p := root
	for i, tc := range testCases {
		p = tc.op(p)
		if p.String() != tc.expected {
			t.Errorf("[%d] Expected %q, got %q", i, tc.expected, p.String())
		}
		if p.Root() != root {
			t.Errorf("[%d] Wrong root: %#v", i, p.Root())
		}
	}
}

func TestPathMultiArg(t *testing.T) {
	testCases := []struct {
		op       func(*Path) *Path
		expected string
	}{
		{
			func(p *Path) *Path { return p },
			"root.first",
		},
		{
			func(p *Path) *Path { return p.Child("second", "third") },
			"root.first.second.third",
		},
		{
			func(p *Path) *Path { return p.Index(0) },
			"root.first.second.third[0]",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first.second.third",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first.second",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root",
		},
	}

	root := NewPath("root", "first")
	p := root
	for i, tc := range testCases {
		p = tc.op(p)
		if p.String() != tc.expected {
			t.Errorf("[%d] Expected %q, got %q", i, tc.expected, p.String())
		}
		if p.Root() != root.Root() {
			t.Errorf("[%d] Wrong root: %#v", i, p.Root())
		}
	}
}

func TestPath2(t *testing.T) {
	testCases := []struct {
		op       func(Path2) Path2
		expected string
	}{
		{
			func(p Path2) Path2 { return p },
			"root",
		},
		{
			func(p Path2) Path2 { return p.Child("first") },
			"root.first",
		},
		{
			func(p Path2) Path2 { return p.Child("second") },
			"root.first.second",
		},
		{
			func(p Path2) Path2 { return p.Index(0) },
			"root.first.second[0]",
		},
		{
			func(p Path2) Path2 { return p.Child("third") },
			"root.first.second[0].third",
		},
		{
			func(p Path2) Path2 { return p.Index(93) },
			"root.first.second[0].third[93]",
		},
		{
			func(p Path2) Path2 { return p.Child("fourth") },
			"root.first.second[0].third[93].fourth",
		},
		{
			func(p Path2) Path2 { return p.Key("k") },
			"root.first.second[0].third[93].fourth[k]",
		},
		{
			func(p Path2) Path2 { return p.Root() },
			"root",
		},
	}

	root := NewPath2("root")
	p := root
	for i, tc := range testCases {
		p = tc.op(p)
		if p.String() != tc.expected {
			t.Errorf("[%d] Expected %q, got %q", i, tc.expected, p.String())
		}
		if p.Root().String() != "root" {
			t.Errorf("[%d] Wrong root: %#v", i, p.Root())
		}
	}
}

func BenchmarkPath(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		root := NewPath("root")
		f1(root, i%8)
	}
}

func BenchmarkPath2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		root := NewPath2("root")
		f2(root, i%8)
	}
}

//go:noinline
func f1(p *Path, n int) string {
	if n == 0 {
		return p.String()
	}
	return f1(p.Child("child"), n-1)
}

//go:noinline
func f2(p Path2, n int) string {
	if n == 0 {
		return p.String()
	}
	return f2(p.Child("child"), n-1)
}

/*
func TestPathMultiArg(t *testing.T) {
	testCases := []struct {
		op       func(*Path) *Path
		expected string
	}{
		{
			func(p *Path) *Path { return p },
			"root.first",
		},
		{
			func(p *Path) *Path { return p.Child("second", "third") },
			"root.first.second.third",
		},
		{
			func(p *Path) *Path { return p.Index(0) },
			"root.first.second.third[0]",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first.second.third",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first.second",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root.first",
		},
		{
			func(p *Path) *Path { return p.parent },
			"root",
		},
	}

	root := NewPath("root", "first")
	p := root
	for i, tc := range testCases {
		p = tc.op(p)
		if p.String() != tc.expected {
			t.Errorf("[%d] Expected %q, got %q", i, tc.expected, p.String())
		}
		if p.Root() != root.Root() {
			t.Errorf("[%d] Wrong root: %#v", i, p.Root())
		}
	}
}
*/
