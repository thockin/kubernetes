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
package enum

// +validateTrue="type E1"
type E1 string

type T1 struct {
	TypeMeta int

	// +validateTrue="field T1.MSE1"
	// +eachKey=+validateTrue="T1.MSE1[keys]"
	// +eachVal=+validateTrue="T1.MSE1[vals]"
	MSE1 map[string]E1 `json:"mse1"`

	// +validateTrue="field T1.MSPE1"
	// +eachKey=+validateTrue="T1.MSPE1[keys]"
	// +eachVal=+validateTrue="T1.MSPE1[vals]"
	MSPE1 map[string]*E1 `json:"mspe1"`

	// +validateTrue="field T1.ME1S"
	// +eachKey=+validateTrue="T1.ME1S[keys]"
	// +eachVal=+validateTrue="T1.ME1S[vals]"
	ME1S map[E1]string `json:"me1s"`

	// +validateTrue="field T1.MPE1S"
	// +eachKey=+validateTrue="T1.MPE1S[keys]"
	// +eachVal=+validateTrue="T1.MPE1S[vals]"
	MPE1S map[*E1]string `json:"mpe1s"`
}
