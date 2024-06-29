# main.tf

resource "aws_lambda_function" "fssou" {
  function_name = local.function_name
  description   = "This is a function for fssou API."
  handler       = "bootstrap"
  filename      = "fssou.zip"
  runtime       = "provided.al2023"
  timeout       = 15
  publish       = true
  source_code_hash = filebase64sha256("fssou.zip")
  role          = aws_iam_role.lambda.arn
  architectures = ["x86_64"]
  vpc_config {
    ipv6_allowed_for_dual_stack = true
    subnet_ids = data.aws_subnets.all.ids
    security_group_ids = [
      aws_security_group.fssou.id
    ]
  }
  environment {
    variables = {
      X_SECRETS_NAME = aws_secretsmanager_secret.x_credentials.name
    }
  }
  tracing_config {
    mode = "Active"
  }
}

resource "aws_lambda_alias" "fssou" {
  name             = "live"
  function_name    = aws_lambda_function.fssou.function_name
  function_version = "$LATEST"
}
