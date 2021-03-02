# Copyright (c) 2020 Gitpod GmbH. All rights reserved.
# Licensed under the MIT License. See License-MIT.txt in the project root for license information.

hostname: ${hostname}
installPodSecurityPolicies: true
imagePrefix: ${image_prefix}
license: ${license}
version: ${version}

components:
  wsDaemon:
    containerRuntime:
      nodeRoots:
      - /var/lib
      - /run/containerd/io.containerd.runtime.v1.linux/moby
