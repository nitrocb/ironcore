/*
 * Copyright (c) 2021 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	time "time"

	networking "github.com/onmetal/onmetal-api/apis/networking"
	clientsetinternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion"
	internalinterfaces "github.com/onmetal/onmetal-api/generated/informers/internalversion/internalinterfaces"
	internalversion "github.com/onmetal/onmetal-api/generated/listers/networking/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkInterfaceBindingInformer provides access to a shared informer and lister for
// NetworkInterfaceBindings.
type NetworkInterfaceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.NetworkInterfaceBindingLister
}

type networkInterfaceBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkInterfaceBindingInformer constructs a new informer for NetworkInterfaceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkInterfaceBindingInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkInterfaceBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkInterfaceBindingInformer constructs a new informer for NetworkInterfaceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkInterfaceBindingInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Networking().NetworkInterfaceBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Networking().NetworkInterfaceBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&networking.NetworkInterfaceBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkInterfaceBindingInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkInterfaceBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkInterfaceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networking.NetworkInterfaceBinding{}, f.defaultInformer)
}

func (f *networkInterfaceBindingInformer) Lister() internalversion.NetworkInterfaceBindingLister {
	return internalversion.NewNetworkInterfaceBindingLister(f.Informer().GetIndexer())
}
