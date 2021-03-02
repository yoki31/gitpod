# Copyright (c) 2020 Gitpod GmbH. All rights reserved.
# Licensed under the MIT License. See License-MIT.txt in the project root for license information.

ingressMode: hosts
hostname: ${hostname}
installation:
  shortname: ${shortname}
components:
  proxy:
    serviceType: LoadBalancer
    serviceAnnotations:
        service.beta.kubernetes.io/aws-load-balancer-type: nlb
        service.beta.kubernetes.io/aws-load-balancer-eip-allocations: ${eip}


branding:
  homepage: ${hostname}
  redirectUrlIfNotAuthenticated: /workspaces/
  redirectUrlAfterLogout: ${hostname}
