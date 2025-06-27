# Multi-stage build for Go Fiber application
FROM golang:1.24.4-alpine AS builder

# Install git and ca-certificates (needed for go mod download)
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser for security
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o {{.ProjectName}} .

# Final stage: create minimal runtime image
FROM scratch

# Import from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy binary
COPY --from=builder /build/{{.ProjectName}} /{{.ProjectName}}

# Use unprivileged user
USER appuser

# Expose port
EXPOSE {{.Port}}
{{- if .WithGRPC}}
EXPOSE 50051
{{- end}}

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/{{.ProjectName}}", "health"] || exit 1

# Run the binary
ENTRYPOINT ["/{{.ProjectName}}"]
