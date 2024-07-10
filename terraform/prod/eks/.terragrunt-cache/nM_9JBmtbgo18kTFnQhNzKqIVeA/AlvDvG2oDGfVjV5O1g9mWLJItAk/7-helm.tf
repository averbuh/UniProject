resource "kubernetes_namespace" "nginx" {
  metadata {
    name = "nginx"
  }

  depends_on = [aws_eks_cluster.this]
}

resource "helm_release" "nginx" {
  name       = "ingress-nginx"
  chart      = "https://github.com/kubernetes/ingress-nginx/releases/download/helm-chart-4.10.1/ingress-nginx-4.10.1.tgz"
  namespace = "nginx"

  set {
    name = "service.type"
    value = "LoadBalancer"
    
  }
  depends_on = [kubernetes_namespace.nginx]
}

resource "kubernetes_namespace" "cert-manager" {
  metadata {
    name = "cert-manager"
  }

  depends_on = [aws_eks_cluster.this]
}
resource "helm_release" "cert-manager" {
  name       = "cert-manager"
  chart = "cert-manager"
  repository = "https://charts.jetstack.io"
  version = "1.15.1"
  namespace = "cert-manager"
  set {
    name = "crds.enabled"
    value = true
    
  }
  
  depends_on = [kubernetes_namespace.cert-manager]
}

# resource "helm_release" "prometheus_operator" {
#   name             = "prometheus-operator"
#   repository       = "https://prometheus-community.github.io/helm-charts"
#   chart            = "kube-prometheus-stack"
#   version          = "61.2.0"
#   namespace        = "monitoring"
#   create_namespace = true
#   cleanup_on_fail  = true

#   set {
#     name = "prometheusSpec.externalLabels.cluster"
#     value = var.eks_name
#   }
#   set {
#     name = "prometheusSpec.externalLabels.environment"
#     value = var.env
#   }

#   values = [
#     file("${path.module}/values.yaml")
#   ]

#   depends_on = [kubernetes_namespace.cert-manager]
# }


# resource "kubernetes_namespace" "alb-ingress-controller" {
#   metadata {
#     name = "alb-ingress-controller"
#   }

#   depends_on = [aws_eks_cluster.this]
# }

# resource "helm_release" "alb_ingress" {
#   name       = "alb-ingress"
#   chart      = "https://aws.github.io/eks-charts"
#   set {
#     name = "clusterName"
#     value = aws_eks_cluster.this.name

#   }
#   depends_on = [kubernetes_namespace.alb-ingress-controller]
#   # set {
#   #   name = "serviceAccount.create"
#   #   value = false

#   # }
#   # set {
#   #   name = "serviceAccount.name"
#   #   value = "alb-ingress-controller"
#   # }
# }
