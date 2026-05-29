output "cluster_endpoint" {
  description = "Endpoint du cluster EKS"
  value       = module.eks.cluster_endpoint
}

output "cluster_name" {
  description = "Nom du cluster EKS"
  value       = module.eks.cluster_name
}

output "db_endpoint" {
  description = "Endpoint RDS PostgreSQL"
  value       = aws_db_instance.postgres.endpoint
}

output "redis_endpoint" {
  description = "Endpoint ElastiCache Redis"
  value       = aws_elasticache_cluster.redis.cache_nodes[0].address
}

output "ecr_api_url" {
  description = "URL du registry ECR pour l'API"
  value       = aws_ecr_repository.api.repository_url
}

output "ecr_worker_url" {
  description = "URL du registry ECR pour le worker"
  value       = aws_ecr_repository.worker.repository_url
}

output "kubeconfig_command" {
  description = "Commande pour configurer kubectl"
  value       = "aws eks update-kubeconfig --region ${var.aws_region} --name ${module.eks.cluster_name}"
}
