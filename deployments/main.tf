variable "access_key" {}
variable "secret_key" {}
variable "profile" {}

variable "name" {}
variable "region" {
  default = "ap-northeast-1"
}
variable "root_domain_name" {}
variable "sub_domain_name" {}
variable "route53_hosted_zone_id" {}
variable "s3_bucket" {}
variable "alb_certificate_arn" {}
variable "cf_certificate_arn" {}
//variable "subnet_public_a" {}
//variable "subnet_public_cidr_a" {}
//variable "subnet_public_c" {}
//variable "subnet_public_cidr_c" {}
//variable "subnet_private_a" {}
//variable "subnet_private_cidr_a" {}
//variable "subnet_private_c" {}
//variable "subnet_private_cidr_c" {}
variable "vpc_id" {}
variable "vpc_cidr" {}
variable "vpc_default_route_table_id" {}
variable "igw_id" {}
variable "iam_task_role_arn" {}
variable "db_user" {}
variable "db_pass" {}

terraform {
  backend "s3" {
    bucket = "voyageapi"
    key = "voyageapi/terraform.tfstate"
    region = "ap-northeast-1"
    profile = "admin"
  }

  required_providers {
    aws = "~> 2.49"
  }
}

provider "aws" {
  access_key = var.access_key
  secret_key = var.secret_key
  region = var.region
}

module "common" {
  source = "./modules/common"
  name = var.name
  region = var.region
  root_domain_name = var.root_domain_name
  sub_domain_name = var.sub_domain_name
  route53_hosted_zone_id = var.route53_hosted_zone_id
  s3_bucket = var.s3_bucket
  iam_task_role_arn = var.iam_task_role_arn
  alb_certificate_arn = var.alb_certificate_arn
  cf_certificate_arn = var.cf_certificate_arn
  //  subnet_public_a       = var.subnet_public_a
  //  subnet_public_cidr_a  = var.subnet_public_cidr_a
  //  subnet_public_c       = var.subnet_public_c
  //  subnet_public_cidr_c  = var.subnet_public_cidr_c
  //  subnet_private_a      = var.subnet_private_a
  //  subnet_private_cidr_a = var.subnet_private_cidr_a
  //  subnet_private_c      = var.subnet_private_c
  //  subnet_private_cidr_c = var.subnet_private_cidr_c
  vpc_id = var.vpc_id
  vpc_cidr = var.vpc_cidr
  vpc_default_route_table_id = var.vpc_default_route_table_id
  igw_id = var.igw_id
  db_user = var.db_user
  db_pass = var.db_pass
}