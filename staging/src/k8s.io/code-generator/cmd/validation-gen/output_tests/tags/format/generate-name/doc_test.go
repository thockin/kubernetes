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

package format

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
)

func Test(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&Struct{
		GenerateNameField:        "foo-bar",
		GenerateNamePtrField:     ptr.To("foo-bar"),
		GenerateNameTypedefField: "foo-bar",
	}).ExpectValid()

	st.Value(&Struct{
		GenerateNameField:        "foo-bar-",
		GenerateNamePtrField:     ptr.To("foo-bar-"),
		GenerateNameTypedefField: "foo-bar-",
	}).ExpectValid()

	st.Value(&Struct{
		GenerateNameField:        "1234",
		GenerateNamePtrField:     ptr.To("1234"),
		GenerateNameTypedefField: "1234",
	}).ExpectValid()

	st.Value(&Struct{
		GenerateNameField:        "-1234-",
		GenerateNamePtrField:     ptr.To("1234-"),
		GenerateNameTypedefField: "1234-",
	}).ExpectValid()

	st.Value(&Struct{
		GenerateNameField:        "",
		GenerateNamePtrField:     ptr.To(""),
		GenerateNameTypedefField: "",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("generateNameField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("generateNamePtrField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("generateNameTypedefField"), nil, "").WithOrigin("format=dns-label"),
	})

	st.Value(&Struct{
		GenerateNameField:        "Not a DNS label",
		GenerateNamePtrField:     ptr.To("Not a DNS label"),
		GenerateNameTypedefField: "Not a DNS label",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("generateNameField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("generateNamePtrField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("generateNameTypedefField"), nil, "").WithOrigin("format=dns-label"),
	})
}
