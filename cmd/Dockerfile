# First stage: Build the Go binary
FROM golang:1.23.1-alpine AS builder

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -v -o devseconnect ./main.go

# Second stage: Create a minimal image for running the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/devseconnect .

# Set environment variables
ENV RABBITMQ_CONNECTION_URL=amqp://devseconnect:admin1234@rabbitmq:5672/

# Command to run the binary
CMD ["./devseconnect"]
