data "aws_iam_policy" "ebs_csi_policy" {
  count = var.enable_ebs ? 1 : 0
  arn = "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"
}

module "irsa-ebs-csi" {
  count = var.enable_ebs ? 1 : 0
  source  = "terraform-aws-modules/iam/aws//modules/iam-assumable-role-with-oidc"
  version = "4.7.0"

  create_role                   = true
  role_name                     = "AmazonEKSTFEBSCSIRole-${var.eks_name}"
  provider_url                  = var.provider_url
  role_policy_arns              = [data.aws_iam_policy.ebs_csi_policy[0].arn]
  oidc_fully_qualified_subjects = ["system:serviceaccount:kube-system:ebs-csi-controller-sa"]
}

resource "aws_eks_addon" "ebs-csi" {
  count = var.enable_ebs ? 1 : 0
  cluster_name             = var.eks_name 
  addon_name               = "aws-ebs-csi-driver"
  addon_version            = "v1.30.0-eksbuild.1"
  service_account_role_arn = module.irsa-ebs-csi[0].iam_role_arn
  tags = {
    "eks_addon" = "ebs-csi"
    "terraform" = "true"
  }
}


