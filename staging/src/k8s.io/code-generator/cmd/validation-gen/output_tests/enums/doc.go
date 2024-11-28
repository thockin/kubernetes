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
package enums

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

type T1 struct {
	TypeMeta int

	E0  E0  `json:"e0"`
	PE0 *E0 `json:"pe0"`

	E1  E1  `json:"e1"`
	PE1 *E1 `json:"pe1"`

	E2  E2  `json:"e2"`
	PE2 *E2 `json:"pe2"`
}

// +k8s:enum
type E0 string // Note: this enum has no values

// +k8s:enum
type E1 string // Note: this enum has 1 value

const (
	E1V1 E1 = "e1v1"
)

// +k8s:enum
type E2 string // Note: this enum has 2 values

const (
	E2V1 E2 = "e2v1"
	E2V2 E2 = "e2v2"
)
