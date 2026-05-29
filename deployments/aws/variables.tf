variable "aws_region" {
  description = "Région AWS"
  type        = string
  default     = "eu-west-3" # Paris
}

variable "db_username" {
  description = "Nom d'utilisateur PostgreSQL"
  type        = string
  sensitive   = true
}

variable "db_password" {
  description = "Mot de passe PostgreSQL"
  type        = string
  sensitive   = true
}
