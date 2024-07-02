# secrets.tf

resource "aws_secretsmanager_secret" "x_credentials" {
  name = "fssou-api-x-credentials"
  description = "Secret for fssou x credentials."
}
