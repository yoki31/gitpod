/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */


locals {
  hostname = trim("${var.subdomain}.${data.aws_route53_zone.gitpod.name}", ".")
  ws_shortname   = trim("ws-${var.gitpod.shortname}", "-")
}

data "aws_route53_zone" "gitpod" {
  name         = var.zone_name
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip
resource "aws_eip" "gitpod" {
    count = length(var.subnet_ids)
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record
resource "aws_route53_record" "gitpod" {
  count = length(var.dns_prefixes)
  zone_id = data.aws_route53_zone.gitpod.zone_id
  name    = trimprefix("${trimsuffix(var.dns_prefixes[count.index], ".")}.${var.subdomain}.${data.aws_route53_zone.gitpod.name}", ".")
  type    = "A"
  ttl = "300"
  records = aws_eip.gitpod.*.public_ip
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record
resource "aws_route53_record" "gitpod_ws" {
  zone_id = data.aws_route53_zone.gitpod.zone_id
  name    = "*.${local.ws_shortname}.${var.subdomain}.${data.aws_route53_zone.gitpod.name}"
  type    = "A"
  ttl = "330"
  records = aws_eip.gitpod.*.public_ip
}
