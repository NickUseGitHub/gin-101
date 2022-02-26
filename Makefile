DB_NAME := go_db
DB_HOST := localhost
DB_USERNAME := go_db
DB_PASSWORD := go_db
DB_PORT := 5432

migrate-create:
	migrate create -ext sql -dir src/migrations -seq $(MIGRATION_NAME)

migrate-up:
	migrate -path src/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path src/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down