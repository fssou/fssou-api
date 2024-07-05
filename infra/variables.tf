# variables.tf

variable "aws_region" {
  type = string
}

variable "gh_repo_name" {
  type = string
}

variable "gh_repo_id" {
  type = string
}

variable "secret_credentials_x" {
  type      = string
  sensitive = true
  validation {
    condition = can(jsonencode(var.secret_credentials_x))
    error_message = "The value of secrets_x must be a valid JSON string."
  }
}
