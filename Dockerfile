# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.18

WORKDIR /app

# Add non root user for security
RUN adduser -D -g '' appuser

# Copy only necessary files from builder
COPY --from=builder /app/main .
COPY templates/ templates/
COPY static/ static/

# Use non root user
USER appuser

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]