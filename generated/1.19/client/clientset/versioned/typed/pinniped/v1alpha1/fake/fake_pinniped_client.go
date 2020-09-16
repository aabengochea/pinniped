/*
Copyright 2020 the Pinniped contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/suzerain-io/pinniped/generated/1.19/client/clientset/versioned/typed/pinniped/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePinnipedV1alpha1 struct {
	*testing.Fake
}

func (c *FakePinnipedV1alpha1) CredentialRequests() v1alpha1.CredentialRequestInterface {
	return &FakeCredentialRequests{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePinnipedV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
