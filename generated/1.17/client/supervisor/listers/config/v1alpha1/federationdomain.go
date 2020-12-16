// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.pinniped.dev/generated/1.17/apis/supervisor/config/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederationDomainLister helps list FederationDomains.
type FederationDomainLister interface {
	// List lists all FederationDomains in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederationDomain, err error)
	// FederationDomains returns an object that can list and get FederationDomains.
	FederationDomains(namespace string) FederationDomainNamespaceLister
	FederationDomainListerExpansion
}

// federationDomainLister implements the FederationDomainLister interface.
type federationDomainLister struct {
	indexer cache.Indexer
}

// NewFederationDomainLister returns a new FederationDomainLister.
func NewFederationDomainLister(indexer cache.Indexer) FederationDomainLister {
	return &federationDomainLister{indexer: indexer}
}

// List lists all FederationDomains in the indexer.
func (s *federationDomainLister) List(selector labels.Selector) (ret []*v1alpha1.FederationDomain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederationDomain))
	})
	return ret, err
}

// FederationDomains returns an object that can list and get FederationDomains.
func (s *federationDomainLister) FederationDomains(namespace string) FederationDomainNamespaceLister {
	return federationDomainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederationDomainNamespaceLister helps list and get FederationDomains.
type FederationDomainNamespaceLister interface {
	// List lists all FederationDomains in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederationDomain, err error)
	// Get retrieves the FederationDomain from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederationDomain, error)
	FederationDomainNamespaceListerExpansion
}

// federationDomainNamespaceLister implements the FederationDomainNamespaceLister
// interface.
type federationDomainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederationDomains in the indexer for a given namespace.
func (s federationDomainNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederationDomain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederationDomain))
	})
	return ret, err
}

// Get retrieves the FederationDomain from the indexer for a given namespace and name.
func (s federationDomainNamespaceLister) Get(name string) (*v1alpha1.FederationDomain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federationdomain"), name)
	}
	return obj.(*v1alpha1.FederationDomain), nil
}
