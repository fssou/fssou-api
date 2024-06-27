# iamsr_policy_base.tf

data "aws_iam_policy_document" "policy_base" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "ec2:DescribeNetworkInterfaces",
      "ec2:CreateNetworkInterface",
      "ec2:DeleteNetworkInterface",
      "ec2:DescribeInstances",
      "ec2:AttachNetworkInterface"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "policy_base" {
  path        = "/iamsr/lambda/"
  name        = "${local.function_name}-base"
  description = "A policy for lambda ${local.function_name}"
  policy      = data.aws_iam_policy_document.policy_base.json
}

resource "aws_iam_role_policy_attachment" "lambda_base" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.policy_base.arn
}
