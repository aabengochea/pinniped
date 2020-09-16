/*
Copyright 2020 the Pinniped contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/suzerain-io/pinniped/generated/1.17/apis/crdpinniped/v1alpha1"
	scheme "github.com/suzerain-io/pinniped/generated/1.17/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CredentialIssuerConfigsGetter has a method to return a CredentialIssuerConfigInterface.
// A group's client should implement this interface.
type CredentialIssuerConfigsGetter interface {
	CredentialIssuerConfigs(namespace string) CredentialIssuerConfigInterface
}

// CredentialIssuerConfigInterface has methods to work with CredentialIssuerConfig resources.
type CredentialIssuerConfigInterface interface {
	Create(*v1alpha1.CredentialIssuerConfig) (*v1alpha1.CredentialIssuerConfig, error)
	Update(*v1alpha1.CredentialIssuerConfig) (*v1alpha1.CredentialIssuerConfig, error)
	UpdateStatus(*v1alpha1.CredentialIssuerConfig) (*v1alpha1.CredentialIssuerConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CredentialIssuerConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.CredentialIssuerConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CredentialIssuerConfig, err error)
	CredentialIssuerConfigExpansion
}

// credentialIssuerConfigs implements CredentialIssuerConfigInterface
type credentialIssuerConfigs struct {
	client rest.Interface
	ns     string
}

// newCredentialIssuerConfigs returns a CredentialIssuerConfigs
func newCredentialIssuerConfigs(c *CrdV1alpha1Client, namespace string) *credentialIssuerConfigs {
	return &credentialIssuerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the credentialIssuerConfig, and returns the corresponding credentialIssuerConfig object, and an error if there is any.
func (c *credentialIssuerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CredentialIssuerConfig, err error) {
	result = &v1alpha1.CredentialIssuerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CredentialIssuerConfigs that match those selectors.
func (c *credentialIssuerConfigs) List(opts v1.ListOptions) (result *v1alpha1.CredentialIssuerConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CredentialIssuerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested credentialIssuerConfigs.
func (c *credentialIssuerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a credentialIssuerConfig and creates it.  Returns the server's representation of the credentialIssuerConfig, and an error, if there is any.
func (c *credentialIssuerConfigs) Create(credentialIssuerConfig *v1alpha1.CredentialIssuerConfig) (result *v1alpha1.CredentialIssuerConfig, err error) {
	result = &v1alpha1.CredentialIssuerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		Body(credentialIssuerConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a credentialIssuerConfig and updates it. Returns the server's representation of the credentialIssuerConfig, and an error, if there is any.
func (c *credentialIssuerConfigs) Update(credentialIssuerConfig *v1alpha1.CredentialIssuerConfig) (result *v1alpha1.CredentialIssuerConfig, err error) {
	result = &v1alpha1.CredentialIssuerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		Name(credentialIssuerConfig.Name).
		Body(credentialIssuerConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *credentialIssuerConfigs) UpdateStatus(credentialIssuerConfig *v1alpha1.CredentialIssuerConfig) (result *v1alpha1.CredentialIssuerConfig, err error) {
	result = &v1alpha1.CredentialIssuerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		Name(credentialIssuerConfig.Name).
		SubResource("status").
		Body(credentialIssuerConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the credentialIssuerConfig and deletes it. Returns an error if one occurs.
func (c *credentialIssuerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *credentialIssuerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched credentialIssuerConfig.
func (c *credentialIssuerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CredentialIssuerConfig, err error) {
	result = &v1alpha1.CredentialIssuerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("credentialissuerconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
