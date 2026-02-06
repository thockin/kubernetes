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

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/v2/codetags"
	"k8s.io/gengo/v2/types"
)

const (
	ifEnabledTag  = "k8s:ifEnabled"
	ifDisabledTag = "k8s:ifDisabled"
)

func init() {
	RegisterTagValidator(&ifTagValidator{true, nil})
	RegisterTagValidator(&ifTagValidator{false, nil})
}

type ifTagValidator struct {
	enabled   bool
	validator TagValidationExtractor
}

func (itv *ifTagValidator) Init(cfg Config) {
	itv.validator = cfg.TagValidator
}

func (itv ifTagValidator) TagName() string {
	if itv.enabled {
		return ifEnabledTag
	}
	return ifDisabledTag
}

var ifEnabledDisabledTagValidScopes = sets.New(ScopeType, ScopeField, ScopeListVal, ScopeMapKey, ScopeMapVal, ScopeConst)

func (ifTagValidator) ValidScopes() sets.Set[Scope] {
	return ifEnabledDisabledTagValidScopes
}

var (
	ifOption = types.Name{Package: libValidationPkg, Name: "IfOption"}
)

func (itv ifTagValidator) GetValidations(context Context, tag codetags.Tag) (Validations, error) {
	optionArg, ok := tag.PositionalArg()
	if !ok {
		return Validations{}, fmt.Errorf("missing required option name positional argument")
	}
	// process is a recursive helper function that handles both immediate and
	// deferred validations from a child tag. It ensures that 'Variables' are not
	// used and properly wraps functions and defers new callbacks for deferred
	// items.
	var process func(Validations) (Validations, error) // because it's recursive
	process = func(in Validations) (Validations, error) {
		result := Validations{}
		result.Variables = append(result.Variables, in.Variables...)
		for _, fn := range in.Functions {
			f := Function(itv.TagName(), fn.Flags, ifOption, optionArg.Value, itv.enabled, WrapperFunction{Function: fn, ObjType: context.Type})
			result.AddFunction(f)
		}
		for _, d := range in.Deferred {
			result.AddDeferred(Deferred(d.Scope, func() (Validations, error) {
				inner, err := d.Callback()
				if err != nil {
					return Validations{}, err
				}
				return process(inner)
			}))
		}
		return result, nil
	}

	validations, err := itv.validator.ExtractTagValidations(context, *tag.ValueTag)
	if err != nil {
		return Validations{}, err
	}
	return process(validations)
}

func (itv ifTagValidator) Docs() TagDoc {
	doc := TagDoc{
		Tag:            itv.TagName(),
		StabilityLevel: TagStabilityLevelAlpha,
		Args: []TagArgDoc{{
			Description: "<option>",
			Type:        codetags.ArgTypeString,
			Required:    true,
		}},
		Scopes: itv.ValidScopes().UnsortedList(),
	}

	doc.PayloadsType = codetags.ValueTypeTag
	doc.PayloadsRequired = true
	if itv.enabled {
		doc.Description = "Declares a validation that only applies when an option is enabled."
		doc.Payloads = []TagPayloadDoc{{
			Description: "<validation-tag>",
			Docs:        "This validation tag will be evaluated only if the validation option is enabled.",
		}}
	} else {
		doc.Description = "Declares a validation that only applies when an option is disabled."
		doc.Payloads = []TagPayloadDoc{{
			Description: "<validation-tag>",
			Docs:        "This validation tag will be evaluated only if the validation option is disabled.",
		}}
	}
	return doc
}
