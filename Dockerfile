# Use the official Golang image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
