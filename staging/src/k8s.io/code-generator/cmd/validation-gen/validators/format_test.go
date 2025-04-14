/*
Copyright 2025 The Kubernetes Authors.

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

package validators

import "testing"

func TestParseFormat(t *testing.T) {
	testCases := []struct {
		input  string
		format string
		arg    string
	}{{
		"foo", "foo", "",
	}, {
		"foo()", "foo", "",
	}, {
		"foo(bar)", "foo", "bar",
	}, {
		"foo(bar qux)", "foo", "bar qux",
	}, {
		// This is terrible, but just not something we care about for now.
		"foo(bar())", "foo(bar())", "",
	}}

	for i, tc := range testCases {
		format, arg := parseFormat(tc.input)
		if format != tc.format || arg != tc.arg {
			t.Errorf("case[%d]: expected (%q, %q), got (%q, %q)", i, tc.format, tc.arg, format, arg)
		}
	}
}
