resource "aws_lb" "alb" {
  name               = "${var.name}-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups = [
  aws_security_group.alb.id]
  subnets = [
  aws_subnet.public_a.id, aws_subnet.public_c.id]

  tags = {
    Name    = "${var.name}-alb"
    Product = var.name
  }
}

# ALB target group

resource "aws_lb_target_group" "api" {
  name        = "${var.name}-alb-tg-api"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.vpc.id
  target_type = "ip"

  health_check {
    interval            = 60
    path                = "/"
    protocol            = "HTTP"
    timeout             = 20
    unhealthy_threshold = 4
    matcher             = 200
  }

  tags = {
    Name    = "${var.name}-alb-target-group-api"
    Product = var.name
  }
}

# ALB Listener

resource "aws_lb_listener" "api" {
  load_balancer_arn = aws_lb.alb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.api.arn
    type             = "forward"
  }
}