job-queue/
├── cmd/
│ └── server/ # Entry point (main.go)
├── internal/ # Application core (encapsulation)
│ ├── api/ # HTTP handlers/controllers
│ │ └── job_handler.go
│ ├── config/ # Configuration (env loader, constants)
│ │ └── config.go
│ ├── models/ # Data models
│ │ └── job.go
│ ├── repository/ # DB layer (Interface + Implementation)
│ │ ├── job_repo.go
│ │ └── job_repo_postgres.go
│ ├── service/ # Business logic (Interface + Implementation)
│ │ └── job_service.go
│ ├── worker/ # Worker pool (Job consumers)
│ │ └── worker_pool.go
│ ├── utils/ # Utilities (logger, error handler)
│ │ ├── logger.go
│ │ └── response.go
│ └── db/ # Database migration & connection
│ └── postgres.go
├── Dockerfile # For containerization
├── go.mod
├── go.sum
└── README.md
