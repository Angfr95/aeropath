# 🛩️ AeroPath

**Plateforme d'apprentissage pour l'aviation générale** — Préparation aux licences PPL, LAPL, et ATPL.

> "Le chemin le plus court entre deux points n'est pas toujours une ligne droite, mais une spirale d'apprentissage."

---

## 🌐 Production

| Service | URL |
|---------|-----|
| **Site** | [https://aeropath.ultizion.com](https://aeropath.ultizion.com) |
| **Grafana** | [https://aeropath.ultizion.com/grafana](https://aeropath.ultizion.com/grafana) |

> ⚠️ Le site force HTTPS. Tout accès HTTP est redirigé automatiquement vers HTTPS.

---

## 📋 Table des matières

- [Architecture](#architecture)
- [Stack technique](#stack-technique)
- [Prérequis](#prérequis)
- [Démarrage rapide](#démarrage-rapide)
- [Docker Compose (dev local)](#docker-compose-dev-local)
- [Monitoring local](#monitoring-local)
- [API](#api)
- [Tests](#tests)
- [Déploiement AWS](#déploiement-aws)
- [Déploiement Kubernetes](#déploiement-kubernetes)
- [CI/CD](#cicd)
- [Sécurité](#sécurité)
- [Contribuer](#contribuer)

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────┐
│                     API Gateway (Gin)                    │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌────────┐  │
│  │  Auth    │  │ Learning │  │  Recomm. │  │  WS    │  │
│  │  HTTP    │  │  HTTP    │  │  HTTP    │  │  Hub   │  │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └───┬────┘  │
│       │              │              │             │       │
│  ┌────▼──────────────▼──────────────▼─────────────▼───┐  │
│  │              Middleware (Auth, RateLimit, CORS)     │  │
│  └───────────────────────┬──────────────────────────────┘  │
└──────────────────────────┼─────────────────────────────────┘
                           │
        ┌──────────────────┼──────────────────┐
        ▼                  ▼                  ▼
   ┌─────────┐      ┌──────────┐       ┌──────────┐
   │PostgreSQL│      │  Redis   │       │   NATS   │
   │  (Métier)│      │  (Cache) │       │ (Events) │
   └─────────┘      └──────────┘       └────┬─────┘
        │                                    │
        ▼                                    ▼
   ┌──────────┐                       ┌──────────┐
   │ClickHouse│                       │  Worker  │
   │(Analytics)│                      │(Consumer)│
   └──────────┘                       └──────────┘

   📊 Monitoring : Prometheus + Grafana + Jaeger
```

### Flux de données

1. **Étudiant** → API Gateway (REST/WebSocket)
2. **API Gateway** → PostgreSQL (données métier)
3. **API Gateway** → Redis (cache, sessions)
4. **API Gateway** → NATS (events : réponses, examens)
5. **Worker** ← NATS (consomme les events)
6. **Worker** → ClickHouse (analytics)
7. **Worker** → Recommandations (spaced repetition)

---

## 🛠️ Stack technique

| Couche | Technologie | Usage |
|--------|------------|-------|
| **Langage** | Go 1.22 | Backend |
| **Framework HTTP** | Gin | API REST |
| **gRPC** | gRPC + Protobuf | Services internes |
| **Base de données** | PostgreSQL 16 | Données métier |
| **Cache** | Redis 7 | Cache, sessions |
| **Message Queue** | NATS + JetStream | Events asynchrones |
| **Analytics** | ClickHouse | Statistiques temps réel |
| **Monitoring** | Prometheus + Grafana | Métriques |
| **Tracing** | Jaeger (OpenTelemetry) | Distributed tracing |
| **Auth** | JWT (HMAC-SHA256) | Authentification |
| **WebSocket** | gorilla/websocket | Notifications temps réel |
| **CI/CD** | GitHub Actions | Build, test, déploiement |
| **IaC** | Terraform | Infrastructure AWS |
| **Orchestration** | Kubernetes (EKS) | Déploiement production |
| **Conteneurisation** | Docker + Docker Compose | Dev local |

---

## 📦 Prérequis

- [Go 1.22+](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/) (pour K8s)
- [Terraform](https://developer.hashicorp.com/terraform/install) (pour AWS)
- [AWS CLI](https://aws.amazon.com/cli/) (pour AWS)
- [golangci-lint](https://golangci-lint.run/) (optionnel)

---

## 🚀 Démarrage rapide

```bash
# 1. Cloner le projet
git clone https://github.com/Angfr95/aeropath.git
cd aeropath

# 2. Démarrer les services (PostgreSQL, Redis, NATS, ClickHouse)
docker compose up -d

# 3. Appliquer les migrations
make migrate

# 4. Lancer le serveur API
make dev

# 5. Lancer le worker (dans un autre terminal)
make worker
```

L'API est accessible sur `http://localhost:8080`.

### Commandes Makefile

```bash
make help                # Affiche toutes les commandes disponibles
make dev                 # Lance le serveur API
make worker              # Lance le worker
make build               # Compile les binaires
make test                # Tests unitaires
make test/cover          # Tests avec couverture
make test/integration    # Tests d'intégration (Docker requis)
make lint                # Analyse statique
make docker/up           # Démarre tous les services Docker
make docker/down         # Arrête les services
make ci                  # Pipeline CI complet
```

---

## 🐳 Docker Compose (dev local)

Le `docker-compose.yml` inclut tous les services nécessaires :

| Service | Image | Port | Description |
|---------|-------|------|-------------|
| **postgres** | postgres:16-alpine | 5432 | Base de données |
| **redis** | redis:7-alpine | 6379 | Cache |
| **nats** | nats:2-alpine | 4222, 8222 | Message queue |
| **clickhouse** | clickhouse/clickhouse-server:24-alpine | 9000, 8123 | Analytics |
| **prometheus** | prom/prometheus:v2.53.0 | 9090 | Métriques |
| **jaeger** | jaegertracing/all-in-one:1.60 | 16686, 4317, 4318 | Tracing distribué |
| **grafana** | grafana/grafana:11.1.0 | 3000 | Dashboards |
| **postgres-exporter** | prometheuscommunity/postgres-exporter | 9187 | Métriques PostgreSQL |
| **redis-exporter** | oliver006/redis_exporter | 9121 | Métriques Redis |
| **nats-exporter** | natsio/prometheus-nats-exporter | 7777 | Métriques NATS |

```bash
# Tout lancer
docker compose up -d

# Monitoring uniquement
docker compose up -d prometheus jaeger grafana postgres-exporter redis-exporter nats-exporter
```

---

## 📊 Monitoring local

Une fois les services démarrés :

| Service | URL | Login |
|---------|-----|-------|
| **Prometheus** | http://localhost:9090 | - |
| **Jaeger** | http://localhost:16686 | - |
| **Grafana** | http://localhost:3000 | `admin` / `admin` |

### Alertes configurées (Prometheus)

| Alerte | Seuil | Sévérité |
|--------|-------|----------|
| **HighErrorRate** | Taux d'erreur > 5% (5 min) | 🔴 Critique |
| **HighLatency** | Latence p99 > 500ms (5 min) | 🟡 Warning |
| **ServiceDown** | Service injoignable (30s) | 🔴 Critique |
| **QueueBacklog** | File NATS > 1000 messages (2 min) | 🟡 Warning |
| **LowExamPassRate** | Taux réussite < 40% (1h) | 🟡 Warning |

### Métriques exposées

```
aeroforge_http_requests_total{method,path,status}
aeroforge_http_request_duration_seconds{method,path}
aeroforge_http_requests_in_flight
aeroforge_questions_answered_total{correct}
aeroforge_active_students
```

---

## 📖 API

### Authentification

```bash
# Inscription
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"pilote@example.com","password":"monmotdepasse","lang":"fr"}'

# Connexion
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"pilote@example.com","password":"monmotdepasse"}'

# → Retourne un token JWT à utiliser dans le header Authorization
```

### Endpoints principaux

| Méthode | Path | Description | Auth |
|---------|------|-------------|------|
| `POST` | `/auth/register` | Inscription | ❌ |
| `POST` | `/auth/login` | Connexion | ❌ |
| `GET` | `/api/students/:id` | Profil étudiant | ✅ |
| `PATCH` | `/api/me/lang` | Changer langue | ✅ |
| `GET` | `/api/lessons` | Liste des leçons | ✅ |
| `GET` | `/api/lessons/:id` | Détail leçon | ✅ |
| `POST` | `/api/lessons` | Créer leçon | ✅ |
| `GET` | `/api/lessons/:id/quiz` | Quiz de leçon | ✅ |
| `GET` | `/api/questions` | Liste questions | ✅ |
| `GET` | `/api/questions/random` | Question aléatoire | ✅ |
| `GET` | `/api/questions/search?q=` | Recherche | ✅ |
| `POST` | `/api/exam/start/:theme` | Démarrer examen | ✅ |
| `POST` | `/api/exam/submit` | Soumettre examen | ✅ |
| `GET` | `/api/recommendations` | Recommandations | ✅ |
| `GET` | `/health` | Healthcheck | ❌ |
| `GET` | `/metrics` | Métriques Prometheus | ❌ |
| `GET` | `/ws` | WebSocket (notifications) | ✅ |

### WebSocket

```javascript
// Connexion
const ws = new WebSocket('ws://localhost:8080/ws');
ws.onmessage = (event) => {
  const msg = JSON.parse(event.data);
  console.log('📩', msg);
};

// Rejoindre une room
ws.send(JSON.stringify({ type: 'room:join', payload: { room: 'exam:123' } }));

// Quitter une room
ws.send(JSON.stringify({ type: 'room:leave', payload: { room: 'exam:123' } }));
```

---

## 🧪 Tests

```bash
# Tests unitaires (rapides)
make test

# Tests avec couverture
make test/cover
# → Ouvre coverage.html dans le navigateur

# Tests d'intégration (nécessite Docker)
docker compose up -d
make test/integration

# Tous les tests
make ci
```

### Structure des tests

```
internal/
├── auth/           → Tests unitaires (JWT, hash)
├── learning/       → Tests unitaires (engine, handlers mockés)
├── persistence/
│   ├── postgres/   → Tests unitaires (repos mockés)
│   ├── redis/      → Tests d'intégration (Redis requis)
│   └── clickhouse/ → Tests d'intégration (ClickHouse requis)
├── events/         → Tests d'intégration (NATS requis)
├── transport/
│   ├── grpc/       → Tests d'intégration (serveur gRPC)
│   └── middleware/ → Tests unitaires (rate limiting)
```

---

## ☁️ Déploiement AWS

L'infrastructure AWS est gérée avec **Terraform** et déploie sur **EKS (Elastic Kubernetes Service)**.

### Architecture AWS

```
┌─────────────────────────────────────────────────────┐
│                    AWS Cloud                         │
│                                                      │
│  ┌──────────────────────────────────────────────┐   │
│  │  VPC (10.0.0.0/16)                           │   │
│  │  ┌──────────────────────────────────────┐    │   │
│  │  │  EKS Cluster (Kubernetes 1.30)       │    │   │
│  │  │  ┌──────────┐  ┌──────────┐         │    │   │
│  │  │  │ API      │  │ Worker   │         │    │   │
│  │  │  │ Gateway  │  │ (x2)     │         │    │   │
│  │  │  │ (x3)     │  │          │         │    │   │
│  │  │  └──────────┘  └──────────┘         │    │   │
│  │  │  ┌──────────┐  ┌──────────┐         │    │   │
│  │  │  │ Jaeger   │  │Prometheus│         │    │   │
│  │  │  └──────────┘  └──────────┘         │    │   │
│  │  └──────────────────────────────────────┘    │   │
│  │  ┌──────────────┐  ┌──────────────────┐      │   │
│  │  │ RDS PostgreSQL│  │ ElastiCache Redis│      │   │
│  │  │ (db.t3.medium)│  │ (cache.t3.micro) │      │   │
│  │  └──────────────┘  └──────────────────┘      │   │
│  └──────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────┘
```

### Ressources AWS créées

| Service | Ressource | Spécifications |
|---------|-----------|----------------|
| **Réseau** | VPC | 10.0.0.0/16, 3 AZ, NAT Gateway |
| **Base de données** | RDS PostgreSQL 16 | db.t3.medium, 20GB gp3, backup 7j |
| **Cache** | ElastiCache Redis 7 | cache.t3.micro |
| **Kubernetes** | EKS 1.30 | 2-6 nodes t3.medium |
| **Registry** | ECR | 2 repositories (api, worker) |

### Déploiement

**Prérequis :** AWS CLI configuré (`aws configure`) + Terraform installé

```bash
# Déploiement automatique (tout-en-un)
cd deployments/aws
chmod +x deploy.sh
./deploy.sh

# Ou étape par étape :
terraform init
terraform apply -var="db_password=monMotDePasse"
```

Le script `deploy.sh` fait automatiquement :
1. ✅ Crée l'infrastructure (VPC, RDS, Redis, EKS, ECR)
2. ✅ Build & push les images Docker sur ECR
3. ✅ Configure kubectl
4. ✅ Déploie les manifests Kubernetes
5. ✅ Vérifie que tout est opérationnel

### Monitoring après déploiement

```bash
# Prometheus
kubectl port-forward -n aeropath svc/prometheus 9090:9090

# Grafana
kubectl port-forward -n aeropath svc/grafana 3000:3000

# Jaeger
kubectl port-forward -n aeropath svc/jaeger 16686:16686
```

---

## ☸️ Déploiement Kubernetes (générique)

Les manifests Kubernetes sont dans le dossier `k8s/` :

```bash
# Appliquer les manifests
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/secrets.yaml
kubectl apply -f k8s/infrastructure.yaml
kubectl apply -f k8s/api-gateway.yaml
kubectl apply -f k8s/worker.yaml
kubectl apply -f k8s/ingress.yaml

# Vérifier l'état
kubectl get pods -n aeropath -w
kubectl get svc -n aeropath
```

### Architecture K8s

- **API Gateway** : 3 réplicas (HPA : 3-10 pods, CPU 70%)
- **Worker** : 2 réplicas (HPA : 2-5 pods, CPU 70%)
- **Redis** : 1 replica
- **NATS** : 3 replicas (JetStream activé)
- **ClickHouse** : StatefulSet (10GB storage)
- **Prometheus** : Auto-scraping via annotations
- **Jaeger** : Tracing distribué
- **Ingress** : TLS avec Let's Encrypt

---

## 🔄 CI/CD

### GitHub Actions — CI (`.github/workflows/ci.yml`)

Déclenché sur `push` (main, develop) et `pull request` :

1. **Lint** : golangci-lint
2. **Test** : Tests unitaires avec PostgreSQL
3. **Build** : Compilation des binaires
4. **Docker** : Build & push sur GitHub Container Registry

### GitHub Actions — CD (`.github/workflows/cd.yml`)

Déclenché sur `push` vers `main` ou tag `v*` :

1. **Docker** : Build & push sur Docker Hub
2. **Deploy** : Déploiement SSH sur le serveur de production

### Secrets GitHub requis

Pour le CD, configurer dans **Settings > Secrets and variables > Actions** :

| Secret | Description |
|--------|-------------|
| `DOCKER_USERNAME` | Nom d'utilisateur Docker Hub |
| `DOCKER_PASSWORD` | Token Docker Hub |
| `DEPLOY_HOST` | IP du serveur de production |
| `DEPLOY_USER` | Utilisateur SSH |
| `DEPLOY_SSH_KEY` | Clé privée SSH |

---

## 🔒 Sécurité

- **Authentification** : JWT HMAC-SHA256
- **Rate Limiting** : Token Bucket (30 req/s par défaut, 5 req/s pour auth)
- **Validation** : Binding Gin + validation des entrées
- **CORS** : Configurable via middleware
- **SQL Injection** : PostgreSQL paramétré (pgx)
- **Dépendances** : Audit régulier (`make security`)
- **AWS** : RDS et Redis dans des sous-réseaux privés, accès uniquement depuis EKS

---

## 🤝 Contribuer

1. Fork le projet
2. Crée une branche (`git checkout -b feature/ma-feature`)
3. Commit (`git commit -m 'feat: ajoute ma feature'`)
4. Push (`git push origin feature/ma-feature`)
5. Ouvre une Pull Request

### Conventions de code

- Suivre [Effective Go](https://go.dev/doc/effective_go)
- Tests obligatoires pour toute nouvelle feature
- Documenter les handlers avec des commentaires DDIA
- Utiliser `gofmt` avant chaque commit

---

## 📚 Références

- [DDIA - Designing Data-Intensive Applications](https://dataintensive.net/)
- [Go Documentation](https://go.dev/doc/)
- [Gin Web Framework](https://gin-gonic.com/)
- [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream)
- [ClickHouse Documentation](https://clickhouse.com/docs)
- [Terraform AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
- [AWS EKS](https://aws.amazon.com/eks/)

---

## 📄 Licence

MIT License — Voir [LICENSE](LICENSE) pour plus de détails.
