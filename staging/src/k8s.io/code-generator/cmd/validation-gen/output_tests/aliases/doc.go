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

// This is a test package.
package aliases

// +validateTrue="type E1"
type E1 string

// +validateTrue="type E2"
type E2 int

// +validateTrue="type E3"
type E3 E1

// +validateTrue="type E4"
type E4 T2

// +validateTrue="type T1"
type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.E1"
	E1 E1 `json:"e1"`
	// +validateTrue="field T1.PE1"
	PE1 *E1 `json:"pe1"`

	// +validateTrue="field T1.E2"
	E2 E2 `json:"e2"`
	// +validateTrue="field T1.PE2"
	PE2 *E2 `json:"pe2"`

	// +validateTrue="field T1.E3"
	E3 E3 `json:"e3"`
	// +validateTrue="field T1.PE3"
	PE3 *E3 `json:"pe3"`

	// +validateTrue="field T1.E4"
	E4 E4 `json:"e4"`
	// +validateTrue="field T1.PE4"
	PE4 *E4 `json:"pe4"`

	// +validateTrue="field T1.T2"
	T2 T2 `json:"t2"`
	// +validateTrue="field T1.PT2"
	PT2 *T2 `json:"pt2"`
}

// +validateTrue="type T2"
type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
}
