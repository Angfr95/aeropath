# 🛩️ AeroPath

**Plateforme d'apprentissage pour l'aviation générale** — Préparation aux licences PPL, LAPL, et ATPL.

> "Le chemin le plus court entre deux points n'est pas toujours une ligne droite, mais une spirale d'apprentissage."

---

## 📋 Table des matières

- [Architecture](#architecture)
- [Stack technique](#stack-technique)
- [Prérequis](#prérequis)
- [Démarrage rapide](#démarrage-rapide)
- [API](#api)
- [Tests](#tests)
- [Déploiement](#déploiement)
- [Monitoring](#monitoring)
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
| **Orchestration** | Kubernetes | Déploiement production |
| **Conteneurisation** | Docker + Docker Compose | Dev local |

---

## 📦 Prérequis

- [Go 1.22+](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/) (pour K8s)
- [golangci-lint](https://golangci-lint.run/) (optionnel)

---

## 🚀 Démarrage rapide

```bash
# 1. Cloner le projet
git clone https://github.com/aeropath/aeropath.git
cd aeropath

# 2. Démarrer les services (PostgreSQL, Redis, NATS, ClickHouse)
make docker/up

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
make help          # Affiche toutes les commandes disponibles
make dev           # Lance le serveur API
make worker        # Lance le worker
make build         # Compile les binaires
make test          # Tests unitaires
make test/cover    # Tests avec couverture
make test/integration  # Tests d'intégration (Docker requis)
make lint          # Analyse statique
make docker/up     # Démarre tous les services Docker
make docker/down   # Arrête les services
make ci            # Pipeline CI complet
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
make docker/up
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

## ☸️ Déploiement Kubernetes

```bash
# Appliquer les manifests
make deploy

# Vérifier l'état
make deploy/status

# Architecture K8s :
# - 3 replicas API Gateway (HPA : 3-10 pods)
# - 2 replicas Worker (HPA : 2-5 pods)
# - Redis, NATS (3 replicas), ClickHouse (StatefulSet)
# - Prometheus + Jaeger
# - Ingress avec TLS (Let's Encrypt)
```

---

## 📊 Monitoring

- **Prometheus** : Métriques à `/metrics` (port 9090)
- **Jaeger** : Tracing distribué (port 16686)
- **Healthcheck** : `/health` et `/ready`

### Métriques exposées

```
aeropath_http_requests_total{method,path,status}
aeropath_http_request_duration_seconds{method,path}
aeropath_questions_answered_total{correct}
aeropath_exams_completed_total{passed}
aeropath_active_connections
aeropath_cache_hits_total
aeropath_cache_misses_total
```

---

## 🔒 Sécurité

- **Authentification** : JWT HMAC-SHA256
- **Rate Limiting** : Token Bucket (30 req/s par défaut, 5 req/s pour auth)
- **Validation** : Binding Gin + validation des entrées
- **CORS** : Configurable via middleware
- **SQL Injection** : PostgreSQL paramétré (pgx)
- **Dépendances** : Audit régulier (`make security`)

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

---

## 📄 Licence

MIT License — Voir [LICENSE](LICENSE) pour plus de détails.
