package cmd

import (
	"fmt"
	"path/filepath"
)

func generateGoFiberProject(projectName, domainName, gcpProject string, grpc bool) error {
	if err := createProjectDir(projectName); err != nil {
		return err
	}

	// Set default domain if not provided
	if domainName == "" {
		domainName = "item"
	}

	// Generate go.mod for Fiber
	goMod := replaceTemplateVars(`module {{PROJECT_NAME}}

go 1.24.4

require (
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/gofiber/swagger v1.1.0
	github.com/swaggo/swag v1.16.2
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gorm.io/gorm v1.25.12
	gorm.io/driver/postgres v1.5.9
	cloud.google.com/go/monitoring v1.20.4
	cloud.google.com/go/logging v1.11.0
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
)`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "go.mod"), goMod); err != nil {
		return err
	}

	// Generate main.go for Fiber
	mainGo := generateFiberMainGoFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "cmd/server/main.go"), mainGo); err != nil {
		return err
	}

	// Generate model (same as Gin)
	model := generateModelFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/models/%s.go", domainName)), model); err != nil {
		return err
	}

	// Generate repository (same as Gin)
	repository := generateRepositoryFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/repository/%s_repository.go", domainName)), repository); err != nil {
		return err
	}

	// Generate service (same as Gin)
	service := generateServiceFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/service/%s_service.go", domainName)), service); err != nil {
		return err
	}

	// Generate Fiber handler
	handler := generateFiberHandlerFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/handlers/%s_handler.go", domainName)), handler); err != nil {
		return err
	}

	// Generate Fiber middleware
	middleware := generateFiberMiddlewareFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "internal/middleware/metrics.go"), middleware); err != nil {
		return err
	}

	// Generate config (same as Gin)
	config := generateConfigFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "internal/config/config.go"), config); err != nil {
		return err
	}

	// Generate Fiber router
	router := generateFiberRouterFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "internal/router/router.go"), router); err != nil {
		return err
	}

	// Generate gRPC files if requested (same as Gin)
	if grpc {
		if err := generateGRPCFiles(projectName, domainName, gcpProject); err != nil {
			return err
		}
	}

	// Generate Dockerfile (same as Gin)
	dockerfile := generateDockerfileGo(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "Dockerfile"), dockerfile); err != nil {
		return err
	}

	// Generate Makefile for Fiber
	makefile := generateFiberMakefileGo(projectName, domainName, gcpProject, grpc)
	if err := createFile(filepath.Join(projectName, "Makefile"), makefile); err != nil {
		return err
	}

	// Generate .env.example (same as Gin)
	envExample := generateEnvExampleGo(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, ".env.example"), envExample); err != nil {
		return err
	}

	// Generate README.md for Fiber
	readme := generateFiberReadmeGo(projectName, domainName, gcpProject, grpc)
	if err := createFile(filepath.Join(projectName, "README.md"), readme); err != nil {
		return err
	}

	return nil
}

func generateFiberMainGoFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package main

import (
	"log"
	"{{PROJECT_NAME}}/internal/config"
	"{{PROJECT_NAME}}/internal/router"
	"{{PROJECT_NAME}}/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// @title {{PROJECT_NAME}} API
// @version 1.0
// @description {{PROJECT_NAME}} CRUD API with Fiber framework
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Add global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(middleware.GCPMetrics(cfg.GCPProjectID))

	// Setup routes
	router.SetupRoutes(app, cfg)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}`, projectName, domainName, gcpProject)
}

func generateFiberHandlerFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package handlers

import (
	"strconv"
	"{{PROJECT_NAME}}/internal/models"
	"{{PROJECT_NAME}}/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// {{DOMAIN_TITLE}}Handler handles HTTP requests for {{DOMAIN_LOWER}} operations
type {{DOMAIN_TITLE}}Handler struct {
	service service.{{DOMAIN_TITLE}}Service
}

// New{{DOMAIN_TITLE}}Handler creates a new instance of {{DOMAIN_TITLE}}Handler
func New{{DOMAIN_TITLE}}Handler(service service.{{DOMAIN_TITLE}}Service) *{{DOMAIN_TITLE}}Handler {
	return &{{DOMAIN_TITLE}}Handler{service: service}
}

// Create{{DOMAIN_TITLE}} godoc
// @Summary Create a new {{DOMAIN_LOWER}}
// @Description Create a new {{DOMAIN_LOWER}} with the input payload
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param {{DOMAIN_LOWER}} body models.Create{{DOMAIN_TITLE}}Request true "{{DOMAIN_TITLE}} data"
// @Success 201 {object} models.{{DOMAIN_TITLE}}Response
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s [post]
func (h *{{DOMAIN_TITLE}}Handler) Create{{DOMAIN_TITLE}}(c *fiber.Ctx) error {
	var req models.Create{{DOMAIN_TITLE}}Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	{{DOMAIN_LOWER}}, err := h.service.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON({{DOMAIN_LOWER}})
}

// Get{{DOMAIN_TITLE}}s godoc
// @Summary Get all {{DOMAIN_LOWER}}s
// @Description Get a list of all {{DOMAIN_LOWER}}s
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Success 200 {array} models.{{DOMAIN_TITLE}}Response
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s [get]
func (h *{{DOMAIN_TITLE}}Handler) Get{{DOMAIN_TITLE}}s(c *fiber.Ctx) error {
	{{DOMAIN_LOWER}}s, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON({{DOMAIN_LOWER}}s)
}

// Get{{DOMAIN_TITLE}} godoc
// @Summary Get a {{DOMAIN_LOWER}} by ID
// @Description Get a single {{DOMAIN_LOWER}} by its ID
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param id path string true "{{DOMAIN_TITLE}} ID"
// @Success 200 {object} models.{{DOMAIN_TITLE}}Response
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s/{id} [get]
func (h *{{DOMAIN_TITLE}}Handler) Get{{DOMAIN_TITLE}}(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	{{DOMAIN_LOWER}}, err := h.service.GetByID(id)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON({{DOMAIN_LOWER}})
}

// Update{{DOMAIN_TITLE}} godoc
// @Summary Update a {{DOMAIN_LOWER}}
// @Description Update a {{DOMAIN_LOWER}} by its ID
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param id path string true "{{DOMAIN_TITLE}} ID"
// @Param {{DOMAIN_LOWER}} body models.Update{{DOMAIN_TITLE}}Request true "{{DOMAIN_TITLE}} update data"
// @Success 200 {object} models.{{DOMAIN_TITLE}}Response
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s/{id} [put]
func (h *{{DOMAIN_TITLE}}Handler) Update{{DOMAIN_TITLE}}(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var req models.Update{{DOMAIN_TITLE}}Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	{{DOMAIN_LOWER}}, err := h.service.Update(id, &req)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON({{DOMAIN_LOWER}})
}

// Delete{{DOMAIN_TITLE}} godoc
// @Summary Delete a {{DOMAIN_LOWER}}
// @Description Delete a {{DOMAIN_LOWER}} by its ID
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param id path string true "{{DOMAIN_TITLE}} ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s/{id} [delete]
func (h *{{DOMAIN_TITLE}}Handler) Delete{{DOMAIN_TITLE}}(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	err = h.service.Delete(id)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}`, projectName, domainName, gcpProject)
}

func generateFiberMiddlewareFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package middleware

import (
	"context"
	"log"
	"strconv"
	"time"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GCPMetrics returns a Fiber middleware that sends metrics to Google Cloud Monitoring
func GCPMetrics(projectID string) fiber.Handler {
	if projectID == "" {
		// Return a no-op middleware if no project ID is provided
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	ctx := context.Background()
	
	// Initialize monitoring client
	monitoringClient, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		log.Printf("Failed to create monitoring client: %v", err)
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	// Initialize logging client
	loggingClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create logging client: %v", err)
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	logger := loggingClient.Logger("{{PROJECT_NAME}}-requests")

	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)
		statusCode := c.Response().StatusCode()

		// Send metrics to GCP
		go func() {
			if err := sendFiberMetrics(monitoringClient, projectID, c.Method(), c.Path(), statusCode, duration); err != nil {
				log.Printf("Failed to send metrics: %v", err)
			}
		}()

		// Send logs to GCP
		go func() {
			logger.Log(logging.Entry{
				Severity: getFiberSeverity(statusCode),
				Payload: map[string]interface{}{
					"method":      c.Method(),
					"path":        c.Path(),
					"status_code": statusCode,
					"duration_ms": duration.Milliseconds(),
					"user_agent":  c.Get("User-Agent"),
					"remote_addr": c.IP(),
					"timestamp":   time.Now().UTC(),
				},
			})
		}()

		return err
	}
}

func sendFiberMetrics(client *monitoring.MetricClient, projectID, method, path string, statusCode int, duration time.Duration) error {
	ctx := context.Background()
	
	// Create time series data
	now := time.Now()
	timeSeries := &monitoringpb.TimeSeries{
		Metric: &monitoringpb.Metric{
			Type: "custom.googleapis.com/{{PROJECT_NAME}}/request_duration",
			Labels: map[string]string{
				"method":      method,
				"endpoint":    path,
				"status_code": strconv.Itoa(statusCode),
			},
		},
		Resource: &monitoringpb.MonitoredResource{
			Type: "generic_node",
			Labels: map[string]string{
				"location":  "global",
				"namespace": "{{PROJECT_NAME}}",
				"node_id":   "default",
			},
		},
		Points: []*monitoringpb.Point{
			{
				Interval: &monitoringpb.TimeInterval{
					EndTime: timestamppb.New(now),
				},
				Value: &monitoringpb.TypedValue{
					Value: &monitoringpb.TypedValue_DoubleValue{
						DoubleValue: float64(duration.Milliseconds()),
					},
				},
			},
		},
	}

	// Send time series
	req := &monitoringpb.CreateTimeSeriesRequest{
		Name:       "projects/" + projectID,
		TimeSeries: []*monitoringpb.TimeSeries{timeSeries},
	}

	return client.CreateTimeSeries(ctx, req)
}

func getFiberSeverity(statusCode int) logging.Severity {
	switch {
	case statusCode >= 500:
		return logging.Error
	case statusCode >= 400:
		return logging.Warning
	default:
		return logging.Info
	}
}`, projectName, domainName, gcpProject)
}

func generateFiberRouterFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package router

import (
	"{{PROJECT_NAME}}/internal/config"
	"{{PROJECT_NAME}}/internal/handlers"
	"{{PROJECT_NAME}}/internal/repository"
	"{{PROJECT_NAME}}/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"{{PROJECT_NAME}}/internal/models"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// Initialize database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&models.{{DOMAIN_TITLE}}{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	{{DOMAIN_LOWER}}Repo := repository.New{{DOMAIN_TITLE}}Repository(db)

	// Initialize services
	{{DOMAIN_LOWER}}Service := service.New{{DOMAIN_TITLE}}Service({{DOMAIN_LOWER}}Repo)

	// Initialize handlers
	{{DOMAIN_LOWER}}Handler := handlers.New{{DOMAIN_TITLE}}Handler({{DOMAIN_LOWER}}Service)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API routes
	api := app.Group("/api/v1")
	
	{{DOMAIN_LOWER}}Routes := api.Group("/{{DOMAIN_LOWER}}s")
	{{DOMAIN_LOWER}}Routes.Post("/", {{DOMAIN_LOWER}}Handler.Create{{DOMAIN_TITLE}})
	{{DOMAIN_LOWER}}Routes.Get("/", {{DOMAIN_LOWER}}Handler.Get{{DOMAIN_TITLE}}s)
	{{DOMAIN_LOWER}}Routes.Get("/:id", {{DOMAIN_LOWER}}Handler.Get{{DOMAIN_TITLE}})
	{{DOMAIN_LOWER}}Routes.Put("/:id", {{DOMAIN_LOWER}}Handler.Update{{DOMAIN_TITLE}})
	{{DOMAIN_LOWER}}Routes.Delete("/:id", {{DOMAIN_LOWER}}Handler.Delete{{DOMAIN_TITLE}})
}`, projectName, domainName, gcpProject)
}

func generateFiberMakefileGo(projectName, domainName, gcpProject string, grpc bool) string {
	grpcCommands := ""
	if grpc {
		grpcCommands = `
# gRPC commands
proto-gen: ## Generate protobuf files
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/proto/$(DOMAIN_NAME).proto

proto-install: ## Install protobuf compiler and Go plugins
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
	}

	return replaceTemplateVars(`# {{PROJECT_NAME}} Makefile

.PHONY: help build run test clean docker-build docker-run docker-push deploy proto-gen proto-install

# Variables
PROJECT_NAME={{PROJECT_NAME}}
DOMAIN_NAME={{DOMAIN_LOWER}}
GCP_PROJECT={{GCP_PROJECT}}
IMAGE_NAME=gcr.io/$(GCP_PROJECT)/$(PROJECT_NAME)
VERSION=$(shell git rev-parse --short HEAD)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development commands
deps: ## Download dependencies
	go mod download
	go mod tidy

build: ## Build the application
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(PROJECT_NAME) cmd/server/main.go

run: ## Run the application
	go run cmd/server/main.go

dev: ## Run the application in development mode with hot reload
	go run cmd/server/main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	go fmt ./...

clean: ## Clean build artifacts
	rm -f bin/$(PROJECT_NAME)
	rm -f coverage.out coverage.html

# Swagger commands
swagger-gen: ## Generate Swagger documentation
	swag init -g cmd/server/main.go -o docs

swagger-install: ## Install Swagger CLI
	go install github.com/swaggo/swag/cmd/swag@latest`+grpcCommands+`

# Docker commands
docker-build: ## Build Docker image
	docker build -t $(IMAGE_NAME):$(VERSION) .
	docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

docker-run: ## Run Docker container locally
	docker run -p 8080:8080 --env-file .env $(IMAGE_NAME):latest

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

# Database commands
db-migrate: ## Run database migrations
	@echo "Database migrations would run here"

# Development setup
setup: deps ## Setup development environment
	@echo "Setting up {{PROJECT_NAME}} development environment..."
	@echo "Creating .env file..."
	@echo "PORT=8080" > .env
	@echo "ENVIRONMENT=development" >> .env
	@echo "DATABASE_URL=postgres://user:password@localhost:5432/{{PROJECT_NAME}}?sslmode=disable" >> .env
	@echo "GCP_PROJECT_ID={{GCP_PROJECT}}" >> .env
	@echo "Setup complete!"

logs: ## View application logs
	docker logs -f $(PROJECT_NAME) 2>/dev/null || echo "Container not running"

health: ## Check application health
	curl -f http://localhost:8080/health || echo "Application not responding"`, projectName, domainName, gcpProject)
}

func generateFiberReadmeGo(projectName, domainName, gcpProject string, grpc bool) string {
	grpcDocs := ""
	if grpc {
		grpcDocs = `
### gRPC

This project also includes gRPC support. To generate protobuf files:

'''bash
make proto-install  # Install protobuf tools (first time only)
make proto-gen      # Generate Go files from .proto files
'''

gRPC endpoints:
- 'Create{{DOMAIN_TITLE}}'
- 'Get{{DOMAIN_TITLE}}'
- 'List{{DOMAIN_TITLE}}s'
- 'Update{{DOMAIN_TITLE}}'
- 'Delete{{DOMAIN_TITLE}}'`
	}

	return replaceTemplateVars(`# {{PROJECT_NAME}}

{{PROJECT_NAME}} - Go CRUD API with Fiber framework and GCP integration

## Features

- ⚡ Go with Fiber framework (Express-inspired, ultra fast)
- 🗄️ PostgreSQL integration with GORM
- 📈 GCP Monitoring and Logging
- 🐳 Docker support
- 📚 Swagger API documentation
- 🧪 Testing setup
- 🏗️ Clean architecture
- 🔄 gRPC support (optional)

## Domain: {{DOMAIN_TITLE}}

This API provides CRUD operations for {{DOMAIN_TITLE}} entities.

## Quick Start

### Prerequisites

- Go 1.24.4+
- PostgreSQL
- Docker (optional)
- GCP account (for metrics)

### Installation

1. Clone the repository
2. Install dependencies:
   '''bash
   make deps
   '''

3. Setup environment:
   '''bash
   make setup
   '''

4. Update the '.env' file with your configuration

### Running the Application

Development mode:
'''bash
make dev
'''

Production mode:
'''bash
make build
./bin/{{PROJECT_NAME}}
'''

### API Documentation

Once the application is running, visit:
- Swagger UI: http://localhost:8080/swagger/index.html
- Health check: http://localhost:8080/health

### API Endpoints

- 'GET /api/v1/{{DOMAIN_LOWER}}s' - Get all {{DOMAIN_LOWER}}s
- 'GET /api/v1/{{DOMAIN_LOWER}}s/:id' - Get {{DOMAIN_LOWER}} by ID
- 'POST /api/v1/{{DOMAIN_LOWER}}s' - Create new {{DOMAIN_LOWER}}
- 'PUT /api/v1/{{DOMAIN_LOWER}}s/:id' - Update {{DOMAIN_LOWER}}
- 'DELETE /api/v1/{{DOMAIN_LOWER}}s/:id' - Delete {{DOMAIN_LOWER}}`+grpcDocs+`

### Testing

Run tests:
'''bash
make test
'''

Run tests with coverage:
'''bash
make test-coverage
'''

### Docker

Build and run with Docker:
'''bash
make docker-build
make docker-run
'''

### Deployment

Deploy to Google Cloud Run:
'''bash
make gcp-configure
make deploy
'''

### Generate Swagger docs

'''bash
make swagger-install  # Install swag CLI (first time only)
make swagger-gen      # Generate swagger docs
'''

## Project Structure

'''
cmd/
└── server/
    └── main.go
internal/
├── config/
│   └── config.go
├── handlers/
│   └── {{DOMAIN_LOWER}}_handler.go
├── middleware/
│   └── metrics.go
├── models/
│   └── {{DOMAIN_LOWER}}.go
├── repository/
│   └── {{DOMAIN_LOWER}}_repository.go
├── router/
│   └── router.go
└── service/
    └── {{DOMAIN_LOWER}}_service.go
api/
└── proto/
    └── {{DOMAIN_LOWER}}.proto
'''

## Environment Variables

See '.env.example' for all available environment variables.

## GCP Integration

This project includes GCP Monitoring and Logging integration:

- Custom metrics for request duration
- Structured logging
- Error tracking
- Performance monitoring

## Why Fiber?

Fiber is an Express-inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.

## License

MIT`, projectName, domainName, gcpProject)
}
