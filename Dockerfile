# First stage: Build the Go binary
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -v -o main ./cmd

# Second stage: Create a minimal image for running the Go binary
FROM alpine:latest

# Install the ca-certificates package to have SSL/TLS certificates
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application wil run on (default for Echo is 1323)
EXPOSE 1323

# Command to run the binary
CMD ["./main"]
