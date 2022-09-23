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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	apinetworkingv1 "k8s.io/api/networking/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1 "k8s.io/client-go/applyconfigurations/networking/v1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeNetworkPolicies implements NetworkPolicyInterface
type FakeNetworkPolicies struct {
	Fake *FakeNetworkingV1
	ns   string
}

var networkpoliciesResource = pkgruntimeschema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "networkpolicies"}

var networkpoliciesKind = pkgruntimeschema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1", Kind: "NetworkPolicy"}

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *FakeNetworkPolicies) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apinetworkingv1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewGetAction(networkpoliciesResource, c.ns, name), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *FakeNetworkPolicies) List(ctx context.Context, opts apismetav1.ListOptions) (result *apinetworkingv1.NetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewListAction(networkpoliciesResource, networkpoliciesKind, c.ns, opts), &apinetworkingv1.NetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apinetworkingv1.NetworkPolicyList{ListMeta: obj.(*apinetworkingv1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*apinetworkingv1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested networkPolicies.
func (c *FakeNetworkPolicies) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewWatchAction(networkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkPolicy and creates it.  Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Create(ctx context.Context, networkPolicy *apinetworkingv1.NetworkPolicy, opts apismetav1.CreateOptions) (result *apinetworkingv1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewCreateAction(networkpoliciesResource, c.ns, networkPolicy), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}

// Update takes the representation of a networkPolicy and updates it. Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Update(ctx context.Context, networkPolicy *apinetworkingv1.NetworkPolicy, opts apismetav1.UpdateOptions) (result *apinetworkingv1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateAction(networkpoliciesResource, c.ns, networkPolicy), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPolicies) UpdateStatus(ctx context.Context, networkPolicy *apinetworkingv1.NetworkPolicy, opts apismetav1.UpdateOptions) (*apinetworkingv1.NetworkPolicy, error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateSubresourceAction(networkpoliciesResource, "status", c.ns, networkPolicy), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}

// Delete takes name of the networkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPolicies) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewDeleteActionWithOptions(networkpoliciesResource, c.ns, name, opts), &apinetworkingv1.NetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPolicies) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewDeleteCollectionAction(networkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apinetworkingv1.NetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkPolicy.
func (c *FakeNetworkPolicies) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apinetworkingv1.NetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(networkpoliciesResource, c.ns, name, pt, data, subresources...), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied networkPolicy.
func (c *FakeNetworkPolicies) Apply(ctx context.Context, networkPolicy *applyconfigurationsnetworkingv1.NetworkPolicyApplyConfiguration, opts apismetav1.ApplyOptions) (result *apinetworkingv1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}
	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(networkpoliciesResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeNetworkPolicies) ApplyStatus(ctx context.Context, networkPolicy *applyconfigurationsnetworkingv1.NetworkPolicyApplyConfiguration, opts apismetav1.ApplyOptions) (result *apinetworkingv1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}
	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(networkpoliciesResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data, "status"), &apinetworkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1.NetworkPolicy), err
}
