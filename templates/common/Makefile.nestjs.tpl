# {{.ProjectName}} Makefile

.PHONY: help install build start dev test lint format clean docker-build docker-run docker-push deploy

# Variables
PROJECT_NAME={{.ProjectName}}
GCP_PROJECT={{.GCPProject}}
IMAGE_NAME=gcr.io/$(GCP_PROJECT)/$(PROJECT_NAME)
VERSION=$(shell git rev-parse --short HEAD)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	npm install

build: ## Build the application
	npm run build

start: ## Start the application in production mode
	npm run start:prod

dev: ## Start the application in development mode
	npm run start:dev

test: ## Run tests
	npm run test

test-e2e: ## Run e2e tests
	npm run test:e2e

test-cov: ## Run tests with coverage
	npm run test:cov

lint: ## Run linter
	npm run lint

format: ## Format code
	npm run format

clean: ## Clean build artifacts
	rm -rf dist
	rm -rf node_modules
	rm -rf coverage

# Docker commands
docker-build: ## Build Docker image
	docker build -t $(IMAGE_NAME):$(VERSION) .
	docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

docker-run: ## Run Docker container locally
	docker run -p {{.Port}}:{{.Port}} --env-file .env $(IMAGE_NAME):latest

docker-push: ## Push Docker image to GCR
	docker push $(IMAGE_NAME):$(VERSION)
	docker push $(IMAGE_NAME):latest

# GCP commands
gcp-configure: ## Configure GCP CLI
	gcloud config set project $(GCP_PROJECT)
	gcloud auth configure-docker

deploy: docker-build docker-push ## Deploy to Google Cloud Run
	gcloud run deploy $(PROJECT_NAME) \
		--image $(IMAGE_NAME):$(VERSION) \
		--platform managed \
		--region us-central1 \
		--allow-unauthenticated \
		--set-env-vars GCP_PROJECT_ID=$(GCP_PROJECT)

# Development commands
setup: install ## Setup development environment
	@echo "Setting up {{.ProjectName}} development environment..."
	@echo "Creating .env file..."
	@echo "PORT={{.Port}}" > .env
	@echo "MONGODB_URI={{.DatabaseType}}://localhost:27017/{{.ProjectName}}" >> .env
	@echo "GCP_PROJECT_ID={{.GCPProject}}" >> .env
	@echo "NODE_ENV=development" >> .env
	@echo "Setup complete!"

logs: ## View application logs
	docker logs -f $(PROJECT_NAME) 2>/dev/null || echo "Container not running"

health: ## Check application health
	curl -f http://localhost:{{.Port}}/api || echo "Application not responding"
