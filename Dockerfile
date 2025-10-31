# --- Stage 1: Build the binary ---
FROM golang:1.25-alpine AS builder

# Install git (needed for go mod download)
RUN apk add --no-cache git

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go binary
# -o sets the output binary name
# CGO_ENABLED=0 makes it statically linked (portable)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# --- Stage 2: Create a minimal runtime image ---
FROM alpine:latest

# Install git for runtime use
RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/configure-git .

# If your app needs configuration files or assets, copy them too:
# COPY --from=builder /app/config.yaml .

# Expose any port your app listens on (optional)
EXPOSE 8000

# Configure git and start the app
CMD ["sh", "-c", "./configure-git && ./app"]
