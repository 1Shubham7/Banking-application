output "k8s_master_id" {
  description = "ID of K8s master instance"
  value       = aws_instance.k8s_master.id
}

output "k8s_worker_id" {
  description = "ID of K8s worker instance"
  value       = aws_instance.k8s_worker.id
}

output "k8s_master_private_ip" {
  description = "Private IP of K8s master"
  value       = aws_instance.k8s_master.private_ip
}

output "k8s_worker_private_ip" {
  description = "Private IP of K8s worker"
  value       = aws_instance.k8s_worker.private_ip
}

output "security_group_id" {
  description = "ID of the security group"
  value       = aws_security_group.k8s_sg.id
}

output "ecr_repository_url" {
  description = "URL of ECR repository"
  value       = aws_ecr_repository.smyik.repository_url
}
