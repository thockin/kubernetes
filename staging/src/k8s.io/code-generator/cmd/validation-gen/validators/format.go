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
	"regexp"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/v2/types"
)

const (
	formatTagName = "k8s:format"
)

func init() {
	RegisterTagValidator(formatTagValidator{})
}

type formatTagValidator struct{}

func (formatTagValidator) Init(_ Config) {}

func (formatTagValidator) TagName() string {
	return formatTagName
}

var formatTagValidScopes = sets.New(ScopeAny)

func (formatTagValidator) ValidScopes() sets.Set[Scope] {
	return formatTagValidScopes
}

var (
	// Keep this list alphabetized.
	dnsLabelValidator     = types.Name{Package: libValidationPkg, Name: "DNSLabel"}
	dnsSubdomainValidator = types.Name{Package: libValidationPkg, Name: "DNSSubdomain"}
	ipSloppyValidator     = types.Name{Package: libValidationPkg, Name: "IPSloppy"}
	generateNameValidator = types.Name{Package: libValidationPkg, Name: "GenerateName"}
)

func (formatTagValidator) GetValidations(context Context, _ []string, payload string) (Validations, error) {
	var result Validations
	if formatFunction, err := getFormatValidationFunction(payload, context.Type); err != nil {
		return result, err
	} else {
		result.AddFunction(formatFunction)
	}
	return result, nil
}

func getFormatValidationFunction(format string, objType *types.Type) (FunctionGen, error) {
	// The naming convention for these formats follows the JSON schema style:
	// all lower-case, dashes between words. See
	// https://json-schema.org/draft/2020-12/json-schema-validation#name-defined-formats
	// for more examples.

	// Keep this list alphabetized.
	if format == "dns-label" {
		return Function(formatTagName, DefaultFlags, dnsLabelValidator), nil
	}
	if format == "dns-subdomain" {
		return Function(formatTagName, DefaultFlags, dnsSubdomainValidator), nil
	}
	// FIXME: Decide syntax: generate-name, generateName, k8s.io/generate-name, etc?
	// FIXME: just parse all formats?
	if format, arg := parseFormat(format); format == "generate-name" {
		wrappee, err := getFormatValidationFunction(arg, objType)
		if err != nil {
			return FunctionGen{}, err
		}
		wrapper := WrapperFunction{
			Function: wrappee,
			ObjType:  objType,
		}
		return Function(formatTagName, DefaultFlags, generateNameValidator, wrapper), nil
	}
	if format == "ip-sloppy" {
		return Function(formatTagName, DefaultFlags, ipSloppyValidator), nil
	}
	// TODO: Flesh out the list of validation functions

	return FunctionGen{}, fmt.Errorf("unsupported validation format %q", format)
}

// Matches for ^<anything>(<anything-or-nothing>)$
var functionalFormatRE = regexp.MustCompile(`^([^(]+)\(([^)]*)\)$`)

// parseFormat parses a format string of the form `format(arg)`.
// This is not very robust, but we don't need a lot from it (yet).
func parseFormat(format string) (string, string) {
	if matches := functionalFormatRE.FindStringSubmatch(format); len(matches) == 3 {
		return matches[1], matches[2]
	}
	return format, ""
}

func (ftv formatTagValidator) Docs() TagDoc {
	return TagDoc{
		Tag:         ftv.TagName(),
		Scopes:      ftv.ValidScopes().UnsortedList(),
		Description: "Indicates that a string field has a particular format.",
		Payloads: []TagPayloadDoc{{ // Keep this list alphabetized.
			Description: "dns-label",
			Docs:        "This field holds a DNS label value (approximately RFC 1123).",
		}, {
			Description: "dns-label",
			Docs:        "This field holds a DNS subdomain value (approximately RFC 1123).",
		}, {
			Description: "generate-name(<other-format>)",
			Docs:        "This field holds a value which is validated as <other-format>, and may be used as metadata.generateName.",
		}, {
			Description: "ip-sloppy",
			Docs:        "This field holds an IPv4 or IPv6 address value. IPv4 octets may have leading zeros.",
		}},
	}
}
