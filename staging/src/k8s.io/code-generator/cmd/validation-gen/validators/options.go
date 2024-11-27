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
	"strings"

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
	enabledTagValues, hasEnabledTags := gengo.ExtractCommentTags("+", comments)[ifOptionEnabledTag]
	disabledTagValues, hasDisabledTags := gengo.ExtractCommentTags("+", comments)[ifOptionDisabledTag]
	if !hasEnabledTags && !hasDisabledTags {
		return Validations{}, nil
	}
	var functions []FunctionGen
	var variables []VariableGen
	for _, v := range enabledTagValues {
		optionName, validations, err := o.parseIfOptionsTag(t, v)
		if err != nil {
			return Validations{}, err
		}
		for _, fn := range validations.Functions {
			functions = append(functions, WithCondition(fn, Conditions{OptionEnabled: optionName}))
		}
		variables = append(variables, validations.Variables...)
	}
	for _, v := range disabledTagValues {
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

func (o optionDeclarativeValidator) parseIfOptionsTag(t *types.Type, tagValue string) (string, Validations, error) {
	parts := strings.SplitN(tagValue, "=", 2)
	if len(parts) != 2 {
		return "", Validations{}, fmt.Errorf("invalid value %q for option %q", tagValue, parts[0])
	}
	optionName := parts[0]
	embeddedValidation := parts[1]
	validations, err := o.cfg.EmbedValidator.ExtractValidations(t, []string{embeddedValidation})
	if err != nil {
		return "", Validations{}, err
	}
	return optionName, validations, nil
}

func (optionDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         ifOptionEnabledTag,
		Description: "Declares a validation that only applies when an option is enabled.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<option-name>=<validation-tag>",
			Docs:        "This validation tag will be evaluated only if the validation option is enabled.",
		}},
	}}
}
