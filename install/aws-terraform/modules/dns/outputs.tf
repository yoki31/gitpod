/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

output "values" {
  value = templatefile("${path.module}/templates/values.tpl", {
    hostname  = local.hostname
    shortname = var.gitpod.shortname
    eip       = join(",", aws_eip.gitpod.*.id)
  })
}