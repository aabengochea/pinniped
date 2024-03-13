// Copyright 2024 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"go.pinniped.dev/generated/latest/apis/concierge/authentication/v1alpha1"
	"go.pinniped.dev/test/testlib"
)

func TestConciergeWebhookAuthenticatorStatus_Parallel(t *testing.T) {
	testEnv := testlib.IntegrationEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	t.Cleanup(cancel)

	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "Basic test to see if the WebhookAuthenticator wakes up or not.",
			run: func(t *testing.T) {
				webhookAuthenticator := testlib.CreateTestWebhookAuthenticator(
					ctx,
					t,
					nil,
					v1alpha1.WebhookAuthenticatorPhaseReady)

				testlib.WaitForWebhookAuthenticatorStatusConditions(
					ctx, t,
					webhookAuthenticator.Name,
					allSuccessfulWebhookAuthenticatorConditions())
			},
		}, {
			name: "valid spec with invalid CA in TLS config will result in a WebhookAuthenticator that is not ready",
			run: func(t *testing.T) {
				caBundleString := "invalid base64-encoded data"
				webhookSpec := testEnv.TestWebhook.DeepCopy()
				webhookSpec.TLS = &v1alpha1.TLSSpec{
					CertificateAuthorityData: caBundleString,
				}

				webhookAuthenticator := testlib.CreateTestWebhookAuthenticator(
					ctx,
					t,
					webhookSpec,
					v1alpha1.WebhookAuthenticatorPhaseError)

				testlib.WaitForWebhookAuthenticatorStatusConditions(
					ctx, t,
					webhookAuthenticator.Name,
					replaceSomeConditions(
						allSuccessfulWebhookAuthenticatorConditions(),
						[]metav1.Condition{
							{
								Type:    "Ready",
								Status:  "False",
								Reason:  "NotReady",
								Message: "the WebhookAuthenticator is not ready: see other conditions for details",
							}, {
								Type:    "AuthenticatorValid",
								Status:  "Unknown",
								Reason:  "UnableToValidate",
								Message: "unable to validate; see other conditions for details",
							}, {
								Type:    "TLSConfigurationValid",
								Status:  "False",
								Reason:  "InvalidTLSConfiguration",
								Message: "invalid TLS configuration: illegal base64 data at input byte 7",
							}, {
								Type:    "TLSConnectionNegotiationValid",
								Status:  "Unknown",
								Reason:  "UnableToValidate",
								Message: "unable to validate; see other conditions for details",
							},
						},
					))
			},
		}, {
			name: "valid spec with valid CA in TLS config but does not match issuer server will result in a WebhookAuthenticator that is not ready",
			run: func(t *testing.T) {
				caBundleSomePivotalCA := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURVVENDQWptZ0F3SUJBZ0lWQUpzNStTbVRtaTJXeUI0bGJJRXBXaUs5a1RkUE1BMEdDU3FHU0liM0RRRUIKQ3dVQU1COHhDekFKQmdOVkJBWVRBbFZUTVJBd0RnWURWUVFLREFkUWFYWnZkR0ZzTUI0WERUSXdNRFV3TkRFMgpNamMxT0ZvWERUSTBNRFV3TlRFMk1qYzFPRm93SHpFTE1Ba0dBMVVFQmhNQ1ZWTXhFREFPQmdOVkJBb01CMUJwCmRtOTBZV3d3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRRERZWmZvWGR4Z2NXTEMKZEJtbHB5a0tBaG9JMlBuUWtsVFNXMno1cGcwaXJjOGFRL1E3MXZzMTRZYStmdWtFTGlvOTRZYWw4R01DdVFrbApMZ3AvUEE5N1VYelhQNDBpK25iNXcwRGpwWWd2dU9KQXJXMno2MFRnWE5NSFh3VHk4ME1SZEhpUFVWZ0VZd0JpCmtkNThzdEFVS1Y1MnBQTU1reTJjNy9BcFhJNmRXR2xjalUvaFBsNmtpRzZ5dEw2REtGYjJQRWV3MmdJM3pHZ2IKOFVVbnA1V05DZDd2WjNVY0ZHNXlsZEd3aGc3cnZ4U1ZLWi9WOEhCMGJmbjlxamlrSVcxWFM4dzdpUUNlQmdQMApYZWhKZmVITlZJaTJtZlczNlVQbWpMdnVKaGpqNDIrdFBQWndvdDkzdWtlcEgvbWpHcFJEVm9wamJyWGlpTUYrCkYxdnlPNGMxQWdNQkFBR2pnWU13Z1lBd0hRWURWUjBPQkJZRUZNTWJpSXFhdVkwajRVWWphWDl0bDJzby9LQ1IKTUI4R0ExVWRJd1FZTUJhQUZNTWJpSXFhdVkwajRVWWphWDl0bDJzby9LQ1JNQjBHQTFVZEpRUVdNQlFHQ0NzRwpBUVVGQndNQ0JnZ3JCZ0VGQlFjREFUQVBCZ05WSFJNQkFmOEVCVEFEQVFIL01BNEdBMVVkRHdFQi93UUVBd0lCCkJqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFYbEh4M2tIMDZwY2NDTDlEVE5qTnBCYnlVSytGd2R6T2IwWFYKcmpNaGtxdHVmdEpUUnR5T3hKZ0ZKNXhUR3pCdEtKamcrVU1pczBOV0t0VDBNWThVMU45U2c5SDl0RFpHRHBjVQpxMlVRU0Y4dXRQMVR3dnJIUzIrdzB2MUoxdHgrTEFiU0lmWmJCV0xXQ21EODUzRlVoWlFZekkvYXpFM28vd0p1CmlPUklMdUpNUk5vNlBXY3VLZmRFVkhaS1RTWnk3a25FcHNidGtsN3EwRE91eUFWdG9HVnlkb3VUR0FOdFhXK2YKczNUSTJjKzErZXg3L2RZOEJGQTFzNWFUOG5vZnU3T1RTTzdiS1kzSkRBUHZOeFQzKzVZUXJwNGR1Nmh0YUFMbAppOHNaRkhidmxpd2EzdlhxL3p1Y2JEaHEzQzBhZnAzV2ZwRGxwSlpvLy9QUUFKaTZLQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
				webhookSpec := testEnv.TestWebhook.DeepCopy()
				webhookSpec.TLS = &v1alpha1.TLSSpec{
					CertificateAuthorityData: caBundleSomePivotalCA,
				}

				webhookAuthenticator := testlib.CreateTestWebhookAuthenticator(
					ctx,
					t,
					webhookSpec,
					v1alpha1.WebhookAuthenticatorPhaseError)

				testlib.WaitForWebhookAuthenticatorStatusConditions(
					ctx, t,
					webhookAuthenticator.Name,
					replaceSomeConditions(
						allSuccessfulWebhookAuthenticatorConditions(),
						[]metav1.Condition{
							{
								Type:    "Ready",
								Status:  "False",
								Reason:  "NotReady",
								Message: "the WebhookAuthenticator is not ready: see other conditions for details",
							}, {
								Type:    "AuthenticatorValid",
								Status:  "Unknown",
								Reason:  "UnableToValidate",
								Message: "unable to validate; see other conditions for details",
							}, {
								Type:    "TLSConnectionNegotiationValid",
								Status:  "False",
								Reason:  "UnableToDialServer",
								Message: "cannot dial server: tls: failed to verify certificate: x509: certificate signed by unknown authority",
							},
						},
					))
			},
		}, {
			name: "invalid with unresponsive endpoint will result in a WebhookAuthenticator that is not ready",
			run: func(t *testing.T) {
				caBundleSomePivotalCA := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURVVENDQWptZ0F3SUJBZ0lWQUpzNStTbVRtaTJXeUI0bGJJRXBXaUs5a1RkUE1BMEdDU3FHU0liM0RRRUIKQ3dVQU1COHhDekFKQmdOVkJBWVRBbFZUTVJBd0RnWURWUVFLREFkUWFYWnZkR0ZzTUI0WERUSXdNRFV3TkRFMgpNamMxT0ZvWERUSTBNRFV3TlRFMk1qYzFPRm93SHpFTE1Ba0dBMVVFQmhNQ1ZWTXhFREFPQmdOVkJBb01CMUJwCmRtOTBZV3d3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRRERZWmZvWGR4Z2NXTEMKZEJtbHB5a0tBaG9JMlBuUWtsVFNXMno1cGcwaXJjOGFRL1E3MXZzMTRZYStmdWtFTGlvOTRZYWw4R01DdVFrbApMZ3AvUEE5N1VYelhQNDBpK25iNXcwRGpwWWd2dU9KQXJXMno2MFRnWE5NSFh3VHk4ME1SZEhpUFVWZ0VZd0JpCmtkNThzdEFVS1Y1MnBQTU1reTJjNy9BcFhJNmRXR2xjalUvaFBsNmtpRzZ5dEw2REtGYjJQRWV3MmdJM3pHZ2IKOFVVbnA1V05DZDd2WjNVY0ZHNXlsZEd3aGc3cnZ4U1ZLWi9WOEhCMGJmbjlxamlrSVcxWFM4dzdpUUNlQmdQMApYZWhKZmVITlZJaTJtZlczNlVQbWpMdnVKaGpqNDIrdFBQWndvdDkzdWtlcEgvbWpHcFJEVm9wamJyWGlpTUYrCkYxdnlPNGMxQWdNQkFBR2pnWU13Z1lBd0hRWURWUjBPQkJZRUZNTWJpSXFhdVkwajRVWWphWDl0bDJzby9LQ1IKTUI4R0ExVWRJd1FZTUJhQUZNTWJpSXFhdVkwajRVWWphWDl0bDJzby9LQ1JNQjBHQTFVZEpRUVdNQlFHQ0NzRwpBUVVGQndNQ0JnZ3JCZ0VGQlFjREFUQVBCZ05WSFJNQkFmOEVCVEFEQVFIL01BNEdBMVVkRHdFQi93UUVBd0lCCkJqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFYbEh4M2tIMDZwY2NDTDlEVE5qTnBCYnlVSytGd2R6T2IwWFYKcmpNaGtxdHVmdEpUUnR5T3hKZ0ZKNXhUR3pCdEtKamcrVU1pczBOV0t0VDBNWThVMU45U2c5SDl0RFpHRHBjVQpxMlVRU0Y4dXRQMVR3dnJIUzIrdzB2MUoxdHgrTEFiU0lmWmJCV0xXQ21EODUzRlVoWlFZekkvYXpFM28vd0p1CmlPUklMdUpNUk5vNlBXY3VLZmRFVkhaS1RTWnk3a25FcHNidGtsN3EwRE91eUFWdG9HVnlkb3VUR0FOdFhXK2YKczNUSTJjKzErZXg3L2RZOEJGQTFzNWFUOG5vZnU3T1RTTzdiS1kzSkRBUHZOeFQzKzVZUXJwNGR1Nmh0YUFMbAppOHNaRkhidmxpd2EzdlhxL3p1Y2JEaHEzQzBhZnAzV2ZwRGxwSlpvLy9QUUFKaTZLQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
				webhookSpec := testEnv.TestWebhook.DeepCopy()
				webhookSpec.TLS = &v1alpha1.TLSSpec{
					CertificateAuthorityData: caBundleSomePivotalCA,
				}
				webhookSpec.Endpoint = "https://127.0.0.1:443/some-fake-endpoint"

				webhookAuthenticator := testlib.CreateTestWebhookAuthenticator(
					ctx,
					t,
					webhookSpec,
					v1alpha1.WebhookAuthenticatorPhaseError)

				testlib.WaitForWebhookAuthenticatorStatusConditions(
					ctx, t,
					webhookAuthenticator.Name,
					replaceSomeConditions(
						allSuccessfulWebhookAuthenticatorConditions(),
						[]metav1.Condition{
							{
								Type:    "Ready",
								Status:  "False",
								Reason:  "NotReady",
								Message: "the WebhookAuthenticator is not ready: see other conditions for details",
							}, {
								Type:    "AuthenticatorValid",
								Status:  "Unknown",
								Reason:  "UnableToValidate",
								Message: "unable to validate; see other conditions for details",
							}, {
								Type:    "TLSConnectionNegotiationValid",
								Status:  "False",
								Reason:  "UnableToDialServer",
								Message: "cannot dial server: dial tcp 127.0.0.1:443: connect: connection refused",
							},
						},
					))
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.run(t)
		})
	}
}

func TestConciergeWebhookAuthenticatorCRDValidations_Parallel(t *testing.T) {
	env := testlib.IntegrationEnv(t)
	webhookAuthenticatorClient := testlib.NewConciergeClientset(t).AuthenticationV1alpha1().WebhookAuthenticators()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	t.Cleanup(cancel)

	objectMeta := testlib.ObjectMetaWithRandomName(t, "webhook-authenticator")
	tests := []struct {
		name                 string
		webhookAuthenticator *v1alpha1.WebhookAuthenticator
		wantErr              string
	}{
		{
			name: "endpoint can not be empty string",
			webhookAuthenticator: &v1alpha1.WebhookAuthenticator{
				ObjectMeta: objectMeta,
				Spec: v1alpha1.WebhookAuthenticatorSpec{
					Endpoint: "",
				},
			},
			wantErr: `WebhookAuthenticator.authentication.concierge.` + env.APIGroupSuffix + ` "` + objectMeta.Name + `" is invalid: ` +
				`spec.endpoint: Invalid value: "": spec.endpoint in body should be at least 1 chars long`,
		},
		{
			name: "endpoint must be https",
			webhookAuthenticator: &v1alpha1.WebhookAuthenticator{
				ObjectMeta: objectMeta,
				Spec: v1alpha1.WebhookAuthenticatorSpec{
					Endpoint: "http://www.example.com",
				},
			},
			wantErr: `WebhookAuthenticator.authentication.concierge.` + env.APIGroupSuffix + ` "` + objectMeta.Name + `" is invalid: ` +
				`spec.endpoint: Invalid value: "http://www.example.com": spec.endpoint in body should match '^https://'`,
		},
		{
			name: "minimum valid authenticator",
			webhookAuthenticator: &v1alpha1.WebhookAuthenticator{
				ObjectMeta: testlib.ObjectMetaWithRandomName(t, "webhook"),
				Spec: v1alpha1.WebhookAuthenticatorSpec{
					Endpoint: "https://localhost/webhook-isnt-actually-here",
				},
			},
		},
		{
			name: "valid authenticator can have empty TLS block",
			webhookAuthenticator: &v1alpha1.WebhookAuthenticator{
				ObjectMeta: testlib.ObjectMetaWithRandomName(t, "webhook"),
				Spec: v1alpha1.WebhookAuthenticatorSpec{
					Endpoint: "https://localhost/webhook-isnt-actually-here",
					TLS:      &v1alpha1.TLSSpec{},
				},
			},
		},
		{
			name: "valid authenticator can have empty TLS CertificateAuthorityData",
			webhookAuthenticator: &v1alpha1.WebhookAuthenticator{
				ObjectMeta: testlib.ObjectMetaWithRandomName(t, "jwtauthenticator"),
				Spec: v1alpha1.WebhookAuthenticatorSpec{
					Endpoint: "https://localhost/webhook-isnt-actually-here",
					TLS: &v1alpha1.TLSSpec{
						CertificateAuthorityData: "pretend-this-is-a-certificate",
					},
				},
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, createErr := webhookAuthenticatorClient.Create(ctx, tt.webhookAuthenticator, metav1.CreateOptions{})

			t.Cleanup(func() {
				// delete if it exists
				delErr := webhookAuthenticatorClient.Delete(ctx, tt.webhookAuthenticator.Name, metav1.DeleteOptions{})
				if !errors.IsNotFound(delErr) {
					require.NoError(t, delErr)
				}
			})

			if tt.wantErr != "" {
				wantErr := tt.wantErr
				require.EqualError(t, createErr, wantErr)
			} else {
				require.NoError(t, createErr)
			}
		})
	}
}
func allSuccessfulWebhookAuthenticatorConditions() []metav1.Condition {
	return []metav1.Condition{{
		Type:    "AuthenticatorValid",
		Status:  "True",
		Reason:  "Success",
		Message: "authenticator initialized",
	}, {
		Type:    "EndpointURLValid",
		Status:  "True",
		Reason:  "Success",
		Message: "endpoint is a valid URL",
	}, {
		Type:    "Ready",
		Status:  "True",
		Reason:  "Success",
		Message: "the WebhookAuthenticator is ready",
	}, {
		Type:    "TLSConfigurationValid",
		Status:  "True",
		Reason:  "Success",
		Message: "successfully parsed specified CA bundle",
	}, {
		Type:    "TLSConnectionNegotiationValid",
		Status:  "True",
		Reason:  "Success",
		Message: "tls verified",
	}}
}
