#!/bin/bash
# ============================================================
#  DEPLOY SCRIPT — AeroPath sur AWS EKS
# ============================================================
# Prérequis :
#   - AWS CLI configuré (aws configure)
#   - Terraform installé
#   - kubectl installé
#   - Docker installé
# ============================================================
set -euo pipefail

echo "🚀 Déploiement AeroPath sur AWS EKS"
echo "======================================"

# ─── 1. Variables ──────────────────────────────────────────
AWS_REGION="${AWS_REGION:-eu-west-3}"
DB_USERNAME="${DB_USERNAME:-aeropath}"
DB_PASSWORD="${DB_PASSWORD:-$(openssl rand -base64 32)}"
CLUSTER_NAME="aeropath-cluster"

# ─── 2. Terraform — Infrastructure ─────────────────────────
echo ""
echo "📦 Étape 1/5 : Déploiement de l'infrastructure Terraform..."
cd "$(dirname "$0")"

terraform init
terraform apply -auto-approve \
  -var="aws_region=${AWS_REGION}" \
  -var="db_username=${DB_USERNAME}" \
  -var="db_password=${DB_PASSWORD}"

# Récupérer les outputs
DB_ENDPOINT=$(terraform output -raw db_endpoint)
REDIS_ENDPOINT=$(terraform output -raw redis_endpoint)
ECR_API_URL=$(terraform output -raw ecr_api_url)
ECR_WORKER_URL=$(terraform output -raw ecr_worker_url)

echo "✅ Infrastructure déployée !"
echo "   PostgreSQL: ${DB_ENDPOINT}"
echo "   Redis:      ${REDIS_ENDPOINT}"

# ─── 3. Docker — Build & Push ──────────────────────────────
echo ""
echo "🐳 Étape 2/5 : Build & Push des images Docker..."
cd ../..

# Login ECR
aws ecr get-login-password --region ${AWS_REGION} | \
  docker login --username AWS --password-stdin ${ECR_API_URL%/*}

# Build & Push API
docker build -t aeropath/api:latest -f Dockerfile.api .
docker tag aeropath/api:latest ${ECR_API_URL}:latest
docker push ${ECR_API_URL}:latest

# Build & Push Worker
docker build -t aeropath/worker:latest -f Dockerfile.worker .
docker tag aeropath/worker:latest ${ECR_WORKER_URL}:latest
docker push ${ECR_WORKER_URL}:latest

echo "✅ Images pushées sur ECR !"

# ─── 4. kubectl — Config ───────────────────────────────────
echo ""
echo "🔧 Étape 3/5 : Configuration de kubectl..."
aws eks update-kubeconfig --region ${AWS_REGION} --name ${CLUSTER_NAME}

# ─── 5. Déploiement Kubernetes ─────────────────────────────
echo ""
echo "☸️  Étape 4/5 : Déploiement sur Kubernetes..."

# Namespace
kubectl apply -f k8s/namespace.yaml

# Secrets
kubectl create secret generic aeropath-secrets \
  --namespace aeropath \
  --from-literal=database-url="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_ENDPOINT}/aeropath?sslmode=require" \
  --from-literal=jwt-secret="$(openssl rand -base64 32)" \
  --dry-run=client -o yaml | kubectl apply -f -

# Infrastructure (Redis, NATS, ClickHouse, Jaeger, Prometheus)
kubectl apply -f k8s/infrastructure.yaml

# API Gateway
kubectl apply -f k8s/api-gateway.yaml

# Worker
kubectl apply -f k8s/worker.yaml

# Ingress
kubectl apply -f k8s/ingress.yaml

echo "✅ Déploiement Kubernetes terminé !"

# ─── 6. Vérification ───────────────────────────────────────
echo ""
echo "👀 Étape 5/5 : Vérification du déploiement..."
echo ""
echo "En attente des pods... (30s)"
sleep 30

kubectl get pods -n aeropath
kubectl get svc -n aeropath

echo ""
echo "======================================"
echo "🎉 AeroPath déployé sur AWS EKS !"
echo "======================================"
echo ""
echo "📊 Monitoring :"
echo "   Prometheus : kubectl port-forward -n aeropath svc/prometheus 9090:9090"
echo "   Grafana    : kubectl port-forward -n aeropath svc/grafana 3000:3000"
echo "   Jaeger     : kubectl port-forward -n aeropath svc/jaeger 16686:16686"
echo ""
echo "🌐 API :"
echo "   http://api.aeropath.app (si DNS configuré)"
echo ""
echo "📝 Commandes utiles :"
echo "   kubectl get pods -n aeropath -w"
echo "   kubectl logs -n aeropath deployment/api-gateway -f"
echo "   kubectl logs -n aeropath deployment/worker -f"
