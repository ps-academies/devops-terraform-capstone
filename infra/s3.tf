resource "random_uuid" "random_id" {}

resource "aws_s3_bucket" "frontend_bucket" {
  bucket        = "${var.project_name}-${random_uuid.random_id.id}"
  acl           = "private"
  force_destroy = true
}

resource "aws_s3_bucket" "logging" {
  bucket        = "access-logs-${random_uuid.random_id.id}"
  acl           = "private"
  force_destroy = true
}
