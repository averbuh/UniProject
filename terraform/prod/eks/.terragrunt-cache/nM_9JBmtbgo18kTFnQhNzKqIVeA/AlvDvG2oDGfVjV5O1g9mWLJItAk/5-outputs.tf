output "eks_name" {
  value = aws_eks_cluster.this.name
}

output "cluster_oidc_issuer_url" {
  value = try(aws_eks_cluster.this.identity[0].oidc[0].issuer, null)
} 

output "openid_provider_arn" {
  value = aws_iam_openid_connect_provider.this[0].arn
}
