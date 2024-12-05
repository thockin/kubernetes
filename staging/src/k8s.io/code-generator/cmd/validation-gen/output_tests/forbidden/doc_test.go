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

package forbidden

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
	pointer "k8s.io/utils/ptr"
)

func Test(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&T1{S: "x", PS: pointer.To("x"), PT2: &T2{S: "x"}}).
		ExpectInvalid(
			field.Forbidden(field.NewPath("s"), ""),
			field.Forbidden(field.NewPath("ps"), ""),
			// field.Forbidden(field.NewPath("t2"), ""),
			field.Forbidden(field.NewPath("pt2"), ""),
		)

	st.Value(&T1{S: "", PS: nil, PT2: nil}).
		ExpectValid()
}
