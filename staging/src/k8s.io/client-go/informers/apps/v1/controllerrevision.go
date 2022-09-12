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
	"context"
	"time"

	apiappsv1 "k8s.io/api/apps/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	listersappsv1 "k8s.io/client-go/listers/apps/v1"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// ControllerRevisionInformer provides access to a shared informer and lister for
// ControllerRevisions.
type ControllerRevisionInformer interface {
	Informer() clientgotoolscache.SharedIndexInformer
	Lister() listersappsv1.ControllerRevisionLister
}

type controllerRevisionInformer struct {
	factory          clientgoinformersinternalinterfaces.SharedInformerFactory
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewControllerRevisionInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers) clientgotoolscache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredControllerRevisionInformer(client clientgokubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers clientgotoolscache.Indexers, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) clientgotoolscache.SharedIndexInformer {
	return clientgotoolscache.NewSharedIndexInformer(
		&clientgotoolscache.ListWatch{
			ListFunc: func(options apismetav1.ListOptions) (apimachinerypkgruntime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().ControllerRevisions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().ControllerRevisions(namespace).Watch(context.TODO(), options)
			},
		},
		&apiappsv1.ControllerRevision{},
		resyncPeriod,
		indexers,
	)
}

func (f *controllerRevisionInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) clientgotoolscache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, f.namespace, resyncPeriod, clientgotoolscache.Indexers{clientgotoolscache.NamespaceIndex: clientgotoolscache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *controllerRevisionInformer) Informer() clientgotoolscache.SharedIndexInformer {
	return f.factory.InformerFor(&apiappsv1.ControllerRevision{}, f.defaultInformer)
}

func (f *controllerRevisionInformer) Lister() listersappsv1.ControllerRevisionLister {
	return listersappsv1.NewControllerRevisionLister(f.Informer().GetIndexer())
}
