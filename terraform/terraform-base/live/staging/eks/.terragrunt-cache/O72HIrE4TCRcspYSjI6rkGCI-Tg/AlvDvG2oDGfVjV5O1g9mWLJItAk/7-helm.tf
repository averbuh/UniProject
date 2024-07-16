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

  depends_on = [aws_eks_cluster.this, helm_release.nginx]
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

 data "kubernetes_service" "nginx_ingress" {
  metadata {
    name      = "ingress-nginx-controller"
    namespace = kubernetes_namespace.nginx.metadata[0].name
  }
  depends_on = [ helm_release.nginx ]
}

output "nginx_ingress_hostname" {
  value = data.kubernetes_service.nginx_ingress.status[0].load_balancer[0].ingress[0].hostname
}

locals {
  elb_name = regex("^(.*?)-", data.kubernetes_service.nginx_ingress.status[0].load_balancer[0].ingress[0].hostname)[0] 
}

output "cluster_oidc_url" {
  value = data.aws_eks_cluster.this.identity[0].oidc[0].issuer
}

data "aws_elb" "nginx_ingress" {
  name = local.elb_name
}

output "elb_dns_name" {
  value = data.aws_elb.nginx_ingress.dns_name
}

output "elb_zone_id" {
  value = data.aws_elb.nginx_ingress.zone_id
}

output "elb_instances" {
  value = data.aws_elb.nginx_ingress.instances
}

data "aws_route53_zone" "host" {
  name = "${var.route53_zone_name}"
}

resource "aws_route53_record" "example" {
  zone_id = data.aws_route53_zone.host.zone_id
  name    = "${var.route53_record_name}"
  type    = "A"
  depends_on = [  data.aws_elb.nginx_ingress ] 

  alias {
    name                   = data.aws_elb.nginx_ingress.dns_name
    zone_id                = data.aws_elb.nginx_ingress.zone_id
    evaluate_target_health = true
  }
}

resource "kubernetes_namespace" "apps" {
  metadata {
    name = "apps"
  }
}

resource "kubernetes_manifest" "cert-issuer" {
  manifest = yamldecode(file(var.cert_issuer_manifest))
  depends_on = [ helm_release.cert-manager, kubernetes_namespace.apps ]
}

resource "kubernetes_manifest" "ingress" {
  manifest = yamldecode(file(var.ingress_manifest))
  depends_on = [ data.kubernetes_service.nginx_ingress, kubernetes_namespace.apps ]
}

output "route53_record_name" {
  value = aws_route53_record.example.name
}


locals {
  oidc_id = regex("id/([^/]+)$", aws_iam_openid_connect_provider.this[0].url)[0]
}

output "oidc_id" {
  value = local.oidc_id
}

resource "aws_iam_role" "s3_access_role" {
  name = "s3-access-role-${aws_eks_cluster.this.name}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Federated = aws_iam_openid_connect_provider.this[0].arn
        },
        Action = "sts:AssumeRoleWithWebIdentity",
        Condition = {
          StringEquals = {
            "oidc.eks.eu-central-1.amazonaws.com/id/${local.oidc_id}:aud" = "sts.amazonaws.com",
            "oidc.eks.eu-central-1.amazonaws.com/id/${local.oidc_id}:sub" = "system:serviceaccount:${kubernetes_namespace.apps.metadata[0].name}:s3-access"
          }
        }
      }
    ]
  })
}

# Optional: Attach a policy to the role
resource "aws_iam_role_policy" "s3_access_policy" {
  depends_on = [ aws_iam_role.s3_access_role ]
  name   = "s3-access-policy"
  role   = aws_iam_role.s3_access_role.id
  policy = jsonencode({
  "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Action": [
                  "s3:GetObject",
                  "s3:PutObject",
                  "s3:DeleteObject"
              ],
              "Resource": "arn:aws:s3:::${var.s3_bucket_name}/*"
          },
          {
              "Effect": "Allow",
              "Action": [
                  "s3:ListBucket"
              ],
              "Resource": "arn:aws:s3:::${var.s3_bucket_name}"
          }
      ]
  })
}

resource "kubernetes_service_account" "s3_access" {
  depends_on = [ aws_iam_role_policy.s3_access_policy ]
  metadata {
    name      = "s3-access"
    namespace = kubernetes_namespace.apps.metadata[0].name
    annotations = {
      "eks.amazonaws.com/role-arn" = aws_iam_role.s3_access_role.arn
    }
  }
}


data "aws_ecr_authorization_token" "ecr" {
  registry_id = "975050257492"
}

locals {
  dockerconfigjson = jsonencode({
    auths = {
      "975050257492.dkr.ecr.eu-central-1.amazonaws.com" = {
        username = "AWS"
        password = data.aws_ecr_authorization_token.ecr.password
        email    = "none"
        auth     = base64encode("AWS:${data.aws_ecr_authorization_token.ecr.password}")
      }
    }
  })
}

output "dockerconfigjson_debug" {
  sensitive = true
  value = local.dockerconfigjson
}

resource "kubernetes_secret" "regcred" {
  
  metadata {
    name      = "regcred"
    namespace = "apps"
  }

  data = {
    ".dockerconfigjson" = local.dockerconfigjson
  }

  type = "kubernetes.io/dockerconfigjson"
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
