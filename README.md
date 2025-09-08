# CCIN CLI - ChrisLoarryn's Comprehensive Code Integration & Initialization Tool

🎯 **Advanced CLI for generating modern, production-ready CRUD applications** with colorized output, intelligent validation, and comprehensive help system.

## Quick Start

### From source code

```bash
git clone https://github.com/chrisloarryn/homebrew-ccin
cd homebrew-ccin
go build -o ccin .
```

### Run directly

```bash
go run main.go [command]
```

### Test the enhanced CLI

```bash
# Beautiful help with colors and emojis
./ccin --help

# Detailed version information
./ccin --version

# Framework-specific help
./ccin generate --help
./ccin generate nestjs --help
```

## ✨ Enhanced Features

- 🎨 **Colorized Output**: Beautiful, colored CLI interface with emojis and clear visual hierarchy
- 🛡️ **Input Validation**: Smart validation with helpful error messages and suggestions
- 📖 **Enhanced Help**: Comprehensive, descriptive help messages in English with examples
- 🔧 **Error Handling**: Professional error messages with quick help suggestions
- 🎯 **User Experience**: Intuitive commands with consistent formatting and clear guidance

## Core Features

- 🚀 **Modular Architecture**: Interchangeable generator system with automatic registration
- 📝 **Template Engine**: Intelligent template processing with dynamic variables
- 🎯 **Multiple Frameworks**: NestJS (Node.js 24.2.0), Go 1.25.1 (Gin), Go 1.25.1 (Fiber), Swift 6.1.2 (Vapor 4)
- 📊 **GCP Integration**: Automatic metrics and logging for Google Cloud Platform
- 🐳 **Docker Ready**: Multi-stage Dockerfiles optimized for production
- 📚 **API Documentation**: Automatic Swagger/OpenAPI generation
- 🔄 **gRPC Support**: Optional gRPC communication support for Go projects
- 🏗️ **Clean Architecture**: DDD patterns and best practices implemented
- 📦 **Dynamic Templates**: Paths and file names with variable substitution
- ⚡ **Automatic Makefile**: Build, test and deploy commands for each framework
- 🔧 **Registry Pattern**: Extensible system for adding new generators

## Available Templates

- nestjs — NestJS (Node.js 24.2.0 + TypeScript + MongoDB)
  Example:
  ```bash
  ccin generate nestjs my-api --domain user --gcp-project my-project
  ```
- go-gin — Go 1.25.1 + Gin (REST/gRPC + PostgreSQL + Clean Architecture)
  Example:
  ```bash
  ccin generate go-gin orders-api --domain order --grpc
  ```
- go-fiber — Go 1.25.1 + Fiber (Ultra-fast REST/gRPC + PostgreSQL)
  Example:
  ```bash
  ccin generate go-fiber products-api --domain product --gcp-project prod
  ```
- swift-vapor — Swift 6.1.2 + Vapor 4 (REST + optional gRPC)
  Example:
  ```bash
  ccin generate swift-vapor catalog-api --domain product --grpc
  ```

## Installation

### From Homebrew (Recommended)

```bash
# Coming soon - will be available once published
brew tap chrisloarryn/homebrew-ccin
brew install ccin
```

### From Docker

```bash
# Pull and run directly (when published)
docker run --rm -it -v $(pwd)/output:/output ghcr.io/chrisloarryn/ccin:latest

# Or build locally
git clone https://github.com/chrisloarryn/homebrew-ccin
cd homebrew-ccin
make docker-build-all  # Builds binary + Docker image
make docker-run
```

### From source code

```bash
git clone https://github.com/chrisloarryn/homebrew-ccin
cd homebrew-ccin
go build -o ccin .
# Optionally, move to PATH
sudo mv ccin /usr/local/bin/
```

### Verify installation

```bash
ccin --version
ccin --help
```

## Usage

### Enhanced CLI Experience

The CLI now features beautiful, colorized output with comprehensive help messages:

```bash
# 🎯 Main help - shows overview with colors and emojis
ccin --help

# 📖 Version information with branding
ccin --version

# 🚀 Generate command help - shows available frameworks
ccin generate --help

# 📦 Framework-specific help with detailed descriptions
ccin generate nestjs --help
ccin generate go-gin --help  
ccin generate go-fiber --help
ccin generate swift-vapor --help
```

### Smart Input Validation

The CLI includes intelligent validation with helpful error messages:

```bash
# ❌ Too short project name - gets helpful feedback
ccin generate nestjs a
# Output: ❌ Invalid project name: project name must be at least 2 characters long
#         💡 Use a descriptive name like 'my-api', 'user-service', etc.

# ❌ Missing project name - gets usage guidance
ccin generate nestjs
# Output: Error: accepts 1 arg(s), received 0
#         💡 Quick help: ccin --help or ccin generate --help
```

### Project Generation Commands

```bash
# Generate NestJS project with enhanced output
ccin generate nestjs my-nestjs-api --domain user --gcp-project my-project

# Generate Go project with Gin (with optional gRPC)
ccin generate go-gin my-gin-api --domain product --gcp-project my-project --grpc

# Generate Go project with Fiber (with optional gRPC)  
ccin generate go-fiber my-fiber-api --domain order --gcp-project my-project --grpc

# Example without GCP (basic functionality only)
ccin generate nestjs simple-api --domain item
```

### Command Parameters

#### Global Parameters
- `project-name`: Name of the project to generate (required, minimum 2 characters)
- `--domain, -d`: Domain/entity name (e.g., user, product, order). Default: "item"
- `--gcp-project, -p`: GCP project ID for metrics integration (optional)

#### Framework-Specific Parameters
**For Go projects (Gin/Fiber):**
- `--grpc, -g`: Include gRPC support in addition to REST API

#### Input Validation
- ✅ Project names must be at least 2 characters long
- ✅ Helpful error messages with naming suggestions
- ✅ Clear guidance for required parameters
- ✅ Smart defaults for optional parameters

### Enhanced CLI Output Examples

#### 🎯 Main Help Output
```
🎯 CCIN CLI - ChrisLoarryn's Comprehensive Code Integration & Initialization Tool

✨ Generate production-ready CRUD applications with multiple frameworks:
   • NestJS (Node.js 24.2.0 + TypeScript + MongoDB)
   • Go 1.25.1 + Gin (REST/gRPC + PostgreSQL + Clean Architecture)  
   • Go 1.25.1 + Fiber (Ultra-fast REST/gRPC + PostgreSQL)

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

#### 📦 Framework-Specific Help
```bash
ccin generate nestjs --help
```
```
📦 NESTJS GENERATOR

🎯 What you'll get:
   • NestJS framework with TypeScript
   • MongoDB with Mongoose ODM
   • Swagger/OpenAPI automatic documentation
   • GCP Metrics interceptors (optional)
   • Docker multi-stage production build
   • Jest testing configuration
   • ESLint + Prettier code quality

📋 Example: ccin generate nestjs my-api --domain user --gcp-project my-project
```

#### 🚀 Project Generation Output
```bash
ccin generate nestjs my-api --domain user --gcp-project my-gcp
```
```
🚀 Generating NestJS CRUD project: my-api
📊 Domain: user
☁️  GCP Project: my-gcp
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📝 Processing templates...
✅ NestJS project 'my-api' generated successfully!

🎯 Next steps:
   cd my-api
   npm install
   npm run start:dev

📚 Check the README.md for complete documentation
```

## Generated Projects

### NestJS (Node.js 24.2.0)

Generates a complete project with:
- ✅ **Runtime**: Node.js 24.2.0 Alpine 3.22
- ✅ **Framework**: NestJS with TypeScript
- ✅ **Database**: MongoDB with Mongoose ODM
- ✅ **Documentation**: Automatic Swagger/OpenAPI
- ✅ **Metrics**: Interceptors for Google Cloud Platform
- ✅ **Validation**: class-validator and class-transformer
- ✅ **Container**: Optimized multi-stage Docker
- ✅ **Testing**: Jest with complete configuration
- ✅ **Code Quality**: ESLint, Prettier and development configuration
- ✅ **Build Tools**: Makefile with automated commands

**Project Structure:**
```
my-nestjs-api/
├── Dockerfile                    # Multi-stage con Node.js 24.2.0
├── Makefile                      # Comandos de build, test, deploy
├── package.json                  # Dependencias y scripts
└── src/
    ├── app.module.ts            # Módulo principal
    ├── main.ts                  # Entry point
    ├── common/
    │   └── interceptors/
    │       └── gcp-metrics.interceptor.ts        └── [domain]/               # Example: user, product, order
            ├── dto/
            │   ├── create-[domain].dto.ts
            │   └── update-[domain].dto.ts
            ├── entities/
            │   └── [domain].entity.ts
            ├── [domain].controller.ts
            ├── [domain].service.ts
            └── [domain].module.ts
```

### Go with Gin (Go 1.25.1)

Generates a complete project with:
- ✅ **Runtime**: Go 1.25.1 with a Gin framework
- ✅ **Database**: PostgreSQL with standard SQL
- ✅ **API**: REST endpoints with JSON responses
- ✅ **gRPC**: Optional support for gRPC communication
- ✅ **Metrics**: Middleware for Google Cloud Platform
- ✅ **Architecture**: Clean Architecture with well-defined layers
- ✅ **Container**: Optimized multi-stage Docker
- ✅ **Build Tools**: Makefile with automated commands
- ✅ **Environment**: Configuration via environment variables

**Project Structure:**
```
my-gin-api/
├── Dockerfile                      # Optimized multi-stage build
├── Makefile                        # Build, test, deploy commands
├── go.mod                          # Go dependencies
├── main.go                         # Entry point
├── .env.example                    # Environment variables template
└── internal/
    ├── api/
    │   └── routes.go              # Route configuration
    ├── config/
    │   └── config.go              # Configuration loading
    ├── database/
    │   └── database.go            # DB connection and setup
    ├── handlers/
    │   └── [domain]_handler.go    # HTTP handlers
    ├── models/
    │   └── [domain].go            # Data models
    ├── services/
    │   └── [domain]_service.go    # Business logic
    ├── middleware/
    │   └── metrics.go             # Metrics middleware
    ├── grpc/                      # (Optional with --grpc)
    │   └── server.go              # gRPC server
    └── metrics/                   # (Only with GCP)
        └── metrics.go             # GCP metrics client
```

### Go with Fiber (Go 1.25.1)

Generates a complete project with:
- ✅ **Runtime**: Go 1.25.1 with Fiber framework (ultra-fast)
- ✅ **Performance**: Framework optimized for speed
- ✅ **Database**: PostgreSQL with standard SQL
- ✅ **API**: REST endpoints with fast JSON responses
- ✅ **gRPC**: Optional support for gRPC communication
- ✅ **Metrics**: Middleware for Google Cloud Platform
- ✅ **Architecture**: Clean Architecture with well-defined layers
- ✅ **Container**: Optimized multi-stage Docker
- ✅ **Build Tools**: Makefile with automated commands
- ✅ **CORS**: CORS middleware included

**Project Structure:**
```
my-fiber-api/
├── Dockerfile                      # Optimized multi-stage build
├── Makefile                        # Build, test, deploy commands
├── go.mod                          # Go dependencies
├── main.go                         # Entry point
├── .env.example                    # Environment variables template
└── internal/
    ├── api/
    │   └── routes.go              # Fiber route configuration
    ├── config/
    │   └── config.go              # Configuration loading
    ├── database/
    │   └── database.go            # DB connection and setup
    ├── handlers/
    │   └── [domain]_handler.go    # Fiber handlers
    ├── models/
    │   └── [domain].go            # Data models
    ├── services/
    │   └── [domain]_service.go    # Business logic
    ├── middleware/
    │   └── metrics.go             # Metrics middleware
    ├── grpc/                      # (Optional with --grpc)
    │   └── server.go              # gRPC server
    └── metrics/                   # (Only with GCP)
        └── metrics.go             # GCP metrics client
```

## 🌐 Generated API Endpoints

All projects include standard REST endpoints:

### Common Endpoints
```bash
# Health Check
GET /health

# CRUD Operations (example with domain="product")
GET    /api/v1/products       # List all products
GET    /api/v1/products/:id   # Get product by ID
POST   /api/v1/products       # Create new product
PUT    /api/v1/products/:id   # Update product
DELETE /api/v1/products/:id   # Delete product
```

### Usage Examples

```bash
# Create a product
curl -X POST http://localhost:3000/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name": "iPhone 15", "description": "Latest iPhone model"}'

# Get all products
curl http://localhost:3000/api/v1/products

# Get specific product
curl http://localhost:3000/api/v1/products/1

# Update product
curl -X PUT http://localhost:3000/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "iPhone 15 Pro", "description": "Updated model"}'

# Delete product
curl -X DELETE http://localhost:3000/api/v1/products/1
```

## Included Features

### 🔄 Complete CRUD
Each project generates:
- CREATE: Create new records
- READ: List and get by ID
- UPDATE: Update existing records
- DELETE: Logical deletion

### 📈 GCP Metrics
- Custom metrics for request duration
- Structured logging
- Automatic error tracking
- Performance monitoring

### 🐳 Docker
- Optimized multistage Dockerfile
- Health checks
- Non-root user
- Minimal image size

### 📦 Makefile
Each project includes commands for:
- `make setup`: Initial setup
- `make dev`: Development mode
- `make build`: Build application
- `make test`: Run tests
- `make docker-build`: Build Docker image
- `make deploy`: Deploy to GCP

### 🌐 API Documentation
- Automatic Swagger UI
- OpenAPI specification
- Documented endpoints
- Request/response examples

## Usage Examples

### E-commerce Project
```bash
# Users API
./ccin generate nestjs ecommerce-users --domain user --gcp-project ecommerce-prod

# Products API
./ccin generate go-gin ecommerce-products --domain product --gcp-project ecommerce-prod --grpc

# Orders API
./ccin generate go-fiber ecommerce-orders --domain order --gcp-project ecommerce-prod
```

### Microservices Project
```bash
# Authentication service
./ccin generate nestjs auth-service --domain auth --gcp-project microservices-dev

# Inventory service
./ccin generate go-gin inventory-service --domain item --gcp-project microservices-dev --grpc

# Notifications service
./ccin generate go-fiber notification-service --domain notification --gcp-project microservices-dev
```

## 🚀 Production Best Practices

### Environment Variables
Each project includes a `.env.example` file with necessary variables:

**NestJS:**
```bash
PORT=3000
DATABASE_URL=mongodb://localhost:27017/my-app
NODE_ENV=production
GCP_PROJECT=my-gcp-project
```

**Go (Gin/Fiber):**
```bash
PORT=8080
GRPC_PORT=50051
DATABASE_URL=postgres://user:pass@localhost/mydb?sslmode=disable
GIN_MODE=release  # Para Gin
GCP_PROJECT=my-gcp-project
```

### Production Commands

**Development:**
```bash
# NestJS
npm run start:dev

# Go
make dev
# or
go run main.go
```

**Production with Docker:**
```bash
# Build image
make docker-build

# Run in production
docker run -p 3000:3000 --env-file .env my-app:latest
```

### Monitoring and Metrics

When you specify `--gcp-project`, the project includes:
- ✅ Latency metrics per endpoint
- ✅ Request counters by status code
- ✅ Automatic error metrics
- ✅ Google Cloud Monitoring integration
- ✅ Health checks for Kubernetes

### Implemented Security

- ✅ **Containers**: Non-root user in Docker
- ✅ **Validation**: Input validation on all endpoints
- ✅ **CORS**: Configured by default
- ✅ **Headers**: Security headers included
- ✅ **Environment**: Sensitive variables externalized

## Development

### New Modular Architecture

The CLI now uses a completely modular and extensible architecture:

```
ccin/
├── cmd/
│   ├── root.go                   # Root command
│   └── generate.go               # Generate command with subcommands
├── internal/
│   ├── common/                   # Shared utilities
│   │   ├── generator.go          # Generator interface
│   │   ├── registry.go           # Global generator registry
│   │   └── template.go           # Template processing engine
│   └── generators/               # Framework-specific generators
│       ├── nestjs/               # NestJS generator
│       ├── go-gin/               # Go Gin generator
│       └── go-fiber/             # Go Fiber generator
├── templates/                    # Templates (.tpl) organized
│   ├── common/                   # Shared Dockerfiles and Makefiles
│   ├── nestjs/                   # NestJS-specific templates
│   ├── go-gin/                   # Go Gin-specific templates
│   └── go-fiber/                 # Go Fiber-specific templates
└── main.go                       # Entry point
```

### How to Add a New Framework

**1. Create the Generator:**
```go
// internal/generators/my-framework/generator.go
package myframework

import "github.com/chrisloarryn/ccin/internal/common"

type Generator struct {
    *common.BaseGenerator
}

func NewGenerator() *Generator {
    return &Generator{
        BaseGenerator: common.NewBaseGenerator(
            "my-framework",
            "Generate My Framework CRUD application",
        ),
    }
}

func (g *Generator) Generate(config *common.GeneratorConfig) error {
    // Generation logic
    data := common.PrepareTemplateData(config)
    processor := common.NewTemplateProcessor(config.TemplateDir, config.OutputDir)
    return processor.ProcessDirectory(data)
}

func init() {
    common.Registry.Register(NewGenerator())
}
```

**2. Create Templates:**
```bash
mkdir -p templates/my-framework
# Create .tpl files with variables like {{.ProjectName}}, {{.DomainLower}}, etc.
```

**3. Add Command:**
```go
// In cmd/generate.go
var myFrameworkCmd = &cobra.Command{
    Use:   "my-framework [project-name]",
    Short: "Generate My Framework CRUD application",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        generator, _ := common.Registry.Get("my-framework")
        config := &common.GeneratorConfig{
            ProjectName: args[0],
            // ... other configurations
        }
        generator.Generate(config)
    },
}
```

**4. Import the Generator:**
```go
// In cmd/generate.go
import _ "github.com/chrisloarryn/ccin/internal/generators/my-framework"
```

## Contributing

1. Fork the repository
2. Create a branch for your feature
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Roadmap

### ✅ Completed (v1.0)
- ✅ **Modular Architecture**: Generator system with registry pattern
- ✅ **Dynamic Templates**: Template engine with variables and dynamic paths  
- ✅ **Multi-Framework**: NestJS (Node.js 24.2.0), Go Gin, Go Fiber
- ✅ **Optimized Docker**: Multi-stage builds for production
- ✅ **GCP Integration**: Automatic metrics in Google Cloud Platform
- ✅ **gRPC Support**: Optional gRPC communication for Go
- ✅ **Clean Architecture**: Layer separation and responsibilities
- ✅ **Build Tools**: Automated Makefiles per framework

### 🔄 In Development (v1.1)
- [ ] **Python Support**: FastAPI with async/await
- 🔄 **Rust Support**: Rocket Web framework (In Progress)
- [ ] **Interactive CLI**: Wizard for project configuration
- [ ] **Custom Templates**: User template customization

### 🚀 Next Versions (v2.0+)
- [ ] **Database Migrations**: Automatic migration system
- [ ] **Authentication Middleware**: JWT, OAuth2, RBAC
- [ ] **Rate Limiting**: Request limits middleware
- [ ] **Caching Layer**: Redis/Memcached integration
- [ ] **CI/CD Templates**: GitHub Actions, GitLab CI
- [ ] **Kubernetes**: Manifests and Helm charts
- [ ] **Observability**: Prometheus, Grafana, Jaeger
- [ ] **API Gateway**: Kong, Ambassador integration

## Architecture

ChrisLoarryn CLI uses a modular, template-based architecture that makes it easy to add new frameworks and maintain existing generators.

### Project Structure

```
ccin/
├── cmd/                          # CLI commands
│   ├── root.go                   # Root command
│   └── generate.go               # Generate command with subcommands
├── internal/
│   ├── common/                   # Shared utilities
│   │   ├── generator.go          # Generator interface and base implementation
│   │   ├── registry.go           # Global generator registry
│   │   └── template.go           # Template processing engine
│   └── generators/               # Framework-specific generators
│       ├── nestjs/               # NestJS generator
│       ├── go-gin/               # Go Gin generator
│       └── go-fiber/             # Go Fiber generator
├── templates/                    # Template files (.tpl)
│   ├── common/                   # Shared templates (Dockerfiles, Makefiles)
│   ├── nestjs/                   # NestJS-specific templates
│   ├── go-gin/                   # Go Gin-specific templates
│   └── go-fiber/                 # Go Fiber-specific templates
└── main.go                       # Application entry point
```

### How it Works

1. **Registry Pattern**: All generators register themselves in a global registry during package initialization
2. **Template Engine**: Uses Go's `text/template` to process `.tpl` files with variable substitution
3. **Modular Generators**: Each framework has its own generator implementing the `Generator` interface
4. **Dynamic Paths**: Template file names and directory structures can include variables (e.g., `{{.DomainLower}}`)

### Adding a New Generator

To add support for a new framework:

1. Create a new directory: `internal/generators/my-framework/`
2. Implement the `Generator` interface:
   ```go
   type Generator interface {
       GetName() string
       GetDescription() string
       Generate(config *GeneratorConfig) error
   }
   ```
3. Register your generator in the `init()` function
4. Create template files in `templates/my-framework/`
5. Add command flags and handling in `cmd/generate.go`

### Template Variables

All templates have access to these variables:

- `{{.ProjectName}}` - The name of the project
- `{{.DomainName}}` - The domain entity name (e.g., "user", "product")
- `{{.DomainTitle}}` - Title case domain (e.g., "User", "Product")
- `{{.DomainUpper}}` - Uppercase domain (e.g., "USER", "PRODUCT")
- `{{.DomainLower}}` - Lowercase domain (e.g., "user", "product")
- `{{.GCPProject}}` - GCP project ID for metrics
- `{{.WithGRPC}}` - Boolean indicating gRPC support
- `{{.Port}}` - Application port
- `{{.DatabaseType}}` - Database type (e.g., "postgresql", "mongodb")

## License

MIT License - see [LICENSE](LICENSE) file for more details.

## Author

**Chris Loarryn** - [@chrisloarryn](https://github.com/chrisloarryn)

---

Questions or suggestions? Open an issue! 🚀



### Swift Vapor (Swift 5/6, Vapor 4)

Generates a complete project with:
- ✅ Framework: Swift Vapor 4
- ✅ Runtime: Swift 5.10+ (macOS 13+/Ubuntu 22.04)
- ✅ API: REST endpoints with clean layers (Controllers/Services/Models)
- ✅ gRPC: Optional scaffolding (Proto + placeholders)
- ✅ Health: Built-in health check endpoint
- ✅ Container: Optimized multi-stage Dockerfile
- ✅ Build Tools: Makefile with handy commands
- ✅ Observability: Simple MetricsMiddleware for request timing logs

Example generation command:
```bash
ccin generate swift-vapor catalog-api --domain product --grpc
```

Project Structure:
```
catalog-api/
├── Package.swift                     # SwiftPM manifest
├── Sources/
│   ├── App/
│   │   ├── Controllers/             # HTTP controllers
│   │   ├── Middleware/              # Middlewares (Metrics, etc.)
│   │   ├── Models/                  # Domain models (Codable)
│   │   ├── Services/                # Business logic services
│   │   ├── GRPC/                    # (Optional with --grpc) gRPC placeholders
│   │   ├── configure.swift          # App configuration & bootstrap
│   │   └── routes.swift             # Routes registration
│   └── Run/
│       └── main.swift               # Entry point
├── Proto/                            # (Optional with --grpc) Proto definitions
│   └── product.proto
├── Dockerfile                        # Multi-stage Docker build
├── Makefile                          # Build & run helpers
└── README.md                         # Project documentation
```

Notes:
- The template uses in-memory storage for simplicity. Swap for a persistence layer (e.g., Fluent + PostgreSQL) as needed.
- gRPC requires generating Swift stubs from the included proto; see the project README for commands and setup.
