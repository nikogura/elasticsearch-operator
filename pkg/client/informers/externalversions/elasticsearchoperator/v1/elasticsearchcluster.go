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
	time "time"

	elasticsearchoperatorv1 "github.com/upmc-enterprises/elasticsearch-operator/pkg/apis/elasticsearchoperator/v1"
	versioned "github.com/upmc-enterprises/elasticsearch-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/upmc-enterprises/elasticsearch-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/upmc-enterprises/elasticsearch-operator/pkg/client/listers/elasticsearchoperator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ElasticsearchClusterInformer provides access to a shared informer and lister for
// ElasticsearchClusters.
type ElasticsearchClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ElasticsearchClusterLister
}

type elasticsearchClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewElasticsearchClusterInformer constructs a new informer for ElasticsearchCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewElasticsearchClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredElasticsearchClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredElasticsearchClusterInformer constructs a new informer for ElasticsearchCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredElasticsearchClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EnterprisesV1().ElasticsearchClusters(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EnterprisesV1().ElasticsearchClusters(namespace).Watch(options)
			},
		},
		&elasticsearchoperatorv1.ElasticsearchCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *elasticsearchClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredElasticsearchClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *elasticsearchClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&elasticsearchoperatorv1.ElasticsearchCluster{}, f.defaultInformer)
}

func (f *elasticsearchClusterInformer) Lister() v1.ElasticsearchClusterLister {
	return v1.NewElasticsearchClusterLister(f.Informer().GetIndexer())
}
