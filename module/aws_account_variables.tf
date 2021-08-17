variable "aws_region" {
  description = "AWS region to deploy to (e.g. eu-central-1)"
  type        = string
  default     = ""
}

variable "aws_assume_role" {
  description = "AWS role to use"
  type        = string
  default     = ""
}

variable "aws_allowed_account_ids" {
  description = "Allowed accounts ids"
  type        = list(string)
  default     = []
}
