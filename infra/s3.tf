resource "random_uuid" "random_id" {}

resource "aws_s3_bucket" "frontend" {
  bucket        = "${var.project_name}-${random_uuid.random_id.id}"
  acl           = "private"
  force_destroy = true
}

resource "aws_s3_bucket" "logging" {
  bucket        = "access-logs-${random_uuid.random_id.id}"
  acl           = "private"
  force_destroy = true
}

resource "aws_s3_bucket_policy" "bucket_logging" {
  bucket = aws_s3_bucket.logging.bucket
  policy = data.aws_iam_policy_document.bucket_logging.json
}

data "aws_iam_policy_document" "bucket_logging" {
  statement {
    actions = ["s3:PutObject"]
    resources = [
      "arn:aws:s3:::${aws_s3_bucket.logging.bucket}/*/AWSLogs/${data.aws_caller_identity.current.account_id}/*"
    ]

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${local.elb_account_id[var.region]}:root"]
    }
  }

  statement {
    actions = ["s3:PutObject"]
    resources = [
      "arn:aws:s3:::${aws_s3_bucket.logging.bucket}/*/AWSLogs/${data.aws_caller_identity.current.account_id}/*"
    ]

    principals {
      type        = "Service"
      identifiers = ["delivery.logs.amazonaws.com"]
    }

    condition {
      test     = "StringEquals"
      variable = "s3:x-amz-acl"
      values   = ["bucket-owner-full-control"]
    }
  }

  statement {
    actions   = ["s3:GetBucketAcl"]
    resources = ["arn:aws:s3:::${aws_s3_bucket.logging.bucket}"]

    principals {
      type        = "Service"
      identifiers = ["delivery.logs.amazonaws.com"]
    }
  }
}
