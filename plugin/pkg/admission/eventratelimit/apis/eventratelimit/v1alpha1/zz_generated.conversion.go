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

package v1alpha1

import (
	"unsafe"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*eventratelimit.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_eventratelimit_Configuration(a.(*Configuration), b.(*eventratelimit.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*eventratelimit.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_eventratelimit_Configuration_To_v1alpha1_Configuration(a.(*eventratelimit.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Limit)(nil), (*eventratelimit.Limit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Limit_To_eventratelimit_Limit(a.(*Limit), b.(*eventratelimit.Limit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*eventratelimit.Limit)(nil), (*Limit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_eventratelimit_Limit_To_v1alpha1_Limit(a.(*eventratelimit.Limit), b.(*Limit), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Configuration_To_eventratelimit_Configuration(in *Configuration, out *eventratelimit.Configuration, s conversion.Scope) error {
	out.Limits = *(*[]eventratelimit.Limit)(unsafe.Pointer(&in.Limits))
	return nil
}

// Convert_v1alpha1_Configuration_To_eventratelimit_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_eventratelimit_Configuration(in *Configuration, out *eventratelimit.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_eventratelimit_Configuration(in, out, s)
}

func autoConvert_eventratelimit_Configuration_To_v1alpha1_Configuration(in *eventratelimit.Configuration, out *Configuration, s conversion.Scope) error {
	out.Limits = *(*[]Limit)(unsafe.Pointer(&in.Limits))
	return nil
}

// Convert_eventratelimit_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_eventratelimit_Configuration_To_v1alpha1_Configuration(in *eventratelimit.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_eventratelimit_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_Limit_To_eventratelimit_Limit(in *Limit, out *eventratelimit.Limit, s conversion.Scope) error {
	out.Type = eventratelimit.LimitType(in.Type)
	out.QPS = in.QPS
	out.Burst = in.Burst
	out.CacheSize = in.CacheSize
	return nil
}

// Convert_v1alpha1_Limit_To_eventratelimit_Limit is an autogenerated conversion function.
func Convert_v1alpha1_Limit_To_eventratelimit_Limit(in *Limit, out *eventratelimit.Limit, s conversion.Scope) error {
	return autoConvert_v1alpha1_Limit_To_eventratelimit_Limit(in, out, s)
}

func autoConvert_eventratelimit_Limit_To_v1alpha1_Limit(in *eventratelimit.Limit, out *Limit, s conversion.Scope) error {
	out.Type = LimitType(in.Type)
	out.QPS = in.QPS
	out.Burst = in.Burst
	out.CacheSize = in.CacheSize
	return nil
}

// Convert_eventratelimit_Limit_To_v1alpha1_Limit is an autogenerated conversion function.
func Convert_eventratelimit_Limit_To_v1alpha1_Limit(in *eventratelimit.Limit, out *Limit, s conversion.Scope) error {
	return autoConvert_eventratelimit_Limit_To_v1alpha1_Limit(in, out, s)
}
