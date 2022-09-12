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

package informers

import (
	"reflect"
	"sync"
	"time"

	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgruntime "k8s.io/apimachinery/pkg/runtime"
	pkgruntimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	clientgoinformersadmissionregistration "k8s.io/client-go/informers/admissionregistration"
	clientgoinformersapiserverinternal "k8s.io/client-go/informers/apiserverinternal"
	clientgoinformersapps "k8s.io/client-go/informers/apps"
	clientgoinformersautoscaling "k8s.io/client-go/informers/autoscaling"
	clientgoinformersbatch "k8s.io/client-go/informers/batch"
	clientgoinformerscertificates "k8s.io/client-go/informers/certificates"
	clientgoinformerscoordination "k8s.io/client-go/informers/coordination"
	clientgoinformerscore "k8s.io/client-go/informers/core"
	clientgoinformersdiscovery "k8s.io/client-go/informers/discovery"
	clientgoinformersevents "k8s.io/client-go/informers/events"
	clientgoinformersextensions "k8s.io/client-go/informers/extensions"
	clientgoinformersflowcontrol "k8s.io/client-go/informers/flowcontrol"
	clientgoinformersinternalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgoinformersnetworking "k8s.io/client-go/informers/networking"
	clientgoinformersnode "k8s.io/client-go/informers/node"
	clientgoinformerspolicy "k8s.io/client-go/informers/policy"
	clientgoinformersrbac "k8s.io/client-go/informers/rbac"
	clientgoinformersscheduling "k8s.io/client-go/informers/scheduling"
	clientgoinformersstorage "k8s.io/client-go/informers/storage"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	clientgotoolscache "k8s.io/client-go/tools/cache"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           clientgokubernetes.Interface
	namespace        string
	tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]clientgotoolscache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
	// wg tracks how many goroutines were started.
	wg sync.WaitGroup
	// shuttingDown is true when Shutdown has been called. It may still be running
	// because it needs to wait for goroutines.
	shuttingDown bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[apismetav1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.namespace = namespace
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client clientgokubernetes.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
// Deprecated: Please use NewSharedInformerFactoryWithOptions instead
func NewFilteredSharedInformerFactory(client clientgokubernetes.Interface, defaultResync time.Duration, namespace string, tweakListOptions clientgoinformersinternalinterfaces.TweakListOptionsFunc) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace), WithTweakListOptions(tweakListOptions))
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client clientgokubernetes.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        v1.NamespaceAll,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]clientgotoolscache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.shuttingDown {
		return
	}

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			f.wg.Add(1)
			// We need a new variable in each loop iteration,
			// otherwise the goroutine would use the loop variable
			// and that keeps changing.
			informer := informer
			go func() {
				defer f.wg.Done()
				informer.Run(stopCh)
			}()
			f.startedInformers[informerType] = true
		}
	}
}

func (f *sharedInformerFactory) Shutdown() {
	f.lock.Lock()
	f.shuttingDown = true
	f.lock.Unlock()

	// Will return immediately if there is nothing to wait for.
	f.wg.Wait()
}

func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj apimachinerypkgruntime.Object, newFunc clientgoinformersinternalinterfaces.NewInformerFunc) clientgotoolscache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
//
// It is typically used like this:
//
//	ctx, cancel := context.Background()
//	defer cancel()
//	factory := NewSharedInformerFactory(client, resyncPeriod)
//	defer factory.WaitForStop()    // Returns immediately if nothing was started.
//	genericInformer := factory.ForResource(resource)
//	typedInformer := factory.SomeAPIGroup().V1().SomeType()
//	factory.Start(ctx.Done())          // Start processing these informers.
//	synced := factory.WaitForCacheSync(ctx.Done())
//	for v, ok := range synced {
//	    if !ok {
//	        fmt.Fprintf(os.Stderr, "caches failed to sync: %v", v)
//	        return
//	    }
//	}
//
//	// Creating informers can also be created after Start, but then
//	// Start must be called again:
//	anotherGenericInformer := factory.ForResource(resource)
//	factory.Start(ctx.Done())
type SharedInformerFactory interface {
	clientgoinformersinternalinterfaces.SharedInformerFactory

	// Start initializes all requested informers. They are handled in goroutines
	// which run until the stop channel gets closed.
	Start(stopCh <-chan struct{})

	// Shutdown marks a factory as shutting down. At that point no new
	// informers can be started anymore and Start will return without
	// doing anything.
	//
	// In addition, Shutdown blocks until all goroutines have terminated. For that
	// to happen, the close channel(s) that they were started with must be closed,
	// either before Shutdown gets called or while it is waiting.
	//
	// Shutdown may be called multiple times, even concurrently. All such calls will
	// block until all goroutines have terminated.
	Shutdown()

	// WaitForCacheSync blocks until all started informers' caches were synced
	// or the stop channel gets closed.
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	// ForResource gives generic access to a shared informer of the matching type.
	ForResource(resource pkgruntimeschema.GroupVersionResource) (GenericInformer, error)

	// InternalInformerFor returns the SharedIndexInformer for obj using an internal
	// client.
	InformerFor(obj apimachinerypkgruntime.Object, newFunc clientgoinformersinternalinterfaces.NewInformerFunc) clientgotoolscache.SharedIndexInformer

	Admissionregistration() clientgoinformersadmissionregistration.Interface
	Internal() clientgoinformersapiserverinternal.Interface
	Apps() clientgoinformersapps.Interface
	Autoscaling() clientgoinformersautoscaling.Interface
	Batch() clientgoinformersbatch.Interface
	Certificates() clientgoinformerscertificates.Interface
	Coordination() clientgoinformerscoordination.Interface
	Core() clientgoinformerscore.Interface
	Discovery() clientgoinformersdiscovery.Interface
	Events() clientgoinformersevents.Interface
	Extensions() clientgoinformersextensions.Interface
	Flowcontrol() clientgoinformersflowcontrol.Interface
	Networking() clientgoinformersnetworking.Interface
	Node() clientgoinformersnode.Interface
	Policy() clientgoinformerspolicy.Interface
	Rbac() clientgoinformersrbac.Interface
	Scheduling() clientgoinformersscheduling.Interface
	Storage() clientgoinformersstorage.Interface
}

func (f *sharedInformerFactory) Admissionregistration() clientgoinformersadmissionregistration.Interface {
	return clientgoinformersadmissionregistration.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Internal() clientgoinformersapiserverinternal.Interface {
	return clientgoinformersapiserverinternal.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apps() clientgoinformersapps.Interface {
	return clientgoinformersapps.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Autoscaling() clientgoinformersautoscaling.Interface {
	return clientgoinformersautoscaling.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Batch() clientgoinformersbatch.Interface {
	return clientgoinformersbatch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Certificates() clientgoinformerscertificates.Interface {
	return clientgoinformerscertificates.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Coordination() clientgoinformerscoordination.Interface {
	return clientgoinformerscoordination.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Core() clientgoinformerscore.Interface {
	return clientgoinformerscore.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Discovery() clientgoinformersdiscovery.Interface {
	return clientgoinformersdiscovery.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Events() clientgoinformersevents.Interface {
	return clientgoinformersevents.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Extensions() clientgoinformersextensions.Interface {
	return clientgoinformersextensions.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Flowcontrol() clientgoinformersflowcontrol.Interface {
	return clientgoinformersflowcontrol.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Networking() clientgoinformersnetworking.Interface {
	return clientgoinformersnetworking.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Node() clientgoinformersnode.Interface {
	return clientgoinformersnode.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Policy() clientgoinformerspolicy.Interface {
	return clientgoinformerspolicy.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Rbac() clientgoinformersrbac.Interface {
	return clientgoinformersrbac.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Scheduling() clientgoinformersscheduling.Interface {
	return clientgoinformersscheduling.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Storage() clientgoinformersstorage.Interface {
	return clientgoinformersstorage.New(f, f.namespace, f.tweakListOptions)
}
