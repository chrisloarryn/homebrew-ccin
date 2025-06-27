# {{.ProjectName}}

A Go CRUD API built with Gin framework.

## Features

- ✅ REST API with Gin
- ✅ PostgreSQL database
- ✅ CRUD operations for {{.DomainLower}}s
{{- if .WithGRPC}}
- ✅ gRPC support
{{- end}}
{{- if .GCPProject}}
- ✅ GCP monitoring and metrics
{{- end}}
- ✅ Docker support
- ✅ Environment configuration

## Quick Start

1. **Clone and setup**
   ```bash
   cd {{.ProjectName}}
   cp .env.example .env
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup database**
   ```bash
   # Create PostgreSQL database
   createdb {{.ProjectName}}_dev
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:{{.Port}}`.

{{- if .WithGRPC}}
The gRPC server will be available at `localhost:50051`.
{{- end}}

## API Endpoints

### {{.DomainTitle}} Management
- `GET /api/v1/{{.DomainLower}}s` - Get all {{.DomainLower}}s
- `GET /api/v1/{{.DomainLower}}s/:id` - Get {{.DomainLower}} by ID
- `POST /api/v1/{{.DomainLower}}s` - Create new {{.DomainLower}}
- `PUT /api/v1/{{.DomainLower}}s/:id` - Update {{.DomainLower}}
- `DELETE /api/v1/{{.DomainLower}}s/:id` - Delete {{.DomainLower}}

### Health Check
- `GET /health` - Health check endpoint

## Example Requests

### Create {{.DomainLower}}
```bash
curl -X POST http://localhost:{{.Port}}/api/v1/{{.DomainLower}}s \
  -H "Content-Type: application/json" \
  -d '{"name": "Example {{.DomainTitle}}", "description": "This is an example"}'
```

### Get all {{.DomainLower}}s
```bash
curl http://localhost:{{.Port}}/api/v1/{{.DomainLower}}s
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | HTTP server port | {{.Port}} |
{{- if .WithGRPC}}
| GRPC_PORT | gRPC server port | 50051 |
{{- end}}
| DATABASE_URL | PostgreSQL connection string | postgres://localhost/{{.ProjectName}}_dev?sslmode=disable |
| GIN_MODE | Gin mode (debug/release) | debug |
{{- if .GCPProject}}
| GCP_PROJECT | GCP project ID for metrics | {{.GCPProject}} |
{{- end}}

## Docker

### Build and run with Docker

```bash
# Build the image
docker build -t {{.ProjectName}} .

# Run the container
docker run -p {{.Port}}:{{.Port}} {{.ProjectName}}
```

### Run with Docker Compose

```bash
docker-compose up -d
```

## Development

### Project Structure

```
{{.ProjectName}}/
├── main.go                 # Application entry point
├── internal/
│   ├── api/               # API routes and setup
│   ├── config/            # Configuration management
│   ├── database/          # Database connection and setup
│   ├── handlers/          # HTTP request handlers
│   ├── models/            # Data models
│   ├── services/          # Business logic
{{- if .WithGRPC}}
│   ├── grpc/              # gRPC server and handlers
{{- end}}
{{- if .GCPProject}}
│   ├── metrics/           # GCP metrics integration
{{- end}}
│   └── middleware/        # HTTP middleware
{{- if .WithGRPC}}
├── proto/                 # Protocol buffer definitions
{{- end}}
├── .env.example           # Environment variables template
├── Dockerfile             # Docker configuration
├── Makefile              # Build and development tasks
└── README.md             # This file
```

### Make Commands

```bash
make build          # Build the application
make run            # Run the application
make test           # Run tests
make clean          # Clean build artifacts
make docker-build   # Build Docker image
make docker-run     # Run Docker container
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
