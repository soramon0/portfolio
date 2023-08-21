build:
	@go build -o bin/portfolio

start: build
	@./bin/portfolio

containers-up:
	@docker compose up -d
containers-down:
	@docker compose down

develop:
	@air

generate-sql:
	@sqlc generate

migrate-up:
	@goose -dir ./src/sql/schema/ postgres postgres://postgres:example@localhost:5432/dev_db up 

migrate-down:
	@goose -dir ./src/sql/schema/ postgres postgres://postgres:example@localhost:5432/dev_db down 