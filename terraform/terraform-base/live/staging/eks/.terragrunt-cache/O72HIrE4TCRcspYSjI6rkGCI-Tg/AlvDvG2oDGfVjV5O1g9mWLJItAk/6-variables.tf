variable "env" {
  description = "Environment name."
  type        = string
}

variable "eks_version" {
  description = "Desired Kubernetes master version."
  type        = string
}

variable "eks_name" {
  description = "Name of the cluster."
  type        = string
}

variable "subnet_ids" {
  description = "List of subnet IDs. Must be in at least two different availability zones."
  type        = list(string)
}

variable "node_iam_policies" {
  description = "List of IAM Policies to attach to EKS-managed nodes."
  type        = map(any)
  default = {
    1 = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
    2 = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
    3 = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
    4 = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
  }
}

variable "node_groups" {
  description = "EKS node groups"
  type        = map(any)
}

variable "enable_irsa" {
  description = "Determines whether to create an OpenID Connect Provider for EKS to enable IRSA"
  type        = bool
  default     = true
}

variable "route53_zone_name" {
  description = "Route53 zone name"
  type        = string
}

variable "route53_record_name" {
  description = "Route53 record name"
  type        = string
}

# variable "s3_access_manifest" {
#   description = "S3 manifest path"
#   type        = string
# }

variable "ingress_manifest" {
  description = "Ingress manifest path"
  type        = string 
}

variable "cert_issuer_manifest" {
  description = "Cert issuer manifest path"
  type        = string
}

variable "s3_bucket_name" {
  description = "S3 bucket name"
  type        = string
}