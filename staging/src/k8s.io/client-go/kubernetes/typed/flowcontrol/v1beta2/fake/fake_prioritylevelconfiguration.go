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

	apiflowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsflowcontrolv1beta2 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta2"
	clientgotesting "k8s.io/client-go/testing"
)

// FakePriorityLevelConfigurations implements PriorityLevelConfigurationInterface
type FakePriorityLevelConfigurations struct {
	Fake *FakeFlowcontrolV1beta2
}

var prioritylevelconfigurationsResource = pkgruntimeschema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "prioritylevelconfigurations"}

var prioritylevelconfigurationsKind = pkgruntimeschema.GroupVersionKind{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Kind: "PriorityLevelConfiguration"}

// Get takes name of the priorityLevelConfiguration, and returns the corresponding priorityLevelConfiguration object, and an error if there is any.
func (c *FakePriorityLevelConfigurations) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiflowcontrolv1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootGetAction(prioritylevelconfigurationsResource, name), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors.
func (c *FakePriorityLevelConfigurations) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiflowcontrolv1beta2.PriorityLevelConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootListAction(prioritylevelconfigurationsResource, prioritylevelconfigurationsKind, opts), &apiflowcontrolv1beta2.PriorityLevelConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apiflowcontrolv1beta2.PriorityLevelConfigurationList{ListMeta: obj.(*apiflowcontrolv1beta2.PriorityLevelConfigurationList).ListMeta}
	for _, item := range obj.(*apiflowcontrolv1beta2.PriorityLevelConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested priorityLevelConfigurations.
func (c *FakePriorityLevelConfigurations) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewRootWatchAction(prioritylevelconfigurationsResource, opts))
}

// Create takes the representation of a priorityLevelConfiguration and creates it.  Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *FakePriorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1beta2.PriorityLevelConfiguration, opts apismetav1.CreateOptions) (result *apiflowcontrolv1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootCreateAction(prioritylevelconfigurationsResource, priorityLevelConfiguration), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}

// Update takes the representation of a priorityLevelConfiguration and updates it. Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *FakePriorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1beta2.PriorityLevelConfiguration, opts apismetav1.UpdateOptions) (result *apiflowcontrolv1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateAction(prioritylevelconfigurationsResource, priorityLevelConfiguration), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePriorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1beta2.PriorityLevelConfiguration, opts apismetav1.UpdateOptions) (*apiflowcontrolv1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateSubresourceAction(prioritylevelconfigurationsResource, "status", priorityLevelConfiguration), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}

// Delete takes name of the priorityLevelConfiguration and deletes it. Returns an error if one occurs.
func (c *FakePriorityLevelConfigurations) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewRootDeleteActionWithOptions(prioritylevelconfigurationsResource, name, opts), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePriorityLevelConfigurations) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewRootDeleteCollectionAction(prioritylevelconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &apiflowcontrolv1beta2.PriorityLevelConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched priorityLevelConfiguration.
func (c *FakePriorityLevelConfigurations) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiflowcontrolv1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, name, pt, data, subresources...), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityLevelConfiguration.
func (c *FakePriorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1beta2.PriorityLevelConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1beta2.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, *name, apimachinerypkgtypes.ApplyPatchType, data), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePriorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1beta2.PriorityLevelConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1beta2.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, *name, apimachinerypkgtypes.ApplyPatchType, data, "status"), &apiflowcontrolv1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiflowcontrolv1beta2.PriorityLevelConfiguration), err
}
