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
    name  = string
  })
  default = {
    owner = "099720109477"
    name  = "ubuntu-eks/k8s_1.18/images/*"
  }

}


variable "node_group" {
  type = object({
    name          = string
    instance_type = string
    desired_size  = number
    min_size      = number
    max_size      = number
  })
  default = {
    name          = "gitpod-cluster-nodegroup1"
    instance_type = "m4.large"
    desired_size  = 1
    min_size      = 1
    max_size      = 3
  }
}