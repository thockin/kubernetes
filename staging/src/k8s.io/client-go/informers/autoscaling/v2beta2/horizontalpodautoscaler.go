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

package v2beta2

import (
	"context"
	"time"

	apiautoscalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	listersautoscalingv2beta2 "k8s.io/client-go/listers/autoscaling/v2beta2"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// HorizontalPodAutoscalerInformer provides access to a shared informer and lister for
// HorizontalPodAutoscalers.
type HorizontalPodAutoscalerInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listersautoscalingv2beta2.HorizontalPodAutoscalerLister
}

type horizontalPodAutoscalerInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHorizontalPodAutoscalerInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHorizontalPodAutoscalerInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHorizontalPodAutoscalerInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHorizontalPodAutoscalerInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).Watch(context.TODO(), options)
			},
		},
		&apiautoscalingv2beta2.HorizontalPodAutoscaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *horizontalPodAutoscalerInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerInformer(client, f.namespace, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *horizontalPodAutoscalerInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apiautoscalingv2beta2.HorizontalPodAutoscaler{}, f.defaultInformer)
}

func (f *horizontalPodAutoscalerInformer) Lister() listersautoscalingv2beta2.HorizontalPodAutoscalerLister {
	return listersautoscalingv2beta2.NewHorizontalPodAutoscalerLister(f.Informer().GetIndexer())
}
