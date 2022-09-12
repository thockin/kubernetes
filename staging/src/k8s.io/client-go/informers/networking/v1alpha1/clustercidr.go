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

package v1alpha1

import (
	"context"
	"time"

	apinetworkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	listersnetworkingv1alpha1 "k8s.io/client-go/listers/networking/v1alpha1"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// ClusterCIDRInformer provides access to a shared informer and lister for
// ClusterCIDRs.
type ClusterCIDRInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listersnetworkingv1alpha1.ClusterCIDRLister
}

type clusterCIDRInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
}

// NewClusterCIDRInformer constructs a new informer for ClusterCIDR type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterCIDRInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredClusterCIDRInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterCIDRInformer constructs a new informer for ClusterCIDR type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterCIDRInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().ClusterCIDRs().List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().ClusterCIDRs().Watch(context.TODO(), options)
			},
		},
		&apinetworkingv1alpha1.ClusterCIDR{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterCIDRInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredClusterCIDRInformer(client, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterCIDRInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apinetworkingv1alpha1.ClusterCIDR{}, f.defaultInformer)
}

func (f *clusterCIDRInformer) Lister() listersnetworkingv1alpha1.ClusterCIDRLister {
	return listersnetworkingv1alpha1.NewClusterCIDRLister(f.Informer().GetIndexer())
}
