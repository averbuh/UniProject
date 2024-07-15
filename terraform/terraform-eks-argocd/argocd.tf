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