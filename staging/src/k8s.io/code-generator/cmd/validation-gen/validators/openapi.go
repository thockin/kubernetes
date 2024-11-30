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
	"strconv"

	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/types"
)

func init() {
	AddToRegistry(InitOpenAPIDeclarativeValidator)
}

func InitOpenAPIDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &openAPIDeclarativeValidator{}
}

type openAPIDeclarativeValidator struct{}

const (
	formatTagName    = "k8s:format"
	maxLengthTagName = "k8s:maxLength"
	maxItemsTagName  = "k8s:maxItems"
	minimumTagName   = "k8s:minimum"
)

var (
	ipValidator        = types.Name{Package: libValidationPkg, Name: "IPSloppy"}
	dnsLabelValidator  = types.Name{Package: libValidationPkg, Name: "DNSLabel"}
	maxLengthValidator = types.Name{Package: libValidationPkg, Name: "MaxLength"}
	maxItemsValidator  = types.Name{Package: libValidationPkg, Name: "MaxItems"}
	minimumValidator   = types.Name{Package: libValidationPkg, Name: "Minimum"}
)

func (openAPIDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	var result Validations
	commentTags := gengo.ExtractCommentTags("+", comments)

	if maxLength, found, err := extractOptionalIntValue(commentTags, maxLengthTagName); err != nil {
		return result, err
	} else if found {
		result.AddFunction(Function(maxLengthTagName, DefaultFlags, maxLengthValidator, maxLength))
	}

	if maxItems, found, err := extractOptionalIntValue(commentTags, maxItemsTagName); err != nil {
		return result, err
	} else if found {
		// Note: maxItems short-circuits other validations.
		result.AddFunction(Function(maxItemsTagName, ShortCircuit, maxItemsValidator, maxItems))
	}

	if minimum, found, err := extractOptionalIntValue(commentTags, minimumTagName); err != nil {
		return result, err
	} else if found {
		result.AddFunction(Function(minimumTagName, DefaultFlags, minimumValidator, minimum))
	}

	if formats := commentTags[formatTagName]; len(formats) > 0 {
		if len(formats) > 1 {
			return result, fmt.Errorf("multiple values found for tag %q", formatTagName)
		}
		format := formats[0]
		if formatFunction, err := getFormatValidationFunction(format); err != nil {
			return result, err
		} else if formatFunction == nil {
			return result, fmt.Errorf("internal error: no validation function found for format %q", format)
		} else {
			result.AddFunction(formatFunction)
		}
	}

	return result, nil
}

func extractOptionalIntValue(commentTags map[string][]string, tagName string) (int, bool, error) {
	values, found := commentTags[tagName]
	if !found || len(values) == 0 {
		return 0, false, nil
	}
	if len(values) > 1 {
		return 0, false, fmt.Errorf("multiple values found for tag %q", tagName)
	}
	intVal, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, false, fmt.Errorf("failed to parse value for tag %q: %v", tagName, err)
	}
	return intVal, true, nil
}

func (openAPIDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         formatTagName,
		Description: "Indicates that a string field has a particular format.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "ip-sloppy",
			Docs:        "This field holds an IPv4 or IPv6 address value. IPv4 octets may have leading zeros.",
		}, {
			Description: "dns-label",
			Docs:        "This field holds a DNS label value.",
		}},
	}, {
		Tag:         maxLengthTagName,
		Description: "Indicates that a string field has a limit on its length.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<non-negative integer>",
			Docs:        "This field must be no more than X characters long.",
		}},
	}, {
		Tag:         maxItemsTagName,
		Description: "Indidates that a slice field has a limit on its size.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{
			{
				Description: "<non-negative integer>",
				Docs:        "This field must be no more than X items long.",
			},
		},
	}, {
		Tag:         minimumTagName,
		Description: "Indicates that a numeric field has a minimum value.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<integer>",
			Docs:        "This field must be greater than or equal to x.",
		}},
	}}
}

func getFormatValidationFunction(format string) (FunctionGen, error) {
	// The naming convention for these formats follows the JSON schema style:
	// all lower-case, dashes between words. See
	// https://json-schema.org/draft/2020-12/json-schema-validation#name-defined-formats
	// for more examples.
	if format == "ip-sloppy" {
		return Function(formatTagName, DefaultFlags, ipValidator), nil
	}
	if format == "dns-label" {
		return Function(formatTagName, DefaultFlags, dnsLabelValidator), nil
	}
	// TODO: Flesh out the list of validation functions

	return nil, fmt.Errorf("unsupported validation format %q", format)
}
