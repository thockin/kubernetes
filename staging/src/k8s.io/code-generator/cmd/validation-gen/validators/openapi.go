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
	"k8s.io/gengo/v2/types"
	"k8s.io/kube-openapi/pkg/generators"
)

func init() {
	AddToRegistry(InitOpenAPIDeclarativeValidator)
}

func InitOpenAPIDeclarativeValidator(_ *ValidatorConfig) DeclarativeValidator {
	return &openAPIDeclarativeValidator{}
}

type openAPIDeclarativeValidator struct{}

const (
	//FIXME: If we change this to "k8s:" then kube-openapi's parsing tries to
	//       parse all tags, not just ones it knows. The previous tag is used
	//       in some types.go fields, so just give it a bogus name for now.
	markerPrefix     = "k8s:validation:new:"
	formatTagName    = markerPrefix + ":format"
	maxLengthTagName = markerPrefix + ":maxLength"
	maxItemsTagName  = markerPrefix + ":maxItems"
)

var (
	ipValidator        = types.Name{Package: libValidationPkg, Name: "IP"}
	dnsLabelValidator  = types.Name{Package: libValidationPkg, Name: "DNSLabel"}
	maxLengthValidator = types.Name{Package: libValidationPkg, Name: "MaxLength"}
	maxItemsValidator  = types.Name{Package: libValidationPkg, Name: "MaxItems"}
)

func (openAPIDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	var result Validations
	// Leverage the kube-openapi parser for 'k8s:validation:' validations.
	schema, err := generators.ParseCommentTags(t, comments, "+"+markerPrefix)
	if err != nil {
		return result, err
	}
	if schema.MaxLength != nil {
		result.AddFunction(Function(maxLengthTagName, DefaultFlags, maxLengthValidator, *schema.MaxLength))
	}
	if schema.MaxItems != nil {
		result.AddFunction(Function(maxItemsTagName, ShortCircuit, maxItemsValidator, *schema.MaxItems))
	}
	if len(schema.Format) > 0 {
		formatFunction := FormatValidationFunction(schema.Format)
		if formatFunction != nil {
			result.AddFunction(formatFunction)
		}
	}

	return result, nil
}

func (openAPIDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         formatTagName,
		Description: "Indicates that a string field has a particular format.",
		Contexts:    []TagContext{TagContextType, TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "ip",
			Docs:        "This field holds an IP address value, either IPv4 or IPv6.",
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
	}}
}

func FormatValidationFunction(format string) FunctionGen {
	// The naming convention for these formats follows the JSON schema style:
	// all lower-case, dashes between words. See
	// https://json-schema.org/draft/2020-12/json-schema-validation#name-defined-formats
	// for more examples.
	if format == "ip" {
		return Function(formatTagName, DefaultFlags, ipValidator)
	}
	if format == "dns-label" {
		return Function(formatTagName, DefaultFlags, dnsLabelValidator)
	}
	// TODO: Flesh out the list of validation functions

	return nil // TODO: ignore unsupported formats?
}
