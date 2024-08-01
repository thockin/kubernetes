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
package withfieldvalidations

type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.S"
	S string `json:"s"`
	// +validateTrue="field T1.T2"
	T2 T2 `json:"t2"`
	// +validateTrue="field T1.T3"
	T3 T3 `json:"t3"`

	// +validateTrue="field T1.E1"
	E1 E1 `json:"e1"`
	// +validateTrue="field T1.E2"
	E2 E2 `json:"e2"`
}

// Note: this has validations and is linked into T1.
type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
}

// Note: this has no validations and is linked into T1.
type T3 struct {
	S string `json:"s"`
}

// Note: this has validations and is not linked into T1.
type T4 struct {
	// +validateTrue="field T4.S"
	S string `json:"s"`
}

// Note: this has no validations and is not linked into T1.
type T5 struct {
	S string `json:"s"`
}

// Note: this has validations and is linked into T1.
// +validateTrue="type E1"
type E1 string

// Note: this has no validations and is linked into T1.
type E2 string

// Note: this has validations and is not linked into T1.
// +validateTrue="field type E3"
type E3 string

// Note: this has no validations and is not linked into T1.
type E4 string
