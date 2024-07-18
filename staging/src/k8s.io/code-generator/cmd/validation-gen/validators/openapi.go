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
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/types"
	"k8s.io/kube-openapi/pkg/generators"
)

func init() {
	AddToRegistry(InitOpenAPIDeclarativeValidator)
}

func InitOpenAPIDeclarativeValidator(c *generator.Context) DeclarativeValidator {
	return &openAPIDeclarativeValidator{}
}

type openAPIDeclarativeValidator struct{}

const (
	markerPrefix     = "+k8s:validation:"
	formatTagName    = markerPrefix + ":format"
	maxLengthTagName = markerPrefix + ":maxLength"
)

var (
	ipValidator        = types.Name{Package: libValidationPkg, Name: "IP"}
	dnsLabelValidator  = types.Name{Package: libValidationPkg, Name: "DNSLabel"}
	maxLengthValidator = types.Name{Package: libValidationPkg, Name: "MaxLength"}
)

func (openAPIDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) ([]FunctionGen, error) {
	var v []FunctionGen
	// Leverage the kube-openapi parser for 'k8s:validation:' validations.
	schema, err := generators.ParseCommentTags(t, comments, markerPrefix)
	if err != nil {
		return nil, err
	}
	if schema.MaxLength != nil {
		v = append(v, Function(maxLengthTagName, maxLengthValidator, *schema.MaxLength))
	}
	if len(schema.Format) > 0 {
		formatFunction := FormatValidationFunction(schema.Format)
		if formatFunction != nil {
			v = append(v, formatFunction)
		}
	}

	return v, nil
}

func FormatValidationFunction(format string) FunctionGen {
	// The naming convention for these formats follows the JSON schema style:
	// all lower-case, dashes between words. See
	// https://json-schema.org/draft/2020-12/json-schema-validation#name-defined-formats
	// for more examples.
	if format == "ip" {
		return Function(formatTagName, ipValidator)
	}
	if format == "dns-label" {
		return Function(formatTagName, dnsLabelValidator)
	}
	// TODO: Flesh out the list of validation functions

	return nil // TODO: ignore unsupported formats?
}
