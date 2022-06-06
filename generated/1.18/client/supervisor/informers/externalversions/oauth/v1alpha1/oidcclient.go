// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	oauthv1alpha1 "go.pinniped.dev/generated/1.18/apis/supervisor/oauth/v1alpha1"
	versioned "go.pinniped.dev/generated/1.18/client/supervisor/clientset/versioned"
	internalinterfaces "go.pinniped.dev/generated/1.18/client/supervisor/informers/externalversions/internalinterfaces"
	v1alpha1 "go.pinniped.dev/generated/1.18/client/supervisor/listers/oauth/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OIDCClientInformer provides access to a shared informer and lister for
// OIDCClients.
type OIDCClientInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OIDCClientLister
}

type oIDCClientInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOIDCClientInformer constructs a new informer for OIDCClient type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOIDCClientInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOIDCClientInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOIDCClientInformer constructs a new informer for OIDCClient type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOIDCClientInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OauthV1alpha1().OIDCClients(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OauthV1alpha1().OIDCClients(namespace).Watch(context.TODO(), options)
			},
		},
		&oauthv1alpha1.OIDCClient{},
		resyncPeriod,
		indexers,
	)
}

func (f *oIDCClientInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOIDCClientInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *oIDCClientInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&oauthv1alpha1.OIDCClient{}, f.defaultInformer)
}

func (f *oIDCClientInformer) Lister() v1alpha1.OIDCClientLister {
	return v1alpha1.NewOIDCClientLister(f.Informer().GetIndexer())
}
