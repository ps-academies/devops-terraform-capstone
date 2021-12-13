resource "random_uuid" "random_id" {}

resource "aws_s3_bucket" "frontend" {
  #checkov:skip=CKV_AWS_19:Don't encrypt publicly accessible website
  #checkov:skip=CKV_AWS_20:Website should be publicly accessible
  #checkov:skip=CKV_AWS_21:Don't worry about versioning because website is versioned through git
  #checkov:skip=CKV_AWS_145:Don't encrypt publicly accessible website
  bucket = "${var.project_name}-${random_uuid.random_id.id}"
  acl    = "public-read"

  force_destroy = true

  logging {
    target_bucket = aws_s3_bucket.logging.id
    target_prefix = "${var.project_name}-s3-frontend"
  }

  website {
    index_document = "index.html"
    error_document = "404.html"
  }
}

resource "aws_s3_bucket_policy" "frontend" {
  bucket = aws_s3_bucket.frontend.bucket
  policy = data.aws_iam_policy_document.frontend.json
}

data "aws_iam_policy_document" "frontend" {
  statement {
    sid       = "PublicReadGetObject"
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::${aws_s3_bucket.frontend.bucket}/*"]

    principals {
      type        = "*"
      identifiers = ["*"]
    }

  }
}

resource "aws_s3_bucket_public_access_block" "frontend" {
  bucket = aws_s3_bucket.frontend.id

  block_public_acls   = false
  block_public_policy = false
}



resource "aws_s3_bucket" "logging" {
  #checkov:skip=CKV_AWS_18:This is the logging bucket
  bucket        = "access-logs-${random_uuid.random_id.id}"
  acl           = "private"
  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "logging" {
  bucket = aws_s3_bucket.logging.id

  block_public_acls   = true
  block_public_policy = true
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

