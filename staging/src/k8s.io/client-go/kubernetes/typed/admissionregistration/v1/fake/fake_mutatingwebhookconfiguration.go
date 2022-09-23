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

	apiadmissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsadmissionregistrationv1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeMutatingWebhookConfigurations implements MutatingWebhookConfigurationInterface
type FakeMutatingWebhookConfigurations struct {
	Fake *FakeAdmissionregistrationV1
}

var mutatingwebhookconfigurationsResource = pkgruntimeschema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "v1", Resource: "mutatingwebhookconfigurations"}

var mutatingwebhookconfigurationsKind = pkgruntimeschema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "v1", Kind: "MutatingWebhookConfiguration"}

// Get takes name of the mutatingWebhookConfiguration, and returns the corresponding mutatingWebhookConfiguration object, and an error if there is any.
func (c *FakeMutatingWebhookConfigurations) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiadmissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootGetAction(mutatingwebhookconfigurationsResource, name), &apiadmissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiadmissionregistrationv1.MutatingWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of MutatingWebhookConfigurations that match those selectors.
func (c *FakeMutatingWebhookConfigurations) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiadmissionregistrationv1.MutatingWebhookConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootListAction(mutatingwebhookconfigurationsResource, mutatingwebhookconfigurationsKind, opts), &apiadmissionregistrationv1.MutatingWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apiadmissionregistrationv1.MutatingWebhookConfigurationList{ListMeta: obj.(*apiadmissionregistrationv1.MutatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*apiadmissionregistrationv1.MutatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested mutatingWebhookConfigurations.
func (c *FakeMutatingWebhookConfigurations) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewRootWatchAction(mutatingwebhookconfigurationsResource, opts))
}

// Create takes the representation of a mutatingWebhookConfiguration and creates it.  Returns the server's representation of the mutatingWebhookConfiguration, and an error, if there is any.
func (c *FakeMutatingWebhookConfigurations) Create(ctx context.Context, mutatingWebhookConfiguration *apiadmissionregistrationv1.MutatingWebhookConfiguration, opts apismetav1.CreateOptions) (result *apiadmissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootCreateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &apiadmissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiadmissionregistrationv1.MutatingWebhookConfiguration), err
}

// Update takes the representation of a mutatingWebhookConfiguration and updates it. Returns the server's representation of the mutatingWebhookConfiguration, and an error, if there is any.
func (c *FakeMutatingWebhookConfigurations) Update(ctx context.Context, mutatingWebhookConfiguration *apiadmissionregistrationv1.MutatingWebhookConfiguration, opts apismetav1.UpdateOptions) (result *apiadmissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &apiadmissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiadmissionregistrationv1.MutatingWebhookConfiguration), err
}

// Delete takes name of the mutatingWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeMutatingWebhookConfigurations) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewRootDeleteActionWithOptions(mutatingwebhookconfigurationsResource, name, opts), &apiadmissionregistrationv1.MutatingWebhookConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMutatingWebhookConfigurations) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewRootDeleteCollectionAction(mutatingwebhookconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &apiadmissionregistrationv1.MutatingWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched mutatingWebhookConfiguration.
func (c *FakeMutatingWebhookConfigurations) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiadmissionregistrationv1.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, name, pt, data, subresources...), &apiadmissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiadmissionregistrationv1.MutatingWebhookConfiguration), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied mutatingWebhookConfiguration.
func (c *FakeMutatingWebhookConfigurations) Apply(ctx context.Context, mutatingWebhookConfiguration *applyconfigurationsadmissionregistrationv1.MutatingWebhookConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiadmissionregistrationv1.MutatingWebhookConfiguration, err error) {
	if mutatingWebhookConfiguration == nil {
		return nil, fmt.Errorf("mutatingWebhookConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(mutatingWebhookConfiguration)
	if err != nil {
		return nil, err
	}
	name := mutatingWebhookConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("mutatingWebhookConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, *name, apimachinerypkgtypes.ApplyPatchType, data), &apiadmissionregistrationv1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiadmissionregistrationv1.MutatingWebhookConfiguration), err
}
