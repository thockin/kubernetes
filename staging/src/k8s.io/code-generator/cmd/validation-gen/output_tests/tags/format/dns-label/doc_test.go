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
		DNSLabelField:        "foo-bar",
		DNSLabelPtrField:     ptr.To("foo-bar"),
		DNSLabelTypedefField: "foo-bar",
	}).ExpectValid()

	st.Value(&Struct{
		DNSLabelField:        "1234",
		DNSLabelPtrField:     ptr.To("1234"),
		DNSLabelTypedefField: "1234",
	}).ExpectValid()

	st.Value(&Struct{
		DNSLabelField:        "",
		DNSLabelPtrField:     ptr.To(""),
		DNSLabelTypedefField: "",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("dnsLabelField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("dnsLabelPtrField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("dnsLabelTypedefField"), nil, "").WithOrigin("format=dns-label"),
	})

	st.Value(&Struct{
		DNSLabelField:        "Not a DNS label",
		DNSLabelPtrField:     ptr.To("Not a DNS label"),
		DNSLabelTypedefField: "Not a DNS label",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("dnsLabelField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("dnsLabelPtrField"), nil, "").WithOrigin("format=dns-label"),
		field.Invalid(field.NewPath("dnsLabelTypedefField"), nil, "").WithOrigin("format=dns-label"),
	})
}
