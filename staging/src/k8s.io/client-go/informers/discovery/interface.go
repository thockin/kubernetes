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

package discovery

import (
	informersdiscoveryv1 "k8s.io/client-go/informers/discovery/v1"
	informersdiscoveryv1beta1 "k8s.io/client-go/informers/discovery/v1beta1"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() informersdiscoveryv1.Interface
	// V1beta1 provides access to shared informers for resources in V1beta1.
	V1beta1() informersdiscoveryv1beta1.Interface
}

type group struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f clientgoinformersinternalinterfaces.SharedInformerFactory, namespace string, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1 returns a new informersdiscoveryv1.Interface.
func (g *group) V1() informersdiscoveryv1.Interface {
	return informersdiscoveryv1.New(g.factory, g.namespace, g.tweakListOptions)
}

// V1beta1 returns a new informersdiscoveryv1beta1.Interface.
func (g *group) V1beta1() informersdiscoveryv1beta1.Interface {
	return informersdiscoveryv1beta1.New(g.factory, g.namespace, g.tweakListOptions)
}
