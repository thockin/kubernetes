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

package primitivestruct

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
	scheme.AddValidationFunc(new(T1), func(obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(obj.(*T1), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_T1(obj *T1, fldPath *field.Path) (errs field.ErrorList) {
	// type T1
	if obj != nil {
		errs = append(errs, validate.FixedResult(fldPath, *obj, true, "type T1")...)
	}

	// field T1.TypeMeta has no validation

	// field T1.MST2
	errs = append(errs,
		func(obj map[string]T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.MST2")...)
			for key, val := range obj {
				errs = append(errs,
					func(obj string, fldPath *field.Path) (errs field.ErrorList) {
						errs = append(errs, validate.FixedResult(fldPath, obj, true, "T1.MST2[keys]")...)
						return
					}(key, fldPath)...)
				errs = append(errs,
					func(obj T2, fldPath *field.Path) (errs field.ErrorList) {
						errs = append(errs, validate.FixedResult(fldPath, obj, true, "T1.MST2[vals]")...)
						errs = append(errs, Validate_T2(&obj, fldPath)...)
						return
					}(val, fldPath.Key(key))...)
			}
			return
		}(obj.MST2, fldPath.Child("mst2"))...)

	// field T1.MSPT2
	errs = append(errs,
		func(obj map[string]*T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.MSPT2")...)
			for key, val := range obj {
				errs = append(errs,
					func(obj string, fldPath *field.Path) (errs field.ErrorList) {
						errs = append(errs, validate.FixedResult(fldPath, obj, true, "T1.MSPT2[keys]")...)
						return
					}(key, fldPath)...)
				errs = append(errs,
					func(obj *T2, fldPath *field.Path) (errs field.ErrorList) {
						if obj != nil {
							errs = append(errs, validate.FixedResult(fldPath, *obj, true, "T1.MSPT2[vals]")...)
						}
						if obj != nil {
							errs = append(errs, Validate_T2(obj, fldPath)...)
						}
						return
					}(val, fldPath.Key(key))...)
			}
			return
		}(obj.MSPT2, fldPath.Child("mspt2"))...)

	return errs
}

func Validate_T2(obj *T2, fldPath *field.Path) (errs field.ErrorList) {
	// type T2
	if obj != nil {
		errs = append(errs, validate.FixedResult(fldPath, *obj, true, "type T2")...)
	}

	// field T2.S
	errs = append(errs,
		func(obj string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T2.S")...)
			return
		}(obj.S, fldPath.Child("s"))...)

	return errs
}
