resource "aws_iam_instance_profile" "backend_server" {
  name = "${var.project_name}-profile"
  role = aws_iam_role.role.name
}

resource "aws_iam_role" "role" {
  name               = "${var.project_name}-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "ec2.amazonaws.com",
      ]
    }
  }
}
