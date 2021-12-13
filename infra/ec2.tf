resource "aws_security_group" "backend_server" {
  name        = var.project_name
  description = "Security group for ${var.project_name}"

  vpc_id = aws_vpc.main.id

  ingress {
    description = "https"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "http"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = concat(
      aws_subnet.main.*.cidr_block,
      aws_subnet.public.*.cidr_block,
    )
  }

  ingress {
    description = "ssh"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "postgres-self"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    self        = true
  }

  ingress {
    description = "postgres"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = aws_subnet.main.*.cidr_block
  }

  egress {
    description      = "allow_all"
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

data "aws_ami" "ubuntu" {
  owners      = ["099720109477"] # canonical
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}

resource "aws_instance" "backend_server" {
  count                = length(aws_subnet.main)
  ami                  = data.aws_ami.ubuntu.id
  instance_type        = "t3.nano"
  iam_instance_profile = aws_iam_instance_profile.backend_server.name

  key_name = aws_key_pair.ssh_key.key_name

  subnet_id                   = aws_subnet.main.*.id[count.index]
  vpc_security_group_ids      = [aws_security_group.backend_server.id]
  associate_public_ip_address = true

  monitoring = true

  metadata_options {
    http_endpoint = "enabled"
    http_tokens   = "required"
  }

  root_block_device {
    encrypted = true
  }

  depends_on = [aws_route_table.main]

  tags = var.aws_tags
}

