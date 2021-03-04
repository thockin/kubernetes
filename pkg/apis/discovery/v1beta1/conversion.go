/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import (
	"k8s.io/api/discovery/v1beta1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/kubernetes/pkg/apis/discovery"
)

const (
	//FIXME: core/v1.LabelTopologyZone exists
	labelTopologyZone = "topology.kubernetes.io/zone"
)

func Convert_v1beta1_Endpoint_To_discovery_Endpoint(in *v1beta1.Endpoint, out *discovery.Endpoint, s conversion.Scope) error {
	if err := autoConvert_v1beta1_Endpoint_To_discovery_Endpoint(in, out, s); err != nil {
		return err
	}

	// Copy topology so we can modify it.
	out.DeprecatedTopology = make(map[string]string, len(in.Topology))
	for k, v := range in.Topology {
		out.DeprecatedTopology[k] = v
	}

	// Move zone from the topology map into a field.
	if zone, ok := in.Topology[labelTopologyZone]; ok {
		out.Zone = &zone
		delete(out.DeprecatedTopology, labelTopologyZone)
	}

	return nil
}

func Convert_discovery_Endpoint_To_v1beta1_Endpoint(in *discovery.Endpoint, out *v1beta1.Endpoint, s conversion.Scope) error {
	if err := autoConvert_discovery_Endpoint_To_v1beta1_Endpoint(in, out, s); err != nil {
		return err
	}

	// Copy topology so we can modify it.
	out.Topology = make(map[string]string, len(in.DeprecatedTopology))
	for k, v := range in.DeprecatedTopology {
		out.Topology[k] = v
	}

	// Move zone from a field into the topology map.
	if in.Zone != nil {
		out.Topology[labelTopologyZone] = *in.Zone
	}

	return nil
}
