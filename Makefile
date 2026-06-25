.PHONY: run build docker-up docker-down docker-logs sqlc-gen test tidy

## Run directly on your local machine (you need to start PostgreSQL first using `make docker-db`)
run:
	go run ./cmd/server

## Compiling binary
build:
	CGO_ENABLED=0 go build -o ./bin/server ./cmd/server

## Docker Compose: Starting Postgres and API Simultaneously
docker-up:
	docker compose up --build -d

## Start only PostgreSQL (commonly used during development)
docker-db:
	docker compose up postgres -d

## Stop all services
docker-down:
	docker compose down

## View API Log
docker-logs:
	docker compose logs -f api

## Regenerate SQLC code (SQLC needs to be installed)
## brew install sqlc
sqlc-gen:
	sqlc generate

## Organize go.mod
tidy:
	go mod tidy

## Running Tests
test:
	go test ./...
