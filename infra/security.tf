resource "aws_security_group" "fssou" {
  name        = "fssou"
  description = "Security group for fssou API."
  vpc_id      = data.aws_vpc.default.id
}
