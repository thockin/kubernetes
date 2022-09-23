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

	apinetworkingv1beta1 "k8s.io/api/networking/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1beta1 "k8s.io/client-go/applyconfigurations/networking/v1beta1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeIngressClasses implements IngressClassInterface
type FakeIngressClasses struct {
	Fake *FakeNetworkingV1beta1
}

var ingressclassesResource = pkgruntimeschema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1beta1", Resource: "ingressclasses"}

var ingressclassesKind = pkgruntimeschema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1beta1", Kind: "IngressClass"}

// Get takes name of the ingressClass, and returns the corresponding ingressClass object, and an error if there is any.
func (c *FakeIngressClasses) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apinetworkingv1beta1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootGetAction(ingressclassesResource, name), &apinetworkingv1beta1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1beta1.IngressClass), err
}

// List takes label and field selectors, and returns the list of IngressClasses that match those selectors.
func (c *FakeIngressClasses) List(ctx context.Context, opts apismetav1.ListOptions) (result *apinetworkingv1beta1.IngressClassList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootListAction(ingressclassesResource, ingressclassesKind, opts), &apinetworkingv1beta1.IngressClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apinetworkingv1beta1.IngressClassList{ListMeta: obj.(*apinetworkingv1beta1.IngressClassList).ListMeta}
	for _, item := range obj.(*apinetworkingv1beta1.IngressClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested ingressClasses.
func (c *FakeIngressClasses) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewRootWatchAction(ingressclassesResource, opts))
}

// Create takes the representation of a ingressClass and creates it.  Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *FakeIngressClasses) Create(ctx context.Context, ingressClass *apinetworkingv1beta1.IngressClass, opts apismetav1.CreateOptions) (result *apinetworkingv1beta1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootCreateAction(ingressclassesResource, ingressClass), &apinetworkingv1beta1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1beta1.IngressClass), err
}

// Update takes the representation of a ingressClass and updates it. Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *FakeIngressClasses) Update(ctx context.Context, ingressClass *apinetworkingv1beta1.IngressClass, opts apismetav1.UpdateOptions) (result *apinetworkingv1beta1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateAction(ingressclassesResource, ingressClass), &apinetworkingv1beta1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1beta1.IngressClass), err
}

// Delete takes name of the ingressClass and deletes it. Returns an error if one occurs.
func (c *FakeIngressClasses) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewRootDeleteActionWithOptions(ingressclassesResource, name, opts), &apinetworkingv1beta1.IngressClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressClasses) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewRootDeleteCollectionAction(ingressclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &apinetworkingv1beta1.IngressClassList{})
	return err
}

// Patch applies the patch and returns the patched ingressClass.
func (c *FakeIngressClasses) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apinetworkingv1beta1.IngressClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(ingressclassesResource, name, pt, data, subresources...), &apinetworkingv1beta1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1beta1.IngressClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied ingressClass.
func (c *FakeIngressClasses) Apply(ctx context.Context, ingressClass *applyconfigurationsnetworkingv1beta1.IngressClassApplyConfiguration, opts apismetav1.ApplyOptions) (result *apinetworkingv1beta1.IngressClass, err error) {
	if ingressClass == nil {
		return nil, fmt.Errorf("ingressClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(ingressClass)
	if err != nil {
		return nil, err
	}
	name := ingressClass.Name
	if name == nil {
		return nil, fmt.Errorf("ingressClass.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(ingressclassesResource, *name, apimachinerypkgtypes.ApplyPatchType, data), &apinetworkingv1beta1.IngressClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apinetworkingv1beta1.IngressClass), err
}
