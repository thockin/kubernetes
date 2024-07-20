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
	AddToRegistryPriority(InitValidateTrueDeclarativeValidator)
	AddToRegistryPriority(InitValidateFalseDeclarativeValidator)
}

func InitValidateTrueDeclarativeValidator(c *generator.Context) DeclarativeValidator {
	return &fixedResultDeclarativeValidator{true}
}

func InitValidateFalseDeclarativeValidator(c *generator.Context) DeclarativeValidator {
	return &fixedResultDeclarativeValidator{false}
}

type fixedResultDeclarativeValidator struct {
	result bool
}

const (
	// These tags can take no value or a string, which will be used in the
	// error message.
	validateTrueTagName  = "validateTrue"  // TODO: also support k8s:...
	validateFalseTagName = "validateFalse" // TODO: also support k8s:...
)

var (
	fixedResultValidator = types.Name{Package: libValidationPkg, Name: "FixedResult"}
)

func (v fixedResultDeclarativeValidator) ExtractValidations(field string, t *types.Type, comments []string) ([]FunctionGen, error) {
	var result []FunctionGen

	if v.result {
		vals, fixedTrue := gengo.ExtractCommentTags("+", comments)[validateTrueTagName]
		if fixedTrue {
			for _, v := range vals {
				result = append(result, Function(validateTrueTagName, fixedResultValidator, true, v))
			}
		}
	} else {
		vals, fixedFalse := gengo.ExtractCommentTags("+", comments)[validateFalseTagName]
		if fixedFalse {
			for _, v := range vals {
				result = append(result, Function(validateFalseTagName, fixedResultValidator, false, v))
			}
		}
	}

	return result, nil
}
