# data.tf

data "aws_caller_identity" "current" {}

data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "default" {
  filter {
    name = "vpc-id"
    values = [
      data.aws_vpc.default.id
    ]
  }
  filter {
    name = "default-for-az"
    values = [
      "true"
    ]
  }
}

data "aws_subnets" "private" {
  filter {
    name = "vpc-id"
    values = [
      data.aws_vpc.default.id
    ]
  }
  filter {
    name = "default-for-az"
    values = [
      "false"
    ]
  }
  filter {
    name = "tag:Name"
    values = [
      "private-a",
      "private-b",
      "private-c",
    ]
  }
}
