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
package elidenovalidations

type T1 struct {
	TypeMeta int

	HasTypeVal HasTypeVal `json:"hasTypeVal"`

	HasFieldVal HasFieldVal `json:"hasFieldVal"`

	HasNoVal HasNoVal `json:"hasNoVal"`

	// +validateTrue="field T1.HasNoValFieldVal"
	HasNoValFieldVal HasNoVal `json:"hasNoValFieldVal"`
}

// +validateTrue="type HasTypeVal"
type HasTypeVal struct {
	// Note: no field validation.
	S string `json:"s"`
}

// Note: no type validation.
type HasFieldVal struct {
	// +validateTrue="field HasFieldVal.S"
	S string `json:"s"`
}

// Note: no type validation.
type HasNoVal struct {
	// Note: no field validation.
	S string `json:"s"`
}

// +validateTrue="type HasTypeValNotLinked"
type HasTypeValNotLinked struct {
	// Note: no field validation.
	S string `json:"s"`
}

// Note: no type validation.
type HasFieldValNotLinked struct {
	// +validateTrue="field HasFieldValNotLinked.S"
	S string `json:"s"`
}

// Note: no type validation.
type HasNoValNotLinked struct {
	// Note: no field validation.
	S string `json:"s"`
}
