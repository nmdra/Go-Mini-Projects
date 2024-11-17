# Stage 1: Build the Go application
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application with optimizations
RUN go build -ldflags="-s -w" -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled application from the builder stage
COPY --from=builder /app/main .

# Copy templates and other static files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static 

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./main"]
