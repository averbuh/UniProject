#variable "image_id" {
#  type = string
#}
#
#variable "instance_type" {
#  type = string
#}
#
#variable "availability_zone" {
#  type = list(string)
#  default = ["eu-central-1a"]
#}

variable "vpc_name" {
  description = "Name of VPC"
  type        = string
  default     = "example-vpc"
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "vpc_azs" {
  description = "Availability zones for VPC"
  type        = list(string)
  default     = ["us-west-2a", "us-west-2b", "us-west-2c"]
}

variable "vpc_private_subnets" {
  description = "Private subnets for VPC"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "vpc_public_subnets" {
  description = "Public subnets for VPC"
  type        = list(string)
  default     = ["10.0.101.0/24", "10.0.102.0/24"]
}

variable "vpc_enable_nat_gateway" {
  description = "Enable NAT gateway for VPC"
  type        = bool
  default     = true
}

variable "vpc_tags" {
  description = "Tags to apply to resources created by VPC module"
  type        = map(string)
  default = {
    Terraform   = "true"
    Environment = "dev"
  }
}

variable "region" {
  description = "AWS region"
  type        = string
  default     = "eu-central-1"
}

variable "control_plane_subnets_ids" {
  description = "IDs of the VPC's control plane subnets"
  type        = list(string)
  default     = ["subnet-12345678", "subnet-87654321"]
}

variable "working_subnets_ids" {
  description = "IDs of the VPC's working subnets"
  type        = list(string)
  default     = ["subnet-12345678", "subnet-87654321"]
}



