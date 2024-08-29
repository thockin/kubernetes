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

package pointers

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
	// field T1.TypeMeta has no validation

	// field T1.PS
	errs = append(errs,
		func(obj *string, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T1.PS")...)
			}
			return
		}(obj.PS, safe.Field(oldObj, func(oldObj T1) *string { return oldObj.PS }), fldPath.Child("ps"))...)

	// field T1.PI
	errs = append(errs,
		func(obj *int, oldObj *int, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T1.PI")...)
			}
			return
		}(obj.PI, safe.Field(oldObj, func(oldObj T1) *int { return oldObj.PI }), fldPath.Child("pi"))...)

	// field T1.PB
	errs = append(errs,
		func(obj *bool, oldObj *bool, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T1.PB")...)
			}
			return
		}(obj.PB, safe.Field(oldObj, func(oldObj T1) *bool { return oldObj.PB }), fldPath.Child("pb"))...)

	// field T1.PF
	errs = append(errs,
		func(obj *float64, oldObj *float64, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T1.PF")...)
			}
			return
		}(obj.PF, safe.Field(oldObj, func(oldObj T1) *float64 { return oldObj.PF }), fldPath.Child("pf"))...)

	// field T1.PT2
	errs = append(errs,
		func(obj *T2, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T1.PT2")...)
			}
			if obj != nil {
				errs = append(errs, Validate_T2(opCtx, obj, oldObj, fldPath)...)
			}
			return
		}(obj.PT2, safe.Field(oldObj, func(oldObj T1) *T2 { return oldObj.PT2 }), fldPath.Child("pt2"))...)

	// field T1.AnotherPS has no validation
	// field T1.AnotherPI has no validation
	// field T1.AnotherPB has no validation
	// field T1.AnotherPF has no validation
	return errs
}

func Validate_T2(opCtx operation.Context, obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
	// field T2.PS
	errs = append(errs,
		func(obj *string, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T2.PS")...)
			}
			return
		}(obj.PS, safe.Field(oldObj, func(oldObj T2) *string { return oldObj.PS }), fldPath.Child("ps"))...)

	// field T2.PI
	errs = append(errs,
		func(obj *int, oldObj *int, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T2.PI")...)
			}
			return
		}(obj.PI, safe.Field(oldObj, func(oldObj T2) *int { return oldObj.PI }), fldPath.Child("pi"))...)

	// field T2.PB
	errs = append(errs,
		func(obj *bool, oldObj *bool, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T2.PB")...)
			}
			return
		}(obj.PB, safe.Field(oldObj, func(oldObj T2) *bool { return oldObj.PB }), fldPath.Child("pb"))...)

	// field T2.PF
	errs = append(errs,
		func(obj *float64, oldObj *float64, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, *obj, *oldObj, true, "field T2.PF")...)
			}
			return
		}(obj.PF, safe.Field(oldObj, func(oldObj T2) *float64 { return oldObj.PF }), fldPath.Child("pf"))...)

	return errs
}
