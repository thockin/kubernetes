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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	pkgutilintstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eviction) DeepCopyInto(out *Eviction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.DeleteOptions != nil {
		in, out := &in.DeleteOptions, &out.DeleteOptions
		*out = new(apismetav1.DeleteOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Eviction.
func (in *Eviction) DeepCopy() *Eviction {
	if in == nil {
		return nil
	}
	out := new(Eviction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *Eviction) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudget) DeepCopyInto(out *PodDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudget.
func (in *PodDisruptionBudget) DeepCopy() *PodDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *PodDisruptionBudget) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetList) DeepCopyInto(out *PodDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetList.
func (in *PodDisruptionBudgetList) DeepCopy() *PodDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *PodDisruptionBudgetList) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(pkgutilintstr.IntOrString)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(apismetav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(pkgutilintstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetStatus) DeepCopyInto(out *PodDisruptionBudgetStatus) {
	*out = *in
	if in.DisruptedPods != nil {
		in, out := &in.DisruptedPods, &out.DisruptedPods
		*out = make(map[string]apismetav1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apismetav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetStatus.
func (in *PodDisruptionBudgetStatus) DeepCopy() *PodDisruptionBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetStatus)
	in.DeepCopyInto(out)
	return out
}
