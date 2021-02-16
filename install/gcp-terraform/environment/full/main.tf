/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

locals {
  region = trimsuffix(var.location,local.zone_suffix)
  zone_suffix = regex("-[a-z]$",var.location)
}

resource "google_compute_network" "gitpod" {
  name                    = "gitpod"
  description             = "Gitpod Cluster Network"
  auto_create_subnetworks = false
  project                 = var.project
}

module "kubernetes" {
  source = "../../modules/kubernetes"

  name    = "gitpod"
  network = google_compute_network.gitpod.name
  project = var.project
  location  = var.location
}


module "dns" {
  source = "../../modules/dns"

  project   = var.project
  location  = var.location
  zone_name = var.zone_name
  name      = "gitpod-dns"
  subdomain = var.subdomain

  providers = {
    google     = google
    kubernetes = kubernetes
  }
}


module "certmanager" {
  source = "../../modules/certmanager"

  project = var.project
  email   = var.certificate_email
  domain  = module.dns.hostname

  providers = {
    google     = google
    kubernetes = kubernetes
    helm       = helm
    kubectl    = kubectl
  }
}

module "registry" {
  source = "../../modules/registry"

  name     = var.subdomain
  project  = var.project
  location = var.container_registry.location

  providers = {
    google     = google
    kubernetes = kubernetes
  }
}


module "storage" {
  source = "../../modules/storage"

  name     = var.subdomain
  project  = var.project
  region   = local.region
  location = "EU"
}

# module "database" {
#   source = "../../modules/database"

#   project = var.project
#   name    = var.database.name
#   region  = local.region
#   network = {
#     id   = google_compute_network.gitpod.id
#     name = google_compute_network.gitpod.name
#   }
# }

#
# Gitpod
#

module "gitpod" {
  source = "../../modules/gitpod"

  project            = var.project
  namespace          = var.namespace
  values             = file("values.yaml")
  dns_values         = module.dns.values
  certificate_values = module.certmanager.values
  # database_values    = module.database.values
  registry_values    = module.registry.values
  storage_values     = module.storage.values
  license            = var.license

  gitpod = {
    repository   = var.gitpod_repository
    chart        = var.gitpod_chart
    version      = var.gitpod_version
    image_prefix = "gcr.io/gitpod-io/self-hosted/"
  }

  providers = {
    google     = google
    kubernetes = kubernetes
    helm       = helm
  }
}
