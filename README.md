# Job Queue Service

A robust job queue service built with Go, featuring asynchronous job processing, PostgreSQL persistence, and RESTful API endpoints.

## Features

- ✨ Asynchronous job processing with worker pools
- 🗄️ PostgreSQL-backed persistence
- 🚀 RESTful API endpoints
- 📝 Job status tracking
- ⚡ Configurable worker count and queue size
- 🔄 Automatic job status updates
- 📊 Pagination support for job listing

## Tech Stack

- **Language**: Go 1.23+
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Environment**: Docker support

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 12 or higher
- Docker (optional)

## Project Structure

```
job-queue/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── api/             # HTTP handlers and routes
│   ├── config/          # Configuration management
│   ├── db/              # Database connection
│   ├── models/          # Data models
│   ├── repository/      # Data access layer
│   ├── service/         # Business logic
│   ├── utils/           # Utilities and constants
│   └── worker/          # Worker pool implementation
├── Dockerfile           # Docker configuration
├── go.mod              # Go module file
├── go.sum              # Go module checksums
└── README.md           # This file
```

## Getting Started

### Local Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd job-queue
   ```

2. Copy the sample environment file:

   ```bash
   cp sample.env .env
   ```

3. Update the `.env` file with your PostgreSQL credentials and other configurations:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password_here
   DB_NAME=job_queue
   DB_SSLMODE=disable
   PORT=8080
   ```

4. Install dependencies:

   ```bash
   go mod download
   ```

5. Run the application:
   ```bash
   go run cmd/main.go
   ```

### Docker Setup

1. Build the Docker image:

   ```bash
   docker build -t job-queue .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 --env-file .env job-queue
   ```

## API Documentation

### Submit a Job

```http
POST /api/jobs
Content-Type: application/json

{
    "payload": "job data here"
}
```

**Response**:

```json
{
  "job_id": "uuid-here",
  "status": "pending"
}
```

### Get Job Status

```http
GET /api/jobs/:id
```

**Response**:

```json
{
  "id": "uuid-here",
  "payload": "job data here",
  "status": "completed",
  "result": "Processed: job data here",
  "created_at": "2024-03-20T10:00:00Z",
  "updated_at": "2024-03-20T10:01:00Z"
}
```

### List Jobs

```http
GET /api/jobs?limit=10&offset=0
```

**Response**:

```json
[
  {
    "id": "uuid-1",
    "payload": "job 1",
    "status": "completed",
    "result": "Processed: job 1",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:01:00Z"
  }
  // ... more jobs
]
```

## Job Status Values

- `pending`: Job is queued but not yet processed
- `processing`: Job is currently being processed
- `completed`: Job has been successfully processed
- `failed`: Job processing failed

## Configuration

The following environment variables can be configured:

| Variable    | Description              | Default   |
| ----------- | ------------------------ | --------- |
| DB_HOST     | PostgreSQL host          | localhost |
| DB_PORT     | PostgreSQL port          | 5432      |
| DB_USER     | PostgreSQL username      | postgres  |
| DB_PASSWORD | PostgreSQL password      | password  |
| DB_NAME     | PostgreSQL database name | job_queue |
| DB_SSLMODE  | PostgreSQL SSL mode      | disable   |
| PORT        | Server port              | 8080      |

## Development

### Running Tests

```bash
go test ./...
```

### Code Style

The project follows standard Go code style. Please ensure your code is formatted using:

```bash
go fmt ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

