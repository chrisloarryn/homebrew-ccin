# syntax=docker/dockerfile:1

# Build stage
FROM swift:6.1.2-jammy AS build
WORKDIR /app

# Copy manifest and resolve dependencies first (better layer caching)
COPY Package.swift ./
RUN swift package resolve

# Copy the rest of the sources
COPY Sources ./Sources

# Build in release mode
RUN swift build -c release --product Run

# Production stage
FROM ubuntu:22.04 AS runtime
WORKDIR /run

# Install runtime dependencies (ca-certificates, tzdata)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates

# Copy built binaries and any required libraries
COPY --from=build /app/.build/release/Run /run/Run

# Non-root user
RUN useradd -m vapor && chown -R vapor:vapor /run
USER vapor

# Expose ports
EXPOSE {{.Port}}
{{- if .WithGRPC}}
EXPOSE 50051
{{- end}}

# Healthcheck (simple HTTP ping)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://127.0.0.1:{{.Port}}/health || exit 1

# Run
ENTRYPOINT ["/run/Run"]
