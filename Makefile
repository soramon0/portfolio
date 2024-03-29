.PHONY: all info go-build go-prepare go-dev svelte-build svelete-dev install-web-dependencies clean

info:
	$(info ------------------------------------------)
	$(info -           Portfolio Project          -)
	$(info ------------------------------------------)
	$(info This Makefile helps you manage your projects.)
	$(info )
	$(info Available commands:)
	$(info - go-build:  Build the Golang project.)
	$(info - svelte-build:  Build the SvelteKit project.)
	$(info - all:  Run all commands (SvelteBuild, GoBuild).)
	$(info )
	$(info Usage: make <command>)

all: svelte-build go-build

go-prepare:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest

go-build:
	@echo "=== Building Protfolio Project ==="
	@go build -o bin/portfolio

go-dev:
	@air

# Build the SvelteKit project
svelte-build: install-web-dependencies
	@echo "=== Building SvelteKit Project ==="
	@if command -v pnpm >/dev/null; then \
		pnpm run -C ./src/template build; \
	else \
		npm run --prefix ./src/template build; \
	fi

svelte-dev:
	@if command -v pnpm >/dev/null; then \
		pnpm run -C ./src/template dev; \
	else \
		npm run --prefix ./src/template dev; \
	fi

# Install template dependencies
install-web-dependencies:
	@if command -v pnpm >/dev/null; then \
		pnpm install -C ./src/template; \
	else \
		npm install --prefix ./src/template; \
	fi

# Clean build artifacts
clean:
	@echo "=== Cleaning build artifacts ==="
	@rm -f bin/portfolio
	@if [ -d "./src/template/__svelte_build__" ]; then \
		rm -rf ./src/template/__svelte_build__; \
	fi

start: migrate-up generate-sql all
	@./bin/portfolio

containers-up:
	@docker compose up -d

containers-down:
	@docker compose down


generate-sql:
	@sqlc generate

migrate-up:
	@goose -dir ./src/sql/schema/ postgres postgres://postgres:example@localhost:5432/dev_db up 

migrate-down:
	@goose -dir ./src/sql/schema/ postgres postgres://postgres:example@localhost:5432/dev_db down 
