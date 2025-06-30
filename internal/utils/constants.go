package utils

// Environment variable keys
const (
	EnvDBHost     = "DB_HOST"
	EnvDBPort     = "DB_PORT"
	EnvDBUser     = "DB_USER"
	EnvDBPassword = "DB_PASSWORD"
	EnvDBName     = "DB_NAME"
	EnvDBSSLMode  = "DB_SSLMODE"
	EnvPort       = "PORT"
)

// Default values
const (
	DefaultDBHost     = "localhost"
	DefaultDBPort     = "5432"
	DefaultDBUser     = "postgres"
	DefaultDBPassword = "password"
	DefaultDBName     = "job_queue"
	DefaultDBSSLMode  = "disable"
	DefaultPort       = "8080"
	DefaultWorkers    = 5
	DefaultQueueSize  = 100
)

// Database connection string format
const PostgresDSNFormat = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"

// API Response messages
const (
	MsgPayloadRequired     = "payload is required"
	MsgFailedToSubmitJob   = "failed to submit job"
	MsgFailedToFetchJob    = "failed to fetch job"
	MsgJobNotFound         = "job not found"
	MsgFailedToFetchJobs   = "failed to fetch jobs"
	MsgConfigLoaded        = "Config loaded successfully"
	MsgDBConnected         = "Database connected successfully"
	MsgDBMigrationFailed   = "Failed to migrate database: %v"
	MsgDBConnectionFailed  = "Failed to connect to database: %v"
	MsgServerStarting      = "Server running at %s"
	MsgWorkersStarted      = "Started %d workers"
	MsgWorkerProcessing    = "Worker %d processing job %s"
	MsgWorkerCompleted     = "Worker %d completed job %s"
	MsgWorkerFailed        = "Worker %d failed job %s: %v"
	MsgJobAddedToQueue     = "Adding job %s to queue"
	MsgJobProcessedResult  = "Processed: %s"
)
