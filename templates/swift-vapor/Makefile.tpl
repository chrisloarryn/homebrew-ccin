# Makefile for {{.ProjectName}} (Swift Vapor)

.PHONY: help deps build run test clean docker-build docker-run docker-run-detached fmt

APP_NAME={{.ProjectName}}
VERSION?=latest
DOCKER_IMAGE={{.ProjectName}}:$(VERSION)
{{- if .GCPProject}}
GCP_PROJECT={{.GCPProject}}
DOCKER_REGISTRY=gcr.io/$(GCP_PROJECT)
{{- end}}

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
deps: ## Resolve SwiftPM dependencies
	swift package resolve

fmt: ## Format Swift code (swift-format if installed)
	@which swift-format >/dev/null 2>&1 && swift-format format -i -r Sources || echo "swift-format not installed"

build: ## Build the app (debug)
	swift build

build-release: ## Build the app (release)
	swift build -c release

run: ## Run the app
	swift run Run

watch: ## Run with auto-reload using watchexec if installed
	@which watchexec >/dev/null 2>&1 && watchexec -r -- swift run Run || echo "watchexec not installed"

test: ## Run tests
	swift test -v

# Docker
docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	docker run -p {{.Port}}:{{.Port}} {{- if .WithGRPC}} -p 50051:50051{{- end}} $(DOCKER_IMAGE)

docker-run-detached: ## Run Docker container in background
	docker run -d -p {{.Port}}:{{.Port}} {{- if .WithGRPC}} -p 50051:50051{{- end}} $(DOCKER_IMAGE)

{{- if .GCPProject}}
docker-push: ## Push Docker image to GCR
	docker tag $(DOCKER_IMAGE) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)
	docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)
{{- end}}

clean: ## Clean build artifacts
	rm -rf .build
