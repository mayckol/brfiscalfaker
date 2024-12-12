# Stage 1: Build the Go application
FROM golang:1.23.4-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o brfiscalfaker ./cmd/brfiscalfaker

# Stage 2: Create the final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/brfiscalfaker .

# Command to run the executable
ENTRYPOINT ["./brfiscalfaker"]
