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
package slices

// +validateTrue="type T1"
type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.LS"
	// +eachVal=+validateTrue="field T1.LS[*]"
	// +eachVal=+required
	LS []string `json:"ls"`
	// +validateTrue="field T1.LPS"
	// +eachVal=+validateTrue="field T1.LPS[*]"
	// +eachVal=+required
	LPS []*string `json:"lps"`

	// +validateTrue="field T1.LT2"
	// +eachVal=+validateTrue="field T1.LT2[*]"
	// +eachVal=+required
	LT2 []T2 `json:"lt2"`
	// +validateTrue="field T1.LPT2"
	// +eachVal=+validateTrue="field T1.LPT2[*]"
	// +eachVal=+required
	LPT2 []*T2 `json:"lpt2"`

	// +validateTrue="field Z"
	Z []string `json:"z"` // Should not have any eachKey, eachVal tags applied
}

// +validateTrue="type T2"
type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
}
