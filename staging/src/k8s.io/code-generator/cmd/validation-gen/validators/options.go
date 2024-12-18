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
	AddToRegistry(InitOptionDeclarativeValidator)
}

func InitOptionDeclarativeValidator(cfg *ValidatorConfig) DeclarativeValidator {
	return &optionDeclarativeValidator{cfg: cfg}
}

type optionDeclarativeValidator struct {
	cfg *ValidatorConfig
}

const (
	ifOptionEnabledTag  = "k8s:ifOptionEnabled"
	ifOptionDisabledTag = "k8s:ifOptionDisabled"
)

func (o optionDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	tags, err := gengo.ExtractFunctionStyleCommentTags("+", []string{ifOptionEnabledTag, ifOptionDisabledTag}, comments)
	if err != nil {
		return Validations{}, err
	}

	enabledTags, hasEnabledTags := tags[ifOptionEnabledTag]
	disabledTags, hasDisabledTags := tags[ifOptionDisabledTag]
	if !hasEnabledTags && !hasDisabledTags {
		return Validations{}, nil
	}

	var functions []FunctionGen
	var variables []VariableGen
	for _, v := range enabledTags {
		optionName, validations, err := o.parseIfOptionsTag(t, v)
		if err != nil {
			return Validations{}, err
		}
		for _, fn := range validations.Functions {
			functions = append(functions, WithCondition(fn, Conditions{OptionEnabled: optionName}))
		}
		variables = append(variables, validations.Variables...)
	}
	for _, v := range disabledTags {
		optionName, validations, err := o.parseIfOptionsTag(t, v)
		if err != nil {
			return Validations{}, err
		}
		for _, fn := range validations.Functions {
			functions = append(functions, WithCondition(fn, Conditions{OptionDisabled: optionName}))
		}
		variables = append(variables, validations.Variables...)
	}
	return Validations{
		Functions: functions,
		Variables: variables,
	}, nil
}

func (o optionDeclarativeValidator) parseIfOptionsTag(t *types.Type, tag gengo.Tag) (string, Validations, error) {
	if len(tag.Args) != 1 {
		return "", Validations{}, fmt.Errorf("tag %q requires 1 argument", tag.Name)
	}
	validations, err := o.cfg.EmbedValidator.ExtractValidations(t, []string{tag.Value})
	if err != nil {
		return "", Validations{}, err
	}
	return tag.Args[0], validations, nil
}

func (optionDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         fmt.Sprintf("%s(<option-name>)", ifOptionEnabledTag),
		Description: "Declares a validation that only applies when an option is enabled.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<validation-tag>",
			Docs:        "This validation tag will be evaluated only if the validation option is enabled.",
		}},
	}, {
		Tag:         fmt.Sprintf("%s(<option-name>)", ifOptionDisabledTag),
		Description: "Declares a validation that only applies when an option is disabled.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<validation-tag>",
			Docs:        "This validation tag will be evaluated only if the validation option is disabled.",
		}},
	}}
}
