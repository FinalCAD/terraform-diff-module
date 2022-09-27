## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.12 |
| <a name="requirement_external"></a> [external](#requirement\_external) | >= 2.1 |
| <a name="requirement_null"></a> [null](#requirement\_null) | >= 3.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |
| <a name="provider_external"></a> [external](#provider\_external) | >= 2.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_secretsmanager_secret_version.secret_manager_current_version](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/secretsmanager_secret_version) | data source |
| [external_external.diff_secret](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aws_allowed_account_ids"></a> [aws\_allowed\_account\_ids](#input\_aws\_allowed\_account\_ids) | Allowed accounts ids | `list(string)` | `[]` | no |
| <a name="input_aws_assume_role"></a> [aws\_assume\_role](#input\_aws\_assume\_role) | AWS role to use | `string` | `""` | no |
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | AWS region to deploy to (e.g. eu-central-1) | `string` | `""` | no |
| <a name="input_aws_secret_id_current"></a> [aws\_secret\_id\_current](#input\_aws\_secret\_id\_current) | AWS Secret Id | `string` | `""` | no |
| <a name="input_secret_manager_new_version"></a> [secret\_manager\_new\_version](#input\_secret\_manager\_new\_version) | Updated secret value as a map of value | `map(any)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_env_vars_change"></a> [env\_vars\_change](#output\_env\_vars\_change) | n/a |
