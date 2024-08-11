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
package structprimitive

// +validateTrue="type T1"
type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.MT2S"
	// +eachKey=+validateTrue="T1.MT2S[keys]"
	// +eachVal=+validateTrue="T1.MT2S[vals]"
	MT2S map[T2]string `json:"mt2s"`

	// +validateTrue="field T1.MPT2S"
	// +eachKey=+validateTrue="T1.MPT2S[keys]"
	// +eachVal=+validateTrue="T1.MPT2S[vals]"
	MPT2S map[*T2]string `json:"mpt2s"`
}

// +validateTrue="type T2"
type T2 struct {
	// +validateTrue="field T2.S"
	S string `json:"s"`
}
