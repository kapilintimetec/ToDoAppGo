FROM golang:1.22.3-alpine

RUN apk add --no-cache git && go install github.com/cespare/reflex@latest

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application binary
RUN go build -o main ./app/cmd/app

# Expose the application port
EXPOSE 8080

# Use reflex for hot reloading during development
CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "./app/cmd/app/main.go"]
