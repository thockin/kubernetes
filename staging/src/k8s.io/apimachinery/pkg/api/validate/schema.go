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
	"golang.org/x/exp/constraints"
	"k8s.io/apimachinery/pkg/api/operation"
	"k8s.io/apimachinery/pkg/api/validate/content"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// MaxLength verifies that the specified value is not longer than max
// characters.
func MaxLength(_ operation.Context, fldPath *field.Path, value, _ *string, max int) field.ErrorList {
	if value == nil {
		return nil
	}
	if len(*value) > max {
		return field.ErrorList{field.Invalid(fldPath, *value, content.MaxLenError(max))}
	}
	return nil
}

// Minimum verifies that the specified value is greater than or equal to min.
func Minimum[T constraints.Integer](_ operation.Context, fldPath *field.Path, value, _ *T, min T) field.ErrorList {
	if value == nil {
		return nil
	}
	if *value < min {
		return field.ErrorList{field.Invalid(fldPath, *value, content.MinError(min))}
	}
	return nil
}

// RequiredValue verifies that the specified value is not the zero-value for
// its type.
func RequiredValue[T comparable](_ operation.Context, fldPath *field.Path, value, _ *T) field.ErrorList {
	var zero T
	if *value != zero {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "")}
}

// RequiredPointer verifies that the specified pointer is not nil.
func RequiredPointer[T any](_ operation.Context, fldPath *field.Path, value, _ *T) field.ErrorList {
	if value != nil {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "")}
}

// RequiredSlice verifies that the specified slice is not empty.
func RequiredSlice[T any](_ operation.Context, fldPath *field.Path, value, _ []T) field.ErrorList {
	if len(value) > 0 {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "")}
}

// RequiredMap verifies that the specified map is not empty.
func RequiredMap[K comparable, T any](_ operation.Context, fldPath *field.Path, value, _ map[K]T) field.ErrorList {
	if len(value) > 0 {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "")}
}

// ForbiddenValue verifies that the specified value is the zero-value for its
// type.
func ForbiddenValue[T comparable](_ operation.Context, fldPath *field.Path, value, _ *T) field.ErrorList {
	var zero T
	if *value == zero {
		return nil
	}
	return field.ErrorList{field.Forbidden(fldPath, "")}
}

// ForbiddenPointer verifies that the specified pointer is nil.
func ForbiddenPointer[T any](_ operation.Context, fldPath *field.Path, value, _ *T) field.ErrorList {
	if value == nil {
		return nil
	}
	return field.ErrorList{field.Forbidden(fldPath, "")}
}

// ForbiddenSlice verifies that the specified slice is empty.
func ForbiddenSlice[T any](_ operation.Context, fldPath *field.Path, value, _ []T) field.ErrorList {
	if len(value) == 0 {
		return nil
	}
	return field.ErrorList{field.Forbidden(fldPath, "")}
}

// RequiredMap verifies that the specified map is empty.
func ForbiddenMap[K comparable, T any](_ operation.Context, fldPath *field.Path, value, _ map[K]T) field.ErrorList {
	if len(value) == 0 {
		return nil
	}
	return field.ErrorList{field.Forbidden(fldPath, "")}
}

// OptionalValue verifies that the specified value is not the zero-value for
// its type. This is identical to RequiredValue, but the caller should treat an
// error here as an indication that the optional value was not specified.
func OptionalValue[T comparable](_ operation.Context, fldPath *field.Path, value, _ *T) field.ErrorList {
	var zero T
	if *value != zero {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "optional value was not specified")}
}

// OptionalPointer verifies that the specified pointer is not nil. This is
// identical to RequiredPointer, but the caller should treat an error here as an
// indication that the optional value was not specified.
func OptionalPointer[T any](_ operation.Context, fldPath *field.Path, value, _ *T) field.ErrorList {
	if value != nil {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "optional value was not specified")}
}

// OptionalSlice verifies that the specified slice is not empty. This is
// identical to RequiredSlice, but the caller should treat an error here as an
// indication that the optional value was not specified.
func OptionalSlice[T any](_ operation.Context, fldPath *field.Path, value, _ []T) field.ErrorList {
	if len(value) > 0 {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "optional value was not specified")}
}

// OptionalMap verifies that the specified map is not empty. This is identical
// to RequiredMap, but the caller should treat an error here as an indication that
// the optional value was not specified.
func OptionalMap[K comparable, T any](_ operation.Context, fldPath *field.Path, value, _ map[K]T) field.ErrorList {
	if len(value) > 0 {
		return nil
	}
	return field.ErrorList{field.Required(fldPath, "optional value was not specified")}
}

// MaxItems verifies that the specified slice is not longer than max items.
func MaxItems[T any](_ operation.Context, fldPath *field.Path, value, _ []T, max int) field.ErrorList {
	if len(value) > max {
		return field.ErrorList{field.TooMany(fldPath, len(value), max)}
	}
	return nil
}
