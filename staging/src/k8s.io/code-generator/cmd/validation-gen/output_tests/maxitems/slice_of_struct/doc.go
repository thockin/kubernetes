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

// +k8s:validation-gen=TypeMeta
// +k8s:validation-gen-scheme-registry=k8s.io/code-generator/cmd/validation-gen/testscheme.Scheme

// This is a test package.
package sliceofstruct

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

// Maxitems
type M struct {
	TypeMeta int

	// slice-of-struct
	// +k8s:validation:new:maxItems=1
	M0 []S `json:"m0"`

	// typedef-of-slice-of-struct
	// Validation on type only.
	M1 SSliceLimited `json:"m1"`
	// Validation on field only.
	// +k8s:validation:new:maxItems=1
	M2 SSlice `json:"m2"`
	// Validation on both type and field.
	// +k8s:validation:new:maxItems=1
	M3 SSliceLimited `json:"m3"`

	// slice-of-pointer-to-struct
	// +k8s:validation:new:maxItems=1
	M4 []*S `json:"m4"`

	// typedef-of-slice-of-pointer-to-struct
	// Validation on type only.
	M5 SPtrSliceLimited `json:"m5"`
	// Validation on field only.
	// +k8s:validation:new:maxItems=1
	M6 SPtrSlice `json:"m6"`
	// Validation on both type and field.
	// +k8s:validation:new:maxItems=1
	M7 SPtrSliceLimited `json:"m7"`
}

type S struct{}

// Note: no limit here
type SSlice []S

// +k8s:validation:new:maxItems=2
type SSliceLimited []S

// +k8s:validation:new:maxItems=2
type SPtrSliceLimited []*S

// Note: no limit here
type SPtrSlice []*S
