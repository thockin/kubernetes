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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/api/discovery/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/discovery"
)

var (
	zoneA = "zoneA"
)

func TestEndpointTopologyConverstion(t *testing.T) {
	largeTopology := createLargeTopologyMaps(false)
	largeTopologyWithStandardKeys := createLargeTopologyMaps(true)

	testcases := []struct {
		desc             string
		external         v1beta1.EndpointSlice
		internal         discovery.EndpointSlice
		expectPruning    bool
		withStandardKeys bool
	}{
		{
			desc:     "no topology",
			external: v1beta1.EndpointSlice{},
			internal: discovery.EndpointSlice{},
		},
		{
			desc: "topology without zone",
			external: v1beta1.EndpointSlice{
				Endpoints: []v1beta1.Endpoint{
					{
						Topology: map[string]string{
							"key1": "val1",
							"key2": "val2",
							"key3": "val3",
						},
					},
					{
						Topology: map[string]string{
							"key1": "val4",
							"key2": "val5",
							"key3": "val6",
						},
					},
				},
			},
			internal: discovery.EndpointSlice{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						v1beta1Topology: `{"s":["key1","key2","key3","val1","val2","val3","val4","val5","val6"],"t":[{"0":3,"1":4,"2":5},{"0":6,"1":7,"2":8}]}`,
					},
				},
				Endpoints: []discovery.Endpoint{{}, {}},
			},
		},
		{
			desc: "topology with zone",
			external: v1beta1.EndpointSlice{
				Endpoints: []v1beta1.Endpoint{
					{
						Topology: map[string]string{
							"key1":                        "val1",
							"key2":                        "val2",
							"key3":                        "val3",
							"topology.kubernetes.io/zone": zoneA,
						},
					},
					{
						Topology: map[string]string{
							"key1":                        "val4",
							"key2":                        "val5",
							"key3":                        "val6",
							"topology.kubernetes.io/zone": zoneA,
						},
					},
				},
			},
			internal: discovery.EndpointSlice{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						v1beta1Topology: `{"s":["key1","key2","key3","val1","val2","val3","val4","val5","val6"],"t":[{"0":3,"1":4,"2":5},{"0":6,"1":7,"2":8}]}`,
					},
				},
				Endpoints: []discovery.Endpoint{
					{Zone: &zoneA},
					{Zone: &zoneA},
				},
			},
		},
		{
			desc: "topology with only zone",
			external: v1beta1.EndpointSlice{
				Endpoints: []v1beta1.Endpoint{
					{
						Topology: map[string]string{
							"topology.kubernetes.io/zone": zoneA,
						},
					},
					{
						Topology: map[string]string{
							"topology.kubernetes.io/zone": zoneA,
						},
					},
				},
			},
			internal: discovery.EndpointSlice{
				Endpoints: []discovery.Endpoint{
					{Zone: &zoneA},
					{Zone: &zoneA},
				},
			},
		},
		{
			desc:          "topology maps are too large for an annotation",
			external:      largeTopology,
			expectPruning: true,
		},
		{
			desc:             "topology maps are too large for an annotation with standardKeys",
			external:         largeTopologyWithStandardKeys,
			withStandardKeys: true,
			expectPruning:    true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			convertedInternal := discovery.EndpointSlice{}
			require.NoError(t, Convert_v1beta1_EndpointSlice_To_discovery_EndpointSlice(&tc.external, &convertedInternal, nil))
			assert.GreaterOrEqual(t, 256*(1<<10), sizeOfAnnotations(convertedInternal.Annotations))

			if !tc.expectPruning {
				assert.Equal(t, tc.internal.Endpoints, convertedInternal.Endpoints, "v1beta1.EndpointSlice -> discovery.EndpointSlice")

				convertedV1beta1 := v1beta1.EndpointSlice{}
				require.NoError(t, Convert_discovery_EndpointSlice_To_v1beta1_EndpointSlice(&tc.internal, &convertedV1beta1, nil))
				assert.Equal(t, tc.external, convertedV1beta1, "discovery.EndpointSlice -> v1beta1.EndpointSlice")
			} else {
				if tc.withStandardKeys {
					var condensedTop EndpointSliceTopology
					assert.NoError(t, json.Unmarshal([]byte(convertedInternal.Annotations[v1beta1Topology]), &condensedTop))
					assert.Contains(t, condensedTop.Strings, "kubernetes.io/hostname")
					assert.Contains(t, condensedTop.Strings, "topology.kubernetes.io/region")
				}
			}
		})
	}
}

func TestEndpointZoneConverstion(t *testing.T) {
	testcases := []struct {
		desc             string
		external         v1beta1.Endpoint
		expectedExternal v1beta1.Endpoint
		internal         discovery.Endpoint
	}{
		{
			desc:             "no topology field",
			external:         v1beta1.Endpoint{},
			expectedExternal: v1beta1.Endpoint{},
			internal:         discovery.Endpoint{},
		},
		{
			desc: "non empty topology map, but no zone",
			external: v1beta1.Endpoint{
				Topology: map[string]string{
					"key1": "val1",
				},
			},
			expectedExternal: v1beta1.Endpoint{},
			internal:         discovery.Endpoint{},
		},
		{
			desc: "non empty topology map, with zone",
			external: v1beta1.Endpoint{
				Topology: map[string]string{
					"key1":                        "val1",
					"topology.kubernetes.io/zone": zoneA,
				},
			},
			expectedExternal: v1beta1.Endpoint{
				Topology: map[string]string{
					"topology.kubernetes.io/zone": zoneA,
				},
			},
			internal: discovery.Endpoint{Zone: &zoneA},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			convertedInternal := discovery.Endpoint{}
			require.NoError(t, Convert_v1beta1_Endpoint_To_discovery_Endpoint(&tc.external, &convertedInternal, nil))
			assert.Equal(t, tc.internal, convertedInternal, "v1beta1.Endpoint -> discovery.Endpoint")

			convertedV1beta1 := v1beta1.Endpoint{}
			require.NoError(t, Convert_discovery_Endpoint_To_v1beta1_Endpoint(&tc.internal, &convertedV1beta1, nil))
			assert.Equal(t, tc.expectedExternal, convertedV1beta1, "discovery.Endpoint -> v1beta1.Endpoint")
		})
	}
}

func createLargeTopologyMaps(withStandardKeys bool) v1beta1.EndpointSlice {
	longString := "abcdefghijklmnopqrstuvwxyz"
	var keys []string
	numKeys := 16
	if withStandardKeys {
		keys = []string{"kubernetes.io/hostname", "topology.kubernetes.io/region"}
		numKeys = 13
	}
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("k%02d-%s", i, longString)
		keys = append(keys, key)
	}

	var endpoints []v1beta1.Endpoint
	for i := 0; i < 1000; i++ {
		topologyMap := make(map[string]string)
		for j, key := range keys {
			val := fmt.Sprintf("ep%03d-v%02d-%s", i, j, longString)
			topologyMap[key] = val

		}
		endpoint := v1beta1.Endpoint{Topology: topologyMap}
		endpoints = append(endpoints, endpoint)
	}

	return v1beta1.EndpointSlice{Endpoints: endpoints}
}
