/*
Copyright 2024 The Kubernetes Authors.

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

package validate

import (
	"context"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/api/operation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestDNS1123Label(t *testing.T) {
	ctx := context.Background()
	fldPath := field.NewPath("test")

	testCases := []struct {
		name     string
		input    string
		wantErrs field.ErrorList
	}{{
		name:     "valid label",
		input:    "valid-label",
		wantErrs: nil,
	}, {
		name:     "valid single character label",
		input:    "a",
		wantErrs: nil,
	}, {
		name:     "valid label with numbers",
		input:    "123-abc",
		wantErrs: nil,
	}, {
		name:  "invalid: uppercase characters",
		input: "Invalid-Label",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}, {
		name:  "invalid: starts with dash",
		input: "-invalid-label",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}, {
		name:  "invalid: ends with dash",
		input: "invalid-label-",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}, {
		name:  "invalid: contains dots",
		input: "invalid.label",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}, {
		name:  "invalid: contains special characters",
		input: "invalid@label",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}, {
		name:  "invalid: too long",
		input: "a" + strings.Repeat("b", 62) + "c", // 64 characters
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}, {
		name:     "valid: max length",
		input:    "a" + strings.Repeat("b", 61) + "c", // 63 characters
		wantErrs: nil,
	}, {
		name:  "invalid: empty string",
		input: "",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-label"),
		},
	}}

	matcher := field.ErrorMatcher{}.ByType().ByField().ByOrigin()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.input
			gotErrs := DNSLabel(ctx, operation.Operation{}, fldPath, &value, nil)

			matcher.Test(t, tc.wantErrs, gotErrs)
		})
	}
}

func TestDNS1123Subdomain(t *testing.T) {
	ctx := context.Background()
	fldPath := field.NewPath("test")

	testCases := []struct {
		name     string
		input    string
		wantErrs field.ErrorList
	}{{
		name:     "valid single label",
		input:    "valid-label",
		wantErrs: nil,
	}, {
		name:     "valid subdomain",
		input:    "this-is.a-valid.subdomain",
		wantErrs: nil,
	}, {
		name:     "valid single character elements",
		input:    "a.b.c",
		wantErrs: nil,
	}, {
		name:     "valid elements with numbers",
		input:    "123.abc-123.456-def",
		wantErrs: nil,
	}, {
		name:     "all number elements",
		input:    "1.2.3.4",
		wantErrs: nil,
	}, {
		name:  "invalid: uppercase characters",
		input: "Invalid.Subdomain",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name:  "invalid: starts with dash",
		input: "this-is.-an-invalid.subdomain",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name:  "invalid: ends with dash",
		input: "this-is.an-invalid-.subdomain",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name:  "invalid: contains double dots",
		input: "invalid..subdomain",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name:  "invalid: contains special characters",
		input: "inv@lid.subdoma!n",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name:  "invalid: too long single label",
		input: "a" + strings.Repeat("b", 252) + "c", // 254 characters
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name: "invalid: too long multiple labels",
		input: strings.Join([]string{
			strings.Repeat("a", 60), // 61 with the "."
			strings.Repeat("b", 60), // 122 with the "."
			strings.Repeat("c", 60), // 183 with the "."
			strings.Repeat("d", 60), // 244 with the "."
			strings.Repeat("e", 10), // 254 characters
		}, "."),
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}, {
		name:     "valid: max length single label",     // supported for compat
		input:    "a" + strings.Repeat("b", 251) + "c", // 253 characters
		wantErrs: nil,
	}, {
		name:  "invalid: empty string",
		input: "",
		wantErrs: field.ErrorList{
			field.Invalid(fldPath, nil, "").WithOrigin("format=dns-subdomain"),
		},
	}}

	matcher := field.ErrorMatcher{}.ByType().ByField().ByOrigin()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.input
			gotErrs := DNSSubdomain(ctx, operation.Operation{}, fldPath, &value, nil)

			matcher.Test(t, tc.wantErrs, gotErrs)
		})
	}
}

func TestMaskTrailingDash(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{{
		name:     "empty",
		input:    "",
		expected: "",
	}, {
		name:     "dash",
		input:    "-",
		expected: "-",
	}, {
		name:     "no dash",
		input:    "foo",
		expected: "foo",
	}, {
		name:     "leading dash",
		input:    "-foo",
		expected: "-foo",
	}, {
		name:     "trailing dash",
		input:    "foo-",
		expected: "fox",
	}, {
		name:     "one byte with trailing dash",
		input:    "b-",
		expected: "x",
	}, {
		name:     "multiple trailing dash",
		input:    "foo---",
		expected: "foo-x",
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maskTrailingDash(tc.input)
			if result != tc.expected {
				t.Errorf("expected: %q, got: %q", tc.expected, result)
			}
		})
	}
}

func TestGenerateName(t *testing.T) {
	ctx := context.Background()
	fldPath := field.NewPath("test")

	testCases := []struct {
		name     string
		input    string
		expected string
	}{{
		name:     "empty",
		input:    "",
		expected: "",
	}, {
		name:     "dash",
		input:    "-",
		expected: "-",
	}, {
		name:     "no dash",
		input:    "foo",
		expected: "foo",
	}, {
		name:     "leading dash",
		input:    "-foo",
		expected: "-foo",
	}, {
		name:     "trailing dash",
		input:    "foo-",
		expected: "fox",
	}, {
		name:     "one byte with trailing dash",
		input:    "b-",
		expected: "x",
	}, {
		name:     "multiple trailing dash",
		input:    "foo---",
		expected: "foo-x",
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wasCalled := false
			fn := func(_ context.Context, _ operation.Operation, _ *field.Path, val, _ *string) field.ErrorList {
				wasCalled = true
				if *val != tc.expected {
					t.Errorf("expected: %q, got: %q", tc.expected, *val)
				}
				return nil
			}

			value := tc.input
			gotErrs := GenerateName(ctx, operation.Operation{}, fldPath, &value, nil, fn)
			if len(gotErrs) != 0 {
				t.Errorf("unexpected errors: %v", gotErrs)
			}
			if wasCalled != true {
				t.Errorf("function was not called")
			}
		})
	}
}
