# Docker image for CCIN CLI
FROM alpine:3.18

# Install ca-certificates and other dependencies
RUN apk --no-cache add ca-certificates curl git

WORKDIR /root/

# Copy the binary (you need to build it first with: go build -o ccin .)
COPY ccin .

# Copy templates directory (needed for code generation)
COPY templates/ ./templates/

# Make it executable and create symbolic link
RUN chmod +x ./ccin && \
    ln -s /root/ccin /usr/local/bin/ccin

# Set the entrypoint
ENTRYPOINT ["ccin"]

# Default command
CMD ["--help"]

# Metadata
LABEL org.opencontainers.image.title="CCIN CLI"
LABEL org.opencontainers.image.description="ChrisLoarryn's Comprehensive Code Integration & Initialization Tool"
LABEL org.opencontainers.image.vendor="Chris Loarryn"
LABEL org.opencontainers.image.source="https://github.com/chrisloarryn/homebrew-ccin"
LABEL org.opencontainers.image.documentation="https://github.com/chrisloarryn/homebrew-ccin/blob/main/README.md"
