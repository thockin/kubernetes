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
package minimum

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/validate/content"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
)

func Test(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&T1{
		PI:  ptr.To(0),
		PU:  ptr.To(uint(0)),
		PAI: ptr.To(AI(0)),
	}).ExpectInvalid(
		field.Invalid(field.NewPath("i"), 0, content.MinError(1)),
		field.Invalid(field.NewPath("pi"), 0, content.MinError(1)),
		field.Invalid(field.NewPath("i16"), 0, content.MinError(1)),
		field.Invalid(field.NewPath("i32"), 0, content.MinError(1)),
		field.Invalid(field.NewPath("i64"), 0, content.MinError(1)),
		field.Invalid(field.NewPath("u"), uint(0), content.MinError(1)),
		field.Invalid(field.NewPath("pu"), uint(0), content.MinError(1)),
		field.Invalid(field.NewPath("u16"), uint(0), content.MinError(1)),
		field.Invalid(field.NewPath("u32"), uint(0), content.MinError(1)),
		field.Invalid(field.NewPath("u64"), uint(0), content.MinError(1)),
		field.Invalid(field.NewPath("ai"), 0, content.MinError(1)),
		field.Invalid(field.NewPath("pai"), 0, content.MinError(1)),
	)

	st.Value(&T1{
		I: 1, PI: ptr.To(1), I16: 1, I32: 1, I64: 1,
		U: 1, PU: ptr.To(uint(1)), U16: 1, U32: 1, U64: 1,
		AI: 1, PAI: ptr.To(AI(1)),
	}).ExpectValid()
}
