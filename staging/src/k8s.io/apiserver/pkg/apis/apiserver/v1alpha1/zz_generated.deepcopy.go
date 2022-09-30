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
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionConfiguration) DeepCopyInto(out *AdmissionConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]AdmissionPluginConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionConfiguration.
func (in *AdmissionConfiguration) DeepCopy() *AdmissionConfiguration {
	if in == nil {
		return nil
	}
	out := new(AdmissionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *AdmissionConfiguration) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionPluginConfiguration) DeepCopyInto(out *AdmissionPluginConfiguration) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(apimachinerypkgruntime.Unknown)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionPluginConfiguration.
func (in *AdmissionPluginConfiguration) DeepCopy() *AdmissionPluginConfiguration {
	if in == nil {
		return nil
	}
	out := new(AdmissionPluginConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationConfiguration) DeepCopyInto(out *AuthenticationConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.JWT != nil {
		in, out := &in.JWT, &out.JWT
		*out = make([]JWTAuthenticator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationConfiguration.
func (in *AuthenticationConfiguration) DeepCopy() *AuthenticationConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuthenticationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthenticationConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationConfiguration) DeepCopyInto(out *AuthorizationConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Authorizers != nil {
		in, out := &in.Authorizers, &out.Authorizers
		*out = make([]AuthorizerConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationConfiguration.
func (in *AuthorizationConfiguration) DeepCopy() *AuthorizationConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuthorizationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthorizationConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizerConfiguration) DeepCopyInto(out *AuthorizerConfiguration) {
	*out = *in
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizerConfiguration.
func (in *AuthorizerConfiguration) DeepCopy() *AuthorizerConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuthorizerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimMappings) DeepCopyInto(out *ClaimMappings) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Groups.DeepCopyInto(&out.Groups)
	out.UID = in.UID
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make([]ExtraMapping, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimMappings.
func (in *ClaimMappings) DeepCopy() *ClaimMappings {
	if in == nil {
		return nil
	}
	out := new(ClaimMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimOrExpression) DeepCopyInto(out *ClaimOrExpression) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimOrExpression.
func (in *ClaimOrExpression) DeepCopy() *ClaimOrExpression {
	if in == nil {
		return nil
	}
	out := new(ClaimOrExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimValidationRule) DeepCopyInto(out *ClaimValidationRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimValidationRule.
func (in *ClaimValidationRule) DeepCopy() *ClaimValidationRule {
	if in == nil {
		return nil
	}
	out := new(ClaimValidationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(Transport)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressSelection) DeepCopyInto(out *EgressSelection) {
	*out = *in
	in.Connection.DeepCopyInto(&out.Connection)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressSelection.
func (in *EgressSelection) DeepCopy() *EgressSelection {
	if in == nil {
		return nil
	}
	out := new(EgressSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressSelectorConfiguration) DeepCopyInto(out *EgressSelectorConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.EgressSelections != nil {
		in, out := &in.EgressSelections, &out.EgressSelections
		*out = make([]EgressSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressSelectorConfiguration.
func (in *EgressSelectorConfiguration) DeepCopy() *EgressSelectorConfiguration {
	if in == nil {
		return nil
	}
	out := new(EgressSelectorConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *EgressSelectorConfiguration) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraMapping) DeepCopyInto(out *ExtraMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraMapping.
func (in *ExtraMapping) DeepCopy() *ExtraMapping {
	if in == nil {
		return nil
	}
	out := new(ExtraMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issuer) DeepCopyInto(out *Issuer) {
	*out = *in
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issuer.
func (in *Issuer) DeepCopy() *Issuer {
	if in == nil {
		return nil
	}
	out := new(Issuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTAuthenticator) DeepCopyInto(out *JWTAuthenticator) {
	*out = *in
	in.Issuer.DeepCopyInto(&out.Issuer)
	if in.ClaimValidationRules != nil {
		in, out := &in.ClaimValidationRules, &out.ClaimValidationRules
		*out = make([]ClaimValidationRule, len(*in))
		copy(*out, *in)
	}
	in.ClaimMappings.DeepCopyInto(&out.ClaimMappings)
	if in.UserValidationRules != nil {
		in, out := &in.UserValidationRules, &out.UserValidationRules
		*out = make([]UserValidationRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTAuthenticator.
func (in *JWTAuthenticator) DeepCopy() *JWTAuthenticator {
	if in == nil {
		return nil
	}
	out := new(JWTAuthenticator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixedClaimOrExpression) DeepCopyInto(out *PrefixedClaimOrExpression) {
	*out = *in
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixedClaimOrExpression.
func (in *PrefixedClaimOrExpression) DeepCopy() *PrefixedClaimOrExpression {
	if in == nil {
		return nil
	}
	out := new(PrefixedClaimOrExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPTransport) DeepCopyInto(out *TCPTransport) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPTransport.
func (in *TCPTransport) DeepCopy() *TCPTransport {
	if in == nil {
		return nil
	}
	out := new(TCPTransport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfiguration) DeepCopyInto(out *TracingConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.TracingConfiguration.DeepCopyInto(&out.TracingConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfiguration.
func (in *TracingConfiguration) DeepCopy() *TracingConfiguration {
	if in == nil {
		return nil
	}
	out := new(TracingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *TracingConfiguration) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transport) DeepCopyInto(out *Transport) {
	*out = *in
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(TCPTransport)
		(*in).DeepCopyInto(*out)
	}
	if in.UDS != nil {
		in, out := &in.UDS, &out.UDS
		*out = new(UDSTransport)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transport.
func (in *Transport) DeepCopy() *Transport {
	if in == nil {
		return nil
	}
	out := new(Transport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDSTransport) DeepCopyInto(out *UDSTransport) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDSTransport.
func (in *UDSTransport) DeepCopy() *UDSTransport {
	if in == nil {
		return nil
	}
	out := new(UDSTransport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserValidationRule) DeepCopyInto(out *UserValidationRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserValidationRule.
func (in *UserValidationRule) DeepCopy() *UserValidationRule {
	if in == nil {
		return nil
	}
	out := new(UserValidationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfiguration) DeepCopyInto(out *WebhookConfiguration) {
	*out = *in
	out.AuthorizedTTL = in.AuthorizedTTL
	out.UnauthorizedTTL = in.UnauthorizedTTL
	out.Timeout = in.Timeout
	in.ConnectionInfo.DeepCopyInto(&out.ConnectionInfo)
	if in.MatchConditions != nil {
		in, out := &in.MatchConditions, &out.MatchConditions
		*out = make([]WebhookMatchCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfiguration.
func (in *WebhookConfiguration) DeepCopy() *WebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConnectionInfo) DeepCopyInto(out *WebhookConnectionInfo) {
	*out = *in
	if in.KubeConfigFile != nil {
		in, out := &in.KubeConfigFile, &out.KubeConfigFile
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConnectionInfo.
func (in *WebhookConnectionInfo) DeepCopy() *WebhookConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(WebhookConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookMatchCondition) DeepCopyInto(out *WebhookMatchCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookMatchCondition.
func (in *WebhookMatchCondition) DeepCopy() *WebhookMatchCondition {
	if in == nil {
		return nil
	}
	out := new(WebhookMatchCondition)
	in.DeepCopyInto(out)
	return out
}
