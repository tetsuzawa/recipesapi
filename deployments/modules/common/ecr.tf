# API
resource "aws_ecr_repository" "api" {
  name = "${var.name}-api"

  provisioner "local-exec" {
    command = "./push-image.sh ${var.name} ${aws_ecr_repository.api.repository_url}"
  }
}