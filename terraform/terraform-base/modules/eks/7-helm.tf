resource "kubernetes_namespace" "nginx" {
  metadata {
    name = "nginx"
  }

  depends_on = [aws_eks_cluster.this]
}

resource "helm_release" "argocd" {
  name       = "ingress-nginx"
  chart      = "https://github.com/kubernetes/ingress-nginx/releases/download/helm-chart-4.10.1/ingress-nginx-4.10.1.tgz"
  set {
    name = "service.type"
    value = "LoadBalancer"
    
  }
  depends_on = [kubernetes_namespace.nginx]
  
}
