# CCIN CLI Usage Examples

This guide shows different use cases of the enhanced CCIN CLI for generating production-ready CRUD applications with beautiful, colorized output.

## 🚀 Getting Started

### First Steps with Enhanced CLI

```bash
# Beautiful help with colors and emojis
ccin --help

# Detailed version information 
ccin --version

# Framework overview with descriptions
ccin generate --help

# Specific framework help
ccin generate nestjs --help
```

## 🎨 Enhanced CLI Features

### Colorized Help Output

```bash
# Beautiful main help with emojis and colors
ccin --help
```

**Example Output:**
```
🎯 CCIN CLI - ChrisLoarryn's Comprehensive Code Integration & Initialization Tool

✨ Generate production-ready CRUD applications with multiple frameworks:
   • NestJS (Node.js 24.2.0 + TypeScript + MongoDB)
   • Go + Gin (REST/gRPC + PostgreSQL + Clean Architecture)
   • Go + Fiber (Ultra-fast REST/gRPC + PostgreSQL)

🎁 What you get:
   ✅ Complete CRUD operations
   ✅ Production-ready Docker configuration
   ✅ GCP metrics & monitoring integration
   ✅ Automated Makefiles for all workflows
   ✅ API documentation (Swagger/OpenAPI)
   ✅ Clean Architecture patterns
   ✅ Template-based code generation

🚀 Quick start: ccin generate --help
```

### Smart Error Handling

The CLI provides intelligent validation with helpful suggestions:

```bash
# Project name too short
ccin generate nestjs a
# Output:
# ❌ Invalid project name: project name must be at least 2 characters long
# 💡 Use a descriptive name like 'my-api', 'user-service', etc.

# Missing required argument
ccin generate nestjs
# Output:
# ❌ Command Error: accepts 1 arg(s), received 0
# 💡 Quick help: ccin --help or ccin generate --help
# 📚 Documentation: https://github.com/chrisloarryn/homebrew-ccin

# Framework-specific validation
ccin generate go-gin my-api
# Output:
# 🚀 Generating Go Gin CRUD project: my-api
# 📊 Domain: item
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# 📝 Processing templates...
# ✅ Go Gin project 'my-api' generated successfully!
```

### Version Information

```bash
ccin --version
```
**Output:**
```
🎯 CCIN CLI - ChrisLoarryn's Comprehensive Code Integration & Initialization Tool
Version: 1.0.0
Author: Chris Loarryn (@chrisloarryn)
Repository: https://github.com/chrisloarryn/homebrew-ccin

✨ Generate production-ready CRUD applications with modern frameworks!
```

## 📦 Basic Projects

### Simple User API

```bash
# NestJS with MongoDB and enhanced output
ccin generate nestjs user-api --domain user --gcp-project my-project

# Go Gin with PostgreSQL  
ccin generate go-gin user-api --domain user --gcp-project my-project

# Go Fiber with PostgreSQL (ultra-fast)
ccin generate go-fiber user-api --domain user --gcp-project my-project
```

**Example Enhanced Output:**
```
🚀 Generating NestJS CRUD project: user-api
📊 Domain: user
☁️  GCP Project: my-project
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📝 Processing templates...
✅ NestJS project 'user-api' generated successfully!

🎯 Next steps:
   cd user-api
   npm install
   npm run start:dev

📚 Check the README.md for complete documentation
```

## 🛒 E-commerce Project

### Microservices Architecture

```bash
# Users Service (NestJS with validation feedback)
ccin generate nestjs ecommerce-users --domain user --gcp-project ecommerce-prod

# Products Service (Go Gin + gRPC)
ccin generate go-gin ecommerce-products --domain product --gcp-project ecommerce-prod --grpc

# Orders Service (Go Fiber)
ccin generate go-fiber ecommerce-orders --domain order --gcp-project ecommerce-prod

# Inventory Service (Go Gin + gRPC)
ccin generate go-gin ecommerce-inventory --domain inventory --gcp-project ecommerce-prod --grpc

# Notifications Service (NestJS)
ccin generate nestjs ecommerce-notifications --domain notification --gcp-project ecommerce-prod
```

## 💼 SaaS Project

### Task Management System

```bash
# Authentication API with enhanced validation
ccin generate nestjs task-auth --domain auth --gcp-project task-saas-prod

# Projects API (Go Gin + gRPC for performance)
ccin generate go-gin task-projects --domain project --gcp-project task-saas-prod --grpc

# Task Management API (Ultra-fast Fiber)
ccin generate go-fiber task-management --domain task --gcp-project task-saas-prod

# Reports API
ccin generate go-gin task-reports --domain report --gcp-project task-saas-prod
```

**Smart Input Validation Example:**
```bash
# Invalid project name - gets helpful feedback
ccin generate nestjs a
# Output: ❌ Invalid project name: project name must be at least 2 characters long
#         💡 Use a descriptive name like 'my-api', 'user-service', etc.
```

## 🌐 IoT Project

### Sensor Platform

```bash
# Device Management API (Fiber for speed)
ccin generate go-fiber iot-devices --domain device --gcp-project iot-platform-prod

# Metrics Collection API (Gin + gRPC for high performance)
ccin generate go-gin iot-metrics --domain metric --gcp-project iot-platform-prod --grpc

# Alerts Service (NestJS for complex logic)
ccin generate nestjs iot-alerts --domain alert --gcp-project iot-platform-prod

# Configuration API
ccin generate go-gin iot-config --domain config --gcp-project iot-platform-prod
```

## 💰 FinTech Project

### Payment System

```bash
# Accounts Management API (NestJS for complex business logic)
ccin generate nestjs fintech-accounts --domain account --gcp-project fintech-prod

# Transactions API (Fiber for high-performance processing)
ccin generate go-fiber fintech-transactions --domain transaction --gcp-project fintech-prod

# Risk Analysis API (Gin + gRPC for real-time processing)
ccin generate go-gin fintech-risk --domain risk --gcp-project fintech-prod --grpc

# Compliance Reports API
ccin generate go-gin fintech-compliance --domain compliance --gcp-project fintech-prod
```

## 🏥 HealthTech Project

### Medical Management System

```bash
# Patient Management API (NestJS for data validation)
ccin generate nestjs health-patients --domain patient --gcp-project health-tech-prod

# Appointments API (Go Gin for scheduling logic)
ccin generate go-gin health-appointments --domain appointment --gcp-project health-tech-prod

# Medical Records API (Fiber for high security performance)
ccin generate go-fiber health-records --domain record --gcp-project health-tech-prod

# Telemedicine API (Gin + gRPC for video streaming)
ccin generate go-gin health-telemedicine --domain session --gcp-project health-tech-prod --grpc
```

## 🛠️ Development Commands

### After generating a project

Once you've generated your project with the enhanced CLI, use these commands:

```bash
cd [project-name]

# Setup environment
make setup

# Local development with live reload
make dev

# Run tests
make test

# Build for production
make build

# Build Docker image
make docker-build

# Deploy to GCP
make deploy
```

### For Go projects with gRPC

```bash
# Install protobuf tools (first time only)
make proto-install

# Generate code from .proto files
make proto-gen

# Build project
make build
```

### For NestJS projects

```bash
# Install dependencies
make install

# Generate Swagger documentation
npm run build

# Run in watch mode
make dev
```

## 🏗️ Architecture Patterns

### Microservices with gRPC

Use gRPC for internal communication between services:

```bash
# Main Service (API Gateway) - NestJS
ccin generate nestjs api-gateway --domain gateway --gcp-project my-services

# Internal Services - Go with gRPC
ccin generate go-gin user-service --domain user --gcp-project my-services --grpc
ccin generate go-gin product-service --domain product --gcp-project my-services --grpc
ccin generate go-gin order-service --domain order --gcp-project my-services --grpc
```

### Event-Driven Architecture

Use Fiber for high-throughput services:

```bash
# Event Store
ccin generate go-fiber event-store --domain event --gcp-project event-system

# Event Processors
ccin generate go-fiber order-processor --domain orderEvent --gcp-project event-system
ccin generate go-fiber payment-processor --domain paymentEvent --gcp-project event-system
```

### CQRS Pattern

Separate commands from queries:

```bash
# Command Side (write operations)
ccin generate go-gin user-commands --domain userCommand --gcp-project cqrs-system --grpc

# Query Side (read operations)
ccin generate go-fiber user-queries --domain userQuery --gcp-project cqrs-system
```

## 🚀 Productivity Tips

### Bulk Generation Scripts

```bash
#!/bin/bash
# generate-ecommerce.sh

DOMAIN="ecommerce-prod"

echo "Generating E-commerce Microservices..."

ccin generate nestjs auth-service --domain auth --gcp-project $DOMAIN
ccin generate go-gin user-service --domain user --gcp-project $DOMAIN --grpc
ccin generate go-gin product-service --domain product --gcp-project $DOMAIN --grpc
ccin generate go-fiber order-service --domain order --gcp-project $DOMAIN
ccin generate go-gin inventory-service --domain inventory --gcp-project $DOMAIN --grpc
ccin generate nestjs notification-service --domain notification --gcp-project $DOMAIN

echo "✅ All services generated successfully!"
```

### Environment-Based Generation

```bash
# Development
ccin generate nestjs my-api --domain item --gcp-project my-project-dev

# Staging  
ccin generate nestjs my-api-staging --domain item --gcp-project my-project-staging

# Production
ccin generate nestjs my-api-prod --domain item --gcp-project my-project-prod
```

## 📋 Framework Use Cases

### When to use NestJS
- ✅ Complex APIs with many integrations
- ✅ Teams familiar with TypeScript/Angular
- ✅ Applications requiring decorators and DI
- ✅ Rapid prototyping with TypeScript
- ✅ Rich ecosystem and enterprise features

### When to use Go Gin  
- ✅ High-performance APIs
- ✅ Microservices with gRPC
- ✅ Robust backend services
- ✅ Balance between performance and ease of use
- ✅ Lightweight REST APIs

### When to use Go Fiber
- ✅ Ultra-high performance APIs
- ✅ Services with many concurrent connections
- ✅ Real-time event processing
- ✅ Low-latency requirements
- ✅ Express.js-like API in Go

## 🎯 Pro Tips

### Enhanced CLI Usage
```bash
# Use tab completion for better experience
ccin generate [TAB][TAB]

# Check version and features regularly
ccin --version

# Use descriptive project names for better organization
ccin generate nestjs user-management-service --domain user

# Take advantage of the colorized output for better readability
ccin generate go-gin --help | less -R
```

### Project Organization
```bash
# Use consistent naming conventions
ccin generate nestjs company-users-api --domain user
ccin generate go-gin company-products-api --domain product
ccin generate go-fiber company-orders-api --domain order
```

## 🐳 Docker Examples

### Using CCIN CLI with Docker

```bash
# Generate projects using Docker (no local Go installation needed)
docker run --rm -it -v $(pwd)/output:/output ccin-cli:latest \
  generate nestjs my-dockerized-api --domain user --gcp-project my-project

# E-commerce microservices with Docker
docker run --rm -it -v $(pwd)/ecommerce:/output ccin-cli:latest \
  generate nestjs ecommerce-users --domain user --gcp-project ecommerce-prod

docker run --rm -it -v $(pwd)/ecommerce:/output ccin-cli:latest \
  generate go-gin ecommerce-products --domain product --gcp-project ecommerce-prod --grpc

docker run --rm -it -v $(pwd)/ecommerce:/output ccin-cli:latest \
  generate go-fiber ecommerce-orders --domain order --gcp-project ecommerce-prod
```

### Docker Compose Development

```bash
# Start development environment
docker-compose up -d ccin-dev

# Generate projects in development container
docker-compose exec ccin-dev ccin generate nestjs my-api --domain user

# Access the development container
docker-compose exec ccin-dev sh

# Clean up
docker-compose down -v
```

### CI/CD Pipeline Examples

```bash
# GitHub Actions example
docker run --rm -v $GITHUB_WORKSPACE:/workspace -w /workspace \
  ccin-cli:latest generate nestjs ci-api --domain service

# GitLab CI example  
docker run --rm -v $CI_PROJECT_DIR:/workspace -w /workspace \
  ccin-cli:latest generate go-gin pipeline-api --domain pipeline --grpc
```
