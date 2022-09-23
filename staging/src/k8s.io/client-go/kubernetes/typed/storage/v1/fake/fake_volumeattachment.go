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

// FakeVolumeAttachments implements VolumeAttachmentInterface
type FakeVolumeAttachments struct {
	Fake *FakeStorageV1
}

var volumeattachmentsResource = pkgruntimeschema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "volumeattachments"}

var volumeattachmentsKind = pkgruntimeschema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1", Kind: "VolumeAttachment"}

// Get takes name of the volumeAttachment, and returns the corresponding volumeAttachment object, and an error if there is any.
func (c *FakeVolumeAttachments) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apistoragev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootGetAction(volumeattachmentsResource, name), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors.
func (c *FakeVolumeAttachments) List(ctx context.Context, opts apismetav1.ListOptions) (result *apistoragev1.VolumeAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootListAction(volumeattachmentsResource, volumeattachmentsKind, opts), &apistoragev1.VolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apistoragev1.VolumeAttachmentList{ListMeta: obj.(*apistoragev1.VolumeAttachmentList).ListMeta}
	for _, item := range obj.(*apistoragev1.VolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested volumeAttachments.
func (c *FakeVolumeAttachments) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewRootWatchAction(volumeattachmentsResource, opts))
}

// Create takes the representation of a volumeAttachment and creates it.  Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *FakeVolumeAttachments) Create(ctx context.Context, volumeAttachment *apistoragev1.VolumeAttachment, opts apismetav1.CreateOptions) (result *apistoragev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootCreateAction(volumeattachmentsResource, volumeAttachment), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}

// Update takes the representation of a volumeAttachment and updates it. Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *FakeVolumeAttachments) Update(ctx context.Context, volumeAttachment *apistoragev1.VolumeAttachment, opts apismetav1.UpdateOptions) (result *apistoragev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateAction(volumeattachmentsResource, volumeAttachment), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeAttachments) UpdateStatus(ctx context.Context, volumeAttachment *apistoragev1.VolumeAttachment, opts apismetav1.UpdateOptions) (*apistoragev1.VolumeAttachment, error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootUpdateSubresourceAction(volumeattachmentsResource, "status", volumeAttachment), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}

// Delete takes name of the volumeAttachment and deletes it. Returns an error if one occurs.
func (c *FakeVolumeAttachments) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewRootDeleteActionWithOptions(volumeattachmentsResource, name, opts), &apistoragev1.VolumeAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeAttachments) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewRootDeleteCollectionAction(volumeattachmentsResource, listOpts)

	_, err := c.Fake.Invokes(action, &apistoragev1.VolumeAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched volumeAttachment.
func (c *FakeVolumeAttachments) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apistoragev1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(volumeattachmentsResource, name, pt, data, subresources...), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied volumeAttachment.
func (c *FakeVolumeAttachments) Apply(ctx context.Context, volumeAttachment *applyconfigurationsstoragev1.VolumeAttachmentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apistoragev1.VolumeAttachment, err error) {
	if volumeAttachment == nil {
		return nil, fmt.Errorf("volumeAttachment provided to Apply must not be nil")
	}
	data, err := json.Marshal(volumeAttachment)
	if err != nil {
		return nil, err
	}
	name := volumeAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("volumeAttachment.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(volumeattachmentsResource, *name, apimachinerypkgtypes.ApplyPatchType, data), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeVolumeAttachments) ApplyStatus(ctx context.Context, volumeAttachment *applyconfigurationsstoragev1.VolumeAttachmentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apistoragev1.VolumeAttachment, err error) {
	if volumeAttachment == nil {
		return nil, fmt.Errorf("volumeAttachment provided to Apply must not be nil")
	}
	data, err := json.Marshal(volumeAttachment)
	if err != nil {
		return nil, err
	}
	name := volumeAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("volumeAttachment.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewRootPatchSubresourceAction(volumeattachmentsResource, *name, apimachinerypkgtypes.ApplyPatchType, data, "status"), &apistoragev1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apistoragev1.VolumeAttachment), err
}
