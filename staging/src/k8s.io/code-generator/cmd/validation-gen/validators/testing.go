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
	"strings"

	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/types"
)

func init() {
	AddToRegistry(InitValidateTrueDeclarativeValidator)
	AddToRegistry(InitValidateFalseDeclarativeValidator)
	AddToRegistry(InitValidateErrorDeclarativeValidator)
}

func InitValidateTrueDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &fixedResultDeclarativeValidator{true}
}

func InitValidateFalseDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &fixedResultDeclarativeValidator{false}
}

type fixedResultDeclarativeValidator struct {
	result bool
}

func InitValidateErrorDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &errorDeclarativeValidator{}
}

type errorDeclarativeValidator struct {
}

const (
	validateTrueTagName  = "k8s:validateTrue"
	validateFalseTagName = "k8s:validateFalse"

	// This tag always returns an error from ExtractValidations.
	validateErrorTagName = "k8s:validateError"
)

var (
	fixedResultValidator = types.Name{Package: libValidationPkg, Name: "FixedResult"}
)

func (v fixedResultDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	var result Validations

	if v.result {
		vals := gengo.ExtractCommentTags("+", comments)[validateTrueTagName]
		for _, val := range vals {
			tag, err := v.parseTagVal(val)
			if err != nil {
				return result, fmt.Errorf("can't extract +%s tag: %w", validateTrueTagName, err)
			}
			result.AddFunction(GenericFunction(validateTrueTagName, tag.flags, fixedResultValidator, tag.typeArgs, true, tag.msg))
		}
	} else {
		vals := gengo.ExtractCommentTags("+", comments)[validateFalseTagName]
		for _, val := range vals {
			tag, err := v.parseTagVal(val)
			if err != nil {
				return result, fmt.Errorf("can't extract +%s tag: %w", validateFalseTagName, err)
			}
			result.AddFunction(GenericFunction(validateFalseTagName, tag.flags, fixedResultValidator, tag.typeArgs, false, tag.msg))
		}
	}

	return result, nil
}

func (v fixedResultDeclarativeValidator) Docs() []TagDoc {
	if v.result {
		return []TagDoc{{
			Tag:         validateTrueTagName,
			Description: "Always passes validation (useful for testing).",
			Contexts:    []TagContext{TagContextType, TagContextField},
			Payloads: []TagPayloadDoc{{
				Description: "<none>",
				Docs:        "",
			}, {
				Description: "<quoted-string>",
				Docs:        "The generated code will include this string.",
			}, {
				Description: "<json-object>",
				Docs:        "",
				Schema: []TagPayloadSchema{{
					Key:   "flags",
					Value: "<list-of-flag-string>",
					Docs:  `values: ShortCircuit, NonError`,
				}, {
					Key:   "msg",
					Value: "<string>",
					Docs:  "The generated code will include this string.",
				}, {
					Key:   "typeArg",
					Value: "<string>",
					Docs:  "The type arg in generated code (must be the value-type, not pointer).",
				}},
			}},
		}}
	} else {
		return []TagDoc{{
			Tag:         validateFalseTagName,
			Description: "Always fails validation (useful for testing).",
			Contexts:    []TagContext{TagContextType, TagContextField},
			Payloads: []TagPayloadDoc{{
				Description: "<none>",
				Docs:        "",
			}, {
				Description: "<quoted-string>",
				Docs:        "The generated code will include this string.",
			}, {
				Description: "<json-object>",
				Docs:        "",
				Schema: []TagPayloadSchema{{
					Key:   "flags",
					Value: "<list-of-flag-string>",
					Docs:  `values: ShortCircuit, NonError`,
				}, {
					Key:   "msg",
					Value: "<string>",
					Docs:  "The generated code will include this string.",
				}},
			}},
		}}
	}
}

type tagVal struct {
	flags    FunctionFlags
	msg      string
	typeArgs []types.Name
}

func (_ fixedResultDeclarativeValidator) parseTagVal(in string) (tagVal, error) {
	type payload struct {
		Flags   []string `json:"flags"`
		Msg     string   `json:"msg"`
		TypeArg string   `json:"typeArg,omitempty"`
	}
	// We expect either a string (maybe empty) or a JSON object.
	if len(in) == 0 {
		return tagVal{}, nil
	}
	var pl payload
	if err := json.Unmarshal([]byte(in), &pl); err != nil {
		s := ""
		if err := json.Unmarshal([]byte(in), &s); err != nil {
			return tagVal{}, fmt.Errorf("error parsing JSON value: %v (%q)", err, in)
		}
		return tagVal{msg: s}, nil
	}
	// The msg field is required in JSON mode.
	if pl.Msg == "" {
		return tagVal{}, fmt.Errorf("JSON msg is required")
	}
	var flags FunctionFlags
	for _, fl := range pl.Flags {
		switch fl {
		case "ShortCircuit":
			flags |= ShortCircuit
		case "NonError":
			flags |= NonError
		default:
			return tagVal{}, fmt.Errorf("unknown flag: %q", fl)
		}
	}
	var typeArgs []types.Name
	if tn := pl.TypeArg; len(tn) > 0 {
		if !strings.HasPrefix(tn, "*") {
			tn = "*" + tn // We always need the pointer type.
		}
		typeArgs = []types.Name{{Package: "", Name: tn}}
	}

	return tagVal{flags, pl.Msg, typeArgs}, nil
}

func (v errorDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	var result Validations
	vals, found := gengo.ExtractCommentTags("+", comments)[validateErrorTagName]
	if found {
		return result, fmt.Errorf("forced error: %q", vals)
	}
	return result, nil
}

func (errorDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         validateErrorTagName,
		Description: "Always fails code generation (useful for testing).",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<string>",
			Docs:        "This string will be included in the error message.",
		}},
	}}
}
