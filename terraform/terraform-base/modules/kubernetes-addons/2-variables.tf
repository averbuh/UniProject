variable "env" {
  description = "Environment name."
  type        = string
}

variable "eks_name" {
  description = "Name of the cluster."
  type        = string
}

variable "enable_ebs" {
  description = "Determines whether to deploy ebs addon"
  type = bool
  default = false
  
}

variable "provider_url" {
  type = string 
}

variable "enable_cluster_autoscaler" {
  description = "Determines whether to deploy cluster autoscaler"
  type        = bool
  default     = false
}

variable "cluster_autoscaler_helm_verion" {
  description = "Cluster Autoscaler Helm verion"
  type        = string
}

variable "openid_provider_arn" {
  description = "IAM Openid Connect Provider ARN"
  type        = string
}
# variable "s3_access_manifest" {
#   description = "S3 manifest path"
#   type        = string
# }

# variable "ingress_manifest" {
#   description = "Ingress manifest path"
#   type        = string 
# }

# variable "cert_issuer_manifest" {
#   description = "Cert issuer manifest path"
#   type        = string
# }

