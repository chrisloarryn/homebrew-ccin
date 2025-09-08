# {{.ProjectName}} — Rust Axum Backend (REST {{ if .WithGRPC }}+ gRPC {{ end }})

Stack: axum (REST) + {{ if .WithGRPC }}tonic (gRPC) + {{ end }}tokio + hyper + tower (+ tower-http)

Clean structure with separation of concerns:

- src/
  - http/: routers, handlers
  - grpc/: gRPC service ({{ if .WithGRPC }}enabled{{ else }}scaffolded{{ end }})
  - domain/: models
  - services/: business logic
  - middleware/: common layers (tracing, cors)
  - main.rs: boot HTTP {{ if .WithGRPC }}and gRPC {{ end }}servers
- proto/: Protobuf contracts {{ if .WithGRPC }}(compiled with tonic-build){{ end }}

## Run

```bash
cargo run
```

## REST Examples

- GET /health
- GET /api/{{.DomainLower}}
- POST /api/{{.DomainLower}}

{{ if .WithGRPC }}
## gRPC

Proto: proto/{{.DomainLower}}.proto

Run server with REST and gRPC concurrently on Tokio runtime.
{{ end }}
