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
// +k8s:validation-gen-scheme-registry=k8s.io/code-generator/cmd/validation-gen/testscheme.Scheme
// +k8s:validation-gen-test-fixture=validateFalse

// This is a test package.
package typedef

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

// +k8s:validateFalse="type T1"
type T1 struct {
	TypeMeta int

	// +k8s:validateFalse="field T1.MSAMSS"
	// +k8s:eachKey=+k8s:validateFalse="T1.MSAMSS[keys]"
	// +k8s:eachVal=+k8s:validateFalse="T1.MSAMSS[vals]"
	MSAMSS map[string]AMSS `json:"msamss"`

	// +k8s:validateTrue="field T1.MSPAMSS"
	// +k8s:eachKey=+k8s:validateTrue="T1.MSPAMSS[keys]"
	// +k8s:eachVal=+k8s:validateTrue="T1.MSPAMSS[vals]"
	MSPAMSS map[string]*AMSS `json:"mspamss"`
}

// +k8s:validateFalse="type AMSS"
// +k8s:eachKey=+k8s:validateFalse="AMSS[keys]"
// +k8s:eachVal=+k8s:validateFalse="AMSS[vals]"
type AMSS map[string]string
