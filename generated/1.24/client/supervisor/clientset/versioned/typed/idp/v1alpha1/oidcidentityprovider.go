// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.pinniped.dev/generated/1.24/apis/supervisor/idp/v1alpha1"
	scheme "go.pinniped.dev/generated/1.24/client/supervisor/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OIDCIdentityProvidersGetter has a method to return a OIDCIdentityProviderInterface.
// A group's client should implement this interface.
type OIDCIdentityProvidersGetter interface {
	OIDCIdentityProviders(namespace string) OIDCIdentityProviderInterface
}

// OIDCIdentityProviderInterface has methods to work with OIDCIdentityProvider resources.
type OIDCIdentityProviderInterface interface {
	Create(ctx context.Context, oIDCIdentityProvider *v1alpha1.OIDCIdentityProvider, opts v1.CreateOptions) (*v1alpha1.OIDCIdentityProvider, error)
	Update(ctx context.Context, oIDCIdentityProvider *v1alpha1.OIDCIdentityProvider, opts v1.UpdateOptions) (*v1alpha1.OIDCIdentityProvider, error)
	UpdateStatus(ctx context.Context, oIDCIdentityProvider *v1alpha1.OIDCIdentityProvider, opts v1.UpdateOptions) (*v1alpha1.OIDCIdentityProvider, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OIDCIdentityProvider, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OIDCIdentityProviderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OIDCIdentityProvider, err error)
	OIDCIdentityProviderExpansion
}

// oIDCIdentityProviders implements OIDCIdentityProviderInterface
type oIDCIdentityProviders struct {
	client rest.Interface
	ns     string
}

// newOIDCIdentityProviders returns a OIDCIdentityProviders
func newOIDCIdentityProviders(c *IDPV1alpha1Client, namespace string) *oIDCIdentityProviders {
	return &oIDCIdentityProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the oIDCIdentityProvider, and returns the corresponding oIDCIdentityProvider object, and an error if there is any.
func (c *oIDCIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OIDCIdentityProvider, err error) {
	result = &v1alpha1.OIDCIdentityProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OIDCIdentityProviders that match those selectors.
func (c *oIDCIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OIDCIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OIDCIdentityProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oIDCIdentityProviders.
func (c *oIDCIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a oIDCIdentityProvider and creates it.  Returns the server's representation of the oIDCIdentityProvider, and an error, if there is any.
func (c *oIDCIdentityProviders) Create(ctx context.Context, oIDCIdentityProvider *v1alpha1.OIDCIdentityProvider, opts v1.CreateOptions) (result *v1alpha1.OIDCIdentityProvider, err error) {
	result = &v1alpha1.OIDCIdentityProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oIDCIdentityProvider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a oIDCIdentityProvider and updates it. Returns the server's representation of the oIDCIdentityProvider, and an error, if there is any.
func (c *oIDCIdentityProviders) Update(ctx context.Context, oIDCIdentityProvider *v1alpha1.OIDCIdentityProvider, opts v1.UpdateOptions) (result *v1alpha1.OIDCIdentityProvider, err error) {
	result = &v1alpha1.OIDCIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		Name(oIDCIdentityProvider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oIDCIdentityProvider).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *oIDCIdentityProviders) UpdateStatus(ctx context.Context, oIDCIdentityProvider *v1alpha1.OIDCIdentityProvider, opts v1.UpdateOptions) (result *v1alpha1.OIDCIdentityProvider, err error) {
	result = &v1alpha1.OIDCIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		Name(oIDCIdentityProvider.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oIDCIdentityProvider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the oIDCIdentityProvider and deletes it. Returns an error if one occurs.
func (c *oIDCIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oIDCIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched oIDCIdentityProvider.
func (c *oIDCIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OIDCIdentityProvider, err error) {
	result = &v1alpha1.OIDCIdentityProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("oidcidentityproviders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
