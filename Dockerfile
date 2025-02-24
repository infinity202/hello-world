# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install necessary build tools and security updates
RUN apk update && \
    apk add --no-cache git ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with security and optimization flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s \
    -X main.version=$(git describe --tags --always) \
    -X main.buildTime=$(date -u '+%Y-%m-%d_%H:%M:%S')" \
    -o main .

# Final stage
FROM alpine:3.18

# Add security updates and required packages
RUN apk update && \
    apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# Add non root user for security
RUN adduser -D -g '' appuser && \
    mkdir -p /app/templates /app/static && \
    chown -R appuser:appuser /app

# Copy only necessary files from builder
COPY --from=builder /app/main .
COPY --chown=appuser:appuser templates/ templates/
COPY --chown=appuser:appuser static/ static/
COPY --chown=appuser:appuser .env.production .env

# Use non root user
USER appuser

# Set environment variables
ENV APP_ENV=production \
    TZ=UTC

# Expose the port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ping || exit 1

# Run the application
CMD ["./main"]