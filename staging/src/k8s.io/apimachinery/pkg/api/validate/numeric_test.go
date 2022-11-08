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

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestGEZ(t *testing.T) {
	cases := []struct {
		val int
		ok  bool
	}{{
		val: -1,
		ok:  false,
	}, {
		val: 0,
		ok:  true,
	}, {
		val: 1,
		ok:  true,
	}}

	for _, tc := range cases {
		var check func(errs field.ErrorList)
		if tc.ok {
			check = func(errs field.ErrorList) {
				if len(errs) != 0 {
					t.Errorf("unexpected failure for input %d: %v", tc.val, errs)
				}
			}
		} else {
			check = func(errs field.ErrorList) {
				if len(errs) == 0 {
					t.Errorf("unexpected success for input %d", tc.val)
				} else if len(errs) > 1 {
					t.Errorf("unexpected failures for input %d: %v", tc.val, errs)
				}
			}
		}
		check(GEZ(tc.val, field.NewPath("fieldname")))
		check(GEZ(int16(tc.val), field.NewPath("fieldname")))
		check(GEZ(int32(tc.val), field.NewPath("fieldname")))
		check(GEZ(int64(tc.val), field.NewPath("fieldname")))
	}
}

func TestGTZ(t *testing.T) {
	cases := []struct {
		val int
		ok  bool
	}{{
		val: -1,
		ok:  false,
	}, {
		val: 0,
		ok:  false,
	}, {
		val: 1,
		ok:  true,
	}}

	for _, tc := range cases {
		var check func(errs field.ErrorList)
		if tc.ok {
			check = func(errs field.ErrorList) {
				if len(errs) != 0 {
					t.Errorf("unexpected failure for input %d: %v", tc.val, errs)
				}
			}
		} else {
			check = func(errs field.ErrorList) {
				if len(errs) == 0 {
					t.Errorf("unexpected success for input %d", tc.val)
				} else if len(errs) > 1 {
					t.Errorf("unexpected failures for input %d: %v", tc.val, errs)
				}
			}
		}
		check(GTZ(tc.val, field.NewPath("fieldname")))
		check(GTZ(int16(tc.val), field.NewPath("fieldname")))
		check(GTZ(int32(tc.val), field.NewPath("fieldname")))
		check(GTZ(int64(tc.val), field.NewPath("fieldname")))
	}
}
