#output "zones" {
#  value = [for az in aws_instance.app_server : upper(az.availability_zone) if az.availability_zone != "eu-central-1b"]
#}
#
#output "amis" {
#  value = [for ami in aws_instance.app_server : ami.ami]
#}


#output "vpc_public_subnets" {
#  description = "IDs of the VPC's public subnets"
#  value       = module.vpc.public_subnets
#}
#
#output "ec2_instance_public_ips" {
#  description = "Public IP addresses of EC2 instances"
#  value       = module.ec2_instances[*].public_ip
#}
#
#output "vpc_private_subnets" {
#  description = "IDs of the VPC's private subnets"
#  value       = module.vpc.private_subnets
#}


output "cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value       = module.eks.cluster_endpoint
}

output "cluster_security_group_id" {
  description = "Security group ids attached to the cluster control plane"
  value       = module.eks.cluster_security_group_id
}

output "region" {
  description = "AWS region"
  value       = var.region
}

output "cluster_name" {
  description = "Kubernetes Cluster Name"
  value       = module.eks.cluster_name
}
