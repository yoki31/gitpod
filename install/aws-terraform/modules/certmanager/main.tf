/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

# https://cert-manager.io/docs/configuration/acme/dns01/route53/


resource "random_id" "certmanager" {
    byte_length = 4
}

# 
resource "aws_iam_role" "certmanager" {
  name = "certmanager-role-${random_id.certmanager.hex}"

  assume_role_policy = <<-EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "${var.gitpod-node-arn}"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}



#
# AWS IAM role policy 'dns-manager-role-policy'
# allow 'dns-manager-role' to configure Route53 txt records
#
resource "aws_iam_role_policy" "certmanager" {
  name = "certmanager-role-policy-${random_id.certmanager.hex}"

  policy = <<-EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "route53:GetChange",
      "Resource": "arn:aws:route53:::change/*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ChangeResourceRecordSets",
        "route53:ListResourceRecordSets"
      ],
      "Resource": "arn:aws:route53:::hostedzone/*"
    },
    {
      "Effect": "Allow",
      "Action": "route53:ListHostedZonesByName",
      "Resource": "*"
    }
  ]
}
EOF
  role   = aws_iam_role.certmanager.name
}

#
# Kubernetes Resources
#

resource "kubernetes_namespace" "certmanager" {
  provider = kubernetes
  metadata {
    name = var.certmanager.namespace
  }
}


# https://registry.terraform.io/providers/hashicorp/helm/latest/docs/resources/release
resource "helm_release" "certmanager" {
  name       = var.certmanager.name
  chart      = var.certmanager.chart
  repository = var.certmanager.repository

  namespace        = kubernetes_namespace.certmanager.metadata[0].name
  create_namespace = false

  wait = true

  set {
    name  = "installCRDs"
    value = var.certmanager.crds
  }
}



#
# Cluster Issuer
#

# https://registry.terraform.io/providers/hashicorp/time/latest/docs/resources/sleep
# waits for CRDS to be installed

resource "time_sleep" "certmanager" {
  create_duration = "300s"

  depends_on = [
    helm_release.certmanager
  ]
}


data "aws_route53_zone" "gitpod" {
  name         = var.zone_name
}


resource "kubectl_manifest" "cluster_issuer" {
  yaml_body = templatefile("${path.module}/templates/clusterissuer.tpl", {
    email = var.email
    region = var.region
    arn = aws_iam_role.certmanager.arn
    zone_id = data.aws_route53_zone.gitpod.zone_id
  })

  depends_on = [
    time_sleep.certmanager,
  ]
}

#
# Certificate
#

locals {
  shortname = trimsuffix("ws-${var.shortname}", "-")
}

resource "kubectl_manifest" "certificate" {
  validate_schema = false

  yaml_body = templatefile("${path.module}/templates/certificate.tpl", {
    name      = var.certificate.name
    namespace = var.certificate.namespace
    domain    = trimprefix("${var.subdomain}.${data.aws_route53_zone.gitpod.name}",".")
    shortname = local.shortname
  })

  depends_on = [
    kubectl_manifest.cluster_issuer,
  ]
}

data "template_file" "values" {
  template = file("${path.module}/templates/values.tpl")

  vars = {
    secret_name     = var.certificate.name
    key_name        = "tls.key"
    chain_name      = "tls.crt"
    full_chain_name = "tls.crt"
  }
}
