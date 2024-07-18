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
	"slices"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// MaxLength verifies that the specified value is not longer than max
// characters.
func MaxLength(fldPath *field.Path, value string, max int) field.ErrorList {
	if len(value) > max {
		return field.ErrorList{field.Invalid(fldPath, value, MaxLenError(max))}
	}
	return nil
}

// Enum verifies that the specified value is one of the valid symbols.
// TODO: scanning the symbols list is O(n) vs the O(1) set.has check used by hand written validation.
func Enum[T ~string](fldPath *field.Path, value T, symbols ...string) field.ErrorList {
	valueString := string(value)
	if !slices.Contains(symbols, valueString) {
		return field.ErrorList{field.NotSupported(fldPath, value, symbols)}
	}
	return nil
}

// Required verifies that the specified value is not the zero-value for its
// type.
func Required[T comparable](fldPath *field.Path, value T) field.ErrorList {
	var zero T
	if value == zero {
		return field.ErrorList{field.Required(fldPath, "")}
	}
	return nil
}
