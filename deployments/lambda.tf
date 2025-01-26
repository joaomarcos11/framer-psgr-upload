provider "aws" {
    region = "us-east-1"
}

resource "aws_lambda_function" "fiap44_framer_psgr_upload" {
    filename                       = "${path.module}/deployment.zip"
    function_name                  = "fiap44_framer_psgr_upload"
    role                           = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/LabRole"
    handler                        = "handleRequest"
    runtime                        = "provided.al2023"
}

data "aws_caller_identity" "current" {}