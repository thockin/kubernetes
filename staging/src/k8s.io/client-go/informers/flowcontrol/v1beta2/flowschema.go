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

// FlowSchemaInformer provides access to a shared informer and lister for
// FlowSchemas.
type FlowSchemaInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listersflowcontrolv1beta2.FlowSchemaLister
}

type flowSchemaInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
}

// NewFlowSchemaInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlowSchemaInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredFlowSchemaInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFlowSchemaInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlowSchemaInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().FlowSchemas().List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().FlowSchemas().Watch(context.TODO(), options)
			},
		},
		&apiflowcontrolv1beta2.FlowSchema{},
		resyncPeriod,
		indexers,
	)
}

func (f *flowSchemaInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredFlowSchemaInformer(client, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *flowSchemaInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apiflowcontrolv1beta2.FlowSchema{}, f.defaultInformer)
}

func (f *flowSchemaInformer) Lister() listersflowcontrolv1beta2.FlowSchemaLister {
	return listersflowcontrolv1beta2.NewFlowSchemaLister(f.Informer().GetIndexer())
}
