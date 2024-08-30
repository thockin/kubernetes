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

package withfieldvalidations

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
			return Validate_T1(opCtx, obj.(*T1), safe.Cast[*T1](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_E1(opCtx operation.Context, obj, oldObj *E1, fldPath *field.Path) (errs field.ErrorList) {
	// type E1
	errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type E1")...)

	return errs
}

func Validate_T1(opCtx operation.Context, obj, oldObj *T1, fldPath *field.Path) (errs field.ErrorList) {
	// field T1.TypeMeta has no validation

	// field T1.S
	errs = append(errs,
		func(obj, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.S")...)
			return
		}(&obj.S, safe.Field(oldObj, func(oldObj *T1) *string { return &oldObj.S }), fldPath.Child("s"))...)

	// field T1.T2
	errs = append(errs,
		func(obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.T2")...)
			errs = append(errs, Validate_T2(opCtx, obj, oldObj, fldPath)...)
			return
		}(&obj.T2, safe.Field(oldObj, func(oldObj *T1) *T2 { return &oldObj.T2 }), fldPath.Child("t2"))...)

	// field T1.T3
	errs = append(errs,
		func(obj, oldObj *T3, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.T3")...)
			return
		}(&obj.T3, safe.Field(oldObj, func(oldObj *T1) *T3 { return &oldObj.T3 }), fldPath.Child("t3"))...)

	// field T1.E1
	errs = append(errs,
		func(obj, oldObj *E1, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.E1")...)
			errs = append(errs, Validate_E1(opCtx, obj, oldObj, fldPath)...)
			return
		}(&obj.E1, safe.Field(oldObj, func(oldObj *T1) *E1 { return &oldObj.E1 }), fldPath.Child("e1"))...)

	// field T1.E2
	errs = append(errs,
		func(obj, oldObj *E2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.E2")...)
			return
		}(&obj.E2, safe.Field(oldObj, func(oldObj *T1) *E2 { return &oldObj.E2 }), fldPath.Child("e2"))...)

	return errs
}

func Validate_T2(opCtx operation.Context, obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
	// field T2.S
	errs = append(errs,
		func(obj, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T2.S")...)
			return
		}(&obj.S, safe.Field(oldObj, func(oldObj *T2) *string { return &oldObj.S }), fldPath.Child("s"))...)

	return errs
}
