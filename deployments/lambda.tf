provider "aws" {
    region = "us-east-1"
}

# data "aws_subnets" "cluster_subnets" {
#     filter {
#         name = "tag:Name"
#         values = ["fiap44-private-us-east-1a", "fiap44-private-us-east-1b"]
#     }
# }

# data "aws_vpc" "vpc" {
#   filter {
#     name   = "tag:Name"
#     values = ["fiap44-vpc"]
#   }
# }

# data "aws_db_instance" "fiap44-db" {
#   db_instance_identifier = "fiap44-db-clientes"
# }

# resource "aws_default_security_group" "default_security_group" {
#   vpc_id = data.aws_vpc.vpc.id

#     ingress {
#         protocol  = -1
#         self      = true
#         from_port = 0
#         to_port   = 0
#     }

#     egress {
#         from_port   = 0
#         to_port     = 0
#         protocol    = "-1"
#         cidr_blocks = ["0.0.0.0/0"]
#     }

#     tags = {
#         Name = "fiap44-default-security-group"
#     }
# }

resource "aws_lambda_function" "auth_lambda_func" {
    filename                       = "${path.module}/deployment.zip"
    function_name                  = "fiap44_framer_psgr_upload"
    role                           = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/LabRole"
    handler                        = "handleRequest"
    runtime                        = "provided.al2023"

    # vpc_config {
    #   subnet_ids = data.aws_subnets.cluster_subnets.ids
    #   security_group_ids = [ aws_default_security_group.default_security_group.id ]
    # }

    # environment {
    #     variables = {
    #         DB_USER = data.aws_db_instance.fiap44-db.master_username
    #         DB_PASS = "senha1ABC"
    #         DB_HOST = data.aws_db_instance.fiap44-db.address
    #         DB_PORT = data.aws_db_instance.fiap44-db.port
    #         DB_NAME = data.aws_db_instance.fiap44-db.db_name
    #     }
    # }
}

data "aws_caller_identity" "current" {}