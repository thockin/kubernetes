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
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/types"
)

func init() {
	AddToRegistry(InitRequiredDeclarativeValidator)
	AddToRegistry(InitOptionalDeclarativeValidator)
}

func InitRequiredDeclarativeValidator(c *generator.Context) DeclarativeValidator {
	return &requiredDeclarativeValidator{}
}

type requiredDeclarativeValidator struct{}

const (
	requiredTagName = "required" // TODO: also support k8s:required
)

var (
	requiredValidator = types.Name{Package: libValidationPkg, Name: "Required"}
)

func (requiredDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	_, required := gengo.ExtractCommentTags("+", comments)[requiredTagName]
	if !required {
		return Validations{}, nil
	}
	return Validations{Functions: []FunctionGen{Function(requiredTagName, ShortCircuit, requiredValidator)}}, nil
}

func (requiredDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         requiredTagName,
		Description: "Indicates that a field is required to be specified.",
		Contexts:    []TagContext{TagContextType, TagContextField},
	}}
}

func InitOptionalDeclarativeValidator(c *generator.Context) DeclarativeValidator {
	return &optionalDeclarativeValidator{}
}

type optionalDeclarativeValidator struct{}

const (
	optionalTagName = "optional" // TODO: also support k8s:optional
)

var (
	optionalValidator = types.Name{Package: libValidationPkg, Name: "Optional"}
)

func (optionalDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	_, optional := gengo.ExtractCommentTags("+", comments)[optionalTagName]
	if !optional {
		return Validations{}, nil
	}
	return Validations{Functions: []FunctionGen{Function(optionalTagName, ShortCircuit|NonError, optionalValidator)}}, nil
}

func (optionalDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         optionalTagName,
		Description: "Indicates that a field is optional.",
		Contexts:    []TagContext{TagContextType, TagContextField},
	}}
}
