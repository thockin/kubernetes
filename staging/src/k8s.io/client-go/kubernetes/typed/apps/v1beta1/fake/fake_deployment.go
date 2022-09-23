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

	apiappsv1beta1 "k8s.io/api/apps/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkglabels "k8s.io/apimachinery/pkg/labels"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"
	clientgotesting "k8s.io/client-go/testing"
)

// FakeDeployments implements DeploymentInterface
type FakeDeployments struct {
	Fake *FakeAppsV1beta1
	ns   string
}

var deploymentsResource = pkgruntimeschema.GroupVersionResource{Group: "apps", Version: "v1beta1", Resource: "deployments"}

var deploymentsKind = pkgruntimeschema.GroupVersionKind{Group: "apps", Version: "v1beta1", Kind: "Deployment"}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *FakeDeployments) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiappsv1beta1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewGetAction(deploymentsResource, c.ns, name), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *FakeDeployments) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiappsv1beta1.DeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewListAction(deploymentsResource, deploymentsKind, c.ns, opts), &apiappsv1beta1.DeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := clientgotesting.ExtractFromListOptions(opts)
	if label == nil {
		label = apimachinerypkglabels.Everything()
	}
	list := &apiappsv1beta1.DeploymentList{ListMeta: obj.(*apiappsv1beta1.DeploymentList).ListMeta}
	for _, item := range obj.(*apiappsv1beta1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested deployments.
func (c *FakeDeployments) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	return c.Fake.
		InvokesWatch(clientgotesting.NewWatchAction(deploymentsResource, c.ns, opts))

}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Create(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.CreateOptions) (result *apiappsv1beta1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewCreateAction(deploymentsResource, c.ns, deployment), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Update(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.UpdateOptions) (result *apiappsv1beta1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateAction(deploymentsResource, c.ns, deployment), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeployments) UpdateStatus(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.UpdateOptions) (*apiappsv1beta1.Deployment, error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewUpdateSubresourceAction(deploymentsResource, "status", c.ns, deployment), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *FakeDeployments) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(clientgotesting.NewDeleteActionWithOptions(deploymentsResource, c.ns, name, opts), &apiappsv1beta1.Deployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployments) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	action := clientgotesting.NewDeleteCollectionAction(deploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiappsv1beta1.DeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deployment.
func (c *FakeDeployments) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiappsv1beta1.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(deploymentsResource, c.ns, name, pt, data, subresources...), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied deployment.
func (c *FakeDeployments) Apply(ctx context.Context, deployment *applyconfigurationsappsv1beta1.DeploymentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.Deployment, err error) {
	if deployment == nil {
		return nil, fmt.Errorf("deployment provided to Apply must not be nil")
	}
	data, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}
	name := deployment.Name
	if name == nil {
		return nil, fmt.Errorf("deployment.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(deploymentsResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDeployments) ApplyStatus(ctx context.Context, deployment *applyconfigurationsappsv1beta1.DeploymentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.Deployment, err error) {
	if deployment == nil {
		return nil, fmt.Errorf("deployment provided to Apply must not be nil")
	}
	data, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}
	name := deployment.Name
	if name == nil {
		return nil, fmt.Errorf("deployment.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(clientgotesting.NewPatchSubresourceAction(deploymentsResource, c.ns, *name, apimachinerypkgtypes.ApplyPatchType, data, "status"), &apiappsv1beta1.Deployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiappsv1beta1.Deployment), err
}
