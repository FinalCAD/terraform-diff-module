output "env_vars_change" {
  value = data.external.diff_secret.result
}
