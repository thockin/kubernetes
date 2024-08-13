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

package elidenovalidations

import (
	fmt "fmt"

	validate "k8s.io/apimachinery/pkg/api/validate"
	runtime "k8s.io/apimachinery/pkg/runtime"
	field "k8s.io/apimachinery/pkg/util/validation/field"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *runtime.Scheme) error {
	scheme.AddValidationFunc((*T1)(nil), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(obj.(*T1), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_HasFieldVal(obj *HasFieldVal, fldPath *field.Path) (errs field.ErrorList) {
	// field HasFieldVal.S
	errs = append(errs,
		func(obj string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field HasFieldVal.S")...)
			return
		}(obj.S, fldPath.Child("s"))...)

	return errs
}

func Validate_HasTypeVal(obj *HasTypeVal, fldPath *field.Path) (errs field.ErrorList) {
	// type HasTypeVal
	if obj != nil {
		errs = append(errs, validate.FixedResult(fldPath, *obj, true, "type HasTypeVal")...)
	}

	// field HasTypeVal.S has no validation
	return errs
}

func Validate_T1(obj *T1, fldPath *field.Path) (errs field.ErrorList) {
	// field T1.TypeMeta has no validation

	// field T1.HasTypeVal
	errs = append(errs,
		func(obj HasTypeVal, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, Validate_HasTypeVal(&obj, fldPath)...)
			return
		}(obj.HasTypeVal, fldPath.Child("hasTypeVal"))...)

	// field T1.HasFieldVal
	errs = append(errs,
		func(obj HasFieldVal, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, Validate_HasFieldVal(&obj, fldPath)...)
			return
		}(obj.HasFieldVal, fldPath.Child("hasFieldVal"))...)

	// field T1.HasNoVal has no validation

	// field T1.HasNoValFieldVal
	errs = append(errs,
		func(obj HasNoVal, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.HasNoValFieldVal")...)
			return
		}(obj.HasNoValFieldVal, fldPath.Child("hasNoValFieldVal"))...)

	return errs
}
