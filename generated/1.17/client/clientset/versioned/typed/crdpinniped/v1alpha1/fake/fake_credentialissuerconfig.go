/*
Copyright 2020 the Pinniped contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/suzerain-io/pinniped/generated/1.17/apis/crdpinniped/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCredentialIssuerConfigs implements CredentialIssuerConfigInterface
type FakeCredentialIssuerConfigs struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var credentialissuerconfigsResource = schema.GroupVersionResource{Group: "crd.pinniped.dev", Version: "v1alpha1", Resource: "credentialissuerconfigs"}

var credentialissuerconfigsKind = schema.GroupVersionKind{Group: "crd.pinniped.dev", Version: "v1alpha1", Kind: "CredentialIssuerConfig"}

// Get takes name of the credentialIssuerConfig, and returns the corresponding credentialIssuerConfig object, and an error if there is any.
func (c *FakeCredentialIssuerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CredentialIssuerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(credentialissuerconfigsResource, c.ns, name), &v1alpha1.CredentialIssuerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuerConfig), err
}

// List takes label and field selectors, and returns the list of CredentialIssuerConfigs that match those selectors.
func (c *FakeCredentialIssuerConfigs) List(opts v1.ListOptions) (result *v1alpha1.CredentialIssuerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(credentialissuerconfigsResource, credentialissuerconfigsKind, c.ns, opts), &v1alpha1.CredentialIssuerConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CredentialIssuerConfigList{ListMeta: obj.(*v1alpha1.CredentialIssuerConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.CredentialIssuerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested credentialIssuerConfigs.
func (c *FakeCredentialIssuerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(credentialissuerconfigsResource, c.ns, opts))

}

// Create takes the representation of a credentialIssuerConfig and creates it.  Returns the server's representation of the credentialIssuerConfig, and an error, if there is any.
func (c *FakeCredentialIssuerConfigs) Create(credentialIssuerConfig *v1alpha1.CredentialIssuerConfig) (result *v1alpha1.CredentialIssuerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(credentialissuerconfigsResource, c.ns, credentialIssuerConfig), &v1alpha1.CredentialIssuerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuerConfig), err
}

// Update takes the representation of a credentialIssuerConfig and updates it. Returns the server's representation of the credentialIssuerConfig, and an error, if there is any.
func (c *FakeCredentialIssuerConfigs) Update(credentialIssuerConfig *v1alpha1.CredentialIssuerConfig) (result *v1alpha1.CredentialIssuerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(credentialissuerconfigsResource, c.ns, credentialIssuerConfig), &v1alpha1.CredentialIssuerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuerConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCredentialIssuerConfigs) UpdateStatus(credentialIssuerConfig *v1alpha1.CredentialIssuerConfig) (*v1alpha1.CredentialIssuerConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(credentialissuerconfigsResource, "status", c.ns, credentialIssuerConfig), &v1alpha1.CredentialIssuerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuerConfig), err
}

// Delete takes name of the credentialIssuerConfig and deletes it. Returns an error if one occurs.
func (c *FakeCredentialIssuerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(credentialissuerconfigsResource, c.ns, name), &v1alpha1.CredentialIssuerConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCredentialIssuerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(credentialissuerconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CredentialIssuerConfigList{})
	return err
}

// Patch applies the patch and returns the patched credentialIssuerConfig.
func (c *FakeCredentialIssuerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CredentialIssuerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(credentialissuerconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CredentialIssuerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuerConfig), err
}
