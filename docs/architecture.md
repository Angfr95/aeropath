# AeroForge — Architecture

## Vue d'ensemble

AeroForge (anciennement Aeropath) est une plateforme d'entraînement aéronautique adaptative orientée PPL/ATPL/CPL/BPL. Elle utilise un système de **spaced repetition** (leitner + SM-2) et de **mastery scoring** pour personnaliser l'apprentissage de chaque étudiant.

## Structure du projet

```
aeropath/
├── cmd/
│   ├── api-gateway/          # Point d'entrée API HTTP (Gin)
│   └── worker/               # Worker pour tâches asynchrones
├── internal/
│   ├── auth/                 # Service d'authentification JWT
│   ├── domain/               # Modèles métier et interfaces (Question, Lesson, Student, etc.)
│   ├── learning/             # Moteur d'apprentissage (spaced repetition, mastery)
│   ├── persistence/
│   │   ├── postgres/         # Repositories PostgreSQL (questions, leçons, étudiants, etc.)
│   │   ├── redis/            # Cache Redis (sessions, rate limiting)
│   │   └── clickhouse/       # Analytics ClickHouse (logs de réponses, agrégations)
│   ├── transport/
│   │   ├── http/             # Handlers HTTP (CRUD questions, leçons, quiz, stats, etc.)
│   │   ├── middleware/       # Middleware (auth JWT, CORS, rate limiting)
│   │   └── grpc/             # Serveur gRPC (learning, analytics, recommendation, monitoring)
│   ├── analytics/            # Moteur d'analytics (stats admin, tendances)
│   ├── monitoring/           # Monitoring (Prometheus metrics, tracing OpenTelemetry, alerting)
│   ├── recommendation/       # Moteur de recommandations adaptatives
│   ├── events/               # Event-driven (producer/consumer Kafka-like, schémas d'événements)
│   └── proto/                # Définitions Protobuf (learning, analytics, recommendation, monitoring)
├── web/
│   └── frontend-react/       # Application React (Vite + Tailwind + i18n)
│       ├── src/
│       │   ├── components/   # Composants réutilisables (LanguageSwitcher, etc.)
│       │   ├── pages/        # Pages (Dashboard, Questions, Quiz, Lessons, History, Stats, Recommendations)
│       │   ├── lib/          # API client (api.ts)
│       │   └── locales/      # Traductions (fr.json, en.json)
│       └── public/           # PWA (manifest.json, service worker, icons)
├── dashboards/
│   ├── prometheus/           # Configuration Prometheus (scrape, alerting rules)
│   └── grafana/              # Dashboard Grafana (JSON model)
├── deployments/
│   └── kubernetes/           # Manifests Kubernetes (api-gateway, worker, etc.)
├── scripts/
│   ├── migrations/           # Migrations SQL (PostgreSQL + ClickHouse)
│   └── seed/                 # Données de seed (questions, leçons)
├── tests/
│   ├── integration/          # Tests d'intégration
│   └── load/                 # Tests de charge (k6)
├── docs/                     # Documentation
├── .github/
│   └── workflows/            # CI/CD (lint, test, build, docker, deploy)
├── Makefile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Stack technique

| Couche | Technologie |
|--------|------------|
| **Langage** | Go 1.22+ |
| **Framework HTTP** | Gin |
| **Base de données principale** | PostgreSQL 16 |
| **Cache** | Redis 7 |
| **Analytics** | ClickHouse |
| **Sérialisation** | Protobuf + gRPC |
| **Authentification** | JWT (access + refresh tokens) |
| **Frontend** | React 18 + TypeScript + Vite + Tailwind CSS |
| **Internationalisation** | i18next (fr, en) |
| **PWA** | Service Worker + Manifest |
| **Monitoring** | Prometheus + Grafana + OpenTelemetry |
| **Alerting** | Alertmanager |
| **Orchestration** | Docker Compose / Kubernetes |
| **CI/CD** | GitHub Actions |
| **WebSocket** | Hub/Client pattern (real-time notifications) |

## Flux de données

```
┌─────────────┐     HTTP/WS      ┌──────────────┐     SQL      ┌────────────┐
│   Frontend   │ ──────────────> │  API Gateway  │ ──────────> │ PostgreSQL │
│   (React)    │ <────────────── │    (Gin)      │ <────────── │            │
└─────────────┘                  │               │             └────────────┘
                                 │               │     Redis   ┌────────────┐
                                 │               │ ──────────> │   Redis    │
                                 │               │ <────────── │            │
                                 │               │             └────────────┘
                                 │               │  ClickHouse ┌────────────┐
                                 │               │ ──────────> │ ClickHouse │
                                 │               │             └────────────┘
                                 │               │
                                 │    gRPC       │
                                 │ ────────────> │
                                 │ <──────────── │
                                 └──────────────┘
                                       │
                              ┌────────┴────────┐
                              │   WebSocket Hub  │
                              │  (real-time)     │
                              └─────────────────┘
```

## API Endpoints

### Questions
- `GET /api/questions` — Liste paginée des questions
- `GET /api/questions/:id` — Détail d'une question
- `POST /api/questions` — Créer une question
- `PUT /api/questions/:id` — Mettre à jour une question
- `DELETE /api/questions/:id` — Supprimer une question
- `GET /api/questions/search?q=` — Rechercher des questions
- `GET /api/questions/random` — Question aléatoire
- `POST /api/questions/answer` — Vérifier une réponse
- `GET /api/questions/by-license/:license` — Questions par licence
- `GET /api/questions/by-category/:category` — Questions par catégorie
- `GET /api/questions/by-theme/:theme` — Questions par thème
- `GET /api/questions/by-difficulty/:level` — Questions par difficulté
- `GET /api/questions/by-subtopic/:subtopic` — Questions par sous-thème
- `GET /api/questions/count` — Nombre total de questions

### Leçons
- `GET /api/lessons` — Liste paginée des leçons
- `GET /api/lessons/:id` — Détail d'une leçon
- `POST /api/lessons` — Créer une leçon
- `PUT /api/lessons/:id` — Mettre à jour une leçon
- `DELETE /api/lessons/:id` — Supprimer une leçon
- `GET /api/lessons/by-license/:license` — Leçons par licence
- `GET /api/lessons/by-category/:category` — Leçons par catégorie

### Étudiants
- `GET /api/students` — Liste des étudiants
- `GET /api/students/:id` — Détail d'un étudiant
- `POST /api/students` — Créer un étudiant
- `PUT /api/students/:id` — Mettre à jour un étudiant
- `DELETE /api/students/:id` — Supprimer un étudiant

### Historique
- `GET /api/history` — Historique des réponses (pagíné)

### Recommandations
- `GET /api/recommendations` — Recommandations personnalisées (progression, weak topics, due cards, mastery by license)

### Statistiques
- `GET /api/admin/stats` — Statistiques administrateur (total questions, leçons, étudiants, précision, répartition par licence/catégorie)

### WebSocket
- `GET /api/ws` — Connexion WebSocket (notifications en temps réel)
- `GET /api/ws/info` — Informations sur le hub WebSocket

### Authentification
- `POST /api/auth/login` — Connexion
- `POST /api/auth/register` — Inscription
- `POST /api/auth/refresh` — Rafraîchir le token

## Moteur de recommandation adaptative

Le système utilise un algorithme hybride :

1. **Spaced Repetition (SM-2 modifié)** : Chaque question a un intervalle de révision qui augmente avec les bonnes réponses et diminue avec les mauvaises.
2. **Mastery Score** : Score de maîtrise par thème (0-100%) basé sur le ratio bonnes/mauvaises réponses pondéré par la difficulté.
3. **Weak Topics** : Thèmes avec un score < 70%, priorisés (high < 40%, medium < 70%, low ≥ 70%).
4. **Progression globale** : Moyenne pondérée des scores de maîtrise.
5. **Due Cards** : Questions dont la date de révision est dépassée, triées par priorité.

## Monitoring

- **Prometheus** : Métriques HTTP (requêtes, latence, erreurs), métriques métier (questions répondues, taux de succès)
- **Grafana** : Dashboard avec panels pour les métriques en temps réel
- **Alertmanager** : Alertes configurées (taux d'erreur > 5%, latence > 500ms, uptime < 99%)
- **OpenTelemetry** : Tracing distribué avec export Jaeger

## Événements

Système event-driven avec producer/consumer :
- `question.answered` — Quand un étudiant répond à une question
- `lesson.completed` — Quand une leçon est terminée
- `student.registered` — Quand un nouvel étudiant s'inscrit
