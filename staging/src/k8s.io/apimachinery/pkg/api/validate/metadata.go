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

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// HasLabelValue validates that metadata has a label with key = value.
func HasLabelValue(metadata metav1.ObjectMeta, fldPath *field.Path, key, value string) field.ErrorList {
	actualValue, found := metadata.Labels[key]
	if !found {
		return field.ErrorList{field.Required(fldPath.Child("labels").Key(key), fmt.Sprintf("must be '%s'", value))}
	}
	if actualValue != value {
		return field.ErrorList{field.Invalid(fldPath.Child("labels").Key(key), actualValue, fmt.Sprintf("must be '%s'", value))}
	}
	return nil
}
