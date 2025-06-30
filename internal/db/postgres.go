package db

import (
	"fmt"
	"log"

	"job-queue/internal/config"
	"job-queue/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		utils.PostgresDSNFormat,
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println(utils.MsgDBConnected)
	return db, nil
}
