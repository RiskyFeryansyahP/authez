package config

import (
	"database/sql"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"

	_ "github.com/go-sql-driver/mysql"    // mysql driver
	_ "github.com/joho/godotenv/autoload" // autoload environment variables
)

// ConfigMap is configuration option for authentication
type ConfigMap struct {
	DB *Database
	GoogleOauth *oauth2.Config
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
		GoogleOauth: &oauth2.Config{
			ClientID:     os.Getenv("ClientID"),
			ClientSecret: os.Getenv("ClientSecret"),
			Endpoint:     google.Endpoint,
			RedirectURL:  "http://localhost:8080/auth/google/callback",
			Scopes:       []string{
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
	}
}

// SetConnection set a new connection database
func (c *ConfigMap) SetConnection() (*sql.DB, error) {
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:8889)/%s", c.DB.DBUsername, c.DB.DBPassword, c.DB.DBHost, c.DB.DBName)

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	return db, nil
}
