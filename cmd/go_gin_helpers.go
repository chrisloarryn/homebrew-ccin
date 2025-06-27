package cmd

import (
	"fmt"
	"path/filepath"
)

func generateMiddlewareFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package middleware

import (
	"context"
	"log"
	"strconv"
	"time"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GCPMetrics returns a Gin middleware that sends metrics to Google Cloud Monitoring
func GCPMetrics(projectID string) gin.HandlerFunc {
	if projectID == "" {
		// Return a no-op middleware if no project ID is provided
		return gin.HandlerFunc(func(c *gin.Context) {
			c.Next()
		})
	}

	ctx := context.Background()
	
	// Initialize monitoring client
	monitoringClient, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		log.Printf("Failed to create monitoring client: %v", err)
		return gin.HandlerFunc(func(c *gin.Context) {
			c.Next()
		})
	}

	// Initialize logging client
	loggingClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create logging client: %v", err)
		return gin.HandlerFunc(func(c *gin.Context) {
			c.Next()
		})
	}

	logger := loggingClient.Logger("{{PROJECT_NAME}}-requests")

	return gin.HandlerFunc(func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)
		statusCode := c.Writer.Status()

		// Send metrics to GCP
		go func() {
			if err := sendMetrics(monitoringClient, projectID, c.Request.Method, c.FullPath(), statusCode, duration); err != nil {
				log.Printf("Failed to send metrics: %v", err)
			}
		}()

		// Send logs to GCP
		go func() {
			logger.Log(logging.Entry{
				Severity: getSeverity(statusCode),
				Payload: map[string]interface{}{
					"method":      c.Request.Method,
					"path":        c.FullPath(),
					"status_code": statusCode,
					"duration_ms": duration.Milliseconds(),
					"user_agent":  c.Request.UserAgent(),
					"remote_addr": c.ClientIP(),
					"timestamp":   time.Now().UTC(),
				},
			})
		}()
	})
}

func sendMetrics(client *monitoring.MetricClient, projectID, method, path string, statusCode int, duration time.Duration) error {
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
		Name:       fmt.Sprintf("projects/%s", projectID),
		TimeSeries: []*monitoringpb.TimeSeries{timeSeries},
	}

	return client.CreateTimeSeries(ctx, req)
}

func getSeverity(statusCode int) logging.Severity {
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

func generateConfigFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Port         string
	Environment  string
	DatabaseURL  string
	GCPProjectID string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/{{PROJECT_NAME}}?sslmode=disable"),
		GCPProjectID: getEnv("GCP_PROJECT_ID", "{{GCP_PROJECT}}"),
	}
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}`, projectName, domainName, gcpProject)
}

func generateRouterFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package router

import (
	"{{PROJECT_NAME}}/internal/config"
	"{{PROJECT_NAME}}/internal/handlers"
	"{{PROJECT_NAME}}/internal/repository"
	"{{PROJECT_NAME}}/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"{{PROJECT_NAME}}/internal/models"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(r *gin.Engine, cfg *config.Config) {
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
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api/v1")
	{
		{{DOMAIN_LOWER}}Routes := api.Group("/{{DOMAIN_LOWER}}s")
		{
			{{DOMAIN_LOWER}}Routes.POST("", {{DOMAIN_LOWER}}Handler.Create{{DOMAIN_TITLE}})
			{{DOMAIN_LOWER}}Routes.GET("", {{DOMAIN_LOWER}}Handler.Get{{DOMAIN_TITLE}}s)
			{{DOMAIN_LOWER}}Routes.GET("/:id", {{DOMAIN_LOWER}}Handler.Get{{DOMAIN_TITLE}})
			{{DOMAIN_LOWER}}Routes.PUT("/:id", {{DOMAIN_LOWER}}Handler.Update{{DOMAIN_TITLE}})
			{{DOMAIN_LOWER}}Routes.DELETE("/:id", {{DOMAIN_LOWER}}Handler.Delete{{DOMAIN_TITLE}})
		}
	}
}`, projectName, domainName, gcpProject)
}

func generateGRPCFiles(projectName, domainName, gcpProject string) error {
	// Generate proto file
	protoContent := replaceTemplateVars(`syntax = "proto3";

package {{DOMAIN_LOWER}};
option go_package = "./pb";

import "google/protobuf/timestamp.proto";

service {{DOMAIN_TITLE}}Service {
  rpc Create{{DOMAIN_TITLE}}(Create{{DOMAIN_TITLE}}Request) returns ({{DOMAIN_TITLE}}Response);
  rpc Get{{DOMAIN_TITLE}}(Get{{DOMAIN_TITLE}}Request) returns ({{DOMAIN_TITLE}}Response);
  rpc List{{DOMAIN_TITLE}}s(List{{DOMAIN_TITLE}}sRequest) returns (List{{DOMAIN_TITLE}}sResponse);
  rpc Update{{DOMAIN_TITLE}}(Update{{DOMAIN_TITLE}}Request) returns ({{DOMAIN_TITLE}}Response);
  rpc Delete{{DOMAIN_TITLE}}(Delete{{DOMAIN_TITLE}}Request) returns (DeleteResponse);
}

message {{DOMAIN_TITLE}} {
  string id = 1;
  string name = 2;
  optional string description = 3;
  bool is_active = 4;
  google.protobuf.Timestamp created_at = 5;
  google.protobuf.Timestamp updated_at = 6;
}

message Create{{DOMAIN_TITLE}}Request {
  string name = 1;
  optional string description = 2;
  optional bool is_active = 3;
}

message Get{{DOMAIN_TITLE}}Request {
  string id = 1;
}

message List{{DOMAIN_TITLE}}sRequest {}

message Update{{DOMAIN_TITLE}}Request {
  string id = 1;
  optional string name = 2;
  optional string description = 3;
  optional bool is_active = 4;
}

message Delete{{DOMAIN_TITLE}}Request {
  string id = 1;
}

message {{DOMAIN_TITLE}}Response {
  {{DOMAIN_TITLE}} {{DOMAIN_LOWER}} = 1;
}

message List{{DOMAIN_TITLE}}sResponse {
  repeated {{DOMAIN_TITLE}} {{DOMAIN_LOWER}}s = 1;
}

message DeleteResponse {
  bool success = 1;
}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("api/proto/%s.proto", domainName)), protoContent); err != nil {
		return err
	}

	// Generate gRPC server
	grpcServer := replaceTemplateVars(`package grpc

import (
	"context"
	"{{PROJECT_NAME}}/internal/service"
	"{{PROJECT_NAME}}/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// {{DOMAIN_TITLE}}Server implements the gRPC {{DOMAIN_TITLE}}Service
type {{DOMAIN_TITLE}}Server struct {
	pb.Unimplemented{{DOMAIN_TITLE}}ServiceServer
	service service.{{DOMAIN_TITLE}}Service
}

// New{{DOMAIN_TITLE}}Server creates a new gRPC server for {{DOMAIN_LOWER}} operations
func New{{DOMAIN_TITLE}}Server(service service.{{DOMAIN_TITLE}}Service) *{{DOMAIN_TITLE}}Server {
	return &{{DOMAIN_TITLE}}Server{service: service}
}

// Create{{DOMAIN_TITLE}} creates a new {{DOMAIN_LOWER}}
func (s *{{DOMAIN_TITLE}}Server) Create{{DOMAIN_TITLE}}(ctx context.Context, req *pb.Create{{DOMAIN_TITLE}}Request) (*pb.{{DOMAIN_TITLE}}Response, error) {
	// Convert protobuf request to service request
	serviceReq := &models.Create{{DOMAIN_TITLE}}Request{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	{{DOMAIN_LOWER}}, err := s.service.Create(serviceReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create {{DOMAIN_LOWER}}: %v", err)
	}

	return &pb.{{DOMAIN_TITLE}}Response{
		{{DOMAIN_TITLE}}: s.modelToPB({{DOMAIN_LOWER}}),
	}, nil
}

// Get{{DOMAIN_TITLE}} retrieves a {{DOMAIN_LOWER}} by ID
func (s *{{DOMAIN_TITLE}}Server) Get{{DOMAIN_TITLE}}(ctx context.Context, req *pb.Get{{DOMAIN_TITLE}}Request) (*pb.{{DOMAIN_TITLE}}Response, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID format: %v", err)
	}

	{{DOMAIN_LOWER}}, err := s.service.GetByID(id)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			return nil, status.Errorf(codes.NotFound, "{{DOMAIN_LOWER}} not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get {{DOMAIN_LOWER}}: %v", err)
	}

	return &pb.{{DOMAIN_TITLE}}Response{
		{{DOMAIN_TITLE}}: s.modelToPB({{DOMAIN_LOWER}}),
	}, nil
}

// List{{DOMAIN_TITLE}}s retrieves all {{DOMAIN_LOWER}}s
func (s *{{DOMAIN_TITLE}}Server) List{{DOMAIN_TITLE}}s(ctx context.Context, req *pb.List{{DOMAIN_TITLE}}sRequest) (*pb.List{{DOMAIN_TITLE}}sResponse, error) {
	{{DOMAIN_LOWER}}s, err := s.service.GetAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list {{DOMAIN_LOWER}}s: %v", err)
	}

	pb{{DOMAIN_TITLE}}s := make([]*pb.{{DOMAIN_TITLE}}, len({{DOMAIN_LOWER}}s))
	for i, {{DOMAIN_LOWER}} := range {{DOMAIN_LOWER}}s {
		pb{{DOMAIN_TITLE}}s[i] = s.modelToPB(&{{DOMAIN_LOWER}})
	}

	return &pb.List{{DOMAIN_TITLE}}sResponse{
		{{DOMAIN_TITLE}}s: pb{{DOMAIN_TITLE}}s,
	}, nil
}

// Update{{DOMAIN_TITLE}} updates an existing {{DOMAIN_LOWER}}
func (s *{{DOMAIN_TITLE}}Server) Update{{DOMAIN_TITLE}}(ctx context.Context, req *pb.Update{{DOMAIN_TITLE}}Request) (*pb.{{DOMAIN_TITLE}}Response, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID format: %v", err)
	}

	serviceReq := &models.Update{{DOMAIN_TITLE}}Request{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	{{DOMAIN_LOWER}}, err := s.service.Update(id, serviceReq)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			return nil, status.Errorf(codes.NotFound, "{{DOMAIN_LOWER}} not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update {{DOMAIN_LOWER}}: %v", err)
	}

	return &pb.{{DOMAIN_TITLE}}Response{
		{{DOMAIN_TITLE}}: s.modelToPB({{DOMAIN_LOWER}}),
	}, nil
}

// Delete{{DOMAIN_TITLE}} deletes a {{DOMAIN_LOWER}}
func (s *{{DOMAIN_TITLE}}Server) Delete{{DOMAIN_TITLE}}(ctx context.Context, req *pb.Delete{{DOMAIN_TITLE}}Request) (*pb.DeleteResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID format: %v", err)
	}

	err = s.service.Delete(id)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			return nil, status.Errorf(codes.NotFound, "{{DOMAIN_LOWER}} not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete {{DOMAIN_LOWER}}: %v", err)
	}

	return &pb.DeleteResponse{Success: true}, nil
}

// modelToPB converts a service model to protobuf format
func (s *{{DOMAIN_TITLE}}Server) modelToPB({{DOMAIN_LOWER}} *models.{{DOMAIN_TITLE}}Response) *pb.{{DOMAIN_TITLE}} {
	pb{{DOMAIN_TITLE}} := &pb.{{DOMAIN_TITLE}}{
		Id:        {{DOMAIN_LOWER}}.ID.String(),
		Name:      {{DOMAIN_LOWER}}.Name,
		IsActive:  {{DOMAIN_LOWER}}.IsActive,
		CreatedAt: timestamppb.New({{DOMAIN_LOWER}}.CreatedAt),
		UpdatedAt: timestamppb.New({{DOMAIN_LOWER}}.UpdatedAt),
	}

	if {{DOMAIN_LOWER}}.Description != nil {
		pb{{DOMAIN_TITLE}}.Description = {{DOMAIN_LOWER}}.Description
	}

	return pb{{DOMAIN_TITLE}}
}`, projectName, domainName, gcpProject)

	return createFile(filepath.Join(projectName, fmt.Sprintf("internal/grpc/%s_server.go", domainName)), grpcServer)
}

func generateDockerfileGo(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Change ownership of the app directory
RUN chown -R appuser:appgroup /root/
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./main"]`, projectName, domainName, gcpProject)
}

func generateMakefileGo(projectName, domainName, gcpProject string, grpc bool) string {
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

func generateEnvExampleGo(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`# Application Configuration
PORT=8080
ENVIRONMENT=development

# Database Configuration
DATABASE_URL=postgres://user:password@localhost:5432/{{PROJECT_NAME}}?sslmode=disable

# GCP Configuration
GCP_PROJECT_ID={{GCP_PROJECT}}`, projectName, domainName, gcpProject)
}

func generateReadmeGo(projectName, domainName, gcpProject string, grpc bool) string {
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

{{PROJECT_NAME}} - Go CRUD API with Gin framework and GCP integration

## Features

- 🚀 Go with Gin framework
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

## License

MIT`, projectName, domainName, gcpProject)
}
