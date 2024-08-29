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

package primitives_ptr_ok

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
	scheme.AddValidationFunc((*T1)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(opCtx, obj.(*T1), safe.Cast[T1](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_T1(opCtx operation.Context, obj, oldObj *T1, fldPath *field.Path) (errs field.ErrorList) {
	// field T1.S
	errs = append(errs,
		func(obj string, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, &obj, oldObj, true, "T1.S")...)
			return
		}(obj.S, safe.Field(oldObj, func(oldObj T1) *string { return &oldObj.S }), fldPath.Child("s"))...)

	// field T1.I
	errs = append(errs,
		func(obj int, oldObj *int, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, &obj, oldObj, true, "T1.I")...)
			return
		}(obj.I, safe.Field(oldObj, func(oldObj T1) *int { return &oldObj.I }), fldPath.Child("i"))...)

	// field T1.B
	errs = append(errs,
		func(obj bool, oldObj *bool, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, &obj, oldObj, true, "T1.B")...)
			return
		}(obj.B, safe.Field(oldObj, func(oldObj T1) *bool { return &oldObj.B }), fldPath.Child("b"))...)

	// field T1.F
	errs = append(errs,
		func(obj float64, oldObj *float64, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, &obj, oldObj, true, "T1.F")...)
			return
		}(obj.F, safe.Field(oldObj, func(oldObj T1) *float64 { return &oldObj.F }), fldPath.Child("f"))...)

	return errs
}
