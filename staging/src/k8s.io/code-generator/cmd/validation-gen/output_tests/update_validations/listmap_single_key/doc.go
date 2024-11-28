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

// +k8s:validation-gen=*
// +k8s:validation-gen-scheme-registry=k8s.io/code-generator/cmd/validation-gen/testscheme.Scheme
// +k8s:validation-gen-test-fixture=validateFalse

// This is a test package.
package listmap_single_key

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

type T1 struct {
	// +k8s:listType=map
	// +k8s:listMapKey=k
	LM1 []M1 `json:"lm1"`

	// +k8s:listType=map
	// +k8s:listMapKey=k
	LM2 []M2 `json:"lm2"`

	// +k8s:listType=map
	// +k8s:listMapKey=k
	LM3 []M3 `json:"lm3"`

	// +k8s:listType=map
	// +k8s:listMapKey=k
	LM4 []M4 `json:"lm4"`
}

type M1 struct {
	// +k8s:validateFalse="M1.K"
	K string `json:"k"`

	// +k8s:validateFalse="M1.S"
	S string `json:"s"`
}

type M2 struct {
	M1 // embedded, no JSON tag
}

type M3 struct {
	M1 `json:",inline"` // embedded
}

type M4 struct {
	M2
}
