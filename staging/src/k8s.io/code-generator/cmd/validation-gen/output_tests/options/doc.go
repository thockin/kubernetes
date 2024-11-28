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

// This is a test package.
package options

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

type T1 struct {
	TypeMeta int

	// +k8s:ifOptionEnabled=FeatureX=+k8s:validateFalse="field T1.S1"
	S1 string `json:"s1"`

	// +k8s:ifOptionDisabled=FeatureX=+k8s:validateFalse="field T1.S2"
	S2 string `json:"s2"`

	// +k8s:ifOptionEnabled=FeatureX=+k8s:validateFalse="field T1.S3.FeatureX"
	// +k8s:ifOptionDisabled=FeatureY=+k8s:validateFalse="field T1.S3.FeatureY"
	S3 string `json:"s3"`
}
