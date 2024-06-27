# iamsr_policy_secrets.tf

data "aws_iam_policy_document" "policy_secrets" {
  statement {
    effect = "Allow"
    actions = [
      "secretsmanager:*"
    ]
    resources = [
        aws_secretsmanager_secret.x_credentials.arn
    ]
  }
}

resource "aws_iam_policy" "policy_secrets" {
  path        = "/iamsr/lambda/"
  name        = "${local.function_name}-secrets"
  description = "A policy for lambda ${local.function_name}"
  policy      = data.aws_iam_policy_document.policy_secrets.json
}

resource "aws_iam_role_policy_attachment" "lambda_secrets" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.policy_secrets.arn
}
