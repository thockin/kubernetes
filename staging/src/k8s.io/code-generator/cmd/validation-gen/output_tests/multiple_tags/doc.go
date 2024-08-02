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
package multipletags

// +validateTrue="type T1 #1"
// +validateTrue="type T1 #2"
// +validateTrue="type T1 #3"
type T1 struct {
	TypeMeta int
	// +validateTrue="field T1.S #1"
	// +validateTrue="field T1.S #2"
	// +validateTrue="field T1.S #3"
	S string `json:"s"`
	// +validateTrue="field T1.T2 #1"
	// +validateTrue="field T1.T2 #2"
	// +validateTrue="field T1.T2 #3"
	T2 T2 `json:"t2"`
}

// +validateTrue="type T2 #1"
// +validateTrue="type T2 #2"
type T2 struct {
	// +validateTrue="field T2.S #1"
	// +validateTrue="field T2.S #2"
	S string `json:"s"`
}
