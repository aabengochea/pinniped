---
title: "Pinniped Documentation"
cascade:
  layout: docs
menu:
  docs:
    name: Overview
    weight: 1
---

![Pinniped Logo](/docs/img/pinniped_logo.svg)

## Overview

Pinniped provides identity services to Kubernetes.

Pinniped allows cluster administrators to easily plug in external identity
providers (IDPs) into Kubernetes clusters. This is achieved via a uniform
install procedure across all types and origins of Kubernetes clusters,
declarative configuration via Kubernetes APIs, enterprise-grade integrations
with IDPs, and distribution-specific integration strategies.

### Example Use Cases

* Your team uses a large enterprise IDP, and has many clusters that they
  manage. Pinniped provides:
  * Seamless and robust integration with the IDP
  * Easy installation across clusters of any type and origin
  * A simplified login flow across all clusters
* Your team shares a single cluster. Pinniped provides:
  * Simple configuration to integrate an IDP
  * Individual, revocable identities

### Architecture

Pinniped offers credential exchange to enable a user to exchange an external IDP
credential for a short-lived, cluster-specific credential. Pinniped supports various
IDP types and implements different integration strategies for various Kubernetes
distributions to make authentication possible.

To learn more, see [docs/architecture](/docs/architecture).

<img src="docs/img/pinniped_architecture_concierge_supervisor.svg" alt="Pinniped Architecture Sketch" width="300px"/>

## Trying Pinniped

Care to kick the tires? It's easy to [install and try Pinniped](/docs/demo).

## Discussion

Got a question, comment, or idea? Please don't hesitate to reach out via the GitHub [Discussions](https://github.com/vmware-tanzu/pinniped/discussions) tab at the top of this page.

## Contributions

Contributions are welcome. Before contributing, please see the [contributing guide](https://github.com/vmware-tanzu/pinniped/blob/main/CONTRIBUTING.md).

## Reporting Security Vulnerabilities

Please follow the procedure described in [SECURITY.md](https://github.com/vmware-tanzu/pinniped/blob/main/SECURITY.md).

## License

Pinniped is open source and licensed under Apache License Version 2.0. See [LICENSE](https://github.com/vmware-tanzu/pinniped/blob/main/LICENSE).

Copyright 2020 the Pinniped contributors. All Rights Reserved.
