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

package recursive

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
	scheme.AddValidationFunc((E1)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_E1(opCtx, obj.(E1), safe.Cast[E1](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc((E2)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_E2(opCtx, obj.(E2), safe.Cast[E2](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc((*T1)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(opCtx, obj.(*T1), safe.Cast[*T1](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc((*T2)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T2(opCtx, obj.(*T2), safe.Cast[*T2](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_E1(opCtx operation.Context, obj, oldObj E1, fldPath *field.Path) (errs field.ErrorList) {
	// type E1
	errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type E1")...)

	for i, val := range obj {
		errs = append(errs,
			func(obj, oldObj E1, fldPath *field.Path) (errs field.ErrorList) {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type E1 values")...)
				errs = append(errs, Validate_E1(opCtx, obj, oldObj, fldPath)...)
				return
			}(val, nil, fldPath.Index(i))...)
	}
	return errs
}

func Validate_E2(opCtx operation.Context, obj, oldObj E2, fldPath *field.Path) (errs field.ErrorList) {
	// type E2
	errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type E2")...)

	for i, val := range obj {
		errs = append(errs,
			func(obj, oldObj E2, fldPath *field.Path) (errs field.ErrorList) {
				errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type E2 values")...)
				errs = append(errs, Validate_E2(opCtx, obj, oldObj, fldPath)...)
				return
			}(*val, nil, fldPath.Index(i))...)
	}
	return errs
}

func Validate_T1(opCtx operation.Context, obj, oldObj *T1, fldPath *field.Path) (errs field.ErrorList) {
	// type T1
	errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type T1")...)

	// field T1.PT1
	errs = append(errs,
		func(obj, oldObj *T1, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.PT1")...)
			errs = append(errs, Validate_T1(opCtx, obj, oldObj, fldPath)...)
			return
		}(obj.PT1, safe.Field(oldObj, func(oldObj *T1) *T1 { return oldObj.PT1 }), fldPath.Child("pt1"))...)

	// field T1.T2
	errs = append(errs,
		func(obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.T2")...)
			errs = append(errs, Validate_T2(opCtx, obj, oldObj, fldPath)...)
			return
		}(&obj.T2, safe.Field(oldObj, func(oldObj *T1) *T2 { return &oldObj.T2 }), fldPath.Child("t2"))...)

	// field T1.PT2
	errs = append(errs,
		func(obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T1.PT2")...)
			errs = append(errs, Validate_T2(opCtx, obj, oldObj, fldPath)...)
			return
		}(obj.PT2, safe.Field(oldObj, func(oldObj *T1) *T2 { return oldObj.PT2 }), fldPath.Child("pt2"))...)

	return errs
}

func Validate_T2(opCtx operation.Context, obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
	// type T2
	errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "type T2")...)

	// field T2.PT1
	errs = append(errs,
		func(obj, oldObj *T1, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T2.PT1")...)
			errs = append(errs, Validate_T1(opCtx, obj, oldObj, fldPath)...)
			return
		}(obj.PT1, safe.Field(oldObj, func(oldObj *T2) *T1 { return oldObj.PT1 }), fldPath.Child("pt1"))...)

	// field T2.PT2
	errs = append(errs,
		func(obj, oldObj *T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(opCtx, fldPath, obj, oldObj, true, "field T2.PT2")...)
			errs = append(errs, Validate_T2(opCtx, obj, oldObj, fldPath)...)
			return
		}(obj.PT2, safe.Field(oldObj, func(oldObj *T2) *T2 { return oldObj.PT2 }), fldPath.Child("pt2"))...)

	return errs
}
