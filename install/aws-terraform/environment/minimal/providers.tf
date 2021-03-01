/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

# https://registry.terraform.io/providers/hashicorp/aws/latest
provider "aws" {
  # Load from env: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_DEFAULT_REGION
  region = var.region
}

# https://registry.terraform.io/providers/hashicorp/kubernetes/latest
provider "kubernetes" {
  host                   = module.kubernetes.cluster.endpoint
  cluster_ca_certificate = base64decode(module.kubernetes.cluster.certificate_authority.0.data)
  token                  = module.kubernetes.cluster_auth.token
}

# https://registry.terraform.io/providers/hashicorp/helm/latest
provider "helm" {
  kubernetes {
    host                   = module.kubernetes.cluster.endpoint
    cluster_ca_certificate = base64decode(module.kubernetes.cluster.certificate_authority.0.data)
    token                  = module.kubernetes.cluster_auth.token
  }
}

# https://registry.terraform.io/providers/gavinbunney/kubectl/latest
provider "kubectl" {
  host                   = module.kubernetes.cluster.endpoint
  cluster_ca_certificate = base64decode(module.kubernetes.cluster.certificate_authority.0.data)
  token                  = module.kubernetes.cluster_auth.token
}
