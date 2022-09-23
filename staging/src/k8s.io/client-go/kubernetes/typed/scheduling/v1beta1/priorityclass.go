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

	apischedulingv1beta1 "k8s.io/api/scheduling/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsschedulingv1beta1 "k8s.io/client-go/applyconfigurations/scheduling/v1beta1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// PriorityClassesGetter has a method to return a PriorityClassInterface.
// A group's client should implement this interface.
type PriorityClassesGetter interface {
	PriorityClasses() PriorityClassInterface
}

// PriorityClassInterface has methods to work with PriorityClass resources.
type PriorityClassInterface interface {
	Create(ctx context.Context, priorityClass *apischedulingv1beta1.PriorityClass, opts apismetav1.CreateOptions) (*apischedulingv1beta1.PriorityClass, error)
	Update(ctx context.Context, priorityClass *apischedulingv1beta1.PriorityClass, opts apismetav1.UpdateOptions) (*apischedulingv1beta1.PriorityClass, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apischedulingv1beta1.PriorityClass, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apischedulingv1beta1.PriorityClassList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apischedulingv1beta1.PriorityClass, err error)
	Apply(ctx context.Context, priorityClass *applyconfigurationsschedulingv1beta1.PriorityClassApplyConfiguration, opts apismetav1.ApplyOptions) (result *apischedulingv1beta1.PriorityClass, err error)
	PriorityClassExpansion
}

// priorityClasses implements PriorityClassInterface
type priorityClasses struct {
	client clientgorest.Interface
}

// newPriorityClasses returns a PriorityClasses
func newPriorityClasses(c *SchedulingV1beta1Client) *priorityClasses {
	return &priorityClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the priorityClass, and returns the corresponding priorityClass object, and an error if there is any.
func (c *priorityClasses) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apischedulingv1beta1.PriorityClass, err error) {
	result = &apischedulingv1beta1.PriorityClass{}
	err = c.client.Get().
		Resource("priorityclasses").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PriorityClasses that match those selectors.
func (c *priorityClasses) List(ctx context.Context, opts apismetav1.ListOptions) (result *apischedulingv1beta1.PriorityClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apischedulingv1beta1.PriorityClassList{}
	err = c.client.Get().
		Resource("priorityclasses").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested priorityClasses.
func (c *priorityClasses) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("priorityclasses").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a priorityClass and creates it.  Returns the server's representation of the priorityClass, and an error, if there is any.
func (c *priorityClasses) Create(ctx context.Context, priorityClass *apischedulingv1beta1.PriorityClass, opts apismetav1.CreateOptions) (result *apischedulingv1beta1.PriorityClass, err error) {
	result = &apischedulingv1beta1.PriorityClass{}
	err = c.client.Post().
		Resource("priorityclasses").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(priorityClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a priorityClass and updates it. Returns the server's representation of the priorityClass, and an error, if there is any.
func (c *priorityClasses) Update(ctx context.Context, priorityClass *apischedulingv1beta1.PriorityClass, opts apismetav1.UpdateOptions) (result *apischedulingv1beta1.PriorityClass, err error) {
	result = &apischedulingv1beta1.PriorityClass{}
	err = c.client.Put().
		Resource("priorityclasses").
		Name(priorityClass.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(priorityClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the priorityClass and deletes it. Returns an error if one occurs.
func (c *priorityClasses) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("priorityclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *priorityClasses) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("priorityclasses").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched priorityClass.
func (c *priorityClasses) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apischedulingv1beta1.PriorityClass, err error) {
	result = &apischedulingv1beta1.PriorityClass{}
	err = c.client.Patch(pt).
		Resource("priorityclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityClass.
func (c *priorityClasses) Apply(ctx context.Context, priorityClass *applyconfigurationsschedulingv1beta1.PriorityClassApplyConfiguration, opts apismetav1.ApplyOptions) (result *apischedulingv1beta1.PriorityClass, err error) {
	if priorityClass == nil {
		return nil, fmt.Errorf("priorityClass provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(priorityClass)
	if err != nil {
		return nil, err
	}
	name := priorityClass.Name
	if name == nil {
		return nil, fmt.Errorf("priorityClass.Name must be provided to Apply")
	}
	result = &apischedulingv1beta1.PriorityClass{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Resource("priorityclasses").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
