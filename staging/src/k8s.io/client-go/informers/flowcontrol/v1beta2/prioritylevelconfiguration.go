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

package v1beta2

import (
	"context"
	"time"

	apiflowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	listersflowcontrolv1beta2 "k8s.io/client-go/listers/flowcontrol/v1beta2"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// PriorityLevelConfigurationInformer provides access to a shared informer and lister for
// PriorityLevelConfigurations.
type PriorityLevelConfigurationInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listersflowcontrolv1beta2.PriorityLevelConfigurationLister
}

type priorityLevelConfigurationInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
}

// NewPriorityLevelConfigurationInformer constructs a new informer for PriorityLevelConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPriorityLevelConfigurationInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredPriorityLevelConfigurationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPriorityLevelConfigurationInformer constructs a new informer for PriorityLevelConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPriorityLevelConfigurationInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().PriorityLevelConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().PriorityLevelConfigurations().Watch(context.TODO(), options)
			},
		},
		&apiflowcontrolv1beta2.PriorityLevelConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *priorityLevelConfigurationInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredPriorityLevelConfigurationInformer(client, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *priorityLevelConfigurationInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apiflowcontrolv1beta2.PriorityLevelConfiguration{}, f.defaultInformer)
}

func (f *priorityLevelConfigurationInformer) Lister() listersflowcontrolv1beta2.PriorityLevelConfigurationLister {
	return listersflowcontrolv1beta2.NewPriorityLevelConfigurationLister(f.Informer().GetIndexer())
}
