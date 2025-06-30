package config

import (
	"job-queue/internal/utils"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	cfg := &Config{
		DBHost:     getEnv(utils.EnvDBHost, utils.DefaultDBHost),
		DBPort:     getEnv(utils.EnvDBPort, utils.DefaultDBPort),
		DBUser:     getEnv(utils.EnvDBUser, utils.DefaultDBUser),
		DBPassword: getEnv(utils.EnvDBPassword, utils.DefaultDBPassword),
		DBName:     getEnv(utils.EnvDBName, utils.DefaultDBName),
		DBSSLMode:  getEnv(utils.EnvDBSSLMode, utils.DefaultDBSSLMode),
	}
	log.Println(utils.MsgConfigLoaded)
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
