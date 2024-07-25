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
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("No validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_T1(in *T1, fldPath *field.Path) (errs field.ErrorList) {
	errs = append(errs, validate.FixedResult(fldPath, in, true, "type T1")...)
	// TypeMeta

	// LS
	errs = append(errs, validate.FixedResult(fldPath.Child("ls"), in.LS, true, "field T1.LS")...)
	for i, val := range in.LS {
		errs = append(errs, validate.FixedResult(fldPath.Index(i), val, true, "field T1.LS[*]")...)
		errs = append(errs, validate.Required(fldPath.Index(i), val)...)
	}

	// LPS
	errs = append(errs, validate.FixedResult(fldPath.Child("lps"), in.LPS, true, "field T1.LPS")...)
	for i, val := range in.LPS {
		if val != nil {
			errs = append(errs, validate.FixedResult(fldPath.Index(i), *val, true, "field T1.LPS[*]")...)
		}
		errs = append(errs, validate.Required(fldPath.Index(i), val)...)
	}

	// LT2
	errs = append(errs, validate.FixedResult(fldPath.Child("lt2"), in.LT2, true, "field T1.LT2")...)
	for i, val := range in.LT2 {
		errs = append(errs, validate.FixedResult(fldPath.Index(i), val, true, "field T1.LT2[*]")...)
		errs = append(errs, validate.Required(fldPath.Index(i), val)...)
		errs = append(errs, Validate_T2(&val, fldPath.Index(i))...)
	}

	// LPT2
	errs = append(errs, validate.FixedResult(fldPath.Child("lpt2"), in.LPT2, true, "field T1.LPT2")...)
	for i, val := range in.LPT2 {
		if val != nil {
			errs = append(errs, validate.FixedResult(fldPath.Index(i), *val, true, "field T1.LPT2[*]")...)
		}
		errs = append(errs, validate.Required(fldPath.Index(i), val)...)
		if val != nil {
			errs = append(errs, Validate_T2(val, fldPath.Index(i))...)
		}
	}

	// AnotherLS
	for i, val := range in.AnotherLS {
	}

	// AnotherLPS
	for i, val := range in.AnotherLPS {
	}

	// AnotherLT2
	for i, val := range in.AnotherLT2 {
		errs = append(errs, Validate_T2(&val, fldPath.Index(i))...)
	}

	// AnotherLPT2
	for i, val := range in.AnotherLPT2 {
		if val != nil {
			errs = append(errs, Validate_T2(val, fldPath.Index(i))...)
		}
	}

	return errs
}

func Validate_T2(in *T2, fldPath *field.Path) (errs field.ErrorList) {
	errs = append(errs, validate.FixedResult(fldPath, in, true, "type T2")...)
	// S
	errs = append(errs, validate.FixedResult(fldPath.Child("s"), in.S, true, "field T2.S")...)

	return errs
}
