package config

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgresql driver
)

// Database is a database configuration
type Database struct {
	DBHost     string `json:"db_host"`
	DBPassword string `json:"db_password"`
	DBUsername string `json:"db_username"`
	DBName     string `json:"db_name"`
	TableName  string `json:"table_name"`
}

// NewDatabase return a new configmap with new value db configuration
func NewDatabase(db Database) *Database {
	return &Database{
		DBHost:     db.DBHost,
		DBPassword: db.DBPassword,
		DBUsername: db.DBUsername,
		DBName:     db.DBName,
		TableName:  db.TableName,
	}
}

// SetConnectionDB return sql database connection
func (d *Database) SetConnectionDB(typeConnection string) (*sql.DB, error) {
	db, err := typeOfDatabase(typeConnection, d)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)

	return db, nil
}

func typeOfDatabase(typeConnection string, db *Database) (*sql.DB, error) {
	switch typeConnection {
	case "mysql":
		databaseURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db.DBUsername, db.DBPassword, db.DBHost, db.DBName)

		db, err := sql.Open("mysql", databaseURL)
		if err != nil {
			return nil, err
		}

		return db, nil

	case "postgresql":
		databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", db.DBUsername, db.DBPassword, db.DBHost, db.DBName)

		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	err := errors.New("please use option type database")
	return nil, err
}
