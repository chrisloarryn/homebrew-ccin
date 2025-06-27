# Makefile for {{.ProjectName}} (Go Fiber)

.PHONY: help build run test clean docker-build docker-run docker-push fmt vet lint deps

# Variables
APP_NAME={{.ProjectName}}
VERSION?=latest
DOCKER_IMAGE={{.ProjectName}}:$(VERSION)
{{- if .GCPProject}}
GCP_PROJECT={{.GCPProject}}
DOCKER_REGISTRY=gcr.io/$(GCP_PROJECT)
{{- end}}

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
deps: ## Install dependencies
	go mod download
	go mod tidy

fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run golangci-lint
	golangci-lint run

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Build
build: ## Build the application
	go build -o bin/$(APP_NAME) .

build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux .

# Run
run: ## Run the application
	go run main.go

dev: ## Run with air for hot reload (requires air: go install github.com/cosmtrek/air@latest)
	air

# Database
db-migrate: ## Run database migrations (placeholder)
	@echo "Database migrations not implemented yet"

db-reset: ## Reset database (placeholder)
	@echo "Database reset not implemented yet"

# Docker
docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	docker run -p {{.Port}}:{{.Port}} {{- if .WithGRPC}} -p 50051:50051{{- end}} --env-file .env $(DOCKER_IMAGE)

docker-run-detached: ## Run Docker container in background
	docker run -d -p {{.Port}}:{{.Port}} {{- if .WithGRPC}} -p 50051:50051{{- end}} --env-file .env $(DOCKER_IMAGE)

{{- if .GCPProject}}
docker-push: ## Push Docker image to GCR
	docker tag $(DOCKER_IMAGE) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)
	docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)

gcp-deploy: docker-build docker-push ## Build and deploy to GCP
	@echo "Deploying to GCP..."
	@echo "Image: $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)"
{{- end}}

# Cleanup
clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html
	docker image prune -f

clean-all: clean ## Clean everything including Docker images
	docker rmi $(DOCKER_IMAGE) 2>/dev/null || true

# Development tools
install-tools: ## Install development tools
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Production
prod-build: ## Build for production
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/$(APP_NAME) .

# Health check
health: ## Check application health
	curl -f http://localhost:{{.Port}}/health || exit 1
