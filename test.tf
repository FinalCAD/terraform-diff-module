provider "aws" {
  region              = local.aws_region
  allowed_account_ids = local.aws_allowed_account_ids

  assume_role {
    role_arn = local.aws_assume_role
  }
}

locals {
  secret_name             = "terraform/production/eu/authorization-service-api"
  aws_region              = "eu-central-1"
  aws_assume_role         = "arn:aws:iam::862962005596:role/tf-terraform-developer"
  aws_allowed_account_ids = ["862962005596"]

  application_environment = {
    "MassTransit__Enable" = true
  }
}

data "aws_secretsmanager_secret" "secret" {
  name = local.secret_name
}

module "diff_secret" {
  source                     = "./module"
  aws_secret_id_current      = data.aws_secretsmanager_secret.secret.id
  secret_manager_new_version = local.application_environment
}

output "env_vars_change" {
  value = values(module.diff_secret.env_vars_change)
}
