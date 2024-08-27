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

package multiple_discriminated_unions

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
	scheme.AddValidationFunc((*DU)(nil), func(opCtx operation.Context, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_DU(opCtx, obj.(*DU), safe.Cast[DU](oldObj), nil)
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

var unionMembershipForDUunion1 = validate.NewDiscriminatedUnionMembership("du1", [2]string{"u1m1", "U1M1"}, [2]string{"u1m2", "U1M2"})
var unionMembershipForDUunion2 = validate.NewDiscriminatedUnionMembership("du2", [2]string{"u2m1", "U2M1"}, [2]string{"u2m2", "U2M2"})

func Validate_DU(opCtx operation.Context, obj, oldObj *DU, fldPath *field.Path) (errs field.ErrorList) {
	// type DU
	if obj != nil {
		errs = append(errs, validate.DiscriminatedUnion(fldPath, *obj, unionMembershipForDUunion1, obj.DU1, obj.U1M1, obj.U1M2)...)
		errs = append(errs, validate.DiscriminatedUnion(fldPath, *obj, unionMembershipForDUunion2, obj.DU2, obj.U2M1, obj.U2M2)...)
	}

	// field DU.TypeMeta has no validation
	// field DU.DU1 has no validation

	// field DU.U1M1
	errs = append(errs,
		func(obj *M1, oldObj *M1, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, Validate_M1(opCtx, obj, oldObj, fldPath)...)
			}
			return
		}(obj.U1M1, safe.Field(oldObj, func(oldObj DU) *M1 { return oldObj.U1M1 }), fldPath.Child("u1m1"))...)

	// field DU.U1M2
	errs = append(errs,
		func(obj *M2, oldObj *M2, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, Validate_M2(opCtx, obj, oldObj, fldPath)...)
			}
			return
		}(obj.U1M2, safe.Field(oldObj, func(oldObj DU) *M2 { return oldObj.U1M2 }), fldPath.Child("u1m2"))...)

	// field DU.DU2 has no validation

	// field DU.U2M1
	errs = append(errs,
		func(obj *M1, oldObj *M1, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, Validate_M1(opCtx, obj, oldObj, fldPath)...)
			}
			return
		}(obj.U2M1, safe.Field(oldObj, func(oldObj DU) *M1 { return oldObj.U2M1 }), fldPath.Child("u2m1"))...)

	// field DU.U2M2
	errs = append(errs,
		func(obj *M2, oldObj *M2, fldPath *field.Path) (errs field.ErrorList) {
			if obj != nil {
				errs = append(errs, Validate_M2(opCtx, obj, oldObj, fldPath)...)
			}
			return
		}(obj.U2M2, safe.Field(oldObj, func(oldObj DU) *M2 { return oldObj.U2M2 }), fldPath.Child("u2m2"))...)

	return errs
}

func Validate_M1(opCtx operation.Context, obj, oldObj *M1, fldPath *field.Path) (errs field.ErrorList) {
	// type M1
	if obj != nil {
		errs = append(errs, validate.FixedResult(fldPath, *obj, true, "type M1")...)
	}

	// field M1.S
	errs = append(errs,
		func(obj string, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field M1.S")...)
			return
		}(obj.S, safe.Field(oldObj, func(oldObj M1) *string { return &oldObj.S }), fldPath.Child("s"))...)

	return errs
}

func Validate_M2(opCtx operation.Context, obj, oldObj *M2, fldPath *field.Path) (errs field.ErrorList) {
	// type M2
	if obj != nil {
		errs = append(errs, validate.FixedResult(fldPath, *obj, true, "type M2")...)
	}

	// field M2.S
	errs = append(errs,
		func(obj string, oldObj *string, fldPath *field.Path) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(fldPath, obj, true, "field M2.S")...)
			return
		}(obj.S, safe.Field(oldObj, func(oldObj M2) *string { return &oldObj.S }), fldPath.Child("s"))...)

	return errs
}
