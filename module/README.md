## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.12 |
| external | >= 2.1 |
| null | >= 3.1 |

## Providers

| Name | Version |
|------|---------|
| aws | n/a |
| external | >= 2.1 |
| null | >= 3.1 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| aws\_allowed\_account\_ids | Allowed accounts ids | `list(string)` | `[]` | no |
| aws\_assume\_role | AWS role to use | `string` | `""` | no |
| aws\_region | AWS region to deploy to (e.g. eu-central-1) | `string` | `""` | no |
| aws\_secret\_id\_current | AWS Secret Id | `string` | `""` | no |
| secret\_manager\_new\_version | Updated secret value as a map of value | `map(any)` | `{}` | no |

## Outputs

No output.

