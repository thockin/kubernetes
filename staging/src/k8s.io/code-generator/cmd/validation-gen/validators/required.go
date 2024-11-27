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

package validators

import (
	"fmt"

	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/types"
)

func init() {
	AddToRegistry(InitRequiredDeclarativeValidator)
	AddToRegistry(InitForbiddenDeclarativeValidator)
	AddToRegistry(InitOptionalDeclarativeValidator)
}

func InitRequiredDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &requiredDeclarativeValidator{}
}

type requiredDeclarativeValidator struct{}

const (
	requiredTagName = "k8s:required"
)

var (
	requiredValueValidator   = types.Name{Package: libValidationPkg, Name: "RequiredValue"}
	requiredPointerValidator = types.Name{Package: libValidationPkg, Name: "RequiredPointer"}
	requiredSliceValidator   = types.Name{Package: libValidationPkg, Name: "RequiredSlice"}
	requiredMapValidator     = types.Name{Package: libValidationPkg, Name: "RequiredMap"}
)

func (requiredDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	_, required := gengo.ExtractCommentTags("+", comments)[requiredTagName]
	if !required {
		return Validations{}, nil
	}
	// Most validators don't care whether the value they are validating was
	// originally defined as a value-type or a pointer-type in the API.  This
	// one does.  Since Go doesn't do partial specialization of templates, we
	// do manual dispatch here.
	for t.Kind == types.Alias {
		t = t.Underlying
	}
	switch t.Kind {
	case types.Slice:
		return Validations{Functions: []FunctionGen{Function(requiredTagName, ShortCircuit, requiredSliceValidator)}}, nil
	case types.Map:
		return Validations{Functions: []FunctionGen{Function(requiredTagName, ShortCircuit, requiredMapValidator)}}, nil
	case types.Pointer:
		return Validations{Functions: []FunctionGen{Function(requiredTagName, ShortCircuit, requiredPointerValidator)}}, nil
	case types.Struct:
		// The +required tag on a non-pointer struct is only for documentation.
		// We don't perform validation here and defer the validation to
		// the struct's fields.
		return Validations{}, nil
	}
	return Validations{Functions: []FunctionGen{Function(requiredTagName, ShortCircuit, requiredValueValidator)}}, nil
}

func (requiredDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         requiredTagName,
		Description: "Indicates that a field is required to be specified.",
		Contexts:    []TagContext{TagContextType, TagContextField},
	}}
}

func InitForbiddenDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &forbiddenDeclarativeValidator{}
}

type forbiddenDeclarativeValidator struct{}

const (
	forbiddenTagName = "k8s:forbidden"
)

var (
	forbiddenValueValidator   = types.Name{Package: libValidationPkg, Name: "ForbiddenValue"}
	forbiddenPointerValidator = types.Name{Package: libValidationPkg, Name: "ForbiddenPointer"}
	forbiddenSliceValidator   = types.Name{Package: libValidationPkg, Name: "ForbiddenSlice"}
	forbiddenMapValidator     = types.Name{Package: libValidationPkg, Name: "ForbiddenMap"}
)

func (forbiddenDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	_, forbidden := gengo.ExtractCommentTags("+", comments)[forbiddenTagName]
	if !forbidden {
		return Validations{}, nil
	}
	// Most validators don't care whether the value they are validating was
	// originally defined as a value-type or a pointer-type in the API.  This
	// one does.  Since Go doesn't do partial specialization of templates, we
	// do manual dispatch here.
	for t.Kind == types.Alias {
		t = t.Underlying
	}
	// Forbidden is weird.  Each of these emits two checks, which are polar
	// opposites.  If the field fails the forbidden check, it will
	// short-circuit and not run the optional check.  If it passes the
	// forbidden check, it must not be specified, so it will "fail" the
	// optional check and short-circuit (but without error).  Why?  For
	// example, this prevents any further validation from trying to run on a
	// nil pointer.
	switch t.Kind {
	case types.Slice:
		return Validations{
			Functions: []FunctionGen{
				Function(forbiddenTagName, ShortCircuit, forbiddenSliceValidator),
				Function(forbiddenTagName, ShortCircuit|NonError, optionalSliceValidator),
			},
		}, nil
	case types.Map:
		return Validations{
			Functions: []FunctionGen{
				Function(forbiddenTagName, ShortCircuit, forbiddenMapValidator),
				Function(forbiddenTagName, ShortCircuit|NonError, optionalMapValidator),
			},
		}, nil
	case types.Pointer:
		return Validations{
			Functions: []FunctionGen{
				Function(forbiddenTagName, ShortCircuit, forbiddenPointerValidator),
				Function(forbiddenTagName, ShortCircuit|NonError, optionalPointerValidator),
			},
		}, nil
	case types.Struct:
		// The +forbidden tag on a non-pointer struct is not supported.
		// If you encounter this error and believe you have a valid use case
		// for forbiddening a non-pointer struct, please let us know! We need
		// to understand your scenario to determine if we need to adjust
		// this behavior or provide alternative validation mechanisms.
		return Validations{}, fmt.Errorf("non-pointer structs cannot use the %q tag", forbiddenTagName)
	}
	return Validations{
		Functions: []FunctionGen{
			Function(forbiddenTagName, ShortCircuit, forbiddenValueValidator),
			Function(forbiddenTagName, ShortCircuit|NonError, optionalValueValidator),
		},
	}, nil
}

func (forbiddenDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         forbiddenTagName,
		Description: "Indicates that a field is forbidden to be specified.",
		Contexts:    []TagContext{TagContextType, TagContextField},
	}}
}

func InitOptionalDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &optionalDeclarativeValidator{}
}

type optionalDeclarativeValidator struct{}

const (
	optionalTagName = "k8s:optional"
)

var (
	optionalValueValidator   = types.Name{Package: libValidationPkg, Name: "OptionalValue"}
	optionalPointerValidator = types.Name{Package: libValidationPkg, Name: "OptionalPointer"}
	optionalSliceValidator   = types.Name{Package: libValidationPkg, Name: "OptionalSlice"}
	optionalMapValidator     = types.Name{Package: libValidationPkg, Name: "OptionalMap"}
)

func (optionalDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	_, optional := gengo.ExtractCommentTags("+", comments)[optionalTagName]
	if !optional {
		return Validations{}, nil
	}
	// Most validators don't care whether the value they are validating was
	// originally defined as a value-type or a pointer-type in the API.  This
	// one does.  Since Go doesn't do partial specialization of templates, we
	// do manual dispatch here.
	for t.Kind == types.Alias {
		t = t.Underlying
	}
	switch t.Kind {
	case types.Slice:
		return Validations{Functions: []FunctionGen{Function(optionalTagName, ShortCircuit|NonError, optionalSliceValidator)}}, nil
	case types.Map:
		return Validations{Functions: []FunctionGen{Function(optionalTagName, ShortCircuit|NonError, optionalMapValidator)}}, nil
	case types.Pointer:
		return Validations{Functions: []FunctionGen{Function(optionalTagName, ShortCircuit|NonError, optionalPointerValidator)}}, nil
	case types.Struct:
		// Specifying that a non-pointer struct is optional doesn't actually
		// make sense technically almost ever, and is better described as a
		// union inside the struct. It does, however, make sense as
		// documentation.
		return Validations{}, nil
	}
	return Validations{Functions: []FunctionGen{Function(optionalTagName, ShortCircuit|NonError, optionalValueValidator)}}, nil
}

func (optionalDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         optionalTagName,
		Description: "Indicates that a field is optional.",
		Contexts:    []TagContext{TagContextType, TagContextField},
	}}
}
