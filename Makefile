DB_URL=postgres://postgres:pass@localhost:5432/postgres

MIGRATIONS_PATH=./migrations

run:
	@go run cmd/main.go

migrate-create:
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down