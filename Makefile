migrate-up:
	docker-compose -f ./.docker/compose.yml exec api sh -c 'migrate -path src/db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'