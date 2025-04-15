provider "aws" {
  region = var.region
}

resource "aws_sqs_queue" "task_queue" {
  name = "go-task-queue"
  delay_seconds             = 0
  message_retention_seconds = 86400
  visibility_timeout_seconds = 30
}