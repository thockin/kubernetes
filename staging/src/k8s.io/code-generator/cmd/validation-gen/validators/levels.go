/*
Copyright 2025 The Kubernetes Authors.

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
	"regexp"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/v2/codetags"
)

var kubeVersionRegex = regexp.MustCompile(`^1\.\d+$`)

const (
	alphaTagName = "k8s:alpha"
	betaTagName  = "k8s:beta"
)

func init() {
	RegisterTagValidator(&levelTagValidator{tagName: alphaTagName, level: ValidationStabilityLevelAlpha})
	RegisterTagValidator(&levelTagValidator{tagName: betaTagName, level: ValidationStabilityLevelBeta})
}

type levelTagValidator struct {
	validator TagValidationExtractor
	tagName   string
	level     ValidationStabilityLevel
}

func (stv *levelTagValidator) Init(cfg Config) {
	stv.validator = cfg.TagValidator
}

func (stv *levelTagValidator) TagName() string {
	return stv.tagName
}

var stageTagValidScopes = sets.New(ScopeType, ScopeField, ScopeListVal, ScopeMapKey, ScopeMapVal)

func (levelTagValidator) ValidScopes() sets.Set[Scope] {
	return stageTagValidScopes
}

// LateTagValidator indicates that this validator has to run AFTER the listType
// and listMapKey tags.
func (levelTagValidator) LateTagValidator() {}

func (stv *levelTagValidator) GetValidations(context Context, tag codetags.Tag) (Validations, error) {
	if tag.ValueType != codetags.ValueTypeTag || tag.ValueTag == nil {
		return Validations{}, fmt.Errorf("requires a validation tag as its value payload")
	}

	if len(tag.Args) > 1 {
		return Validations{}, fmt.Errorf("at most one optional kubernetes version argument is supported")
	}

	var version string
	if len(tag.Args) == 1 {
		arg := tag.Args[0]
		version = arg.Value
		if !kubeVersionRegex.MatchString(version) {
			return Validations{}, fmt.Errorf("invalid kubernetes version format, expected 1.<minor version>, got %s", version)
		}
	}

	context.StabilityLevel = stv.level
	validations, err := stv.validator.ExtractTagValidations(context, *tag.ValueTag)
	if err != nil {
		return Validations{}, err
	}

	result := Validations{}
	result.Variables = append(result.Variables, validations.Variables...)
	for _, fn := range validations.Functions {
		f := fn
		f.StabilityLevel = stv.level
		result.AddFunction(f)
	}

	return result, nil
}

func (stv *levelTagValidator) Docs() TagDoc {
	doc := TagDoc{
		Tag:            stv.TagName(),
		StabilityLevel: TagStabilityLevelAlpha,
		Scopes:         stv.ValidScopes().UnsortedList(),
		Description:    fmt.Sprintf("Marks the given payload validation as a %s validation of the handwritten validation code. An optional Kubernetes version can be specified.", stv.level),
		Args: []TagArgDoc{{
			Description: "The Kubernetes version (e.g. `1.34`) at which this validation was added.",
			Type:        codetags.ArgTypeString,
			Name:        "since",
		}},
		Payloads: []TagPayloadDoc{{
			Description: "<validation-tag>",
			Docs:        fmt.Sprintf("The validation tag to evaluate as a %s validation.", stv.level),
		}},
		PayloadsType:     codetags.ValueTypeTag,
		PayloadsRequired: true,
	}
	return doc
}
