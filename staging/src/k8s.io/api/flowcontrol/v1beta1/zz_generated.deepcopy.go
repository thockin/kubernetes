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

package v1beta1

import (
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowDistinguisherMethod) DeepCopyInto(out *FlowDistinguisherMethod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowDistinguisherMethod.
func (in *FlowDistinguisherMethod) DeepCopy() *FlowDistinguisherMethod {
	if in == nil {
		return nil
	}
	out := new(FlowDistinguisherMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchema) DeepCopyInto(out *FlowSchema) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchema.
func (in *FlowSchema) DeepCopy() *FlowSchema {
	if in == nil {
		return nil
	}
	out := new(FlowSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *FlowSchema) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaCondition) DeepCopyInto(out *FlowSchemaCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaCondition.
func (in *FlowSchemaCondition) DeepCopy() *FlowSchemaCondition {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaList) DeepCopyInto(out *FlowSchemaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlowSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaList.
func (in *FlowSchemaList) DeepCopy() *FlowSchemaList {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *FlowSchemaList) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaSpec) DeepCopyInto(out *FlowSchemaSpec) {
	*out = *in
	out.PriorityLevelConfiguration = in.PriorityLevelConfiguration
	if in.DistinguisherMethod != nil {
		in, out := &in.DistinguisherMethod, &out.DistinguisherMethod
		*out = new(FlowDistinguisherMethod)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRulesWithSubjects, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaSpec.
func (in *FlowSchemaSpec) DeepCopy() *FlowSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaStatus) DeepCopyInto(out *FlowSchemaStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]FlowSchemaCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaStatus.
func (in *FlowSchemaStatus) DeepCopy() *FlowSchemaStatus {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSubject) DeepCopyInto(out *GroupSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSubject.
func (in *GroupSubject) DeepCopy() *GroupSubject {
	if in == nil {
		return nil
	}
	out := new(GroupSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitResponse) DeepCopyInto(out *LimitResponse) {
	*out = *in
	if in.Queuing != nil {
		in, out := &in.Queuing, &out.Queuing
		*out = new(QueuingConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitResponse.
func (in *LimitResponse) DeepCopy() *LimitResponse {
	if in == nil {
		return nil
	}
	out := new(LimitResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitedPriorityLevelConfiguration) DeepCopyInto(out *LimitedPriorityLevelConfiguration) {
	*out = *in
	in.LimitResponse.DeepCopyInto(&out.LimitResponse)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitedPriorityLevelConfiguration.
func (in *LimitedPriorityLevelConfiguration) DeepCopy() *LimitedPriorityLevelConfiguration {
	if in == nil {
		return nil
	}
	out := new(LimitedPriorityLevelConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonResourcePolicyRule) DeepCopyInto(out *NonResourcePolicyRule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NonResourceURLs != nil {
		in, out := &in.NonResourceURLs, &out.NonResourceURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonResourcePolicyRule.
func (in *NonResourcePolicyRule) DeepCopy() *NonResourcePolicyRule {
	if in == nil {
		return nil
	}
	out := new(NonResourcePolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRulesWithSubjects) DeepCopyInto(out *PolicyRulesWithSubjects) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceRules != nil {
		in, out := &in.ResourceRules, &out.ResourceRules
		*out = make([]ResourcePolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NonResourceRules != nil {
		in, out := &in.NonResourceRules, &out.NonResourceRules
		*out = make([]NonResourcePolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRulesWithSubjects.
func (in *PolicyRulesWithSubjects) DeepCopy() *PolicyRulesWithSubjects {
	if in == nil {
		return nil
	}
	out := new(PolicyRulesWithSubjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfiguration) DeepCopyInto(out *PriorityLevelConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfiguration.
func (in *PriorityLevelConfiguration) DeepCopy() *PriorityLevelConfiguration {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *PriorityLevelConfiguration) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationCondition) DeepCopyInto(out *PriorityLevelConfigurationCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationCondition.
func (in *PriorityLevelConfigurationCondition) DeepCopy() *PriorityLevelConfigurationCondition {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationList) DeepCopyInto(out *PriorityLevelConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PriorityLevelConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationList.
func (in *PriorityLevelConfigurationList) DeepCopy() *PriorityLevelConfigurationList {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *PriorityLevelConfigurationList) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationReference) DeepCopyInto(out *PriorityLevelConfigurationReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationReference.
func (in *PriorityLevelConfigurationReference) DeepCopy() *PriorityLevelConfigurationReference {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationSpec) DeepCopyInto(out *PriorityLevelConfigurationSpec) {
	*out = *in
	if in.Limited != nil {
		in, out := &in.Limited, &out.Limited
		*out = new(LimitedPriorityLevelConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationSpec.
func (in *PriorityLevelConfigurationSpec) DeepCopy() *PriorityLevelConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationStatus) DeepCopyInto(out *PriorityLevelConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PriorityLevelConfigurationCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationStatus.
func (in *PriorityLevelConfigurationStatus) DeepCopy() *PriorityLevelConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueuingConfiguration) DeepCopyInto(out *QueuingConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueuingConfiguration.
func (in *QueuingConfiguration) DeepCopy() *QueuingConfiguration {
	if in == nil {
		return nil
	}
	out := new(QueuingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyRule) DeepCopyInto(out *ResourcePolicyRule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyRule.
func (in *ResourcePolicyRule) DeepCopy() *ResourcePolicyRule {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSubject) DeepCopyInto(out *ServiceAccountSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSubject.
func (in *ServiceAccountSubject) DeepCopy() *ServiceAccountSubject {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(UserSubject)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(GroupSubject)
		**out = **in
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccountSubject)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSubject) DeepCopyInto(out *UserSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSubject.
func (in *UserSubject) DeepCopy() *UserSubject {
	if in == nil {
		return nil
	}
	out := new(UserSubject)
	in.DeepCopyInto(out)
	return out
}
