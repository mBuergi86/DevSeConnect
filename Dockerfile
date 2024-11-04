FROM golang:1.23.2-alpine AS builder
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM alpine:latest

# Copy the binary from the builder stage to the final stage.
COPY --from=builder /main /main

# Set the work directory
WORKDIR /app

# Expose the port the app runs on
EXPOSE 1323

# Run
CMD ["/main"]
