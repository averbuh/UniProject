resource "kubernetes_namespace" "argocd" {
  metadata {
    name = "argocd"
  }

  depends_on = [aws_eks_cluster.test]
}

resource "helm_release" "argocd" {
  name       = "argocd"
  chart      = "https://github.com/argoproj/argo-helm/releases/download/argo-cd-6.7.15/argo-cd-6.7.15.tgz"
  namespace  = "argocd"
  depends_on = [kubernetes_namespace.argocd]
}
