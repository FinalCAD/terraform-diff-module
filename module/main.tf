locals {
  diff_version = trimspace(file("${path.module}/version.txt"))
}

data "aws_secretsmanager_secret_version" "secret_manager_current_version" {
  secret_id = var.aws_secret_id_current
}

data "external" "diff_secret" {
  program = ["sh", "${path.module}/scripts/diff.sh",
    local.diff_version,
    data.aws_secretsmanager_secret_version.secret_manager_current_version.secret_string,
  jsonencode(var.secret_manager_new_version)]
}

resource "null_resource" "deploy_info" {
  triggers = {
    env_change = join("\n", [for key, value in data.external.diff_secret.result : "${key} -> ${value}"])
  }
}
