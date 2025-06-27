# ChrisLoarryn CLI

Un CLI avanzado y potente para generar aplicaciones CRUD con diferentes frameworks, incluyendo interceptores, métricas de GCP, Docker y archivos de configuración.

## Características

- 🚀 **Múltiples Frameworks**: NestJS, Go con Gin, Go con Fiber
- 📊 **Integración GCP**: Métricas y logging automático
- 🐳 **Docker Ready**: Dockerfile optimizado incluido
- 📚 **Documentación API**: Swagger/OpenAPI automático
- 🔄 **gRPC Support**: Opcional para proyectos Go
- 🏗️ **Arquitectura Limpia**: Patrones y mejores prácticas
- 📦 **Template Engine**: Sistema de plantillas flexible
- ⚡ **Makefile**: Comandos automatizados para cada proyecto

## Instalación

### Desde el código fuente

```bash
git clone https://github.com/chrisloarryn/chrisloarryn-cli
cd chrisloarryn-cli
go build -o chrisloarryn-cli main.go
```

### Ejecutar directamente

```bash
go run main.go [command]
```

## Uso

### Comandos disponibles

```bash
# Mostrar ayuda general
./chrisloarryn-cli --help

# Mostrar ayuda de generación
./chrisloarryn-cli generate --help

# Generar proyecto NestJS
./chrisloarryn-cli generate nestjs my-api --domain user --gcp-project my-project

# Generar proyecto Go con Gin (con gRPC)
./chrisloarryn-cli generate go-gin my-api --domain product --gcp-project my-project --grpc

# Generar proyecto Go con Fiber
./chrisloarryn-cli generate go-fiber my-api --domain order --gcp-project my-project
```

### Parámetros

- `project-name`: Nombre del proyecto a generar
- `--domain, -d`: Nombre del dominio/entidad (ej: user, product, order)
- `--gcp-project, -p`: ID del proyecto GCP para métricas
- `--grpc, -g`: Incluir soporte gRPC (solo para proyectos Go)

## Proyectos Generados

### NestJS

Genera un proyecto completo con:
- ✅ TypeScript con NestJS framework
- ✅ MongoDB con Mongoose
- ✅ Swagger/OpenAPI documentation
- ✅ Interceptores para métricas GCP
- ✅ Validación con class-validator
- ✅ Docker multistage
- ✅ Jest testing setup
- ✅ ESLint y Prettier

**Estructura:**
```
src/
├── common/interceptors/gcp-metrics.interceptor.ts
├── [domain]/
│   ├── dto/
│   ├── entities/
│   ├── [domain].controller.ts
│   ├── [domain].service.ts
│   └── [domain].module.ts
├── app.module.ts
└── main.ts
```

### Go con Gin

Genera un proyecto completo con:
- ✅ Go con Gin framework
- ✅ PostgreSQL con GORM
- ✅ Swagger documentation
- ✅ Middleware para métricas GCP
- ✅ gRPC support (opcional)
- ✅ Clean architecture
- ✅ Docker multistage
- ✅ Testing setup

**Estructura:**
```
cmd/server/main.go
internal/
├── config/config.go
├── handlers/[domain]_handler.go
├── middleware/metrics.go
├── models/[domain].go
├── repository/[domain]_repository.go
├── router/router.go
└── service/[domain]_service.go
api/proto/[domain].proto (si gRPC)
```

### Go con Fiber

Genera un proyecto completo con:
- ✅ Go con Fiber framework (ultra-rápido)
- ✅ PostgreSQL con GORM
- ✅ Swagger documentation
- ✅ Middleware para métricas GCP
- ✅ gRPC support (opcional)
- ✅ Clean architecture
- ✅ Docker multistage
- ✅ Testing setup

**Estructura:** Similar a Gin pero optimizada para Fiber

## Funcionalidades Incluidas

### 🔄 CRUD Completo
Cada proyecto genera:
- CREATE: Crear nuevos registros
- READ: Listar y obtener por ID
- UPDATE: Actualizar registros existentes
- DELETE: Eliminación lógica

### 📈 Métricas GCP
- Custom metrics para duración de requests
- Structured logging
- Error tracking automático
- Performance monitoring

### 🐳 Docker
- Dockerfile multistage optimizado
- Health checks
- Non-root user
- Minimal image size

### 📦 Makefile
Cada proyecto incluye comandos para:
- `make setup`: Configuración inicial
- `make dev`: Modo desarrollo
- `make build`: Construir aplicación
- `make test`: Ejecutar tests
- `make docker-build`: Construir imagen Docker
- `make deploy`: Desplegar a GCP

### 🌐 API Documentation
- Swagger UI automático
- OpenAPI specification
- Endpoints documentados
- Ejemplos de request/response

## Ejemplos de Uso

### Proyecto E-commerce
```bash
# API de usuarios
./chrisloarryn-cli generate nestjs ecommerce-users --domain user --gcp-project ecommerce-prod

# API de productos
./chrisloarryn-cli generate go-gin ecommerce-products --domain product --gcp-project ecommerce-prod --grpc

# API de órdenes
./chrisloarryn-cli generate go-fiber ecommerce-orders --domain order --gcp-project ecommerce-prod
```

### Proyecto Microservicios
```bash
# Servicio de autenticación
./chrisloarryn-cli generate nestjs auth-service --domain auth --gcp-project microservices-dev

# Servicio de inventario
./chrisloarryn-cli generate go-gin inventory-service --domain item --gcp-project microservices-dev --grpc

# Servicio de notificaciones
./chrisloarryn-cli generate go-fiber notification-service --domain notification --gcp-project microservices-dev
```

## Desarrollo

### Estructura del CLI

```
cmd/
├── root.go           # Comando raíz
├── generate.go       # Comando generate principal
├── nestjs.go         # Generador NestJS
├── go_gin.go         # Generador Go Gin
├── go_gin_helpers.go # Helpers para Gin
└── go_fiber.go       # Generador Go Fiber
main.go               # Entry point
```

### Agregar un nuevo framework

1. Crear archivo `cmd/new_framework.go`
2. Implementar función `generateNewFrameworkProject`
3. Registrar comando en `cmd/generate.go`
4. Agregar templates y helpers necesarios

## Contribuir

1. Fork el repositorio
2. Crea una branch para tu feature
3. Commit tus cambios
4. Push a la branch
5. Crea un Pull Request

## Roadmap

- [ ] Soporte para Python (FastAPI/Django)
- [ ] Soporte para Rust (Actix/Axum)  
- [ ] Templates personalizables
- [ ] CLI interactivo
- [ ] Integración con CI/CD
- [ ] Database migrations
- [ ] Authentication middleware
- [ ] Rate limiting
- [ ] Caching layer

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

## Licencia

MIT License - ver archivo [LICENSE](LICENSE) para más detalles.

## Autor

**Chris Loarryn** - [@chrisloarryn](https://github.com/chrisloarryn)

---

¿Preguntas o sugerencias? ¡Abre un issue! 🚀
