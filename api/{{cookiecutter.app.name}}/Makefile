.DEFAULT_GOAL := help

# Create Make variables from .env
include .env

# Docker Images
GOLANGCI_LINT_IMAGE := golangci/golangci-lint:latest
OPENAPI_SPEC_VALIDATOR_IMAGE := p1c2u/openapi-spec-validator:latest

.PHONY: help
help: ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: run
run: ## Run the application on the host
	go run main.go

.PHONY: bootstrap
bootstrap: clean build start migrate ## Bootstrap the application

.PHONY: clean
clean: ## Stop and remove containers, volumes
	docker-compose down --volumes

.PHONY: stop
stop: ## Stop and remove containers
	docker-compose down

.PHONY: build
build: ## Build containers
	docker-compose build

.PHONY: start
start: ## Start containers
	docker-compose up --detach database

.PHONY: restart
restart: docker-compose.yaml ## Restart containers
	docker-compose restart

.PHONY: logs
logs: docker-compose.yaml ## Show logs from containers
	docker-compose logs --follow

FLAGS :=
.PHONY: test
test: ## Run tests. Use `make FLAGS=<flags> test` to pass custom flags
	go test -shuffle=on -race -count=1 $(FLAGS) $(COVERAGE) ./...

.PHONY: coverage
coverage: COVERAGE = -cover -coverprofile=coverage.out
coverage: test

.PHONY: migrate
migrate: db-migrate db-migrate-test-data ## Run database migrations

.PHONY: db-migrate
db-migrate:
	@docker-compose run migrate -url="jdbc:mysql://$(MYSQL_HOST):3306/$(MYSQL_DATABASE)?timeout=3600s&collation=utf8mb4_unicode_ci" -user="$(MYSQL_USER)" -password="$(MYSQL_PASSWORD)" -connectRetries=3 -locations="filesystem:/migrations" -outOfOrder=true -ignoreMigrationPatterns='*:missing,*:future' migrate

.PHONY: db-migrate-test-data
db-migrate-test-data:
	@docker-compose run migrate -url="jdbc:mysql://$(MYSQL_HOST):3306/$(MYSQL_DATABASE)?timeout=3600s&collation=utf8mb4_unicode_ci" -user="$(MYSQL_USER)" -password="$(MYSQL_PASSWORD)" -connectRetries=3 -locations="filesystem:/testdata/migrations" -outOfOrder=true -ignoreMigrationPatterns='*:missing,*:future' migrate

.PHONY: go-lint
go-lint: ## Run Go linters
	docker run -v "$(shell pwd):/app" -w /app $(GOLANGCI_LINT_IMAGE) golangci-lint -v --timeout=5m run ./...

.PHONY: openapi-lint
openapi-lint: api ## Run OpenAPI linter
	docker run -v "${PWD}/api/spec/openapi.yaml:/openapi.yaml" --rm $(OPENAPI_SPEC_VALIDATOR_IMAGE) /openapi.yaml

.PHONY: lint
lint: go-lint openapi-lint ## Run all linters