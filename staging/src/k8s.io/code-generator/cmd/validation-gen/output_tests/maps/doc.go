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
package maps

// +validateTrue="type E1"
type E1 string

// +validateTrue="type T1"
type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.MSS"
	// +eachKey=+validateTrue="key T1.MSS[*]"
	// +eachVal=+validateTrue="val T1.MSS[*]"
	MSS map[string]string `json:"mss"`
	// +validateTrue="field T1.MSPS"
	// +eachKey=+validateTrue="key T1.MSPS[*]"
	// +eachVal=+validateTrue="val T1.MSPS[*]"
	MSPS map[string]*string `json:"msps"`
	// +validateTrue="field T1.MPSS"
	// +eachKey=+validateTrue="key T1.MPSS[*]"
	// +eachVal=+validateTrue="val T1.MPSS[*]"
	MPSS map[*string]string `json:"mpss"`
	// +validateTrue="field T1.MPSPS"
	// +eachKey=+validateTrue="key T1.MPSPS[*]"
	// +eachVal=+validateTrue="val T1.MPSPS[*]"
	MPSPS map[*string]*string `json:"mpsps"`

	// +validateTrue="field T1.MST2"
	// +eachKey=+validateTrue="key T1.MST2[*]"
	// +eachVal=+validateTrue="val T1.MST2[*]"
	MST2 map[string]string `json:"mst2"`
	// +validateTrue="field T1.MSPT2"
	// +eachKey=+validateTrue="key T1.MSPT2[*]"
	// +eachVal=+validateTrue="val T1.MSPT2[*]"
	MSPT2 map[string]*string `json:"mspt2"`

	// +validateTrue="field T1.MSE1"
	// +eachKey=+validateTrue="key T1.MSE1[*]"
	// +eachVal=+validateTrue="val T1.MSE1[*]"
	MSE1 map[string]E1 `json:"mse1"`
	// +validateTrue="field T1.ME1S"
	// +eachKey=+validateTrue="key T1.ME1S[*]"
	// +eachVal=+validateTrue="val T1.ME1S[*]"
	ME1S map[E1]string

	// Duplicate types with no validation.
	AnotherMSS   map[string]string   `json:"anothermss"`
	AnotherMSPS  map[string]*string  `json:"anothermsps"`
	AnotherMPSS  map[*string]string  `json:"anothermpss"`
	AnotherMPSPS map[*string]*string `json:"anothermpsps"`
	AnotherMST2  map[string]string   `json:"anothermst2"`
	AnotherMSPT2 map[string]*string  `json:"anothermspt2"`
	AnotherMSE1  map[string]E1       `json:"anothermse1"`
	AnotherME1S  map[E1]string
}

// +validateTrue="type T2"
type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
	// +validateTrue="field T2.PS"
	PS string `json:"ps"`
}
