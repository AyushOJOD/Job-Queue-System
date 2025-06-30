package main

import (
	"fmt"
	"log"
	"os"

	"job-queue/internal/api"
	"job-queue/internal/config"
	"job-queue/internal/db"
	"job-queue/internal/models"
	"job-queue/internal/repository"
	"job-queue/internal/service"
	"job-queue/internal/utils"
	"job-queue/internal/worker"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Load config
	cfg := config.LoadConfig()

	// Initialize DB
	dbConn, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf(utils.MsgDBConnectionFailed, err)
	}

	// Auto migrate Job table
	if err := dbConn.AutoMigrate(&models.Job{}); err != nil {
		log.Fatalf(utils.MsgDBMigrationFailed, err)
	}

	// Initialize repository and service
	jobRepo := repository.NewJobRepository(dbConn)
	jobService := service.NewJobService(jobRepo)

	// Initialize worker pool
	wp := worker.NewWorkerPool(utils.DefaultWorkers, jobService)
	wp.Start()

	// Initialize router
	r := gin.Default()

	// Setup routes
	jobHandler := api.NewJobHandler(jobService, wp)
	api.RegisterRoutes(r, jobHandler)

	// Start server
	port := os.Getenv(utils.EnvPort)
	if port == "" {
		port = utils.DefaultPort
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf(utils.MsgServerStarting, addr)
	r.Run(addr)
}
