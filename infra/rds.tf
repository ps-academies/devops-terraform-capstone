resource "aws_db_subnet_group" "main" {
  name       = "main"
  subnet_ids = aws_subnet.main.*.id
}

resource "random_password" "postgres_admin_password" {
  length  = 32
  special = false
}

resource "random_password" "postgres_app_password" {
  length  = 32
  special = false
}

resource "aws_db_instance" "postgres" {
  #checkov:skip=CKV_AWS_17:Create public IP because we don't have access to private GH Actions runners
  apply_immediately                   = true
  allocated_storage                   = 5
  backup_retention_period             = 7
  backup_window                       = "09:00-10:00" // UTC format of "hh24:mi-hh24:mi"
  ca_cert_identifier                  = "rds-ca-2019" //legacy setting would likely be "rds-ca-2015", opting for newer cert as this is required post March 2020
  copy_tags_to_snapshot               = true
  db_subnet_group_name                = aws_db_subnet_group.main.name
  delete_automated_backups            = true
  deletion_protection                 = false
  engine                              = "postgres"
  engine_version                      = "13.4"
  iam_database_authentication_enabled = true
  identifier                          = var.project_name
  instance_class                      = "db.t3.micro"
  iops                                = 0
  maintenance_window                  = "sun:00:00-sun:01:00" // UTC format of "ddd:hh24:mi-ddd:hh24:mi"
  max_allocated_storage               = 10
  db_name                             = var.project_name
  multi_az                            = true
  option_group_name                   = "default:postgres-13"
  parameter_group_name                = "default.postgres13"
  password                            = random_password.postgres_admin_password.result
  performance_insights_enabled        = true
  port                                = 5432
  publicly_accessible                 = true
  skip_final_snapshot                 = true
  snapshot_identifier                 = null
  storage_encrypted                   = true
  storage_type                        = "gp2"
  username                            = "${var.project_name}_admin"

  vpc_security_group_ids = [
    aws_security_group.backend_server.id
  ]

  lifecycle {
    ignore_changes = [
      snapshot_identifier,
      latest_restorable_time,
    ]
  }
}

