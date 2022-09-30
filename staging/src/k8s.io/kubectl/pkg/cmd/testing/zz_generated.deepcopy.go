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

package testing

import (
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNamespacedType) DeepCopyInto(out *ExternalNamespacedType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNamespacedType.
func (in *ExternalNamespacedType) DeepCopy() *ExternalNamespacedType {
	if in == nil {
		return nil
	}
	out := new(ExternalNamespacedType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ExternalNamespacedType) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNamespacedType2) DeepCopyInto(out *ExternalNamespacedType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNamespacedType2.
func (in *ExternalNamespacedType2) DeepCopy() *ExternalNamespacedType2 {
	if in == nil {
		return nil
	}
	out := new(ExternalNamespacedType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ExternalNamespacedType2) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalType) DeepCopyInto(out *ExternalType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalType.
func (in *ExternalType) DeepCopy() *ExternalType {
	if in == nil {
		return nil
	}
	out := new(ExternalType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ExternalType) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalType2) DeepCopyInto(out *ExternalType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalType2.
func (in *ExternalType2) DeepCopy() *ExternalType2 {
	if in == nil {
		return nil
	}
	out := new(ExternalType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ExternalType2) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalNamespacedType) DeepCopyInto(out *InternalNamespacedType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalNamespacedType.
func (in *InternalNamespacedType) DeepCopy() *InternalNamespacedType {
	if in == nil {
		return nil
	}
	out := new(InternalNamespacedType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *InternalNamespacedType) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalType) DeepCopyInto(out *InternalType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalType.
func (in *InternalType) DeepCopy() *InternalType {
	if in == nil {
		return nil
	}
	out := new(InternalType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *InternalType) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
