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

package v1

import (
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	kubeschedulerconfigv1 "k8s.io/kube-scheduler/config/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *apimachinerypkgruntime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.DefaultPreemptionArgs{}, func(obj interface{}) {
		SetObjectDefaults_DefaultPreemptionArgs(obj.(*kubeschedulerconfigv1.DefaultPreemptionArgs))
	})
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.InterPodAffinityArgs{}, func(obj interface{}) {
		SetObjectDefaults_InterPodAffinityArgs(obj.(*kubeschedulerconfigv1.InterPodAffinityArgs))
	})
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.KubeSchedulerConfiguration{}, func(obj interface{}) {
		SetObjectDefaults_KubeSchedulerConfiguration(obj.(*kubeschedulerconfigv1.KubeSchedulerConfiguration))
	})
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.NodeResourcesBalancedAllocationArgs{}, func(obj interface{}) {
		SetObjectDefaults_NodeResourcesBalancedAllocationArgs(obj.(*kubeschedulerconfigv1.NodeResourcesBalancedAllocationArgs))
	})
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.NodeResourcesFitArgs{}, func(obj interface{}) {
		SetObjectDefaults_NodeResourcesFitArgs(obj.(*kubeschedulerconfigv1.NodeResourcesFitArgs))
	})
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.PodTopologySpreadArgs{}, func(obj interface{}) {
		SetObjectDefaults_PodTopologySpreadArgs(obj.(*kubeschedulerconfigv1.PodTopologySpreadArgs))
	})
	scheme.AddTypeDefaultingFunc(&kubeschedulerconfigv1.VolumeBindingArgs{}, func(obj interface{}) {
		SetObjectDefaults_VolumeBindingArgs(obj.(*kubeschedulerconfigv1.VolumeBindingArgs))
	})
	return nil
}

func SetObjectDefaults_DefaultPreemptionArgs(in *kubeschedulerconfigv1.DefaultPreemptionArgs) {
	SetDefaults_DefaultPreemptionArgs(in)
}

func SetObjectDefaults_InterPodAffinityArgs(in *kubeschedulerconfigv1.InterPodAffinityArgs) {
	SetDefaults_InterPodAffinityArgs(in)
}

func SetObjectDefaults_KubeSchedulerConfiguration(in *kubeschedulerconfigv1.KubeSchedulerConfiguration) {
	SetDefaults_KubeSchedulerConfiguration(in)
}

func SetObjectDefaults_NodeResourcesBalancedAllocationArgs(in *kubeschedulerconfigv1.NodeResourcesBalancedAllocationArgs) {
	SetDefaults_NodeResourcesBalancedAllocationArgs(in)
}

func SetObjectDefaults_NodeResourcesFitArgs(in *kubeschedulerconfigv1.NodeResourcesFitArgs) {
	SetDefaults_NodeResourcesFitArgs(in)
}

func SetObjectDefaults_PodTopologySpreadArgs(in *kubeschedulerconfigv1.PodTopologySpreadArgs) {
	SetDefaults_PodTopologySpreadArgs(in)
}

func SetObjectDefaults_VolumeBindingArgs(in *kubeschedulerconfigv1.VolumeBindingArgs) {
	SetDefaults_VolumeBindingArgs(in)
}
