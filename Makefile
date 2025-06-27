# ChrisLoarryn CLI Makefile

.PHONY: help build install clean test demo

BINARY_NAME=chrisloarryn-cli
VERSION=$(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the CLI binary
	go build $(LDFLAGS) -o $(BINARY_NAME) main.go

install: build ## Install the CLI to GOPATH/bin
	cp $(BINARY_NAME) $(GOPATH)/bin/

clean: ## Clean build artifacts and test projects
	rm -f $(BINARY_NAME)
	rm -rf my-*-api/
	rm -rf test-*/
	rm -rf example-*/

test: ## Run tests
	go test -v ./...

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run

deps: ## Download dependencies
	go mod download
	go mod tidy

# Demo commands
demo-nestjs: build ## Generate a demo NestJS project
	./$(BINARY_NAME) generate nestjs demo-nestjs-api --domain user --gcp-project demo-project
	@echo "✅ Demo NestJS project created in demo-nestjs-api/"

demo-gin: build ## Generate a demo Go Gin project with gRPC
	./$(BINARY_NAME) generate go-gin demo-gin-api --domain product --gcp-project demo-project --grpc
	@echo "✅ Demo Go Gin project created in demo-gin-api/"

demo-fiber: build ## Generate a demo Go Fiber project
	./$(BINARY_NAME) generate go-fiber demo-fiber-api --domain order --gcp-project demo-project
	@echo "✅ Demo Go Fiber project created in demo-fiber-api/"

demo-all: demo-nestjs demo-gin demo-fiber ## Generate all demo projects
	@echo "🎉 All demo projects created successfully!"

# Development
dev: ## Run the CLI in development mode
	go run main.go

# Documentation
docs: ## Generate documentation
	@echo "Documentation is in README.md"

# Release
release: clean test build ## Prepare for release
	@echo "Release candidate built: $(BINARY_NAME)"
	@echo "Version: $(VERSION)"

# Check CLI functionality
check: build ## Test basic CLI functionality
	./$(BINARY_NAME) --help
	./$(BINARY_NAME) generate --help
	@echo "✅ CLI basic functionality test passed"
