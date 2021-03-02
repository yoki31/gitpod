variable "zone_name" {
    type = string
}

variable "subdomain" {
  type    = string
  default = "gitpod"
}

variable "dns_prefixes" {
  type    = list(string)
  default = ["", "*"]
}

variable "gitpod" {
  type = object({
    namespace = string
    shortname = string
  })
  default = {
    namespace = "default"
    shortname = ""
  }
}

variable "vpc_id" {
    type = string
}

variable "subnet_ids" {
    type = list(string)
}
