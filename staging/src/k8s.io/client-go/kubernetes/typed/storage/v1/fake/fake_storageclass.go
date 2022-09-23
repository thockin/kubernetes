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

	apistoragev1 "k8s.io/api/storage/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1 "k8s.io/client-go/applyconfigurations/storage/v1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeStorageClasses implements StorageClassInterface
type FakeStorageClasses struct {
	Fake *FakeStorageV1
}

var storageclassesResource = pkgruntimeschema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}

var storageclassesKind = pkgruntimeschema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1", Kind: "StorageClass"}

// Get takes name of the storageClass, and returns the corresponding storageClass object, and an error if there is any.
func (c *FakeStorageClasses) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apistoragev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootGetAction(storageclassesResource, name), &apistoragev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.StorageClass), err
}

// List takes label and field selectors, and returns the list of StorageClasses that match those selectors.
func (c *FakeStorageClasses) List(ctx context.Context, opts apismetav1.ListOptions) (result *apistoragev1.StorageClassList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootListAction(storageclassesResource, storageclassesKind, opts), &apistoragev1.StorageClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apistoragev1.StorageClassList{ListMeta: obj.(*apistoragev1.StorageClassList).ListMeta}
	for _, item := range obj.(*apistoragev1.StorageClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested storageClasses.
func (c *FakeStorageClasses) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewRootWatchAction(storageclassesResource, opts))
}

// Create takes the representation of a storageClass and creates it.  Returns the server's representation of the storageClass, and an error, if there is any.
func (c *FakeStorageClasses) Create(ctx context.Context, storageClass *apistoragev1.StorageClass, opts apismetav1.CreateOptions) (result *apistoragev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootCreateAction(storageclassesResource, storageClass), &apistoragev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.StorageClass), err
}

// Update takes the representation of a storageClass and updates it. Returns the server's representation of the storageClass, and an error, if there is any.
func (c *FakeStorageClasses) Update(ctx context.Context, storageClass *apistoragev1.StorageClass, opts apismetav1.UpdateOptions) (result *apistoragev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateAction(storageclassesResource, storageClass), &apistoragev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.StorageClass), err
}

// Delete takes name of the storageClass and deletes it. Returns an error if one occurs.
func (c *FakeStorageClasses) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewRootDeleteActionWithOptions(storageclassesResource, name, opts), &apistoragev1.StorageClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageClasses) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewRootDeleteCollectionAction(storageclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &apistoragev1.StorageClassList{})
	return err
}

// Patch applies the patch and returns the patched storageClass.
func (c *FakeStorageClasses) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apistoragev1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(storageclassesResource, name, pt, data, subresources...), &apistoragev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.StorageClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied storageClass.
func (c *FakeStorageClasses) Apply(ctx context.Context, storageClass *applyconfigurationsstoragev1.StorageClassApplyConfiguration, opts apismetav1.ApplyOptions) (result *apistoragev1.StorageClass, err error) {
	if storageClass == nil {
		return nil, fmt.Errorf("storageClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(storageClass)
	if err != nil {
		return nil, err
	}
	name := storageClass.Name
	if name == nil {
		return nil, fmt.Errorf("storageClass.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(storageclassesResource, *name, apimachinerypkgtypes.ApplyPatchType, data), &apistoragev1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.StorageClass), err
}
