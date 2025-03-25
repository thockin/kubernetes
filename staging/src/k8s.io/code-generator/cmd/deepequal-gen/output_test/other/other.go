/*
Copyright 2025 The Kubernetes Authors.

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

// Package other contains test types for testing DeepEqual behavior
package other

// StructWithoutMethod is a simple struct that will get a generated DeepEqual method
type StructWithoutMethod struct {
	StringField string
}

// StructWithMethod is a struct with a manually implemented DeepEqual method
type StructWithMethod struct {
	StringField string
}

// DeepEqual implements custom equality checking
func (in *StructWithMethod) DeepEqual(other *StructWithMethod) bool {
	if in == other {
		return true
	}
	if in == nil || other == nil {
		return false
	}
	return in.StringField == other.StringField
}
