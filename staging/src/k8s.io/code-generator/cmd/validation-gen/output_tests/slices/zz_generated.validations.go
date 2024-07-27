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

package slices

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

	// field T1.LS
	errs = append(errs,
		func(obj []string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.LS")...)
			for i, val := range obj {
				errs = append(errs,
					func(obj string, fldPath *field.Path) (errs field.ErrorList) {
						if e := validate.Required(fldPath, obj); len(e) != 0 {
							errs = append(errs, e...)
							return // fatal
						}
						errs = append(errs, validate.FixedResult(fldPath, obj, true, "val T1.LS[*]")...)
						return
					}(val, fldPath.Index(i))...)
			}
			return
		}(obj.LS, fldPath.Child("ls"))...)

	// field T1.LPS
	errs = append(errs,
		func(obj []*string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.LPS")...)
			for i, val := range obj {
				errs = append(errs,
					func(obj *string, fldPath *field.Path) (errs field.ErrorList) {
						if e := validate.Required(fldPath, obj); len(e) != 0 {
							errs = append(errs, e...)
							return // fatal
						}
						if obj != nil {
							errs = append(errs, validate.FixedResult(fldPath, *obj, true, "val T1.LPS[*]")...)
						}
						return
					}(val, fldPath.Index(i))...)
			}
			return
		}(obj.LPS, fldPath.Child("lps"))...)

	// field T1.LT2
	errs = append(errs,
		func(obj []T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.LT2")...)
			for i, val := range obj {
				errs = append(errs,
					func(obj T2, fldPath *field.Path) (errs field.ErrorList) {
						if e := validate.Required(fldPath, obj); len(e) != 0 {
							errs = append(errs, e...)
							return // fatal
						}
						errs = append(errs, validate.FixedResult(fldPath, obj, true, "val T1.LT2[*]")...)
						errs = append(errs, Validate_T2(&obj, fldPath)...)
						return
					}(val, fldPath.Index(i))...)
			}
			return
		}(obj.LT2, fldPath.Child("lt2"))...)

	// field T1.LPT2
	errs = append(errs,
		func(obj []*T2, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field T1.LPT2")...)
			for i, val := range obj {
				errs = append(errs,
					func(obj *T2, fldPath *field.Path) (errs field.ErrorList) {
						if e := validate.Required(fldPath, obj); len(e) != 0 {
							errs = append(errs, e...)
							return // fatal
						}
						if obj != nil {
							errs = append(errs, validate.FixedResult(fldPath, *obj, true, "val T1.LPT2[*]")...)
						}
						if obj != nil {
							errs = append(errs, Validate_T2(obj, fldPath)...)
						}
						return
					}(val, fldPath.Index(i))...)
			}
			return
		}(obj.LPT2, fldPath.Child("lpt2"))...)

	// field T1.AnotherLS has no validation
	// field T1.AnotherLPS has no validation

	// field T1.AnotherLT2
	errs = append(errs,
		func(obj []T2, fldPath *field.Path) (errs field.ErrorList) {
			for i, val := range obj {
				errs = append(errs,
					func(obj T2, fldPath *field.Path) (errs field.ErrorList) {
						errs = append(errs, Validate_T2(&obj, fldPath)...)
						return
					}(val, fldPath.Index(i))...)
			}
			return
		}(obj.AnotherLT2, fldPath.Child("anotherlt2"))...)

	// field T1.AnotherLPT2
	errs = append(errs,
		func(obj []*T2, fldPath *field.Path) (errs field.ErrorList) {
			for i, val := range obj {
				errs = append(errs,
					func(obj *T2, fldPath *field.Path) (errs field.ErrorList) {
						if obj != nil {
							errs = append(errs, Validate_T2(obj, fldPath)...)
						}
						return
					}(val, fldPath.Index(i))...)
			}
			return
		}(obj.AnotherLPT2, fldPath.Child("anotherlpt2"))...)

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
