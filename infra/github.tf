data "github_actions_public_key" "capstone" {
  repository = var.github_repo
}

resource "github_repository_environment" "capstone" {
  environment = var.environment
  repository  = var.github_repo
}

resource "github_actions_secret" "capstone_aws_access_key_id" {
  repository      = var.github_repo
  secret_name     = "AWS_ACCESS_KEY_ID"
  plaintext_value = var.access_key
}

resource "github_actions_secret" "capstone_aws_secret_access_key" {
  repository      = var.github_repo
  secret_name     = "AWS_SECRET_ACCESS_KEY"
  plaintext_value = var.secret_key
}

resource "github_actions_environment_secret" "capstone_s3_bucket_name" {
  environment     = github_repository_environment.capstone.environment
  repository      = var.github_repo
  secret_name     = "S3_BUCKET_NAME"
  plaintext_value = aws_s3_bucket.frontend.bucket
}

resource "github_actions_environment_secret" "capstone_aws_region" {
  environment     = github_repository_environment.capstone.environment
  repository      = var.github_repo
  secret_name     = "AWS_REGION"
  plaintext_value = var.region
}

resource "github_actions_environment_secret" "ssh_private_key" {
  environment     = github_repository_environment.capstone.environment
  repository      = var.github_repo
  secret_name     = "SSH_PRIVATE_KEY"
  plaintext_value = tls_private_key.ssh.private_key_pem
}