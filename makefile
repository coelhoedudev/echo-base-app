.PHONY: run build test migrate rollback seed dev docker-up docker-down

# Development
run:
	@go run cmd/api/main.go

build:
	@go build -o bin/api cmd/api/main.go

test:
	@go test ./... -v

test-coverage:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

# Database
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

# Development with hot reload
dev:
	@air

# Docker
docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

docker-logs:
	@docker-compose logs -f

# Code quality
lint:
	@golangci-lint run

fmt:
	@gofmt -w .

tidy:
	@go mod tidy

# Build for production
build-prod:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api cmd/api/main.go

help:
	@echo "Available commands:"
	@echo "  run              - Run the application"
	@echo "  dev              - Run with hot reload (air)"
	@echo "  build            - Build the application"
	@echo "  test             - Run tests"
	@echo "  migrate          - Run database migrations"
	@echo "  rollback         - Rollback migrations"
	@echo "  seed             - Run database seeds"
	@echo "  migration - Create new migration file"
	@echo "  docker-up        - Start docker containers"
	@echo "  docker-down      - Stop docker containers"
	@echo "  lint             - Run linter"
	@echo "  fmt              - Format code"