# ============================================================
#  AWS INFRASTRUCTURE — AeroPath
# ============================================================
# Déploiement sur AWS EKS (Kubernetes) avec RDS PostgreSQL
# ============================================================

terraform {
  required_version = ">= 1.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.0"
    }
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = "~> 1.14"
    }
  }

  backend "s3" {
    bucket = "aeropath-terraform-state"
    key    = "infrastructure/terraform.tfstate"
    region = "eu-west-3"
  }
}

provider "aws" {
  region = var.aws_region
}

# ============================================================
#  VPC — Réseau
# ============================================================
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "~> 5.0"

  name = "aeropath-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["${var.aws_region}a", "${var.aws_region}b", "${var.aws_region}c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway     = true
  single_nat_gateway     = true
  enable_dns_hostnames   = true
  enable_dns_support     = true

  tags = {
    Environment = "production"
    Project     = "aeropath"
  }
}

# ============================================================
#  RDS PostgreSQL — Base de données
# ============================================================
resource "aws_db_subnet_group" "aeropath" {
  name       = "aeropath-db-subnet"
  subnet_ids = module.vpc.private_subnets
}

resource "aws_db_instance" "postgres" {
  identifier = "aeropath-postgres"

  engine         = "postgres"
  engine_version = "16.3"
  instance_class = "db.t3.medium"

  db_name  = "aeropath"
  username = var.db_username
  password = var.db_password

  db_subnet_group_name   = aws_db_subnet_group.aeropath.name
  vpc_security_group_ids = [aws_security_group.rds.id]

  allocated_storage     = 20
  max_allocated_storage = 100
  storage_type          = "gp3"
  storage_encrypted     = true

  backup_retention_period = 7
  backup_window          = "03:00-04:00"
  maintenance_window     = "sun:04:00-sun:05:00"

  skip_final_snapshot = false
  final_snapshot_identifier = "aeropath-postgres-final-${formatdate("YYYY-MM-DD-hhmm", timestamp())}"

  deletion_protection = true

  tags = {
    Name        = "aeropath-postgres"
    Environment = "production"
  }
}

# ============================================================
#  ElastiCache Redis — Cache
# ============================================================
resource "aws_elasticache_subnet_group" "aeropath" {
  name       = "aeropath-redis-subnet"
  subnet_ids = module.vpc.private_subnets
}

resource "aws_elasticache_cluster" "redis" {
  cluster_id = "aeropath-redis"

  engine         = "redis"
  engine_version = "7.1"
  node_type      = "cache.t3.micro"
  num_cache_nodes = 1

  subnet_group_name = aws_elasticache_subnet_group.aeropath.name
  security_group_ids = [aws_security_group.redis.id]

  parameter_group_name = "default.redis7"

  tags = {
    Name        = "aeropath-redis"
    Environment = "production"
  }
}

# ============================================================
#  EKS — Cluster Kubernetes
# ============================================================
module "eks" {
  source = "terraform-aws-modules/eks/aws"
  version = "~> 20.0"

  cluster_name    = "aeropath-cluster"
  cluster_version = "1.30"

  cluster_endpoint_public_access = true

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  eks_managed_node_groups = {
    main = {
      desired_size = 2
      min_size     = 2
      max_size     = 6

      instance_types = ["t3.medium"]
      capacity_type  = "ON_DEMAND"

      tags = {
        Environment = "production"
      }
    }
  }

  tags = {
    Environment = "production"
    Project     = "aeropath"
  }
}

# ============================================================
#  Security Groups
# ============================================================
resource "aws_security_group" "rds" {
  name        = "aeropath-rds-sg"
  description = "Security group for RDS PostgreSQL"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [module.eks.node_security_group_id]
    description     = "PostgreSQL from EKS"
  }

  tags = {
    Name = "aeropath-rds-sg"
  }
}

resource "aws_security_group" "redis" {
  name        = "aeropath-redis-sg"
  description = "Security group for ElastiCache Redis"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port       = 6379
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [module.eks.node_security_group_id]
    description     = "Redis from EKS"
  }

  tags = {
    Name = "aeropath-redis-sg"
  }
}

# ============================================================
#  ECR — Container Registry
# ============================================================
resource "aws_ecr_repository" "api" {
  name = "aeropath/api"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "worker" {
  name = "aeropath/worker"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

# ============================================================
#  Outputs
# ============================================================
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
