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
package primitives

type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.S"
	S string `json:"s"`
	// +validateTrue="field T1.I"
	I int `json:"i"`
	// +validateTrue="field T1.B"
	B bool `json:"b"`
	// +validateTrue="field T1.F"
	F float64 `json:"f"`

	// +validateTrue="field T1.T2"
	T2 T2 `json:"t2"`

	// No internal validations.
	T3 T3 `json:"t3"`

	// Duplicate types with no validation.
	AnotherS  string  `json:"anothers"`
	AnotherI  int     `json:"anotheri"`
	AnotherB  bool    `json:"anotherb"`
	AnotherF  float64 `json:"anotherf"`
	AnotherT2 T2      `json:"anothert2"`
}

// Note: This has validations and is linked into the type-graph of T1.
type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
	// +validateTrue="field T2.I"
	I int `json:"i"`
	// +validateTrue="field T2.B"
	B bool `json:"b"`
	// +validateTrue="field T2.F"
	F float64 `json:"f"`
}

// Note: This has no validations and is linked into the type-graph of T1.
type T3 struct {
	S string  `json:"s"`
	I int     `json:"i"`
	B bool    `json:"b"`
	F float64 `json:"f"`
}

// Note: This has validations and is not linked into the type-graph of T1.
type T4 struct {
	// +validateTrue="field T4.S"
	S string `json:"s"`
	// +validateTrue="field T4.I"
	I int `json:"i"`
	// +validateTrue="field T4.B"
	B bool `json:"b"`
	// +validateTrue="field T4.F"
	F float64 `json:"f"`
}

// Note: This has no validations and is not linked into the type-graph of T1.
type T5 struct {
	S string  `json:"s"`
	I int     `json:"i"`
	B bool    `json:"b"`
	F float64 `json:"f"`
}
