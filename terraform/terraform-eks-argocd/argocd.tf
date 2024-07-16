resource "kubernetes_namespace" "argocd" {
  metadata {
    name = "argocd"
  }

  depends_on = [aws_eks_cluster.test]

  lifecycle {
    prevent_destroy = true
  }
}

resource "helm_release" "argocd" {
  name       = "argocd"
  chart      = "https://github.com/argoproj/argo-helm/releases/download/argo-cd-5.36.11/argo-cd-5.36.11.tgz"
  namespace  = "argocd"
  depends_on = [kubernetes_namespace.argocd]

  set {
    name  = "server.service.type"
    value = "LoadBalancer"
  }
  set {
    name = "configs.secret.argocdServerAdminPassword"
    value = "$2a$10$jY88T7TYERd0gihy0ZCN0.6SAaaFyXcN91L/QY9zH1ZbCqTPtfcKa"
  }
  lifecycle {
    prevent_destroy = true
  }
}

data "kubernetes_service" "argocd_server" {
  metadata {
    name      = "argocd-server"
    namespace = kubernetes_namespace.argocd.metadata[0].name
  }
  depends_on = [ helm_release.argocd ]
}

output "argocd_server_hostname" {
  value = data.kubernetes_service.argocd_server.status[0].load_balancer[0].ingress[0].hostname
}

locals {
  elb_name = regex("^(.*?)-", data.kubernetes_service.argocd_server.status[0].load_balancer[0].ingress[0].hostname)[0]
}

output "cluster_oidc_url" {
  value = data.aws_eks_cluster.test.identity[0].oidc[0].issuer
}

data "aws_elb" "argocd-server" {
  name = local.elb_name
}

output "elb_dns_name" {
  value = data.aws_elb.argocd-server.dns_name
}

output "elb_zone_id" {
  value = data.aws_elb.argocd-server.zone_id
}

output "elb_instances" {
  value = data.aws_elb.argocd-server.instances
}

data "aws_route53_zone" "host" {
  name = "averbuchpro.com"
}
resource "aws_route53_record" "example" {
  zone_id = data.aws_route53_zone.host.zone_id
  name    = "argocd.averbuchpro.com"
  type    = "A"
  depends_on = [  data.aws_elb.argocd-server ]

  alias {
    name                   = data.aws_elb.argocd-server.dns_name 
    zone_id                = data.aws_elb.argocd-server.zone_id
    evaluate_target_health = true
  }
}


output "route53_record_name" {
  value = aws_route53_record.example.name
}

# resource "helm_release" "thanos" {
#   name             = "thanos"
#   repository       = "https://charts.bitnami.com/bitnami"
#   chart            = "thanos"
#   version          = "15.7.12"
#   namespace        = "monitoring"
#  create_namespace  = true
#  cleanup_on_fail   = true

#   values = [
#     file("${path.module}/values.yaml")
#   ]

#   depends_on = [ aws_eks_cluster.test ]
# }

# resource "helm_release" "grafana" {
#   name             = "grafana"
#   repository       = "https://charts.bitnami.com/bitnami"
#   chart            = "grafana"
#   version          = "11.3.8"
#   namespace        = "observability"
#  create_namespace  = true
#  cleanup_on_fail   = true

#   values = [
#     file("${path.module}/values.yaml")
#   ]

#   depends_on = [ aws_eks_cluster.test ]
# }