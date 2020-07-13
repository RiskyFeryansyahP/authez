package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/confus1on/authez/internal/model"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3" // sqlite driver
)

func mockConnection(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "file:authez?mode=memory&cache=shared")
	require.NoError(t, err)

	return db
}

func mockData(db *sql.DB) {
	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username VARCHAR, password VARCHAR, fullname TEXT)")
	_ , err := stmt.Exec()
	if err != nil {
		log.Printf("error executing create table: %+v", err)
	}

	stmt, _ = db.Prepare("INSERT INTO users(id, username, password, fullname) VALUES (?, ?, ?, ?)")
	_, err = stmt.Exec(1, "risky", "risky123", "risky feryansyah")
	if err != nil {
		log.Printf("error executing insert table: %+v", err)
	}
}

func TestAuthRepository_FindUser(t *testing.T) {
	db := mockConnection(t)

	mockData(db)

	authRepo := NewAuthRepository(db)

	t.Run("test auth repository to find user", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "risky123",
			TableName: "users",
		}

		_, err := authRepo.FindUser(input)
		require.NoError(t, err)
	})

	t.Run("test find user with non exist table", func(t *testing.T) {
		input := model.InputAuth{
			Username: "",
			Password: "",
		}

		_, err := authRepo.FindUser(input)
		require.Error(t, err)
	})

	t.Run("test find user with empty data", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "",
			Password:  "risky",
			TableName: "users",
		}

		_, err := authRepo.FindUser(input)
		require.Error(t, err)
	})
}
