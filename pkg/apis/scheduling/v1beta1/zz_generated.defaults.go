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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	apischedulingv1beta1 "k8s.io/api/scheduling/v1beta1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *apimachinerypkgruntime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&apischedulingv1beta1.PriorityClass{}, func(obj interface{}) { SetObjectDefaults_PriorityClass(obj.(*apischedulingv1beta1.PriorityClass)) })
	scheme.AddTypeDefaultingFunc(&apischedulingv1beta1.PriorityClassList{}, func(obj interface{}) {
		SetObjectDefaults_PriorityClassList(obj.(*apischedulingv1beta1.PriorityClassList))
	})
	return nil
}

func SetObjectDefaults_PriorityClass(in *apischedulingv1beta1.PriorityClass) {
	SetDefaults_PriorityClass(in)
}

func SetObjectDefaults_PriorityClassList(in *apischedulingv1beta1.PriorityClassList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PriorityClass(a)
	}
}
