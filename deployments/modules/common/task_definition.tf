# API
data "template_file" "api" {
  template = file("${path.root}/task-definitions/api-task-definition.json")

  vars = {
    product              = var.name
    image-url            = aws_ecr_repository.api.repository_url
    cpu                  = var.api_cpu
    memory               = var.api_memory
    cloudwatch_log_group = aws_cloudwatch_log_group.ecs.name
  }
}

resource "aws_ecs_task_definition" "api" {
  family                   = "${var.name}-api-task"
  container_definitions    = data.template_file.api.rendered
  task_role_arn            = data.aws_arn.task_role.arn
  execution_role_arn       = data.aws_arn.task_role.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.api_cpu
  memory                   = var.api_memory

  provisioner "local-exec" {
    command = "./push-image.sh ${aws_ecr_repository.api.repository_url}"
  }
}