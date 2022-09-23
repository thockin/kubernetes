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

<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
	v1 "k8s.io/api/flowcontrol/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	flowcontrolv1 "k8s.io/client-go/applyconfigurations/flowcontrol/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	flowcontrolv1alpha1 "k8s.io/client-go/applyconfigurations/flowcontrol/v1alpha1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
=======
	apiflowcontrolv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsflowcontrolv1alpha1 "k8s.io/client-go/applyconfigurations/flowcontrol/v1alpha1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
)

// PriorityLevelConfigurationsGetter has a method to return a PriorityLevelConfigurationInterface.
// A group's client should implement this interface.
type PriorityLevelConfigurationsGetter interface {
	PriorityLevelConfigurations() PriorityLevelConfigurationInterface
}

// PriorityLevelConfigurationInterface has methods to work with PriorityLevelConfiguration resources.
type PriorityLevelConfigurationInterface interface {
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
	Create(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.CreateOptions) (*v1.PriorityLevelConfiguration, error)
	Update(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.UpdateOptions) (*v1.PriorityLevelConfiguration, error)
	UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.UpdateOptions) (*v1.PriorityLevelConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PriorityLevelConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PriorityLevelConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PriorityLevelConfiguration, err error)
	Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityLevelConfiguration, err error)
	ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityLevelConfiguration, err error)
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	Create(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.CreateOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	Update(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PriorityLevelConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PriorityLevelConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PriorityLevelConfiguration, err error)
	Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error)
	ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error)
=======
	Create(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1alpha1.PriorityLevelConfiguration, opts apismetav1.CreateOptions) (*apiflowcontrolv1alpha1.PriorityLevelConfiguration, error)
	Update(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1alpha1.PriorityLevelConfiguration, opts apismetav1.UpdateOptions) (*apiflowcontrolv1alpha1.PriorityLevelConfiguration, error)
	UpdateStatus(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1alpha1.PriorityLevelConfiguration, opts apismetav1.UpdateOptions) (*apiflowcontrolv1alpha1.PriorityLevelConfiguration, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apiflowcontrolv1alpha1.PriorityLevelConfiguration, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apiflowcontrolv1alpha1.PriorityLevelConfigurationList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error)
	Apply(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error)
	ApplyStatus(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error)
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	PriorityLevelConfigurationExpansion
}

// priorityLevelConfigurations implements PriorityLevelConfigurationInterface
type priorityLevelConfigurations struct {
	client clientgorest.Interface
}

// newPriorityLevelConfigurations returns a PriorityLevelConfigurations
func newPriorityLevelConfigurations(c *FlowcontrolV1Client) *priorityLevelConfigurations {
	return &priorityLevelConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the priorityLevelConfiguration, and returns the corresponding priorityLevelConfiguration object, and an error if there is any.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PriorityLevelConfiguration, err error) {
	result = &v1.PriorityLevelConfiguration{}
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
=======
func (c *priorityLevelConfigurations) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	err = c.client.Get().
		Resource("prioritylevelconfigurations").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PriorityLevelConfigurationList, err error) {
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PriorityLevelConfigurationList, err error) {
=======
func (c *priorityLevelConfigurations) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfigurationList, err error) {
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
	result = &v1.PriorityLevelConfigurationList{}
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	result = &v1alpha1.PriorityLevelConfigurationList{}
=======
	result = &apiflowcontrolv1alpha1.PriorityLevelConfigurationList{}
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	err = c.client.Get().
		Resource("prioritylevelconfigurations").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
// Watch returns a watch.Interface that watches the requested priorityLevelConfigurations.
func (c *priorityLevelConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
// Watch returns a watch.Interface that watches the requested priorityLevelConfigurations.
func (c *priorityLevelConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
=======
// Watch returns a apimachinerypkgwatch.Interface that watches the requested priorityLevelConfigurations.
func (c *priorityLevelConfigurations) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("prioritylevelconfigurations").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a priorityLevelConfiguration and creates it.  Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.CreateOptions) (result *v1.PriorityLevelConfiguration, err error) {
	result = &v1.PriorityLevelConfiguration{}
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.CreateOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
=======
func (c *priorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1alpha1.PriorityLevelConfiguration, opts apismetav1.CreateOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	err = c.client.Post().
		Resource("prioritylevelconfigurations").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(priorityLevelConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a priorityLevelConfiguration and updates it. Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.UpdateOptions) (result *v1.PriorityLevelConfiguration, err error) {
	result = &v1.PriorityLevelConfiguration{}
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
=======
func (c *priorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1alpha1.PriorityLevelConfiguration, opts apismetav1.UpdateOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	err = c.client.Put().
		Resource("prioritylevelconfigurations").
		Name(priorityLevelConfiguration.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(priorityLevelConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1.PriorityLevelConfiguration, opts metav1.UpdateOptions) (result *v1.PriorityLevelConfiguration, err error) {
	result = &v1.PriorityLevelConfiguration{}
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1alpha1.PriorityLevelConfiguration, opts v1.UpdateOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
=======
func (c *priorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *apiflowcontrolv1alpha1.PriorityLevelConfiguration, opts apismetav1.UpdateOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	err = c.client.Put().
		Resource("prioritylevelconfigurations").
		Name(priorityLevelConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(priorityLevelConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the priorityLevelConfiguration and deletes it. Returns an error if one occurs.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
=======
func (c *priorityLevelConfigurations) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	return c.client.Delete().
		Resource("prioritylevelconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
=======
func (c *priorityLevelConfigurations) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("prioritylevelconfigurations").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched priorityLevelConfiguration.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PriorityLevelConfiguration, err error) {
	result = &v1.PriorityLevelConfiguration{}
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PriorityLevelConfiguration, err error) {
	result = &v1alpha1.PriorityLevelConfiguration{}
=======
func (c *priorityLevelConfigurations) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	err = c.client.Patch(pt).
		Resource("prioritylevelconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityLevelConfiguration.
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityLevelConfiguration, err error) {
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
=======
func (c *priorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
	result = &v1.PriorityLevelConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
=======
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
		Resource("prioritylevelconfigurations").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1.PriorityLevelConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityLevelConfiguration, err error) {
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
func (c *priorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PriorityLevelConfiguration, err error) {
=======
func (c *priorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *applyconfigurationsflowcontrolv1alpha1.PriorityLevelConfigurationApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1alpha1.PriorityLevelConfiguration, err error) {
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}

	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}

<<<<<<< HEAD:staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1/prioritylevelconfiguration.go
	result = &v1.PriorityLevelConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
||||||| parent of e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
	result = &v1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(types.ApplyPatchType).
=======
	result = &apiflowcontrolv1alpha1.PriorityLevelConfiguration{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
>>>>>>> e9771bdbaa6 (codegen client-gen (BROKEN)):staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1/prioritylevelconfiguration.go
		Resource("prioritylevelconfigurations").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
