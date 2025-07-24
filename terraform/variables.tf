variable "cluster_name" {
  type = string
}

variable "cluster_driver" {
  type = string
}

variable "cluster_version" {
  type = string
}

variable "cluster_addons" {
  type = list(string)
}

variable "helm_releases" {
  type = list(object({
    name       = string
    namespace  = string
    chart      = string
    repository = string
    version    = string
    values     = string
  }))
}