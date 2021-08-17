variable "aws_secret_id_current" {
  description = "AWS Secret Id"
  default     = ""
  type        = string
}

variable "secret_manager_new_version" {
  description = "Updated secret value as a map of value"
  default     = {}
  type        = map(any)
}
