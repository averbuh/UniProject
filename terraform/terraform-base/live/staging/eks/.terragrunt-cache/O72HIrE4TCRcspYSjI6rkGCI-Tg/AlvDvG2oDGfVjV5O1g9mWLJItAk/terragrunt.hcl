terraform {
  source = "../../../modules/eks"
}

include "root" {
  path = find_in_parent_folders()
}

include "env" {
  path           = find_in_parent_folders("env.hcl")
  expose         = true
  merge_strategy = "no_merge"
}

inputs = {
  eks_version = "1.27"
  env         = include.env.locals.env
  eks_name    = "demo"
  subnet_ids  = dependency.vpc.outputs.private_subnet_ids

  route53_zone_name = "averbuchpro.com"
  route53_record_name = "stage.api.averbuchpro.com"

  s3_bucket_name = "test-images-vue"

  cert_issuer_manifest = "${get_terragrunt_dir()}/cert-issuer-staging.yaml"
  ingress_manifest = "${get_terragrunt_dir()}/ingress-staging.yaml"
  # s3_access_manifest = "${get_terragrunt_dir()}/s3-serviceAccount.yaml"


  node_groups = {
    general = {
      capacity_type  = "ON_DEMAND"
      instance_types = ["t3.small"]
      scaling_config = {
        desired_size = 3
        max_size     = 10
        min_size     = 0
        
      }
    }
  }
}

dependency "vpc" {
  config_path = "../vpc"

  mock_outputs = {
    private_subnet_ids = ["subnet-1234", "subnet-5678"]
  }
}
