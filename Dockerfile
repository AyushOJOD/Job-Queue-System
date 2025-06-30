FROM golang:1.23-alpine

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN go build -o job-queue ./cmd/main.go

# Expose port
EXPOSE 8080

# Run
CMD ["./job-queue"]
