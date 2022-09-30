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

	apiauthenticationv1 "k8s.io/api/authentication/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	pkgapisaudit "k8s.io/apiserver/pkg/apis/audit"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Event)(nil), (*pkgapisaudit.Event)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_Event_To_audit_Event(a.(*Event), b.(*pkgapisaudit.Event), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.Event)(nil), (*Event)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_Event_To_v1_Event(a.(*pkgapisaudit.Event), b.(*Event), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EventList)(nil), (*pkgapisaudit.EventList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_EventList_To_audit_EventList(a.(*EventList), b.(*pkgapisaudit.EventList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.EventList)(nil), (*EventList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_EventList_To_v1_EventList(a.(*pkgapisaudit.EventList), b.(*EventList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*GroupResources)(nil), (*pkgapisaudit.GroupResources)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_GroupResources_To_audit_GroupResources(a.(*GroupResources), b.(*pkgapisaudit.GroupResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.GroupResources)(nil), (*GroupResources)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_GroupResources_To_v1_GroupResources(a.(*pkgapisaudit.GroupResources), b.(*GroupResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ObjectReference)(nil), (*pkgapisaudit.ObjectReference)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_ObjectReference_To_audit_ObjectReference(a.(*ObjectReference), b.(*pkgapisaudit.ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.ObjectReference)(nil), (*ObjectReference)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_ObjectReference_To_v1_ObjectReference(a.(*pkgapisaudit.ObjectReference), b.(*ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Policy)(nil), (*pkgapisaudit.Policy)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_Policy_To_audit_Policy(a.(*Policy), b.(*pkgapisaudit.Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.Policy)(nil), (*Policy)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_Policy_To_v1_Policy(a.(*pkgapisaudit.Policy), b.(*Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PolicyList)(nil), (*pkgapisaudit.PolicyList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_PolicyList_To_audit_PolicyList(a.(*PolicyList), b.(*pkgapisaudit.PolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.PolicyList)(nil), (*PolicyList)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_PolicyList_To_v1_PolicyList(a.(*pkgapisaudit.PolicyList), b.(*PolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PolicyRule)(nil), (*pkgapisaudit.PolicyRule)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1_PolicyRule_To_audit_PolicyRule(a.(*PolicyRule), b.(*pkgapisaudit.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisaudit.PolicyRule)(nil), (*PolicyRule)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_audit_PolicyRule_To_v1_PolicyRule(a.(*pkgapisaudit.PolicyRule), b.(*PolicyRule), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Event_To_audit_Event(in *Event, out *pkgapisaudit.Event, s apimachinerypkgconversion.Scope) error {
	out.Level = pkgapisaudit.Level(in.Level)
	out.AuditID = apimachinerypkgtypes.UID(in.AuditID)
	out.Stage = pkgapisaudit.Stage(in.Stage)
	out.RequestURI = in.RequestURI
	out.Verb = in.Verb
	out.User = in.User
	out.ImpersonatedUser = (*apiauthenticationv1.UserInfo)(unsafe.Pointer(in.ImpersonatedUser))
	out.SourceIPs = *(*[]string)(unsafe.Pointer(&in.SourceIPs))
	out.UserAgent = in.UserAgent
	out.ObjectRef = (*pkgapisaudit.ObjectReference)(unsafe.Pointer(in.ObjectRef))
	out.ResponseStatus = (*apismetav1.Status)(unsafe.Pointer(in.ResponseStatus))
	out.RequestObject = (*apimachinerypkgruntime.Unknown)(unsafe.Pointer(in.RequestObject))
	out.ResponseObject = (*apimachinerypkgruntime.Unknown)(unsafe.Pointer(in.ResponseObject))
	out.RequestReceivedTimestamp = in.RequestReceivedTimestamp
	out.StageTimestamp = in.StageTimestamp
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_v1_Event_To_audit_Event is an autogenerated conversion function.
func Convert_v1_Event_To_audit_Event(in *Event, out *pkgapisaudit.Event, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_Event_To_audit_Event(in, out, s)
}

func autoConvert_audit_Event_To_v1_Event(in *pkgapisaudit.Event, out *Event, s apimachinerypkgconversion.Scope) error {
	out.Level = Level(in.Level)
	out.AuditID = apimachinerypkgtypes.UID(in.AuditID)
	out.Stage = Stage(in.Stage)
	out.RequestURI = in.RequestURI
	out.Verb = in.Verb
	out.User = in.User
	out.ImpersonatedUser = (*apiauthenticationv1.UserInfo)(unsafe.Pointer(in.ImpersonatedUser))
	out.SourceIPs = *(*[]string)(unsafe.Pointer(&in.SourceIPs))
	out.UserAgent = in.UserAgent
	out.ObjectRef = (*ObjectReference)(unsafe.Pointer(in.ObjectRef))
	out.ResponseStatus = (*apismetav1.Status)(unsafe.Pointer(in.ResponseStatus))
	out.RequestObject = (*apimachinerypkgruntime.Unknown)(unsafe.Pointer(in.RequestObject))
	out.ResponseObject = (*apimachinerypkgruntime.Unknown)(unsafe.Pointer(in.ResponseObject))
	out.RequestReceivedTimestamp = in.RequestReceivedTimestamp
	out.StageTimestamp = in.StageTimestamp
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_audit_Event_To_v1_Event is an autogenerated conversion function.
func Convert_audit_Event_To_v1_Event(in *pkgapisaudit.Event, out *Event, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_Event_To_v1_Event(in, out, s)
}

func autoConvert_v1_EventList_To_audit_EventList(in *EventList, out *pkgapisaudit.EventList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]pkgapisaudit.Event)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_EventList_To_audit_EventList is an autogenerated conversion function.
func Convert_v1_EventList_To_audit_EventList(in *EventList, out *pkgapisaudit.EventList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_EventList_To_audit_EventList(in, out, s)
}

func autoConvert_audit_EventList_To_v1_EventList(in *pkgapisaudit.EventList, out *EventList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Event)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_audit_EventList_To_v1_EventList is an autogenerated conversion function.
func Convert_audit_EventList_To_v1_EventList(in *pkgapisaudit.EventList, out *EventList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_EventList_To_v1_EventList(in, out, s)
}

func autoConvert_v1_GroupResources_To_audit_GroupResources(in *GroupResources, out *pkgapisaudit.GroupResources, s apimachinerypkgconversion.Scope) error {
	out.Group = in.Group
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_v1_GroupResources_To_audit_GroupResources is an autogenerated conversion function.
func Convert_v1_GroupResources_To_audit_GroupResources(in *GroupResources, out *pkgapisaudit.GroupResources, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_GroupResources_To_audit_GroupResources(in, out, s)
}

func autoConvert_audit_GroupResources_To_v1_GroupResources(in *pkgapisaudit.GroupResources, out *GroupResources, s apimachinerypkgconversion.Scope) error {
	out.Group = in.Group
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_audit_GroupResources_To_v1_GroupResources is an autogenerated conversion function.
func Convert_audit_GroupResources_To_v1_GroupResources(in *pkgapisaudit.GroupResources, out *GroupResources, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_GroupResources_To_v1_GroupResources(in, out, s)
}

func autoConvert_v1_ObjectReference_To_audit_ObjectReference(in *ObjectReference, out *pkgapisaudit.ObjectReference, s apimachinerypkgconversion.Scope) error {
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = apimachinerypkgtypes.UID(in.UID)
	out.APIGroup = in.APIGroup
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.Subresource = in.Subresource
	return nil
}

// Convert_v1_ObjectReference_To_audit_ObjectReference is an autogenerated conversion function.
func Convert_v1_ObjectReference_To_audit_ObjectReference(in *ObjectReference, out *pkgapisaudit.ObjectReference, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_ObjectReference_To_audit_ObjectReference(in, out, s)
}

func autoConvert_audit_ObjectReference_To_v1_ObjectReference(in *pkgapisaudit.ObjectReference, out *ObjectReference, s apimachinerypkgconversion.Scope) error {
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = apimachinerypkgtypes.UID(in.UID)
	out.APIGroup = in.APIGroup
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.Subresource = in.Subresource
	return nil
}

// Convert_audit_ObjectReference_To_v1_ObjectReference is an autogenerated conversion function.
func Convert_audit_ObjectReference_To_v1_ObjectReference(in *pkgapisaudit.ObjectReference, out *ObjectReference, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_ObjectReference_To_v1_ObjectReference(in, out, s)
}

func autoConvert_v1_Policy_To_audit_Policy(in *Policy, out *pkgapisaudit.Policy, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]pkgapisaudit.PolicyRule)(unsafe.Pointer(&in.Rules))
	out.OmitStages = *(*[]pkgapisaudit.Stage)(unsafe.Pointer(&in.OmitStages))
	out.OmitManagedFields = in.OmitManagedFields
	return nil
}

// Convert_v1_Policy_To_audit_Policy is an autogenerated conversion function.
func Convert_v1_Policy_To_audit_Policy(in *Policy, out *pkgapisaudit.Policy, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_Policy_To_audit_Policy(in, out, s)
}

func autoConvert_audit_Policy_To_v1_Policy(in *pkgapisaudit.Policy, out *Policy, s apimachinerypkgconversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]PolicyRule)(unsafe.Pointer(&in.Rules))
	out.OmitStages = *(*[]Stage)(unsafe.Pointer(&in.OmitStages))
	out.OmitManagedFields = in.OmitManagedFields
	return nil
}

// Convert_audit_Policy_To_v1_Policy is an autogenerated conversion function.
func Convert_audit_Policy_To_v1_Policy(in *pkgapisaudit.Policy, out *Policy, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_Policy_To_v1_Policy(in, out, s)
}

func autoConvert_v1_PolicyList_To_audit_PolicyList(in *PolicyList, out *pkgapisaudit.PolicyList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]pkgapisaudit.Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_PolicyList_To_audit_PolicyList is an autogenerated conversion function.
func Convert_v1_PolicyList_To_audit_PolicyList(in *PolicyList, out *pkgapisaudit.PolicyList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_PolicyList_To_audit_PolicyList(in, out, s)
}

func autoConvert_audit_PolicyList_To_v1_PolicyList(in *pkgapisaudit.PolicyList, out *PolicyList, s apimachinerypkgconversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_audit_PolicyList_To_v1_PolicyList is an autogenerated conversion function.
func Convert_audit_PolicyList_To_v1_PolicyList(in *pkgapisaudit.PolicyList, out *PolicyList, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_PolicyList_To_v1_PolicyList(in, out, s)
}

func autoConvert_v1_PolicyRule_To_audit_PolicyRule(in *PolicyRule, out *pkgapisaudit.PolicyRule, s apimachinerypkgconversion.Scope) error {
	out.Level = pkgapisaudit.Level(in.Level)
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.UserGroups = *(*[]string)(unsafe.Pointer(&in.UserGroups))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.Resources = *(*[]pkgapisaudit.GroupResources)(unsafe.Pointer(&in.Resources))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	out.OmitStages = *(*[]pkgapisaudit.Stage)(unsafe.Pointer(&in.OmitStages))
	out.OmitManagedFields = (*bool)(unsafe.Pointer(in.OmitManagedFields))
	return nil
}

// Convert_v1_PolicyRule_To_audit_PolicyRule is an autogenerated conversion function.
func Convert_v1_PolicyRule_To_audit_PolicyRule(in *PolicyRule, out *pkgapisaudit.PolicyRule, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1_PolicyRule_To_audit_PolicyRule(in, out, s)
}

func autoConvert_audit_PolicyRule_To_v1_PolicyRule(in *pkgapisaudit.PolicyRule, out *PolicyRule, s apimachinerypkgconversion.Scope) error {
	out.Level = Level(in.Level)
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.UserGroups = *(*[]string)(unsafe.Pointer(&in.UserGroups))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.Resources = *(*[]GroupResources)(unsafe.Pointer(&in.Resources))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	out.OmitStages = *(*[]Stage)(unsafe.Pointer(&in.OmitStages))
	out.OmitManagedFields = (*bool)(unsafe.Pointer(in.OmitManagedFields))
	return nil
}

// Convert_audit_PolicyRule_To_v1_PolicyRule is an autogenerated conversion function.
func Convert_audit_PolicyRule_To_v1_PolicyRule(in *pkgapisaudit.PolicyRule, out *PolicyRule, s apimachinerypkgconversion.Scope) error {
	return autoConvert_audit_PolicyRule_To_v1_PolicyRule(in, out, s)
}
