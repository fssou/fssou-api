resource "aws_lambda_function" "fssou" {
  function_name     = local.function_name
  description       = "This is a function for fssou API."
  handler           = "bootstrap"
  filename          = "fssou.zip"
  runtime           = "provided.al2023"
  source_code_hash  = filebase64sha256("fssou.zip")
  role              = aws_iam_role.lambda.arn
  architectures     = ["x86_64"]
  vpc_config {
    subnet_ids = data.aws_subnets.all.ids
    security_group_ids = [
      aws_security_group.fssou.id
    ]
  }
  environment {
    variables = {
    }
  }
}
