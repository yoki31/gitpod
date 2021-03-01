/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

locals {
  kubernetes = {
    cluster_name   = "gitpod${var.project == "" ? "" : "-${var.project}"}"
    version        = "1.17"
    min_node_count = 1
    max_node_count = 3
    instance_type  = "m4.large"
  }
  vpc = {
    name = "gitpod${var.project == "" ? "" : "-${var.project}"}"
  }
  config_output_path = pathexpand("~/.kube/config")
  gitpod = {
    namespace   = "default"
    valuesFiles = []
  }
}

module "kubernetes" {
  source = "../../modules/kubernetes"

  providers = {
    aws = aws
  }
}

module "certmanager" {
  source = "../../modules/certmanager"

  zone_name = var.dns_zone_name
  email   = var.certificate_email
  gitpod-node-arn = module.kubernetes.worker_iam_role_arn
  region = var.region

  providers = {
    aws     = aws
    kubernetes = kubernetes
    helm       = helm
    kubectl    = kubectl
  }

  depends_on = [
    module.kubernetes,
  ]
}


module "gitpod" {
  source = "../../modules/gitpod"

  values             = file("values.yaml")
  domain = var.dns_zone_name
  certificate_values = module.certmanager.values
  gitpod = {
    repository   = null
    chart        = "/workspace/gitpod/chart"
    version      = "0.7.0-beta1"
    image_prefix = "gcr.io/gitpod-io/self-hosted/"
  }
  depends_on = [
    module.kubernetes,
  ]
}

module "dns" {
    source = "../../modules/dns"

    loadbalancer = module.gitpod.loadbalancer    
    zone_name = var.dns_zone_name
    vpc_id = module.kubernetes.vpc_id
    subnet_ids = module.kubernetes.subnet_ids
}
