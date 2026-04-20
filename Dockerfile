# ============================================================
# Stage 1: Build the Go binary
# ============================================================
FROM golang:1.25-alpine AS builder

# Install git (required for go mod download with private repos)
RUN apk add --no-cache git

WORKDIR /app

# Copy go module files first for better Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked binary
# NOTE: CGO_ENABLED=0 means SQLite driver will NOT work in this container.
#       Use PostgreSQL, MySQL, or SQL Server when deploying with Docker.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/server .

# ============================================================
# Stage 2: Create a minimal production image
# ============================================================
FROM alpine:latest

# Install CA certificates for HTTPS calls and timezone data
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Copy .env.example as a reference (actual .env should be mounted or use env vars)
COPY --from=builder /app/.env.example .env.example

# Expose the application port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./server"]
