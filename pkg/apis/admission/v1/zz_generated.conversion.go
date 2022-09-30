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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	"unsafe"

	apiadmissionv1 "k8s.io/api/admission/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	pkgapisadmission "k8s.io/kubernetes/pkg/apis/admission"
	apisauthenticationv1 "k8s.io/kubernetes/pkg/apis/authentication/v1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*apiadmissionv1.AdmissionRequest)(nil), (*pkgapisadmission.AdmissionRequest)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_AdmissionRequest_To_admission_AdmissionRequest(a.(*apiadmissionv1.AdmissionRequest), b.(*pkgapisadmission.AdmissionRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisadmission.AdmissionRequest)(nil), (*apiadmissionv1.AdmissionRequest)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_admission_AdmissionRequest_To_v1_AdmissionRequest(a.(*pkgapisadmission.AdmissionRequest), b.(*apiadmissionv1.AdmissionRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiadmissionv1.AdmissionResponse)(nil), (*pkgapisadmission.AdmissionResponse)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_AdmissionResponse_To_admission_AdmissionResponse(a.(*apiadmissionv1.AdmissionResponse), b.(*pkgapisadmission.AdmissionResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisadmission.AdmissionResponse)(nil), (*apiadmissionv1.AdmissionResponse)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_admission_AdmissionResponse_To_v1_AdmissionResponse(a.(*pkgapisadmission.AdmissionResponse), b.(*apiadmissionv1.AdmissionResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiadmissionv1.AdmissionReview)(nil), (*pkgapisadmission.AdmissionReview)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_AdmissionReview_To_admission_AdmissionReview(a.(*apiadmissionv1.AdmissionReview), b.(*pkgapisadmission.AdmissionReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisadmission.AdmissionReview)(nil), (*apiadmissionv1.AdmissionReview)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_admission_AdmissionReview_To_v1_AdmissionReview(a.(*pkgapisadmission.AdmissionReview), b.(*apiadmissionv1.AdmissionReview), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_AdmissionRequest_To_admission_AdmissionRequest(in *apiadmissionv1.AdmissionRequest, out *pkgapisadmission.AdmissionRequest, s apimachinerypkgconversion.Scope) error {
	out.UID = apimachinerypkgtypes.UID(in.UID)
	out.Kind = in.Kind
	out.Resource = in.Resource
	out.SubResource = in.SubResource
	out.RequestKind = (*apismetav1.GroupVersionKind)(unsafe.Pointer(in.RequestKind))
	out.RequestResource = (*apismetav1.GroupVersionResource)(unsafe.Pointer(in.RequestResource))
	out.RequestSubResource = in.RequestSubResource
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Operation = pkgapisadmission.Operation(in.Operation)
	if err := apisauthenticationv1.Convert_v1_UserInfo_To_authentication_UserInfo(&in.UserInfo, &out.UserInfo, s); err != nil {
		return err
	}
	if err := apimachinerypkgruntime.Convert_runtime_RawExtension_To_runtime_Object(&in.Object, &out.Object, s); err != nil {
		return err
	}
	if err := apimachinerypkgruntime.Convert_runtime_RawExtension_To_runtime_Object(&in.OldObject, &out.OldObject, s); err != nil {
		return err
	}
	out.DryRun = (*bool)(unsafe.Pointer(in.DryRun))
	if err := apimachinerypkgruntime.Convert_runtime_RawExtension_To_runtime_Object(&in.Options, &out.Options, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_AdmissionRequest_To_admission_AdmissionRequest is an autogenerated conversion function.
func Convert_v1_AdmissionRequest_To_admission_AdmissionRequest(in *apiadmissionv1.AdmissionRequest, out *pkgapisadmission.AdmissionRequest, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_AdmissionRequest_To_admission_AdmissionRequest(in, out, s)
}

func autoConvert_admission_AdmissionRequest_To_v1_AdmissionRequest(in *pkgapisadmission.AdmissionRequest, out *apiadmissionv1.AdmissionRequest, s apimachinerypkgconversion.Scope) error {
	out.UID = apimachinerypkgtypes.UID(in.UID)
	out.Kind = in.Kind
	out.Resource = in.Resource
	out.SubResource = in.SubResource
	out.RequestKind = (*apismetav1.GroupVersionKind)(unsafe.Pointer(in.RequestKind))
	out.RequestResource = (*apismetav1.GroupVersionResource)(unsafe.Pointer(in.RequestResource))
	out.RequestSubResource = in.RequestSubResource
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Operation = apiadmissionv1.Operation(in.Operation)
	if err := apisauthenticationv1.Convert_authentication_UserInfo_To_v1_UserInfo(&in.UserInfo, &out.UserInfo, s); err != nil {
		return err
	}
	if err := apimachinerypkgruntime.Convert_runtime_Object_To_runtime_RawExtension(&in.Object, &out.Object, s); err != nil {
		return err
	}
	if err := apimachinerypkgruntime.Convert_runtime_Object_To_runtime_RawExtension(&in.OldObject, &out.OldObject, s); err != nil {
		return err
	}
	out.DryRun = (*bool)(unsafe.Pointer(in.DryRun))
	if err := apimachinerypkgruntime.Convert_runtime_Object_To_runtime_RawExtension(&in.Options, &out.Options, s); err != nil {
		return err
	}
	return nil
}

// Convert_admission_AdmissionRequest_To_v1_AdmissionRequest is an autogenerated conversion function.
func Convert_admission_AdmissionRequest_To_v1_AdmissionRequest(in *pkgapisadmission.AdmissionRequest, out *apiadmissionv1.AdmissionRequest, s apimachinerypkgconversion.Scope) error {
	return autoConvert_admission_AdmissionRequest_To_v1_AdmissionRequest(in, out, s)
}

func autoConvert_v1_AdmissionResponse_To_admission_AdmissionResponse(in *apiadmissionv1.AdmissionResponse, out *pkgapisadmission.AdmissionResponse, s apimachinerypkgconversion.Scope) error {
	out.UID = apimachinerypkgtypes.UID(in.UID)
	out.Allowed = in.Allowed
	out.Result = (*apismetav1.Status)(unsafe.Pointer(in.Result))
	out.Patch = *(*[]byte)(unsafe.Pointer(&in.Patch))
	out.PatchType = (*pkgapisadmission.PatchType)(unsafe.Pointer(in.PatchType))
	out.AuditAnnotations = *(*map[string]string)(unsafe.Pointer(&in.AuditAnnotations))
	out.Warnings = *(*[]string)(unsafe.Pointer(&in.Warnings))
	return nil
}

// Convert_v1_AdmissionResponse_To_admission_AdmissionResponse is an autogenerated conversion function.
func Convert_v1_AdmissionResponse_To_admission_AdmissionResponse(in *apiadmissionv1.AdmissionResponse, out *pkgapisadmission.AdmissionResponse, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_AdmissionResponse_To_admission_AdmissionResponse(in, out, s)
}

func autoConvert_admission_AdmissionResponse_To_v1_AdmissionResponse(in *pkgapisadmission.AdmissionResponse, out *apiadmissionv1.AdmissionResponse, s apimachinerypkgconversion.Scope) error {
	out.UID = apimachinerypkgtypes.UID(in.UID)
	out.Allowed = in.Allowed
	out.Result = (*apismetav1.Status)(unsafe.Pointer(in.Result))
	out.Patch = *(*[]byte)(unsafe.Pointer(&in.Patch))
	out.PatchType = (*apiadmissionv1.PatchType)(unsafe.Pointer(in.PatchType))
	out.AuditAnnotations = *(*map[string]string)(unsafe.Pointer(&in.AuditAnnotations))
	out.Warnings = *(*[]string)(unsafe.Pointer(&in.Warnings))
	return nil
}

// Convert_admission_AdmissionResponse_To_v1_AdmissionResponse is an autogenerated conversion function.
func Convert_admission_AdmissionResponse_To_v1_AdmissionResponse(in *pkgapisadmission.AdmissionResponse, out *apiadmissionv1.AdmissionResponse, s apimachinerypkgconversion.Scope) error {
	return autoConvert_admission_AdmissionResponse_To_v1_AdmissionResponse(in, out, s)
}

func autoConvert_v1_AdmissionReview_To_admission_AdmissionReview(in *apiadmissionv1.AdmissionReview, out *pkgapisadmission.AdmissionReview, s apimachinerypkgconversion.Scope) error {
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(pkgapisadmission.AdmissionRequest)
		if err := Convert_v1_AdmissionRequest_To_admission_AdmissionRequest(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Request = nil
	}
	out.Response = (*pkgapisadmission.AdmissionResponse)(unsafe.Pointer(in.Response))
	return nil
}

// Convert_v1_AdmissionReview_To_admission_AdmissionReview is an autogenerated conversion function.
func Convert_v1_AdmissionReview_To_admission_AdmissionReview(in *apiadmissionv1.AdmissionReview, out *pkgapisadmission.AdmissionReview, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_AdmissionReview_To_admission_AdmissionReview(in, out, s)
}

func autoConvert_admission_AdmissionReview_To_v1_AdmissionReview(in *pkgapisadmission.AdmissionReview, out *apiadmissionv1.AdmissionReview, s apimachinerypkgconversion.Scope) error {
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(apiadmissionv1.AdmissionRequest)
		if err := Convert_admission_AdmissionRequest_To_v1_AdmissionRequest(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Request = nil
	}
	out.Response = (*apiadmissionv1.AdmissionResponse)(unsafe.Pointer(in.Response))
	return nil
}

// Convert_admission_AdmissionReview_To_v1_AdmissionReview is an autogenerated conversion function.
func Convert_admission_AdmissionReview_To_v1_AdmissionReview(in *pkgapisadmission.AdmissionReview, out *apiadmissionv1.AdmissionReview, s apimachinerypkgconversion.Scope) error {
	return autoConvert_admission_AdmissionReview_To_v1_AdmissionReview(in, out, s)
}
