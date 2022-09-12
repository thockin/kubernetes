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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// PodDisruptionBudgets returns a PodDisruptionBudgetInformer.
	PodDisruptionBudgets() PodDisruptionBudgetInformer
}

type version struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f clientgoinformersinternalinterfaces.SharedInformerFactory, namespace string, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// PodDisruptionBudgets returns a PodDisruptionBudgetInformer.
func (v *version) PodDisruptionBudgets() PodDisruptionBudgetInformer {
	return &podDisruptionBudgetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
