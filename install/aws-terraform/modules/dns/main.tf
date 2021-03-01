/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */


locals {
  hostname = trimsuffix("${var.subdomain}.${data.aws_route53_zone.gitpod.name}", ".")
  shortname   = trimsuffix("ws-${var.gitpod.shortname}", "-")
}

data "aws_route53_zone" "gitpod" {
  name         = var.zone_name
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/elb
data "aws_elb" "gitpod" {
  name = substr(var.loadbalancer,0,32)
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record
resource "aws_route53_record" "gitpod" {
  count = length(var.dns_prefixes)
  zone_id = data.aws_route53_zone.gitpod.zone_id
  name    = trimprefix("${trimsuffix(var.dns_prefixes[count.index], ".")}.${var.subdomain}.${data.aws_route53_zone.gitpod.name}", ".")
  type    = "A"
  alias {
    name                   = data.aws_elb.gitpod.dns_name
    zone_id                = data.aws_elb.gitpod.zone_id
    evaluate_target_health = true
  }
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record
resource "aws_route53_record" "gitpod_ws" {
  zone_id = data.aws_route53_zone.gitpod.zone_id
  name    = "*.${local.shortname}.${var.subdomain}.${data.aws_route53_zone.gitpod.name}"
  type    = "A"
  alias {
    name                   = data.aws_elb.gitpod.dns_name
    zone_id                = data.aws_elb.gitpod.zone_id
    evaluate_target_health = true
  }
}
