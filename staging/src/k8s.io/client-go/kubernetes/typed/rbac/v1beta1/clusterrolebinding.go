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

package v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	apirbacv1beta1 "k8s.io/api/rbac/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsrbacv1beta1 "k8s.io/client-go/applyconfigurations/rbac/v1beta1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// ClusterRoleBindingsGetter has a method to return a ClusterRoleBindingInterface.
// A group's client should implement this interface.
type ClusterRoleBindingsGetter interface {
	ClusterRoleBindings() ClusterRoleBindingInterface
}

// ClusterRoleBindingInterface has methods to work with ClusterRoleBinding resources.
type ClusterRoleBindingInterface interface {
	Create(ctx context.Context, clusterRoleBinding *apirbacv1beta1.ClusterRoleBinding, opts apismetav1.CreateOptions) (*apirbacv1beta1.ClusterRoleBinding, error)
	Update(ctx context.Context, clusterRoleBinding *apirbacv1beta1.ClusterRoleBinding, opts apismetav1.UpdateOptions) (*apirbacv1beta1.ClusterRoleBinding, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apirbacv1beta1.ClusterRoleBinding, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apirbacv1beta1.ClusterRoleBindingList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apirbacv1beta1.ClusterRoleBinding, err error)
	Apply(ctx context.Context, clusterRoleBinding *applyconfigurationsrbacv1beta1.ClusterRoleBindingApplyConfiguration, opts apismetav1.ApplyOptions) (result *apirbacv1beta1.ClusterRoleBinding, err error)
	ClusterRoleBindingExpansion
}

// clusterRoleBindings implements ClusterRoleBindingInterface
type clusterRoleBindings struct {
	client clientgorest.Interface
}

// newClusterRoleBindings returns a ClusterRoleBindings
func newClusterRoleBindings(c *RbacV1beta1Client) *clusterRoleBindings {
	return &clusterRoleBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterRoleBinding, and returns the corresponding clusterRoleBinding object, and an error if there is any.
func (c *clusterRoleBindings) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apirbacv1beta1.ClusterRoleBinding, err error) {
	result = &apirbacv1beta1.ClusterRoleBinding{}
	err = c.client.Get().
		Resource("clusterrolebindings").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterRoleBindings that match those selectors.
func (c *clusterRoleBindings) List(ctx context.Context, opts apismetav1.ListOptions) (result *apirbacv1beta1.ClusterRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apirbacv1beta1.ClusterRoleBindingList{}
	err = c.client.Get().
		Resource("clusterrolebindings").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested clusterRoleBindings.
func (c *clusterRoleBindings) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterrolebindings").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterRoleBinding and creates it.  Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *clusterRoleBindings) Create(ctx context.Context, clusterRoleBinding *apirbacv1beta1.ClusterRoleBinding, opts apismetav1.CreateOptions) (result *apirbacv1beta1.ClusterRoleBinding, err error) {
	result = &apirbacv1beta1.ClusterRoleBinding{}
	err = c.client.Post().
		Resource("clusterrolebindings").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(clusterRoleBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterRoleBinding and updates it. Returns the server's representation of the clusterRoleBinding, and an error, if there is any.
func (c *clusterRoleBindings) Update(ctx context.Context, clusterRoleBinding *apirbacv1beta1.ClusterRoleBinding, opts apismetav1.UpdateOptions) (result *apirbacv1beta1.ClusterRoleBinding, err error) {
	result = &apirbacv1beta1.ClusterRoleBinding{}
	err = c.client.Put().
		Resource("clusterrolebindings").
		Name(clusterRoleBinding.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(clusterRoleBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterRoleBinding and deletes it. Returns an error if one occurs.
func (c *clusterRoleBindings) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterrolebindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterRoleBindings) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterrolebindings").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterRoleBinding.
func (c *clusterRoleBindings) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apirbacv1beta1.ClusterRoleBinding, err error) {
	result = &apirbacv1beta1.ClusterRoleBinding{}
	err = c.client.Patch(pt).
		Resource("clusterrolebindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterRoleBinding.
func (c *clusterRoleBindings) Apply(ctx context.Context, clusterRoleBinding *applyconfigurationsrbacv1beta1.ClusterRoleBindingApplyConfiguration, opts apismetav1.ApplyOptions) (result *apirbacv1beta1.ClusterRoleBinding, err error) {
	if clusterRoleBinding == nil {
		return nil, fmt.Errorf("clusterRoleBinding provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterRoleBinding)
	if err != nil {
		return nil, err
	}
	name := clusterRoleBinding.Name
	if name == nil {
		return nil, fmt.Errorf("clusterRoleBinding.Name must be provided to Apply")
	}
	result = &apirbacv1beta1.ClusterRoleBinding{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Resource("clusterrolebindings").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
