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
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestHasLabelValue(t *testing.T) {
	input := metav1.ObjectMeta{
		Name:      "name",
		Namespace: "ns",
		Labels: map[string]string{
			"key":   "val",
			"other": "blah",
		},
	}

	// found
	if errs := HasLabelValue(input, field.NewPath("metadata"), "key", "val"); len(errs) != 0 {
		t.Errorf("unexpected failure: %v", errs)
	}

	// not found
	if errs := HasLabelValue(input, field.NewPath("metadata"), "badkey", "val"); len(errs) != 1 {
		if len(errs) == 0 {
			t.Errorf("unexpected success")
		} else {
			t.Errorf("unexpected failures: %v", errs)
		}
	}

	// wrong value
	if errs := HasLabelValue(input, field.NewPath("metadata"), "key", "badval"); len(errs) != 1 {
		if len(errs) == 0 {
			t.Errorf("unexpected success")
		} else {
			t.Errorf("unexpected failures: %v", errs)
		}
	}
}
