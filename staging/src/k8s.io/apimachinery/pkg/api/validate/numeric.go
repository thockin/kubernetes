/*
Copyright 2022 The Kubernetes Authors.

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

package validate

import "k8s.io/apimachinery/pkg/util/validation/field"

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// GEZ validates that the specified value is greater than or equal to 0.
func GEZ[T signed](value T, fldPath *field.Path) field.ErrorList {
	if value >= 0 {
		return nil
	}
	return field.ErrorList{field.Invalid(fldPath, value, `must be greater than or equal to 0`)}
}
