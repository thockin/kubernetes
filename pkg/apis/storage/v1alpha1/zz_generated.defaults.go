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

package v1alpha1

import (
	"k8s.io/api/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/core/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1alpha1.VolumeAttachment{}, func(obj interface{}) { SetObjectDefaults_VolumeAttachment(obj.(*v1alpha1.VolumeAttachment)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.VolumeAttachmentList{}, func(obj interface{}) { SetObjectDefaults_VolumeAttachmentList(obj.(*v1alpha1.VolumeAttachmentList)) })
	return nil
}

func SetObjectDefaults_VolumeAttachment(in *v1alpha1.VolumeAttachment) {
	if in.Spec.Source.InlineVolumeSpec != nil {
		v1.SetDefaults_ResourceList(&in.Spec.Source.InlineVolumeSpec.Capacity)
		if in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.HostPath != nil {
			v1.SetDefaults_HostPathVolumeSource(in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.HostPath)
		}
		if in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.RBD != nil {
			v1.SetDefaults_RBDPersistentVolumeSource(in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.RBD)
		}
		if in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.ISCSI != nil {
			v1.SetDefaults_ISCSIPersistentVolumeSource(in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.ISCSI)
		}
		if in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.AzureDisk != nil {
			v1.SetDefaults_AzureDiskVolumeSource(in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.AzureDisk)
		}
		if in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.ScaleIO != nil {
			v1.SetDefaults_ScaleIOPersistentVolumeSource(in.Spec.Source.InlineVolumeSpec.PersistentVolumeSource.ScaleIO)
		}
	}
}

func SetObjectDefaults_VolumeAttachmentList(in *v1alpha1.VolumeAttachmentList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VolumeAttachment(a)
	}
}
