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
package typedef

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

// Maxitems
type M struct {
	TypeMeta int

	// typedef-of-slice-of-primitive
	// Validation on type only.
	M0 IntSliceLimited `json:"m0"`
	// Validation on field only.
	// +k8s:validation:new:maxItems=1
	M1 IntSlice `json:"m1"`
	// Validation on both type and field.
	// +k8s:validation:new:maxItems=1
	M2 IntSliceLimited `json:"m2"`

	// typedef-of-slice-of-pointer-to-primitive
	// Validation on type only.
	M3 IntPtrSliceLimited `json:"m3"`
	// Validation on field only.
	// +k8s:validation:new:maxItems=1
	M4 IntPtrSlice `json:"m4"`
	// Validation on both type and field.
	// +k8s:validation:new:maxItems=1
	M5 IntPtrSliceLimited `json:"m5"`
}

// Note: no limit here
type IntSlice []int

// +k8s:validation:new:maxItems=2
type IntSliceLimited []int

// Note: no limit here
type IntPtrSlice []*int

// +k8s:validation:new:maxItems=2
type IntPtrSliceLimited []*int
