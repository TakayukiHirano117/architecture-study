.PHONY: up down restart logs build migrate migrate-up migrate-down migrate-force help

# Docker commands (ルートから直接実行)
up:
	docker compose -f backend/.docker/compose.yml up

down:
	docker compose -f backend/.docker/compose.yml down

restart:
	docker compose -f backend/.docker/compose.yml restart

logs:
	docker compose -f backend/.docker/compose.yml logs -f api

build:
	docker compose -f backend/.docker/compose.yml build --no-cache

# Migration commands
migrate-up:
	docker compose -f backend/.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'

migrate-down:
	docker compose -f backend/.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" down 1'

migrate-force:
	docker compose -f backend/.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" force $(VERSION)'

migrate: migrate-up

# Help
help:
	@echo "=== Available commands ==="
	@echo "  make up             - Start all Docker containers"
	@echo "  make down           - Stop and remove containers"
	@echo "  make restart        - Restart containers"
	@echo "  make logs           - Show api container logs"
	@echo "  make build          - Rebuild Docker images"
	@echo ""
	@echo "  make migrate        - Run database migrations (up)"
	@echo "  make migrate-up     - Run all pending migrations"
	@echo "  make migrate-down   - Rollback last migration"
	@echo "  make migrate-force  - Force migration to specific version"
	@echo ""
	@echo "  make help           - Show this help message"
	@echo ""
	@echo "All commands can be run from project root."
