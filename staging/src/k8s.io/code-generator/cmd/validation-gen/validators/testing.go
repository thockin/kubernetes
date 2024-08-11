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
	"encoding/json"
	"fmt"

	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/types"
)

func init() {
	AddToRegistry(InitValidateTrueDeclarativeValidator)
	AddToRegistry(InitValidateFalseDeclarativeValidator)
	AddToRegistry(InitValidateErrorDeclarativeValidator)
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

func InitValidateErrorDeclarativeValidator(c *generator.Context) DeclarativeValidator {
	return &errorDeclarativeValidator{}
}

type errorDeclarativeValidator struct {
}

const (
	// These tags can take no value or a quoted string or a JSON object, which will be used in the
	// error message.  The JSON object schema is:
	//   {
	//     "flags": <list-of-string>  # optional: "PtrOK" or "IsFatal"
	//     "msg":   <string>          # required
	//   }
	validateTrueTagName  = "validateTrue"  // TODO: also support k8s:...
	validateFalseTagName = "validateFalse" // TODO: also support k8s:...

	// This tag always returns an error from ExtractValidations.
	validateErrorTagName = "validateError" // TODO: also support k8s:...
)

var (
	fixedResultValidator = types.Name{Package: libValidationPkg, Name: "FixedResult"}
)

func (v fixedResultDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) ([]FunctionGen, error) {
	var result []FunctionGen

	if v.result {
		vals := gengo.ExtractCommentTags("+", comments)[validateTrueTagName]
		for _, val := range vals {
			flags, msg, err := v.parseTagVal(val)
			if err != nil {
				return nil, fmt.Errorf("can't extract +%s tag: %w", validateTrueTagName, err)
			}
			result = append(result, Function(validateTrueTagName, flags, fixedResultValidator, true, msg))
		}
	} else {
		vals := gengo.ExtractCommentTags("+", comments)[validateFalseTagName]
		for _, val := range vals {
			flags, msg, err := v.parseTagVal(val)
			if err != nil {
				return nil, fmt.Errorf("can't extract +%s tag: %w", validateFalseTagName, err)
			}
			result = append(result, Function(validateFalseTagName, flags, fixedResultValidator, false, msg))
		}
	}

	return result, nil
}

func (_ fixedResultDeclarativeValidator) parseTagVal(in string) (FunctionFlags, string, error) {
	type payload struct {
		Flags []string `json:"flags"`
		Msg   string   `json:"msg"`
	}
	// We expect either a string (maybe empty) or a JSON object.
	if len(in) == 0 {
		return 0, "", nil
	}
	var pl payload
	if err := json.Unmarshal([]byte(in), &pl); err != nil {
		s := ""
		if err := json.Unmarshal([]byte(in), &s); err != nil {
			return 0, "", fmt.Errorf("error parsing JSON value: %v (%q)", err, in)
		}
		return 0, s, nil
	}
	// The msg field is required in JSON mode.
	if pl.Msg == "" {
		return 0, "", fmt.Errorf("JSON msg is required")
	}
	var flags FunctionFlags
	for _, fl := range pl.Flags {
		switch fl {
		case "IsFatal":
			flags |= IsFatal
		case "PtrOK":
			flags |= PtrOK
		default:
			return 0, "", fmt.Errorf("unknown flag: %q", fl)
		}
	}

	return flags, pl.Msg, nil
}

func (v errorDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) ([]FunctionGen, error) {
	vals, found := gengo.ExtractCommentTags("+", comments)[validateErrorTagName]
	if found {
		return nil, fmt.Errorf("forced error: %q", vals)
	}
	return nil, nil
}
