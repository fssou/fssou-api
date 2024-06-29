# security.tf

resource "aws_security_group" "fssou" {
  name        = "fssou"
  description = "Security group for fssou API."
  vpc_id      = data.aws_vpc.default.id
  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = ["0.0.0.0/0", "::/0"]
    description = "Allow all traffic out."
  }
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0", "::/0"]
    description = "Allow all traffic out."
  }
}
