package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // autoload environment variables
)

// ConfigMap is configuration option for authentication
type ConfigMap struct {
	DB *Database
}

// NewConfigMap return a configmap with a default value in table name (`users`)
func NewConfigMap() *ConfigMap {
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")

	return &ConfigMap{
		DB: &Database{
			DBHost:     dbHost,
			DBPassword: dbPassword,
			DBUsername: dbUsername,
			DBName:     dbName,
			TableName:  "users",
		},
	}
}
