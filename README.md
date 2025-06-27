# ChrisLoarryn CLI

An advanced and powerful CLI for generating CRUD applications with modular architecture, template-based approach and support for multiple modern frameworks.

## Features

- 🚀 **Modular Architecture**: Interchangeable generator system with automatic registration
- 📝 **Template Engine**: Intelligent template processing with dynamic variables
- 🎯 **Multiple Frameworks**: NestJS (Node.js 24.2.0), Go with Gin, Go with Fiber
- 📊 **GCP Integration**: Automatic metrics and logging for Google Cloud Platform
- 🐳 **Docker Ready**: Multi-stage Dockerfiles optimized for production
- 📚 **API Documentation**: Automatic Swagger/OpenAPI generation
- 🔄 **gRPC Support**: Optional gRPC communication support for Go projects
- 🏗️ **Clean Architecture**: DDD patterns and best practices implemented
- 📦 **Dynamic Templates**: Paths and file names with variable substitution
- ⚡ **Automatic Makefile**: Build, test and deploy commands for each framework
- 🔧 **Registry Pattern**: Extensible system for adding new generators

## Installation

### From source code

```bash
git clone https://github.com/chrisloarryn/chrisloarryn-cli
cd chrisloarryn-cli
go build -o chrisloarryn-cli .
```

### Run directly

```bash
go run main.go [command]
```

## Usage

### Available commands

```bash
# Show general help
./chrisloarryn-cli --help

# Show generation help
./chrisloarryn-cli generate --help

# Generate NestJS project with Node.js 24.2.0
./chrisloarryn-cli generate nestjs my-nestjs-api --domain user --gcp-project my-project

# Generate Go project with Gin (with optional gRPC)
./chrisloarryn-cli generate go-gin my-gin-api --domain product --gcp-project my-project --grpc

# Generate Go project with Fiber (with optional gRPC)
./chrisloarryn-cli generate go-fiber my-fiber-api --domain order --gcp-project my-project --grpc

# Example without GCP (basic functionality only)
./chrisloarryn-cli generate nestjs simple-api --domain item
```

### Global Parameters

- `project-name`: Name of the project to generate (required)
- `--domain, -d`: Domain/entity name (e.g., user, product, order). Default: "item"
- `--gcp-project, -p`: GCP project ID for metrics (optional)

### Specific Parameters

**For Go projects (Gin/Fiber):**
- `--grpc, -g`: Include gRPC support in addition to REST API

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

### Go with Gin (Go 1.24.4)

Generates a complete project with:
- ✅ **Runtime**: Go 1.24.4 with Gin framework
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

### Go with Fiber (Go 1.24.4)

Generates a complete project with:
- ✅ **Runtime**: Go 1.24.4 with Fiber framework (ultra-fast)
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
./chrisloarryn-cli generate nestjs ecommerce-users --domain user --gcp-project ecommerce-prod

# Products API
./chrisloarryn-cli generate go-gin ecommerce-products --domain product --gcp-project ecommerce-prod --grpc

# Orders API
./chrisloarryn-cli generate go-fiber ecommerce-orders --domain order --gcp-project ecommerce-prod
```

### Microservices Project
```bash
# Authentication service
./chrisloarryn-cli generate nestjs auth-service --domain auth --gcp-project microservices-dev

# Inventory service
./chrisloarryn-cli generate go-gin inventory-service --domain item --gcp-project microservices-dev --grpc

# Notifications service
./chrisloarryn-cli generate go-fiber notification-service --domain notification --gcp-project microservices-dev
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
chrisloarryn-cli/
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

import "github.com/chrisloarryn/chrisloarryn-cli/internal/common"

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
import _ "github.com/chrisloarryn/chrisloarryn-cli/internal/generators/my-framework"
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
chrisloarryn-cli/
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
