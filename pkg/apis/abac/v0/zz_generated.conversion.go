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

package v0

import (
	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	pkgapisabac "k8s.io/kubernetes/pkg/apis/abac"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddConversionFunc((*Policy)(nil), (*pkgapisabac.Policy)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v0_Policy_To_abac_Policy(a.(*Policy), b.(*pkgapisabac.Policy), scope)
	}); err != nil {
		return err
	}
	return nil
}
