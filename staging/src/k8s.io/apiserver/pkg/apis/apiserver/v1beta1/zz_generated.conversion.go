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

package v1beta1

import (
	"unsafe"

	apimachinerypkgconversion "k8s.io/apimachinery/pkg/conversion"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	pkgapisapiserver "k8s.io/apiserver/pkg/apis/apiserver"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *apimachinerypkgruntime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Connection)(nil), (*pkgapisapiserver.Connection)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_Connection_To_apiserver_Connection(a.(*Connection), b.(*pkgapisapiserver.Connection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.Connection)(nil), (*Connection)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_Connection_To_v1beta1_Connection(a.(*pkgapisapiserver.Connection), b.(*Connection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.EgressSelection)(nil), (*EgressSelection)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_EgressSelection_To_v1beta1_EgressSelection(a.(*pkgapisapiserver.EgressSelection), b.(*EgressSelection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EgressSelectorConfiguration)(nil), (*pkgapisapiserver.EgressSelectorConfiguration)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(a.(*EgressSelectorConfiguration), b.(*pkgapisapiserver.EgressSelectorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.EgressSelectorConfiguration)(nil), (*EgressSelectorConfiguration)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_EgressSelectorConfiguration_To_v1beta1_EgressSelectorConfiguration(a.(*pkgapisapiserver.EgressSelectorConfiguration), b.(*EgressSelectorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TCPTransport)(nil), (*pkgapisapiserver.TCPTransport)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_TCPTransport_To_apiserver_TCPTransport(a.(*TCPTransport), b.(*pkgapisapiserver.TCPTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.TCPTransport)(nil), (*TCPTransport)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_TCPTransport_To_v1beta1_TCPTransport(a.(*pkgapisapiserver.TCPTransport), b.(*TCPTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TLSConfig)(nil), (*pkgapisapiserver.TLSConfig)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_TLSConfig_To_apiserver_TLSConfig(a.(*TLSConfig), b.(*pkgapisapiserver.TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.TLSConfig)(nil), (*TLSConfig)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_TLSConfig_To_v1beta1_TLSConfig(a.(*pkgapisapiserver.TLSConfig), b.(*TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Transport)(nil), (*pkgapisapiserver.Transport)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_Transport_To_apiserver_Transport(a.(*Transport), b.(*pkgapisapiserver.Transport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.Transport)(nil), (*Transport)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_Transport_To_v1beta1_Transport(a.(*pkgapisapiserver.Transport), b.(*Transport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UDSTransport)(nil), (*pkgapisapiserver.UDSTransport)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_UDSTransport_To_apiserver_UDSTransport(a.(*UDSTransport), b.(*pkgapisapiserver.UDSTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*pkgapisapiserver.UDSTransport)(nil), (*UDSTransport)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_apiserver_UDSTransport_To_v1beta1_UDSTransport(a.(*pkgapisapiserver.UDSTransport), b.(*UDSTransport), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*EgressSelection)(nil), (*pkgapisapiserver.EgressSelection)(nil), func(a, b interface{}, scope apimachinerypkgconversion.Scope) error {
		return Convert_v1beta1_EgressSelection_To_apiserver_EgressSelection(a.(*EgressSelection), b.(*pkgapisapiserver.EgressSelection), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Connection_To_apiserver_Connection(in *Connection, out *pkgapisapiserver.Connection, s apimachinerypkgconversion.Scope) error {
	out.ProxyProtocol = pkgapisapiserver.ProtocolType(in.ProxyProtocol)
	out.Transport = (*pkgapisapiserver.Transport)(unsafe.Pointer(in.Transport))
	return nil
}

// Convert_v1beta1_Connection_To_apiserver_Connection is an autogenerated conversion function.
func Convert_v1beta1_Connection_To_apiserver_Connection(in *Connection, out *pkgapisapiserver.Connection, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_Connection_To_apiserver_Connection(in, out, s)
}

func autoConvert_apiserver_Connection_To_v1beta1_Connection(in *pkgapisapiserver.Connection, out *Connection, s apimachinerypkgconversion.Scope) error {
	out.ProxyProtocol = ProtocolType(in.ProxyProtocol)
	out.Transport = (*Transport)(unsafe.Pointer(in.Transport))
	return nil
}

// Convert_apiserver_Connection_To_v1beta1_Connection is an autogenerated conversion function.
func Convert_apiserver_Connection_To_v1beta1_Connection(in *pkgapisapiserver.Connection, out *Connection, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_Connection_To_v1beta1_Connection(in, out, s)
}

func autoConvert_v1beta1_EgressSelection_To_apiserver_EgressSelection(in *EgressSelection, out *pkgapisapiserver.EgressSelection, s apimachinerypkgconversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1beta1_Connection_To_apiserver_Connection(&in.Connection, &out.Connection, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_apiserver_EgressSelection_To_v1beta1_EgressSelection(in *pkgapisapiserver.EgressSelection, out *EgressSelection, s apimachinerypkgconversion.Scope) error {
	out.Name = in.Name
	if err := Convert_apiserver_Connection_To_v1beta1_Connection(&in.Connection, &out.Connection, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiserver_EgressSelection_To_v1beta1_EgressSelection is an autogenerated conversion function.
func Convert_apiserver_EgressSelection_To_v1beta1_EgressSelection(in *pkgapisapiserver.EgressSelection, out *EgressSelection, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_EgressSelection_To_v1beta1_EgressSelection(in, out, s)
}

func autoConvert_v1beta1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(in *EgressSelectorConfiguration, out *pkgapisapiserver.EgressSelectorConfiguration, s apimachinerypkgconversion.Scope) error {
	if in.EgressSelections != nil {
		in, out := &in.EgressSelections, &out.EgressSelections
		*out = make([]pkgapisapiserver.EgressSelection, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_EgressSelection_To_apiserver_EgressSelection(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.EgressSelections = nil
	}
	return nil
}

// Convert_v1beta1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration is an autogenerated conversion function.
func Convert_v1beta1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(in *EgressSelectorConfiguration, out *pkgapisapiserver.EgressSelectorConfiguration, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_EgressSelectorConfiguration_To_apiserver_EgressSelectorConfiguration(in, out, s)
}

func autoConvert_apiserver_EgressSelectorConfiguration_To_v1beta1_EgressSelectorConfiguration(in *pkgapisapiserver.EgressSelectorConfiguration, out *EgressSelectorConfiguration, s apimachinerypkgconversion.Scope) error {
	if in.EgressSelections != nil {
		in, out := &in.EgressSelections, &out.EgressSelections
		*out = make([]EgressSelection, len(*in))
		for i := range *in {
			if err := Convert_apiserver_EgressSelection_To_v1beta1_EgressSelection(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.EgressSelections = nil
	}
	return nil
}

// Convert_apiserver_EgressSelectorConfiguration_To_v1beta1_EgressSelectorConfiguration is an autogenerated conversion function.
func Convert_apiserver_EgressSelectorConfiguration_To_v1beta1_EgressSelectorConfiguration(in *pkgapisapiserver.EgressSelectorConfiguration, out *EgressSelectorConfiguration, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_EgressSelectorConfiguration_To_v1beta1_EgressSelectorConfiguration(in, out, s)
}

func autoConvert_v1beta1_TCPTransport_To_apiserver_TCPTransport(in *TCPTransport, out *pkgapisapiserver.TCPTransport, s apimachinerypkgconversion.Scope) error {
	out.URL = in.URL
	out.TLSConfig = (*pkgapisapiserver.TLSConfig)(unsafe.Pointer(in.TLSConfig))
	return nil
}

// Convert_v1beta1_TCPTransport_To_apiserver_TCPTransport is an autogenerated conversion function.
func Convert_v1beta1_TCPTransport_To_apiserver_TCPTransport(in *TCPTransport, out *pkgapisapiserver.TCPTransport, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_TCPTransport_To_apiserver_TCPTransport(in, out, s)
}

func autoConvert_apiserver_TCPTransport_To_v1beta1_TCPTransport(in *pkgapisapiserver.TCPTransport, out *TCPTransport, s apimachinerypkgconversion.Scope) error {
	out.URL = in.URL
	out.TLSConfig = (*TLSConfig)(unsafe.Pointer(in.TLSConfig))
	return nil
}

// Convert_apiserver_TCPTransport_To_v1beta1_TCPTransport is an autogenerated conversion function.
func Convert_apiserver_TCPTransport_To_v1beta1_TCPTransport(in *pkgapisapiserver.TCPTransport, out *TCPTransport, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_TCPTransport_To_v1beta1_TCPTransport(in, out, s)
}

func autoConvert_v1beta1_TLSConfig_To_apiserver_TLSConfig(in *TLSConfig, out *pkgapisapiserver.TLSConfig, s apimachinerypkgconversion.Scope) error {
	out.CABundle = in.CABundle
	out.ClientKey = in.ClientKey
	out.ClientCert = in.ClientCert
	return nil
}

// Convert_v1beta1_TLSConfig_To_apiserver_TLSConfig is an autogenerated conversion function.
func Convert_v1beta1_TLSConfig_To_apiserver_TLSConfig(in *TLSConfig, out *pkgapisapiserver.TLSConfig, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_TLSConfig_To_apiserver_TLSConfig(in, out, s)
}

func autoConvert_apiserver_TLSConfig_To_v1beta1_TLSConfig(in *pkgapisapiserver.TLSConfig, out *TLSConfig, s apimachinerypkgconversion.Scope) error {
	out.CABundle = in.CABundle
	out.ClientKey = in.ClientKey
	out.ClientCert = in.ClientCert
	return nil
}

// Convert_apiserver_TLSConfig_To_v1beta1_TLSConfig is an autogenerated conversion function.
func Convert_apiserver_TLSConfig_To_v1beta1_TLSConfig(in *pkgapisapiserver.TLSConfig, out *TLSConfig, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_TLSConfig_To_v1beta1_TLSConfig(in, out, s)
}

func autoConvert_v1beta1_Transport_To_apiserver_Transport(in *Transport, out *pkgapisapiserver.Transport, s apimachinerypkgconversion.Scope) error {
	out.TCP = (*pkgapisapiserver.TCPTransport)(unsafe.Pointer(in.TCP))
	out.UDS = (*pkgapisapiserver.UDSTransport)(unsafe.Pointer(in.UDS))
	return nil
}

// Convert_v1beta1_Transport_To_apiserver_Transport is an autogenerated conversion function.
func Convert_v1beta1_Transport_To_apiserver_Transport(in *Transport, out *pkgapisapiserver.Transport, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_Transport_To_apiserver_Transport(in, out, s)
}

func autoConvert_apiserver_Transport_To_v1beta1_Transport(in *pkgapisapiserver.Transport, out *Transport, s apimachinerypkgconversion.Scope) error {
	out.TCP = (*TCPTransport)(unsafe.Pointer(in.TCP))
	out.UDS = (*UDSTransport)(unsafe.Pointer(in.UDS))
	return nil
}

// Convert_apiserver_Transport_To_v1beta1_Transport is an autogenerated conversion function.
func Convert_apiserver_Transport_To_v1beta1_Transport(in *pkgapisapiserver.Transport, out *Transport, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_Transport_To_v1beta1_Transport(in, out, s)
}

func autoConvert_v1beta1_UDSTransport_To_apiserver_UDSTransport(in *UDSTransport, out *pkgapisapiserver.UDSTransport, s apimachinerypkgconversion.Scope) error {
	out.UDSName = in.UDSName
	return nil
}

// Convert_v1beta1_UDSTransport_To_apiserver_UDSTransport is an autogenerated conversion function.
func Convert_v1beta1_UDSTransport_To_apiserver_UDSTransport(in *UDSTransport, out *pkgapisapiserver.UDSTransport, s apimachinerypkgconversion.Scope) error {
	return autoConvert_v1beta1_UDSTransport_To_apiserver_UDSTransport(in, out, s)
}

func autoConvert_apiserver_UDSTransport_To_v1beta1_UDSTransport(in *pkgapisapiserver.UDSTransport, out *UDSTransport, s apimachinerypkgconversion.Scope) error {
	out.UDSName = in.UDSName
	return nil
}

// Convert_apiserver_UDSTransport_To_v1beta1_UDSTransport is an autogenerated conversion function.
func Convert_apiserver_UDSTransport_To_v1beta1_UDSTransport(in *pkgapisapiserver.UDSTransport, out *UDSTransport, s apimachinerypkgconversion.Scope) error {
	return autoConvert_apiserver_UDSTransport_To_v1beta1_UDSTransport(in, out, s)
}
