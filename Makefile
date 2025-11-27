.PHONY: migrate-up migrate-down gomock-generate-all lint-docker lint-fix-docker

migrate-up:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'

migrate-down:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" down 1'

# gomock生成コマンド
gomock-generate-all:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'go generate ./...'

# Lint
lint-docker:
	docker-compose -f ./.docker/compose.yml exec api golangci-lint run ./...

lint-fix-docker:
	docker-compose -f ./.docker/compose.yml exec api golangci-lint run --fix ./...