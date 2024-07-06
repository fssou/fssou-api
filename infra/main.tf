# main.tf

resource "aws_lambda_function" "fssou" {
  function_name    = local.function_name
  description      = "This is a function for fssou API."
  handler          = "bootstrap"
  filename         = "fssou.zip"
  runtime          = "provided.al2023"
  timeout          = 15
  publish          = true
  source_code_hash = fileexists("fssou.zip") ? filebase64sha256("fssou.zip") : null
  role             = aws_iam_role.lambda.arn
  architectures = ["x86_64"]
  environment {
    variables = {
      SECRET_CREDENTIALS_X = var.secret_credentials_x
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

resource "aws_lambda_function_url" "fssou" {
  authorization_type = "NONE"
  function_name      = aws_lambda_function.fssou.function_name
  cors {
    max_age           = 3000
    allow_credentials = true
    expose_headers = ["*"]
    allow_origins = ["*"]
    allow_methods = ["GET", "POST", "PUT", "DELETE", "OPTIONS"]
    allow_headers = ["*"]
  }
}
