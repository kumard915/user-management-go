# Use the official Golang image as a base
FROM golang:1.23.4-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod ./
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main .

# Use a lightweight image to run the Go application
FROM alpine:latest

# Set working directory in the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Set the default command to run the application
CMD ["/app/main"]

# Expose port 8080 (change if your app uses a different port)
EXPOSE 8080
