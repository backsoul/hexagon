# Use the official Golang image
FROM golang:1.17-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal alpine image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary file from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the Go application will run
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
