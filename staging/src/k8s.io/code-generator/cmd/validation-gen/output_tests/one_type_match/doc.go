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
package onetypematch

type T1 struct {
	TypeMeta int
	// +validateTrue="field T1.S"
	S  string `json:"s"`
	T2 T2     `json:"t2"`
	T3 T3     `json:"t3"`
	T4 T4     `json:"t4"`
	T5 T5     `json:"t5"`
	T6 T6     `json:"t6"`
}

type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
}

// This type should not be generated. It has no validations.
type T3 struct {
	S string `json:"s"`
}

// This type should be generated, it has a child type that has validations.
type T4 struct {
	t2 T2 `json:"t2"`
}

// This type should be generated, it has a map value type that has validations.
type T5 struct {
	t2 map[string]T2 `json:"t2"`
}

// This type should be generated, it has a list items that has validations.
type T6 struct {
	t2 []T2 `json:"t2"`
}

type private struct {
	S string
}
