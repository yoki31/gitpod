# Copyright (c) 2020 Gitpod GmbH. All rights reserved.
# Licensed under the MIT License. See License-MIT.txt in the project root for license information.

apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-issuer
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: "${email}"
    privateKeySecretRef:
      name: letsencrypt-key
    solvers:
      - dns01:
          route53:
            region: ${region}
            hostedZoneID: ${zone_id}
            role: "${arn}"
