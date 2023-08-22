# Portfolio Server

Go server that powers my portfolio website.

## Quick Start

### Install Required Tools

- Hot reloading using [Air](https://github.com/cosmtrek/air)
  - `go install github.com/cosmtrek/air@latest`
- [Docker](https://docs.docker.com/engine/install/) for database in development
- SQL generation using [sqlc](https://https://github.com/sqlc-dev/sqlc)
  - `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- Manage migrations using [Goose](https://github.com/pressly/goose)
  - `go install github.com/pressly/goose/v3/cmd/goose@latest`

### Run Server

- Create .env file
  - `cp .env.example .env`
- Install dependencies
  - `go mod download`
- Start containers
  - `make containers-up`
- Run migrations
  - `make migrate-up`
- Run app with hot reload or in production
  - `make go-dev`
  - `make start`
