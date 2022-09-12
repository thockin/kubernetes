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

	apiextensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	listersextensionsv1beta1 "k8s.io/client-go/listers/extensions/v1beta1"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// NetworkPolicyInformer provides access to a shared informer and lister for
// NetworkPolicies.
type NetworkPolicyInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listersextensionsv1beta1.NetworkPolicyLister
}

type networkPolicyInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkPolicyInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkPolicyInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().NetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().NetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&apiextensionsv1beta1.NetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkPolicyInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredNetworkPolicyInformer(client, f.namespace, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkPolicyInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apiextensionsv1beta1.NetworkPolicy{}, f.defaultInformer)
}

func (f *networkPolicyInformer) Lister() listersextensionsv1beta1.NetworkPolicyLister {
	return listersextensionsv1beta1.NewNetworkPolicyLister(f.Informer().GetIndexer())
}
