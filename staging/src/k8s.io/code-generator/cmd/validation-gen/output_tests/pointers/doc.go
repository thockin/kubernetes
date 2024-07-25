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
package pointers

type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.PS"
	PS *string `json:"ps"`
	// +validateTrue="field T1.PI"
	PI *int `json:"pi"`
	// +validateTrue="field T1.PB"
	PB *bool `json:"pb"`
	// +validateTrue="field T1.PF"
	PF *float64 `json:"pf"`

	// +validateTrue="field T1.PT2"
	PT2 *T2 `json:"pt2"`

	// Duplicate types with no validation.
	AnotherPS  *string  `json:"aotherps"`
	AnotherPI  *int     `json:"aotherpi"`
	AnotherPB  *bool    `json:"aotherpb"`
	AnotherPF  *float64 `json:"aotherpf"`
	AnotherPT2 *T2      `json:"aotherpt2"`
}

type T2 struct {
	// +validateTrue="field T2.PS"
	PS *string `json:"ps"`
	// +validateTrue="field T2.PI"
	PI *int `json:"pi"`
	// +validateTrue="field T2.PB"
	PB *bool `json:"pb"`
	// +validateTrue="field T2.PF"
	PF *float64 `json:"pf"`
}
