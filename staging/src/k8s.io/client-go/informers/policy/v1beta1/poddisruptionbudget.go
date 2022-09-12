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

package v1beta1

import (
	"context"
	"time"

	apipolicyv1beta1 "k8s.io/api/policy/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	listerspolicyv1beta1 "k8s.io/client-go/listers/policy/v1beta1"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// PodDisruptionBudgetInformer provides access to a shared informer and lister for
// PodDisruptionBudgets.
type PodDisruptionBudgetInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listerspolicyv1beta1.PodDisruptionBudgetLister
}

type podDisruptionBudgetInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodDisruptionBudgetInformer constructs a new informer for PodDisruptionBudget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodDisruptionBudgetInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredPodDisruptionBudgetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodDisruptionBudgetInformer constructs a new informer for PodDisruptionBudget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodDisruptionBudgetInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1beta1().PodDisruptionBudgets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1beta1().PodDisruptionBudgets(namespace).Watch(context.TODO(), options)
			},
		},
		&apipolicyv1beta1.PodDisruptionBudget{},
		resyncPeriod,
		indexers,
	)
}

func (f *podDisruptionBudgetInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredPodDisruptionBudgetInformer(client, f.namespace, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podDisruptionBudgetInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apipolicyv1beta1.PodDisruptionBudget{}, f.defaultInformer)
}

func (f *podDisruptionBudgetInformer) Lister() listerspolicyv1beta1.PodDisruptionBudgetLister {
	return listerspolicyv1beta1.NewPodDisruptionBudgetLister(f.Informer().GetIndexer())
}
