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

package v1alpha1

import (
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditAnnotation) DeepCopyInto(out *AuditAnnotation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditAnnotation.
func (in *AuditAnnotation) DeepCopy() *AuditAnnotation {
	if in == nil {
		return nil
	}
	out := new(AuditAnnotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressionWarning) DeepCopyInto(out *ExpressionWarning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressionWarning.
func (in *ExpressionWarning) DeepCopy() *ExpressionWarning {
	if in == nil {
		return nil
	}
	out := new(ExpressionWarning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchCondition) DeepCopyInto(out *MatchCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchCondition.
func (in *MatchCondition) DeepCopy() *MatchCondition {
	if in == nil {
		return nil
	}
	out := new(MatchCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchResources) DeepCopyInto(out *MatchResources) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(apismetav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectSelector != nil {
		in, out := &in.ObjectSelector, &out.ObjectSelector
		*out = new(apismetav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceRules != nil {
		in, out := &in.ResourceRules, &out.ResourceRules
		*out = make([]NamedRuleWithOperations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExcludeResourceRules != nil {
		in, out := &in.ExcludeResourceRules, &out.ExcludeResourceRules
		*out = make([]NamedRuleWithOperations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPolicy != nil {
		in, out := &in.MatchPolicy, &out.MatchPolicy
		*out = new(MatchPolicyType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchResources.
func (in *MatchResources) DeepCopy() *MatchResources {
	if in == nil {
		return nil
	}
	out := new(MatchResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedRuleWithOperations) DeepCopyInto(out *NamedRuleWithOperations) {
	*out = *in
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.RuleWithOperations.DeepCopyInto(&out.RuleWithOperations)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedRuleWithOperations.
func (in *NamedRuleWithOperations) DeepCopy() *NamedRuleWithOperations {
	if in == nil {
		return nil
	}
	out := new(NamedRuleWithOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamKind) DeepCopyInto(out *ParamKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamKind.
func (in *ParamKind) DeepCopy() *ParamKind {
	if in == nil {
		return nil
	}
	out := new(ParamKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamRef) DeepCopyInto(out *ParamRef) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ParameterNotFoundAction != nil {
		in, out := &in.ParameterNotFoundAction, &out.ParameterNotFoundAction
		*out = new(ParameterNotFoundActionType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamRef.
func (in *ParamRef) DeepCopy() *ParamRef {
	if in == nil {
		return nil
	}
	out := new(ParamRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeChecking) DeepCopyInto(out *TypeChecking) {
	*out = *in
	if in.ExpressionWarnings != nil {
		in, out := &in.ExpressionWarnings, &out.ExpressionWarnings
		*out = make([]ExpressionWarning, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeChecking.
func (in *TypeChecking) DeepCopy() *TypeChecking {
	if in == nil {
		return nil
	}
	out := new(TypeChecking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicy) DeepCopyInto(out *ValidatingAdmissionPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicy.
func (in *ValidatingAdmissionPolicy) DeepCopy() *ValidatingAdmissionPolicy {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ValidatingAdmissionPolicy) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicyBinding) DeepCopyInto(out *ValidatingAdmissionPolicyBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicyBinding.
func (in *ValidatingAdmissionPolicyBinding) DeepCopy() *ValidatingAdmissionPolicyBinding {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicyBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ValidatingAdmissionPolicyBinding) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicyBindingList) DeepCopyInto(out *ValidatingAdmissionPolicyBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidatingAdmissionPolicyBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicyBindingList.
func (in *ValidatingAdmissionPolicyBindingList) DeepCopy() *ValidatingAdmissionPolicyBindingList {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicyBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ValidatingAdmissionPolicyBindingList) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicyBindingSpec) DeepCopyInto(out *ValidatingAdmissionPolicyBindingSpec) {
	*out = *in
	if in.ParamRef != nil {
		in, out := &in.ParamRef, &out.ParamRef
		*out = new(ParamRef)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchResources != nil {
		in, out := &in.MatchResources, &out.MatchResources
		*out = new(MatchResources)
		(*in).DeepCopyInto(*out)
	}
	if in.ValidationActions != nil {
		in, out := &in.ValidationActions, &out.ValidationActions
		*out = make([]ValidationAction, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicyBindingSpec.
func (in *ValidatingAdmissionPolicyBindingSpec) DeepCopy() *ValidatingAdmissionPolicyBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicyBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicyList) DeepCopyInto(out *ValidatingAdmissionPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidatingAdmissionPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicyList.
func (in *ValidatingAdmissionPolicyList) DeepCopy() *ValidatingAdmissionPolicyList {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *ValidatingAdmissionPolicyList) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicySpec) DeepCopyInto(out *ValidatingAdmissionPolicySpec) {
	*out = *in
	if in.ParamKind != nil {
		in, out := &in.ParamKind, &out.ParamKind
		*out = new(ParamKind)
		**out = **in
	}
	if in.MatchConstraints != nil {
		in, out := &in.MatchConstraints, &out.MatchConstraints
		*out = new(MatchResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Validations != nil {
		in, out := &in.Validations, &out.Validations
		*out = make([]Validation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(FailurePolicyType)
		**out = **in
	}
	if in.AuditAnnotations != nil {
		in, out := &in.AuditAnnotations, &out.AuditAnnotations
		*out = make([]AuditAnnotation, len(*in))
		copy(*out, *in)
	}
	if in.MatchConditions != nil {
		in, out := &in.MatchConditions, &out.MatchConditions
		*out = make([]MatchCondition, len(*in))
		copy(*out, *in)
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make([]Variable, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicySpec.
func (in *ValidatingAdmissionPolicySpec) DeepCopy() *ValidatingAdmissionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingAdmissionPolicyStatus) DeepCopyInto(out *ValidatingAdmissionPolicyStatus) {
	*out = *in
	if in.TypeChecking != nil {
		in, out := &in.TypeChecking, &out.TypeChecking
		*out = new(TypeChecking)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingAdmissionPolicyStatus.
func (in *ValidatingAdmissionPolicyStatus) DeepCopy() *ValidatingAdmissionPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ValidatingAdmissionPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Validation) DeepCopyInto(out *Validation) {
	*out = *in
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(apismetav1.StatusReason)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Validation.
func (in *Validation) DeepCopy() *Validation {
	if in == nil {
		return nil
	}
	out := new(Validation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}
