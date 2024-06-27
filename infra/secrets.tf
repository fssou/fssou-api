# secrets.tf

resource "aws_secretsmanager_secret" "x_credentials" {
  name = "fssou-x-credentials"
  description = "Secret for fssou x credentials."
}
