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

package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	apidiscoveryv1 "k8s.io/api/discovery/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsdiscoveryv1 "k8s.io/client-go/applyconfigurations/discovery/v1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// EndpointSlicesGetter has a method to return a EndpointSliceInterface.
// A group's client should implement this interface.
type EndpointSlicesGetter interface {
	EndpointSlices(namespace string) EndpointSliceInterface
}

// EndpointSliceInterface has methods to work with EndpointSlice resources.
type EndpointSliceInterface interface {
	Create(ctx context.Context, endpointSlice *apidiscoveryv1.EndpointSlice, opts apismetav1.CreateOptions) (*apidiscoveryv1.EndpointSlice, error)
	Update(ctx context.Context, endpointSlice *apidiscoveryv1.EndpointSlice, opts apismetav1.UpdateOptions) (*apidiscoveryv1.EndpointSlice, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apidiscoveryv1.EndpointSlice, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apidiscoveryv1.EndpointSliceList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apidiscoveryv1.EndpointSlice, err error)
	Apply(ctx context.Context, endpointSlice *applyconfigurationsdiscoveryv1.EndpointSliceApplyConfiguration, opts apismetav1.ApplyOptions) (result *apidiscoveryv1.EndpointSlice, err error)
	EndpointSliceExpansion
}

// endpointSlices implements EndpointSliceInterface
type endpointSlices struct {
	client clientgorest.Interface
	ns     string
}

// newEndpointSlices returns a EndpointSlices
func newEndpointSlices(c *DiscoveryV1Client, namespace string) *endpointSlices {
	return &endpointSlices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the endpointSlice, and returns the corresponding endpointSlice object, and an error if there is any.
func (c *endpointSlices) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apidiscoveryv1.EndpointSlice, err error) {
	result = &apidiscoveryv1.EndpointSlice{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpointslices").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EndpointSlices that match those selectors.
func (c *endpointSlices) List(ctx context.Context, opts apismetav1.ListOptions) (result *apidiscoveryv1.EndpointSliceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apidiscoveryv1.EndpointSliceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested endpointSlices.
func (c *endpointSlices) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a endpointSlice and creates it.  Returns the server's representation of the endpointSlice, and an error, if there is any.
func (c *endpointSlices) Create(ctx context.Context, endpointSlice *apidiscoveryv1.EndpointSlice, opts apismetav1.CreateOptions) (result *apidiscoveryv1.EndpointSlice, err error) {
	result = &apidiscoveryv1.EndpointSlice{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(endpointSlice).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a endpointSlice and updates it. Returns the server's representation of the endpointSlice, and an error, if there is any.
func (c *endpointSlices) Update(ctx context.Context, endpointSlice *apidiscoveryv1.EndpointSlice, opts apismetav1.UpdateOptions) (result *apidiscoveryv1.EndpointSlice, err error) {
	result = &apidiscoveryv1.EndpointSlice{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpointslices").
		Name(endpointSlice.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(endpointSlice).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the endpointSlice and deletes it. Returns an error if one occurs.
func (c *endpointSlices) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpointslices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpointSlices) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpointslices").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched endpointSlice.
func (c *endpointSlices) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apidiscoveryv1.EndpointSlice, err error) {
	result = &apidiscoveryv1.EndpointSlice{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("endpointslices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied endpointSlice.
func (c *endpointSlices) Apply(ctx context.Context, endpointSlice *applyconfigurationsdiscoveryv1.EndpointSliceApplyConfiguration, opts apismetav1.ApplyOptions) (result *apidiscoveryv1.EndpointSlice, err error) {
	if endpointSlice == nil {
		return nil, fmt.Errorf("endpointSlice provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(endpointSlice)
	if err != nil {
		return nil, err
	}
	name := endpointSlice.Name
	if name == nil {
		return nil, fmt.Errorf("endpointSlice.Name must be provided to Apply")
	}
	result = &apidiscoveryv1.EndpointSlice{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Namespace(c.ns).
		Resource("endpointslices").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
