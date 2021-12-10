resource "aws_iam_instance_profile" "backend_server" {
  name = "${var.project_name}-profile"
  role = aws_iam_role.project.name
}

resource "aws_iam_role" "project" {
  name               = "${var.project_name}-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "assume_role" {
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

resource "aws_iam_role_policy_attachment" "ssm" {
  role       = aws_iam_role.project.name
  policy_arn = aws_iam_policy.ssm.arn
}

resource "aws_iam_policy" "ssm" {
  name        = "ssm-policy"
  description = "Allow role to access ssm"
  policy      = data.aws_iam_policy_document.ssm.json
}

data "aws_iam_policy_document" "ssm" {
  statement {
    actions = [
      "ssm:SendCommand",
    ]
    resources = [
      "arn:aws:ssm:${var.region}::document/AWS-RunShellScript"
    ]
  }

  statement {
    actions = [
      "ec2messages:*",
      "ssmmessages:*",
      "ssm:*",
    ]
    resources = [
      "*"
    ]
  }
}