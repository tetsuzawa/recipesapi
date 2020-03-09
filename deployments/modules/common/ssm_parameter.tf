resource "aws_ssm_parameter" "api_port" {
  name  = "API_PORT"
  type  = "SecureString"
  value = "80"
}