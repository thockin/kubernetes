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

	apiappsv1beta1 "k8s.io/api/apps/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// StatefulSetsGetter has a method to return a StatefulSetInterface.
// A group's client should implement this interface.
type StatefulSetsGetter interface {
	StatefulSets(namespace string) StatefulSetInterface
}

// StatefulSetInterface has methods to work with StatefulSet resources.
type StatefulSetInterface interface {
	Create(ctx context.Context, statefulSet *apiappsv1beta1.StatefulSet, opts apismetav1.CreateOptions) (*apiappsv1beta1.StatefulSet, error)
	Update(ctx context.Context, statefulSet *apiappsv1beta1.StatefulSet, opts apismetav1.UpdateOptions) (*apiappsv1beta1.StatefulSet, error)
	UpdateStatus(ctx context.Context, statefulSet *apiappsv1beta1.StatefulSet, opts apismetav1.UpdateOptions) (*apiappsv1beta1.StatefulSet, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apiappsv1beta1.StatefulSet, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apiappsv1beta1.StatefulSetList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiappsv1beta1.StatefulSet, err error)
	Apply(ctx context.Context, statefulSet *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.StatefulSet, err error)
	ApplyStatus(ctx context.Context, statefulSet *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.StatefulSet, err error)
	StatefulSetExpansion
}

// statefulSets implements StatefulSetInterface
type statefulSets struct {
	client clientgorest.Interface
	ns     string
}

// newStatefulSets returns a StatefulSets
func newStatefulSets(c *AppsV1beta1Client, namespace string) *statefulSets {
	return &statefulSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the statefulSet, and returns the corresponding statefulSet object, and an error if there is any.
func (c *statefulSets) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiappsv1beta1.StatefulSet, err error) {
	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors.
func (c *statefulSets) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiappsv1beta1.StatefulSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apiappsv1beta1.StatefulSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested statefulSets.
func (c *statefulSets) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a statefulSet and creates it.  Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *statefulSets) Create(ctx context.Context, statefulSet *apiappsv1beta1.StatefulSet, opts apismetav1.CreateOptions) (result *apiappsv1beta1.StatefulSet, err error) {
	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(statefulSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a statefulSet and updates it. Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *statefulSets) Update(ctx context.Context, statefulSet *apiappsv1beta1.StatefulSet, opts apismetav1.UpdateOptions) (result *apiappsv1beta1.StatefulSet, err error) {
	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(statefulSet.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(statefulSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *statefulSets) UpdateStatus(ctx context.Context, statefulSet *apiappsv1beta1.StatefulSet, opts apismetav1.UpdateOptions) (result *apiappsv1beta1.StatefulSet, err error) {
	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(statefulSet.Name).
		SubResource("status").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(statefulSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the statefulSet and deletes it. Returns an error if one occurs.
func (c *statefulSets) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *statefulSets) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched statefulSet.
func (c *statefulSets) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiappsv1beta1.StatefulSet, err error) {
	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("statefulsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied statefulSet.
func (c *statefulSets) Apply(ctx context.Context, statefulSet *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.StatefulSet, err error) {
	if statefulSet == nil {
		return nil, fmt.Errorf("statefulSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(statefulSet)
	if err != nil {
		return nil, err
	}
	name := statefulSet.Name
	if name == nil {
		return nil, fmt.Errorf("statefulSet.Name must be provided to Apply")
	}
	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Namespace(c.ns).
		Resource("statefulsets").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *statefulSets) ApplyStatus(ctx context.Context, statefulSet *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.StatefulSet, err error) {
	if statefulSet == nil {
		return nil, fmt.Errorf("statefulSet provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(statefulSet)
	if err != nil {
		return nil, err
	}

	name := statefulSet.Name
	if name == nil {
		return nil, fmt.Errorf("statefulSet.Name must be provided to Apply")
	}

	result = &apiappsv1beta1.StatefulSet{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Namespace(c.ns).
		Resource("statefulsets").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
