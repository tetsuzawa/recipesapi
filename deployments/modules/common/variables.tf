variable "name" {
  type        = string
  default     = "sample"
  description = "Product name"
}

variable "region" {
  type        = string
  default     = "ap-northeast-1"
  description = "Main region"
}

variable "root_domain_name" {
  type        = string
  default     = "example.com"
  description = "Root domain name"
}

variable "sub_domain_name" {
  type        = string
  default     = "sub"
  description = "Sub domain name"
}

variable "route53_hosted_zone_id" {
  type        = string
  default     = "Z2S***********"
  description = "ID of hosted zone which your domain will be hosted "
}

variable "s3_bucket" {
  type        = string
  default     = "sub.example.com"
  description = "Backend bucket name"
}

variable "vpc_id" {
  type        = string
  default     = "vpc-xxxxxxxx"
  description = "ID of vpc"
}

variable "vpc_cidr" {
  type        = string
  default     = "10.0.0.0/16"
  description = "cidr of vpc"
}

variable "vpc_default_route_table_id" {
  type        = string
  default     = "rtb-xxxxxxxxxxxxxxxxxxx"
  description = "default route table ID"
}

variable "igw_id" {
  type        = string
  default     = "igw-xxxxxxxxxxxxxxxxxxx"
  description = "internet gateway ID"
}
variable "iam_task_role_arn" {
  type        = string
  default     = "arn:aws:iam::*************:role/*****************"
  description = "ARN of IAM task role"
}

variable "alb_certificate_arn" {
  type        = string
  default     = "arn:aws:acm:ap-northeast-1:xxxxxxxxxxxx:certificate/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
  description = "ARN certification of ALB"
}

variable "cf_certificate_arn" {
  default = "arn:aws:acm:us-east-1:xxxxxxxxxxxx:certificate/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
}

variable "db_name" {
  type        = string
  default     = "mydb"
  description = "Name of DB"
}

variable "db_user" {
  type        = string
  default     = "user"
  description = "User name of DB"
}

variable "db_pass" {
  type        = string
  default     = "pass"
  description = "Password of DB"
}

variable "db_storage_size" {
  type        = string
  default     = "20"
  description = "Storage size of DB"
}

variable "front_cpu" {
  type        = string
  default     = "256"
  description = "CPU size of frontend task"
}

variable "front_memory" {
  type        = string
  default     = "512"
  description = "Memory size of frontend task"
}

variable "api_cpu" {
  type        = string
  default     = "256"
  description = "CPU size of api task"
}

variable "api_memory" {
  type        = string
  default     = "512"
  description = "Memory size of api task"
}
