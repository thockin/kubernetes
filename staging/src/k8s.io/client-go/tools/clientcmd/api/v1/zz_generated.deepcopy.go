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
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthInfo) DeepCopyInto(out *AuthInfo) {
	*out = *in
	if in.ClientCertificateData != nil {
		in, out := &in.ClientCertificateData, &out.ClientCertificateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ClientKeyData != nil {
		in, out := &in.ClientKeyData, &out.ClientKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateGroups != nil {
		in, out := &in.ImpersonateGroups, &out.ImpersonateGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateUserExtra != nil {
		in, out := &in.ImpersonateUserExtra, &out.ImpersonateUserExtra
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.AuthProvider != nil {
		in, out := &in.AuthProvider, &out.AuthProvider
		*out = new(AuthProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(ExecConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthInfo.
func (in *AuthInfo) DeepCopy() *AuthInfo {
	if in == nil {
		return nil
	}
	out := new(AuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProviderConfig) DeepCopyInto(out *AuthProviderConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProviderConfig.
func (in *AuthProviderConfig) DeepCopy() *AuthProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AuthProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	in.Preferences.DeepCopyInto(&out.Preferences)
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]NamedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthInfos != nil {
		in, out := &in.AuthInfos, &out.AuthInfos
		*out = make([]NamedAuthInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Contexts != nil {
		in, out := &in.Contexts, &out.Contexts
		*out = make([]NamedContext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new apimachinerypkgruntime.Object.
func (in *Config) DeepCopyObject() apimachinerypkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecConfig) DeepCopyInto(out *ExecConfig) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]ExecEnvVar, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecConfig.
func (in *ExecConfig) DeepCopy() *ExecConfig {
	if in == nil {
		return nil
	}
	out := new(ExecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecEnvVar) DeepCopyInto(out *ExecEnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecEnvVar.
func (in *ExecEnvVar) DeepCopy() *ExecEnvVar {
	if in == nil {
		return nil
	}
	out := new(ExecEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedAuthInfo) DeepCopyInto(out *NamedAuthInfo) {
	*out = *in
	in.AuthInfo.DeepCopyInto(&out.AuthInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedAuthInfo.
func (in *NamedAuthInfo) DeepCopy() *NamedAuthInfo {
	if in == nil {
		return nil
	}
	out := new(NamedAuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedCluster) DeepCopyInto(out *NamedCluster) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedCluster.
func (in *NamedCluster) DeepCopy() *NamedCluster {
	if in == nil {
		return nil
	}
	out := new(NamedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedContext) DeepCopyInto(out *NamedContext) {
	*out = *in
	in.Context.DeepCopyInto(&out.Context)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedContext.
func (in *NamedContext) DeepCopy() *NamedContext {
	if in == nil {
		return nil
	}
	out := new(NamedContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedExtension) DeepCopyInto(out *NamedExtension) {
	*out = *in
	in.Extension.DeepCopyInto(&out.Extension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedExtension.
func (in *NamedExtension) DeepCopy() *NamedExtension {
	if in == nil {
		return nil
	}
	out := new(NamedExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preferences) DeepCopyInto(out *Preferences) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preferences.
func (in *Preferences) DeepCopy() *Preferences {
	if in == nil {
		return nil
	}
	out := new(Preferences)
	in.DeepCopyInto(out)
	return out
}
