# Ejemplos de Uso

Esta guía muestra diferentes casos de uso del CLI ChrisLoarryn para generar aplicaciones CRUD.

## Proyectos Básicos

### API Simple de Usuarios

```bash
# NestJS con MongoDB
./ccin generate nestjs user-api --domain user --gcp-project my-project

# Go Gin con PostgreSQL
./ccin generate go-gin user-api --domain user --gcp-project my-project

# Go Fiber con PostgreSQL
./ccin generate go-fiber user-api --domain user --gcp-project my-project
```

## Proyecto E-commerce

### Arquitectura de Microservicios

```bash
# Servicio de Usuarios (NestJS)
./ccin generate nestjs ecommerce-users --domain user --gcp-project ecommerce-prod

# Servicio de Productos (Go Gin + gRPC)
./ccin generate go-gin ecommerce-products --domain product --gcp-project ecommerce-prod --grpc

# Servicio de Órdenes (Go Fiber)
./ccin generate go-fiber ecommerce-orders --domain order --gcp-project ecommerce-prod

# Servicio de Inventario (Go Gin + gRPC)
./ccin generate go-gin ecommerce-inventory --domain inventory --gcp-project ecommerce-prod --grpc

# Servicio de Notificaciones (NestJS)
./ccin generate nestjs ecommerce-notifications --domain notification --gcp-project ecommerce-prod
```

## Proyecto SaaS

### Sistema de Gestión de Tareas

```bash
# API de Autenticación
./ccin generate nestjs task-auth --domain auth --gcp-project task-saas-prod

# API de Proyectos
./ccin generate go-gin task-projects --domain project --gcp-project task-saas-prod --grpc

# API de Tareas
./ccin generate go-fiber task-management --domain task --gcp-project task-saas-prod

# API de Reportes
./ccin generate go-gin task-reports --domain report --gcp-project task-saas-prod
```

## Proyecto IoT

### Plataforma de Sensores

```bash
# API de Dispositivos
./ccin generate go-fiber iot-devices --domain device --gcp-project iot-platform-prod

# API de Métricas (con gRPC para alta performance)
./ccin generate go-gin iot-metrics --domain metric --gcp-project iot-platform-prod --grpc

# API de Alertas
./ccin generate nestjs iot-alerts --domain alert --gcp-project iot-platform-prod

# API de Configuración
./ccin generate go-gin iot-config --domain config --gcp-project iot-platform-prod
```

## Proyecto FinTech

### Sistema de Pagos

```bash
# API de Usuarios/Cuentas
./ccin generate nestjs fintech-accounts --domain account --gcp-project fintech-prod

# API de Transacciones (alta performance con Fiber)
./ccin generate go-fiber fintech-transactions --domain transaction --gcp-project fintech-prod

# API de Análisis de Riesgo (con gRPC)
./ccin generate go-gin fintech-risk --domain risk --gcp-project fintech-prod --grpc

# API de Reportes Regulatorios
./ccin generate go-gin fintech-compliance --domain compliance --gcp-project fintech-prod
```

## Proyecto HealthTech

### Sistema de Gestión Médica

```bash
# API de Pacientes
./ccin generate nestjs health-patients --domain patient --gcp-project health-tech-prod

# API de Citas Médicas
./ccin generate go-gin health-appointments --domain appointment --gcp-project health-tech-prod

# API de Expedientes (alta seguridad)
./ccin generate go-fiber health-records --domain record --gcp-project health-tech-prod

# API de Telemedicina (con gRPC para video)
./ccin generate go-gin health-telemedicine --domain session --gcp-project health-tech-prod --grpc
```

## Comandos de Desarrollo

### Después de generar un proyecto

```bash
cd [project-name]

# Configurar entorno
make setup

# Desarrollo local
make dev

# Ejecutar tests
make test

# Construir para producción
make build

# Construir imagen Docker
make docker-build

# Desplegar a GCP
make deploy
```

### Para proyectos Go con gRPC

```bash
# Instalar herramientas protobuf (primera vez)
make proto-install

# Generar código desde .proto
make proto-gen

# Construir proyecto
make build
```

### Para proyectos NestJS

```bash
# Instalar dependencias
make install

# Generar documentación Swagger
npm run build

# Ejecutar en modo watch
make dev
```

## Patrones de Arquitectura

### Microservicios con gRPC

Use gRPC para comunicación interna entre servicios:

```bash
# Servicio Principal (API Gateway) - NestJS
./ccin generate nestjs api-gateway --domain gateway --gcp-project my-services

# Servicios Internos - Go con gRPC
./ccin generate go-gin user-service --domain user --gcp-project my-services --grpc
./ccin generate go-gin product-service --domain product --gcp-project my-services --grpc
./ccin generate go-gin order-service --domain order --gcp-project my-services --grpc
```

### Event-Driven Architecture

Use Fiber para servicios de alta throughput:

```bash
# Event Store
./ccin generate go-fiber event-store --domain event --gcp-project event-system

# Event Processors
./ccin generate go-fiber order-processor --domain orderEvent --gcp-project event-system
./ccin generate go-fiber payment-processor --domain paymentEvent --gcp-project event-system
```

### CQRS Pattern

Separe comandos de consultas:

```bash
# Command Side (escritura)
./ccin generate go-gin user-commands --domain userCommand --gcp-project cqrs-system --grpc

# Query Side (lectura)
./ccin generate go-fiber user-queries --domain userQuery --gcp-project cqrs-system
```

## Tips de Productividad

### Scripting Bulk Generation

```bash
#!/bin/bash
# generate-ecommerce.sh

DOMAIN="ecommerce-prod"

echo "Generating E-commerce Microservices..."

./ccin generate nestjs auth-service --domain auth --gcp-project $DOMAIN
./ccin generate go-gin user-service --domain user --gcp-project $DOMAIN --grpc
./ccin generate go-gin product-service --domain product --gcp-project $DOMAIN --grpc
./ccin generate go-fiber order-service --domain order --gcp-project $DOMAIN
./ccin generate go-gin inventory-service --domain inventory --gcp-project $DOMAIN --grpc
./ccin generate nestjs notification-service --domain notification --gcp-project $DOMAIN

echo "✅ All services generated successfully!"
```

### Environment Setup

```bash
# Development
./ccin generate nestjs my-api --domain item --gcp-project my-project-dev

# Staging  
./ccin generate nestjs my-api-staging --domain item --gcp-project my-project-staging

# Production
./ccin generate nestjs my-api-prod --domain item --gcp-project my-project-prod
```

## Casos de Uso por Framework

### Cuándo usar NestJS
- APIs complejas con muchas integraciones
- Equipos familiarizados con TypeScript/Angular
- Aplicaciones que requieren decoradores y DI
- Prototipado rápido con TypeScript

### Cuándo usar Go Gin
- APIs de alto rendimiento
- Microservicios con gRPC
- Servicios de backend robustos
- Cuando necesitas balance entre performance y facilidad

### Cuándo usar Go Fiber
- APIs de ultra-alta performance
- Servicios con muchas conexiones concurrentes
- Event processing
- Cuando la velocidad es crítica
