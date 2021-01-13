---
title: "Pinniped Concierge and Supervisor Demo"
cascade:
  layout: docs
---

# Trying Pinniped Supervisor and Concierge

## Prerequisites

1. A Kubernetes cluster of a type supported by Pinniped Concierge as described in [architecture](/docs/architecture).

   Don't have a cluster handy? Consider using [kind](https://kind.sigs.k8s.io/) on your local machine.
   See below for an example of using kind.

1. A Kubernetes cluster of a type supported by Pinniped Supervisor (this can be the same cluster as the above, or different).

1. A kubeconfig that has admin-like privileges on each cluster.

1. An external OIDC identity provider to use as the source of identity for Pinniped.

## Overview

Installing and trying Pinniped on any cluster will consist of the following general steps. See the next section below
for a more specific example, including the commands to use for that case.

1. Install the Pinniped Supervisor. See [deploy/supervisor/README.md](https://github.com/vmware-tanzu/pinniped/blob/main/deploy/supervisor/README.md).
1. Create a
   [`FederationDomain`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-supervisor-config-v1alpha1-federationdomain)
   via the installed Pinniped Supervisor.
1. Create an
   [`OIDCIdentityProvider`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-supervisor-idp-v1alpha1-oidcidentityprovider)
   via the installed Pinniped Supervisor.
1. Install the Pinniped Concierge. See [deploy/concierge/README.md](https://github.com/vmware-tanzu/pinniped/blob/main/deploy/concierge/README.md).
1. Create a
   [`JWTAuthenticator`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-concierge-authentication-v1alpha1-jwtauthenticator)
   via the installed Pinniped Concierge.
1. Download the Pinniped CLI from [Pinniped's github Releases page](https://github.com/vmware-tanzu/pinniped/releases/latest).
1. Generate a kubeconfig using the Pinniped CLI. Run `pinniped get kubeconfig --help` for more information.
1. Run `kubectl` commands using the generated kubeconfig. The Pinniped Supervisor and Concierge will automatically be used for authentication during those commands.

## Example of Deploying on Multiple kind Clusters

[kind](https://kind.sigs.k8s.io) is a tool for creating and managing Kubernetes clusters on your local machine
which uses Docker containers as the cluster's "nodes". This is a convenient way to try out Pinniped on local
non-production clusters.

The following steps will deploy the latest release of Pinniped on kind. It will deploy the Pinniped
Supervisor on one cluster, and the Pinniped Concierge on another cluster. A multi-cluster deployment
strategy is typical for Pinniped. The Pinniped Concierge will use a
[`JWTAuthenticator`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-concierge-authentication-v1alpha1-jwtauthenticator)
to authenticate federated identities from the Supervisor.

1. Install the tools required for the following steps.

   -  [Install kind](https://kind.sigs.k8s.io/docs/user/quick-start/), if not already installed. e.g. `brew install kind` on MacOS.

   - kind depends on Docker. If not already installed, [install Docker](https://docs.docker.com/get-docker/), e.g. `brew cask install docker` on MacOS.

   - This demo requires `kubectl`, which comes with Docker, or can be [installed separately](https://kubernetes.io/docs/tasks/tools/install-kubectl/).

   - This demo requires `openssl`, which is installed on MacOS by default, or can be [installed separately](https://www.openssl.org/).

1. Create a new Kubernetes cluster for the Pinniped Supervisor using `kind create cluster --name pinniped-supervisor`.

1. Create a new Kubernetes cluster for the Pinniped Concierge using `kind create cluster --name pinniped-concierge`.

1. Deploy the Pinniped Supervisor with a valid serving certificate and network path. See
   [deploy/supervisor/README.md](https://github.com/vmware-tanzu/pinniped/blob/main/deploy/supervisor/README.md).

   For purposes of this demo, the following issuer will be used. This issuer is specific to DNS and
   TLS infrastructure set up for this demo.

   ```bash
   issuer=https://my-supervisor.demo.pinniped.dev
   ```

   This demo uses a `Secret` named `my-federation-domain-tls` to provide the serving certificate for
   the
   [`FederationDomain`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-supervisor-config-v1alpha1-federationdomain). The
   serving certificate `Secret` must be of type `kubernetes.io/tls`.

   The CA bundle for this serving
   certificate is assumed to be written, base64-encoded, to a file named
   `/tmp/pinniped-supervisor-ca-bundle-base64-encoded.pem`.

1. Create a
   [`FederationDomain`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-supervisor-config-v1alpha1-federationdomain)
   object to configure the Pinniped Supervisor to issue federated identities.

   ```bash
   cat <<EOF | kubectl create --context kind-pinniped-supervisor --namespace pinniped-supervisor -f -
   apiVersion: config.supervisor.pinniped.dev/v1alpha1
   kind: FederationDomain
   metadata:
     name: my-federation-domain
   spec:
     issuer: $issuer
     tls:
       secretName: my-federation-domain-tls
   EOF
   ```

1. Create a `Secret` with the external OIDC identity provider OAuth 2.0 client credentials named
   `my-oidc-identity-provider-client` in the pinniped-supervisor namespace.

   ```bash
   kubectl create secret generic my-oidc-identity-provider-client \
     --context kind-pinniped-supervisor \
     --namespace pinniped-supervisor \
     --type secrets.pinniped.dev/oidc-client \
     --from-literal=clientID=xxx \
     --from-literal=clientSecret=yyy
   ```

1. Create an
   [`OIDCIdentityProvider`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-supervisor-idp-v1alpha1-oidcidentityprovider)
   object to configure the Pinniped Supervisor to federate identities from an upstream OIDC identity
   provider.

   Replace the `issuer` with your external identity provider's issuer and
   adjust any other configuration on the spec.

   ```bash
   cat <<EOF | kubectl create --context kind-pinniped-supervisor --namespace pinniped-supervisor -f -
   apiVersion: idp.supervisor.pinniped.dev/v1alpha1
   kind: OIDCIdentityProvider
   metadata:
     name: my-oidc-identity-provider
   spec:
     issuer: https://dev-zzz.okta.com/oauth2/default
     claims:
       username: email
     authorizationConfig:
       additionalScopes: ['email']
     client:
       secretName: my-oidc-identity-provider-client
   EOF
   ```

1. Query GitHub's API for the git tag of the latest Pinniped
   [release](https://github.com/vmware-tanzu/pinniped/releases/latest).

   ```bash
   pinniped_version=$(curl https://api.github.com/repos/vmware-tanzu/pinniped/releases/latest -s | jq .name -r)
   ```

   Alternatively, you can manually select [any release version](https://github.com/vmware-tanzu/pinniped/releases)
   of Pinniped.

   ```bash
   # Example of manually choosing a release version...
   pinniped_version=v0.3.0
   ```

1. Deploy the Pinniped Concierge.

   ```bash
   kubectl apply \
     --context kind-pinniped-concierge \
     -f https://github.com/vmware-tanzu/pinniped/releases/download/$pinniped_version/install-pinniped-concierge.yaml
   ```

   The `install-pinniped-concierge.yaml` file includes the default deployment options.
   If you would prefer to customize the available options, please see [deploy/concierge/README.md](https://github.com/vmware-tanzu/pinniped/blob/main/deploy/concierge/README.md)
   for instructions on how to deploy using `ytt`.

1. Generate a random audience value for this cluster.

   ```bash
   audience="$(openssl rand -hex 8)"
   ```

1. Create a
   [`JWTAuthenticator`](https://github.com/vmware-tanzu/pinniped/blob/main/generated/1.20/README.adoc#k8s-api-go-pinniped-dev-generated-1-19-apis-concierge-authentication-v1alpha1-jwtauthenticator)
   object to configure the Pinniped Concierge to authenticate using the Pinniped Supervisor.

    ```bash
    cat <<EOF | kubectl create --context kind-pinniped-concierge --namespace pinniped-concierge -f -
    apiVersion: authentication.concierge.pinniped.dev/v1alpha1
    kind: JWTAuthenticator
    metadata:
      name: my-jwt-authenticator
    spec:
      issuer: $issuer
      audience: $audience
      tls:
        certificateAuthorityData: $(cat /tmp/pinniped-supervisor-ca-bundle-base64-encoded.pem)
    EOF
    ```
1. Download the latest version of the Pinniped CLI binary for your platform
   from Pinniped's [latest release](https://github.com/vmware-tanzu/pinniped/releases/latest).

1. Move the Pinniped CLI binary to your preferred filename and directory. Add the executable bit,
   e.g. `chmod +x /usr/local/bin/pinniped`.

1. Generate a kubeconfig for the current cluster.
   ```bash
   pinniped get kubeconfig \
     --kubeconfig-context kind-pinniped-concierge \
     --concierge-namespace pinniped-concierge \
     > /tmp/pinniped-kubeconfig
   ```

   If you are using MacOS, you may get an error dialog that says
   `“pinniped” cannot be opened because the developer cannot be verified`. Cancel this dialog, open System Preferences,
   click on Security & Privacy, and click the Allow Anyway button next to the Pinniped message.
   Run the above command again and another dialog will appear saying
   `macOS cannot verify the developer of “pinniped”. Are you sure you want to open it?`.
   Click Open to allow the command to proceed.

1. Try using the generated kubeconfig to issue arbitrary `kubectl` commands. The `pinniped` CLI will
   open a browser page that can be used to login to the external OIDC identity provider configured earlier.

   ```bash
   kubectl --kubeconfig /tmp/pinniped-kubeconfig get pods -n pinniped-concierge
   ```

   Because this user has no RBAC permissions on this cluster, the previous command results in an
   error that is similar to
   `Error from server (Forbidden): pods is forbidden: User "pinny" cannot list resource "pods"
   in API group "" in the namespace "pinniped"`, where `pinny` is the username that was used to login
   to the upstream OIDC identity provider. However, this does prove that you are authenticated and
   acting as the `pinny` user.

1. As the admin user, create RBAC rules for the test user to give them permissions to perform actions on the cluster.
   For example, grant the test user permission to view all cluster resources.

   ```bash
   kubectl --context kind-pinniped-concierge create clusterrolebinding pinny-can-read --clusterrole view --user pinny
   ```

1. Use the generated kubeconfig to issue arbitrary `kubectl` commands as the `pinny` user.

   ```bash
   kubectl --kubeconfig /tmp/pinniped-kubeconfig get pods -n pinniped-concierge
   ```

   The user has permission to list pods, so the command succeeds this time.
   Pinniped has provided authentication into the cluster for your `kubectl` command! 🎉

1. Carry on issuing as many `kubectl` commands as you'd like as the `pinny` user.
   Each invocation will use Pinniped for authentication.
   You may find it convenient to set the `KUBECONFIG` environment variable rather than passing `--kubeconfig` to each invocation.

   ```bash
   export KUBECONFIG=/tmp/pinniped-kubeconfig
   kubectl get namespaces
   kubectl get pods -A
   ```

1. Profit! 💰