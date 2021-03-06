/*
Copyright 2019 The Knative Authors

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
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	servingv1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"
	versioned "knative.dev/serving/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/serving/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "knative.dev/serving/pkg/client/listers/serving/v1beta1"
)

// ConfigurationInformer provides access to a shared informer and lister for
// Configurations.
type ConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ConfigurationLister
}

type configurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewConfigurationInformer constructs a new informer for Configuration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredConfigurationInformer constructs a new informer for Configuration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1beta1().Configurations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1beta1().Configurations(namespace).Watch(options)
			},
		},
		&servingv1beta1.Configuration{},
		resyncPeriod,
		indexers,
	)
}

func (f *configurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servingv1beta1.Configuration{}, f.defaultInformer)
}

func (f *configurationInformer) Lister() v1beta1.ConfigurationLister {
	return v1beta1.NewConfigurationLister(f.Informer().GetIndexer())
}
