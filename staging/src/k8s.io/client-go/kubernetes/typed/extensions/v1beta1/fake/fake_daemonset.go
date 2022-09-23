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

	apiextensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsextensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeDaemonSets implements DaemonSetInterface
type FakeDaemonSets struct {
	Fake *FakeExtensionsV1beta1
	ns   string
}

var daemonsetsResource = pkgruntimeschema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "daemonsets"}

var daemonsetsKind = pkgruntimeschema.GroupVersionKind{Group: "extensions", Version: "v1beta1", Kind: "DaemonSet"}

// Get takes name of the daemonSet, and returns the corresponding daemonSet object, and an error if there is any.
func (c *FakeDaemonSets) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiextensionsv1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewGetAction(daemonsetsResource, c.ns, name), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}

// List takes label and field selectors, and returns the list of DaemonSets that match those selectors.
func (c *FakeDaemonSets) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiextensionsv1beta1.DaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewListAction(daemonsetsResource, daemonsetsKind, c.ns, opts), &apiextensionsv1beta1.DaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apiextensionsv1beta1.DaemonSetList{ListMeta: obj.(*apiextensionsv1beta1.DaemonSetList).ListMeta}
	for _, item := range obj.(*apiextensionsv1beta1.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested daemonSets.
func (c *FakeDaemonSets) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewWatchAction(daemonsetsResource, c.ns, opts))

}

// Create takes the representation of a daemonSet and creates it.  Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *FakeDaemonSets) Create(ctx context.Context, daemonSet *apiextensionsv1beta1.DaemonSet, opts apismetav1.CreateOptions) (result *apiextensionsv1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewCreateAction(daemonsetsResource, c.ns, daemonSet), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}

// Update takes the representation of a daemonSet and updates it. Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *FakeDaemonSets) Update(ctx context.Context, daemonSet *apiextensionsv1beta1.DaemonSet, opts apismetav1.UpdateOptions) (result *apiextensionsv1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateAction(daemonsetsResource, c.ns, daemonSet), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDaemonSets) UpdateStatus(ctx context.Context, daemonSet *apiextensionsv1beta1.DaemonSet, opts apismetav1.UpdateOptions) (*apiextensionsv1beta1.DaemonSet, error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateSubresourceAction(daemonsetsResource, "status", c.ns, daemonSet), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}

// Delete takes name of the daemonSet and deletes it. Returns an error if one occurs.
func (c *FakeDaemonSets) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewDeleteActionWithOptions(daemonsetsResource, c.ns, name, opts), &apiextensionsv1beta1.DaemonSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaemonSets) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewDeleteCollectionAction(daemonsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiextensionsv1beta1.DaemonSetList{})
	return err
}

// Patch applies the patch and returns the patched daemonSet.
func (c *FakeDaemonSets) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiextensionsv1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(daemonsetsResource, c.ns, name, pt, data, subresources...), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied daemonSet.
func (c *FakeDaemonSets) Apply(ctx context.Context, daemonSet *applyconfigurationsextensionsv1beta1.DaemonSetApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiextensionsv1beta1.DaemonSet, err error) {
	if daemonSet == nil {
		return nil, fmt.Errorf("daemonSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(daemonSet)
	if err != nil {
		return nil, err
	}
	name := daemonSet.Name
	if name == nil {
		return nil, fmt.Errorf("daemonSet.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(daemonsetsResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDaemonSets) ApplyStatus(ctx context.Context, daemonSet *applyconfigurationsextensionsv1beta1.DaemonSetApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiextensionsv1beta1.DaemonSet, err error) {
	if daemonSet == nil {
		return nil, fmt.Errorf("daemonSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(daemonSet)
	if err != nil {
		return nil, err
	}
	name := daemonSet.Name
	if name == nil {
		return nil, fmt.Errorf("daemonSet.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(daemonsetsResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data, "status"), &apiextensionsv1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1beta1.DaemonSet), err
}
