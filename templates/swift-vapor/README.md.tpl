# {{.ProjectName}}

A Swift backend built with Vapor 4.

## Features

- ✅ REST API with Vapor 4 (clean layers: Controllers/Services/Models)
{{- if .WithGRPC}}
- ✅ gRPC scaffolding (Proto and placeholders included)
{{- end}}
- ✅ Health check endpoint
- ✅ Docker (multi-stage) + Makefile workflows
- ✅ Environment configuration via Swift (no DB by default)

## Prerequisites

- Swift 6.1.2
- Vapor toolbox (optional): `brew install vapor/tap/vapor` or use SwiftPM directly
- Docker (optional for containerized runs)

## Quick Start

1. Clone and setup
   ```bash
   cd {{.ProjectName}}
   ```

2. Resolve dependencies
   ```bash
   swift package resolve
   ```

3. Run the application
   ```bash
   swift run Run
   ```

The API will be available at `http://localhost:{{.Port}}`.
{{- if .WithGRPC}}
The gRPC endpoint (scaffold) is planned at `localhost:50051` (see gRPC section below).
{{- end}}

## API Endpoints

### Health Check
- `GET /health` — returns `OK`

### {{.DomainTitle}} Management
The RESTful routes are mounted under `/api/v1/{{.DomainLower}}`.

- `GET /api/v1/{{.DomainLower}}` — List all
- `GET /api/v1/{{.DomainLower}}/:id` — Get by ID (UUID)
- `POST /api/v1/{{.DomainLower}}` — Create
- `PUT /api/v1/{{.DomainLower}}/:id` — Update by ID
- `DELETE /api/v1/{{.DomainLower}}/:id` — Delete by ID

### Example Requests

Create {{.DomainLower}}:
```bash
curl -X POST http://localhost:{{.Port}}/api/v1/{{.DomainLower}} \
  -H "Content-Type: application/json" \
  -d '{"name": "Example {{.DomainTitle}}"}'
```

List {{.DomainLower}}:
```bash
curl http://localhost:{{.Port}}/api/v1/{{.DomainLower}}
```

## Project Structure

```
{{.ProjectName}}/
├── Package.swift                     # SwiftPM manifest
├── Sources/
│   ├── App/
│   │   ├── Controllers/             # HTTP controllers
│   │   ├── Middleware/              # Middlewares (Metrics, etc.)
│   │   ├── Models/                  # Domain models (Codable)
│   │   ├── Services/                # Business logic services
{{- if .WithGRPC}}
│   │   ├── GRPC/                    # gRPC service placeholders
{{- end}}
│   │   ├── configure.swift          # App configuration & bootstrap
│   │   └── routes.swift             # Routes registration
│   └── Run/
│       └── main.swift               # Entry point
{{- if .WithGRPC}}
├── Proto/
│   └── {{.DomainLower}}.proto       # Example proto definitions
{{- end}}
├── Dockerfile
├── Makefile
└── README.md
```

## Makefile Commands

```bash
make deps             # Resolve SwiftPM dependencies
make build            # Build (debug)
make build-release    # Build (release)
make run              # Run the app
make docker-build     # Build Docker image
make docker-run       # Run Docker container (maps ports)
```

## Docker

Build and run with Docker:
```bash
docker build -t {{.ProjectName}} .
docker run -p {{.Port}}:{{.Port}} {{- if .WithGRPC}} -p 50051:50051{{- end}} {{.ProjectName}}
```

## gRPC (Optional)
{{- if .WithGRPC}}
This template includes:
- Proto file at `Proto/{{.DomainLower}}.proto`
- Placeholder Swift file at `Sources/App/GRPC/{{.DomainTitle}}GRPCService.swift`

To enable gRPC:
1. Install `protoc` and `grpc-swift` plugins.
2. Generate Swift files from proto:
   ```bash
   protoc \
     --swift_out=./Sources/App/GRPC \
     --swift-grpc_out=./Sources/App/GRPC \
     Proto/{{.DomainLower}}.proto
   ```
3. Implement and bind your service provider in `configure.swift` (see `startGRPCServer`).
{{- else}}
You can enable gRPC generation using the CLI flag `--grpc` when generating the project: `ccin generate swift-vapor {{.ProjectName}} --domain {{.DomainLower}} --grpc`.
{{- end}}

## Environment

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | HTTP server port | {{.Port}} |

## Notes
- This template uses in-memory storage for demonstration purposes. Replace with a database or persistence layer as needed (e.g., Fluent + PostgreSQL).
- The MetricsMiddleware logs request durations to help with basic observability.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
