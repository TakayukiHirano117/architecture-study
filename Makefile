.PHONY: migrate-up migrate-down migrate-force gomock-generate-all lint-docker lint-fix-docker format-docker goimports-docker

migrate-up:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'

migrate-down:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" down 1'

migrate-force:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" force $(VERSION)'

# gomock生成コマンド
gomock-generate-all:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'go generate ./...'

# test
test-docker:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'go test $$(go list ./... | grep -v -E "(/mock/|/infra/|/db$$|/customerr)")'

# Lint
lint-docker:
	docker-compose -f ./.docker/compose.yml exec api golangci-lint run ./...

lint-fix-docker:
	docker-compose -f ./.docker/compose.yml exec api golangci-lint run --fix ./...

# Format (gofmt + goimports)
format-docker:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'go fmt ./...'
	docker-compose -f ./.docker/compose.yml exec api sh -c 'goimports -w -local github.com/TakayukiHirano117/architecture-study .'

# goimports のみ実行
goimports-docker:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'goimports -w -local github.com/TakayukiHirano117/architecture-study .'

# tidy
tidy-docker:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'go mod tidy'
