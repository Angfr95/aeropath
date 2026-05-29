.PHONY: dev build test lint vet clean docker-up docker-down migrate \
        proto worker swagger help ci integration-test security-audit \
        docker-build deploy

# ─── Variables ───────────────────────────────────────────────
APP_NAME      := aeropath
CMD_DIR       := ./cmd/api-gateway
WORKER_DIR    := ./cmd/worker
BUILD_DIR     := ./build
GO_FLAGS      := -ldflags="-s -w"
DATABASE_URL  ?= postgres://aeropath:secret@localhost:5432/aeropath?sslmode=disable

# ─── Développement ───────────────────────────────────────────

## dev — Lance le serveur API en mode développement
dev:
	go run $(CMD_DIR)

## worker — Lance le worker (events, analytics, etc.)
worker:
	go run $(WORKER_DIR)

## build — Compile l'API Gateway et le Worker
build:
	@mkdir -p $(BUILD_DIR)
	go build $(GO_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-api $(CMD_DIR)
	go build $(GO_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-worker $(WORKER_DIR)
	@echo "✅ Build terminé : $(BUILD_DIR)/"

## test — Lance tous les tests unitaires
test:
	go test ./... -count=1 -race -timeout 60s

## test/cover — Tests avec couverture
test/cover:
	go test ./... -count=1 -race -coverprofile=coverage.out -timeout 60s
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Rapport de couverture : coverage.html"

## test/integration — Tests d'intégration (nécessite Docker)
test/integration:
	go test -tags=integration ./... -count=1 -v -timeout 120s

## lint — Analyse statique du code
lint:
	golangci-lint run ./... --timeout=5m || true

## vet — Vérifications Go standard
vet:
	go vet ./...

## fmt — Formate le code
fmt:
	go fmt ./...

## tidy — Nettoie les dépendances
tidy:
	go mod tidy
	go mod verify

# ─── Docker ──────────────────────────────────────────────────

## docker/up — Démarre tous les services
docker/up:
	docker compose up -d --wait

## docker/down — Arrête tous les services
docker/down:
	docker compose down

## docker/logs — Logs des services
docker/logs:
	docker compose logs -f

## docker/ps — État des services
docker/ps:
	docker compose ps

## docker/build — Build les images Docker
docker/build:
	docker compose build

# ─── Base de données ─────────────────────────────────────────

## migrate — Applique les migrations PostgreSQL
migrate:
	psql "$(DATABASE_URL)" -f scripts/migrations/001_init.sql

## migrate/reset — Reset complet de la base
migrate/reset:
	psql "$(DATABASE_URL)" -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
	$(MAKE) migrate

# ─── gRPC / Proto ────────────────────────────────────────────

## proto — Génère le code Go depuis les fichiers .proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/proto/*.proto
	@echo "✅ Code gRPC généré"

# ─── CI / Qualité ────────────────────────────────────────────

## ci — Pipeline CI complet
ci: tidy vet fmt test build

## security — Audit de sécurité des dépendances
security:
	go list -json -m all | nancy sleuth || true
	@echo "ℹ️  Installe nancy : go install github.com/sonatype-nexus-community/nancy@latest"

## audit — Vérifie les vulnérabilités Go
audit:
	go list -u -m all 2>/dev/null | grep '\['
	@echo "ℹ️  Vérifie les mises à jour disponibles ci-dessus"

# ─── Déploiement ─────────────────────────────────────────────

## deploy — Déploiement (exemple avec kubectl)
deploy:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -f k8s/
	@echo "✅ Déploiement appliqué"

## deploy/status — État du déploiement
deploy/status:
	kubectl get pods -n aeropath
	kubectl get svc -n aeropath

# ─── Utilitaires ─────────────────────────────────────────────

## swagger — Génère la documentation Swagger
swagger:
	swag init -g cmd/api-gateway/main.go -o cmd/api-gateway/docs
	@echo "✅ Documentation Swagger générée"

## clean — Nettoie les artefacts de build
clean:
	rm -rf $(BUILD_DIR) coverage.out coverage.html tmp/
	go clean -cache
	@echo "✅ Nettoyage terminé"

## help — Affiche cette aide
help:
	@echo "╔══════════════════════════════════════════════╗"
	@echo "║   🛩️  AeroPath — Commandes disponibles      ║"
	@echo "╚══════════════════════════════════════════════╝"
	@echo ""
	@grep -E '^## ' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "📖 Exemples :"
	@echo "  make dev              # Lance le serveur"
	@echo "  make test             # Tests unitaires"
	@echo "  make docker/up        # Démarre tous les services"
	@echo "  make ci               # Pipeline CI complet"
