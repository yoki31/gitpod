/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

variable "certmanager" {
  type = object({
    name       = string
    namespace  = string
    chart      = string
    repository = string
    crds_url   = string
    crds       = bool
  })
  default = {
    name       = "certmanger"
    namespace  = "certmanager"
    chart      = "cert-manager"
    version    = "v1.1.0"
    repository = "https://charts.jetstack.io"
    crds_url   = "https://github.com/jetstack/cert-manager/releases/download/v1.1.0/cert-manager.yaml"
    crds       = true
  }
}

variable "gitpod-node-arn" {
  type = string
}

variable "region" {
  type = string
}

variable "email" {
  type = string
}

variable "zone_name" {
  type = string
}

variable "subdomain" {
  type    = string
  default = "gitpod"
}

variable "shortname" {
  type    = string
  default = ""
}

variable "certificate" {
  type = object({
    name      = string
    namespace = string
  })
  default = {
    name      = "gitpod-certificate"
    namespace = "default"
  }
}
