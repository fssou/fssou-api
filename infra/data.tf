data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "all" {
  id = data.aws_vpc.default.id
}
