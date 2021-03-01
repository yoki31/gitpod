variable "cluster" {
  type = object({
    name         = string
    version      = string
    service_cidr = string
  })
  default = {
    name         = "gitpod-cluster"
    version      = "1.18"
    service_cidr = "172.20.0.0/16"
  }
}

variable "ami" {
    type = object({
        owner = string
        name = string
    })
    default = {
        owner = "099720109477"
        name = "ubuntu-eks/k8s_1.18/images/*"
    }

}