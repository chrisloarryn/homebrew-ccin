# {{.ProjectName}}

A NestJS CRUD API built with TypeScript.

## Features

- ✅ REST API with NestJS (TypeScript)
- ✅ Modular structure (Controller/Service/Module + DTOs/Entities)
- ✅ Health check endpoint
- ✅ Docker (multi-stage) + Makefile workflows
- ✅ Environment-driven configuration
{{- if .GCPProject}}
- ✅ Optional GCP metrics interceptor
{{- end}}

## Quick Start

1. Install dependencies
   ```bash
   cd {{.ProjectName}}
   npm install
   ```

2. Run the application (development)
   ```bash
   npm run start:dev
   ```

The API will be available at `http://localhost:{{.Port}}`.

## API Endpoints

### Health Check
- `GET /health` — returns `OK`

### {{.DomainTitle}} Management
The RESTful routes are mounted under `/api/v1/{{.DomainLower}}s`.

- `GET /api/v1/{{.DomainLower}}s` — List all
- `GET /api/v1/{{.DomainLower}}s/:id` — Get by ID
- `POST /api/v1/{{.DomainLower}}s` — Create
- `PUT /api/v1/{{.DomainLower}}s/:id` — Update by ID
- `DELETE /api/v1/{{.DomainLower}}s/:id` — Delete by ID

### Example Requests

Create {{.DomainLower}}:
```bash
curl -X POST http://localhost:{{.Port}}/api/v1/{{.DomainLower}}s \
  -H "Content-Type: application/json" \
  -d '{"name": "Example {{.DomainTitle}}"}'
```

List {{.DomainLower}}:
```bash
curl http://localhost:{{.Port}}/api/v1/{{.DomainLower}}s
```

## Project Structure

```
{{.ProjectName}}/
├── Dockerfile                     # Multi-stage Docker build
├── Makefile                       # Build & run helpers
├── package.json                   # Scripts and dependencies
└── src/
    ├── main.ts                    # Entry point
    ├── app.module.ts              # Root module
    ├── common/
    │   └── interceptors/
    │       └── gcp-metrics.interceptor.ts   # (Optional) GCP metrics
    └── {{.DomainLower}}/
        ├── dto/
        │   ├── create-{{.DomainLower}}.dto.ts
        │   └── update-{{.DomainLower}}.dto.ts
        ├── entities/
        │   └── {{.DomainLower}}.entity.ts
        ├── {{.DomainLower}}.controller.ts
        ├── {{.DomainLower}}.module.ts
        └── {{.DomainLower}}.service.ts
```

## Makefile Commands

```bash
make help           # Show available targets
make build          # Build the project
make dev            # Run in watch mode
make test           # Run tests (if configured)
make docker-build   # Build Docker image
make docker-run     # Run Docker container
```

## Docker

Build and run with Docker:
```bash
docker build -t {{.ProjectName}} .
docker run -p {{.Port}}:{{.Port}} {{.ProjectName}}
```

## Environment

| Variable | Description | Default |
|----------|-------------|---------|
| PORT     | HTTP server port | {{.Port}} |
{{- if .GCPProject}}
| GCP_PROJECT | GCP project ID for metrics | {{.GCPProject}} |
{{- end}}

## License
This project is licensed under the MIT License - see the LICENSE file for details.
