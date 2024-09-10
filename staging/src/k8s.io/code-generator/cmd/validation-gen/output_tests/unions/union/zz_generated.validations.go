//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by validation-gen. DO NOT EDIT.

package union

import (
	fmt "fmt"

	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	field "k8s.io/apimachinery/pkg/util/validation/field"
	testscheme "k8s.io/code-generator/cmd/validation-gen/testscheme"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *testscheme.Scheme) error {
	scheme.AddValidationFunc((*U)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_U(opCtx, obj.(*U), safe.Cast[*U](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

var unionMembershipForU = validate.NewUnionMembership([2]string{"m1", "M1"}, [2]string{"m2", "M2"})

func Validate_U(opCtx operation.Context, obj, oldObj *U, fldPath *field.Path) (errs field.ErrorList) {
	// type U
	errs = append(errs, validate.Union(opCtx, fldPath, obj, oldObj, unionMembershipForU, obj.M1, obj.M2)...)

	// field U.TypeMeta has no validation

	// field U.M1
	errs = append(errs,
		func(obj, oldObj *M1, fldPath *field.Path) (errs field.ErrorList) {
			if e := validate.Optional(opCtx, fldPath, obj, oldObj); len(e) != 0 {
				return // do not proceed
			}
			return
		}(obj.M1, safe.Field(oldObj, func(oldObj *U) *M1 { return oldObj.M1 }), fldPath.Child("m1"))...)

	// field U.M2
	errs = append(errs,
		func(obj, oldObj *M2, fldPath *field.Path) (errs field.ErrorList) {
			if e := validate.Optional(opCtx, fldPath, obj, oldObj); len(e) != 0 {
				return // do not proceed
			}
			return
		}(obj.M2, safe.Field(oldObj, func(oldObj *U) *M2 { return oldObj.M2 }), fldPath.Child("m2"))...)

	return errs
}
