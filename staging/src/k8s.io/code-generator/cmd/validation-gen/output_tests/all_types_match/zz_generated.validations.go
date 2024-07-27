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

package alltypesmatch

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
	scheme.AddValidationFunc(new(ES), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_ES(obj.(*ES), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc(new(ET1), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_ET1(obj.(*ET1), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc(new(T1), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(obj.(*T1), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc(new(T2), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T2(obj.(*T2), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc(new(T3), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T3(obj.(*T3), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc(new(T4), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T4(obj.(*T4), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_ES(obj *ES, fldPath *field.Path) (errs field.ErrorList) {
	// type ES
	if obj != nil {
		errs = append(errs, validate.FixedResult(fldPath, *obj, true, "type ES")...)
	}

	return errs
}

func Validate_ET1(obj *ET1, fldPath *field.Path) (errs field.ErrorList) {
	// field ET1.S
	errs = append(errs,
		func(obj string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.S")...)
			return
		}(obj.S, fldPath.Child("s"))...)

	// field ET1.T2
	errs = append(errs,
		func(obj T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, Validate_T2(&obj, fldPath)...)
			return
		}(obj.T2, fldPath.Child("t2"))...)

	// field ET1.T3 has no validation
	return errs
}

func Validate_T1(obj *T1, fldPath *field.Path) (errs field.ErrorList) {
	// field T1.S
	errs = append(errs,
		func(obj string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.S")...)
			return
		}(obj.S, fldPath.Child("s"))...)

	// field T1.T2
	errs = append(errs,
		func(obj T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, Validate_T2(&obj, fldPath)...)
			return
		}(obj.T2, fldPath.Child("t2"))...)

	// field T1.T3 has no validation
	return errs
}

func Validate_T2(obj *T2, fldPath *field.Path) (errs field.ErrorList) {
	// field T2.S
	errs = append(errs,
		func(obj string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T2.S")...)
			return
		}(obj.S, fldPath.Child("s"))...)

	return errs
}

func Validate_T3(obj *T3, fldPath *field.Path) (errs field.ErrorList) {
	// field T3.S has no validation
	return errs
}

func Validate_T4(obj *T4, fldPath *field.Path) (errs field.ErrorList) {
	// field T4.S has no validation
	return errs
}
