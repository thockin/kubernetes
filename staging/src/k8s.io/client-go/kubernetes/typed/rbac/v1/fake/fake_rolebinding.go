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

	apirbacv1 "k8s.io/api/rbac/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsrbacv1 "k8s.io/client-go/applyconfigurations/rbac/v1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeRoleBindings implements RoleBindingInterface
type FakeRoleBindings struct {
	Fake *FakeRbacV1
	ns   string
}

var rolebindingsResource = pkgruntimeschema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "rolebindings"}

var rolebindingsKind = pkgruntimeschema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "RoleBinding"}

// Get takes name of the roleBinding, and returns the corresponding roleBinding object, and an error if there is any.
func (c *FakeRoleBindings) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apirbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewGetAction(rolebindingsResource, c.ns, name), &apirbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apirbacv1.RoleBinding), err
}

// List takes label and field selectors, and returns the list of RoleBindings that match those selectors.
func (c *FakeRoleBindings) List(ctx context.Context, opts apismetav1.ListOptions) (result *apirbacv1.RoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewListAction(rolebindingsResource, rolebindingsKind, c.ns, opts), &apirbacv1.RoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apirbacv1.RoleBindingList{ListMeta: obj.(*apirbacv1.RoleBindingList).ListMeta}
	for _, item := range obj.(*apirbacv1.RoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested roleBindings.
func (c *FakeRoleBindings) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewWatchAction(rolebindingsResource, c.ns, opts))

}

// Create takes the representation of a roleBinding and creates it.  Returns the server's representation of the roleBinding, and an error, if there is any.
func (c *FakeRoleBindings) Create(ctx context.Context, roleBinding *apirbacv1.RoleBinding, opts apismetav1.CreateOptions) (result *apirbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewCreateAction(rolebindingsResource, c.ns, roleBinding), &apirbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apirbacv1.RoleBinding), err
}

// Update takes the representation of a roleBinding and updates it. Returns the server's representation of the roleBinding, and an error, if there is any.
func (c *FakeRoleBindings) Update(ctx context.Context, roleBinding *apirbacv1.RoleBinding, opts apismetav1.UpdateOptions) (result *apirbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateAction(rolebindingsResource, c.ns, roleBinding), &apirbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apirbacv1.RoleBinding), err
}

// Delete takes name of the roleBinding and deletes it. Returns an error if one occurs.
func (c *FakeRoleBindings) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewDeleteActionWithOptions(rolebindingsResource, c.ns, name, opts), &apirbacv1.RoleBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleBindings) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewDeleteCollectionAction(rolebindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apirbacv1.RoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched roleBinding.
func (c *FakeRoleBindings) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apirbacv1.RoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(rolebindingsResource, c.ns, name, pt, data, subresources...), &apirbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apirbacv1.RoleBinding), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied roleBinding.
func (c *FakeRoleBindings) Apply(ctx context.Context, roleBinding *applyconfigurationsrbacv1.RoleBindingApplyConfiguration, opts apismetav1.ApplyOptions) (result *apirbacv1.RoleBinding, err error) {
	if roleBinding == nil {
		return nil, fmt.Errorf("roleBinding provided to Apply must not be nil")
	}
	data, err := json.Marshal(roleBinding)
	if err != nil {
		return nil, err
	}
	name := roleBinding.Name
	if name == nil {
		return nil, fmt.Errorf("roleBinding.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(rolebindingsResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data), &apirbacv1.RoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apirbacv1.RoleBinding), err
}
