package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/confus1on/authez/config"
	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
)

// AuthRepository is repository which has a configuration in it
type AuthRepository struct {
	DB *sql.DB
}

// NewAuthRepository initiate configuration and return `AuthRepository struct`
func NewAuthRepository(db *sql.DB) auth.RepositoryAuth {
	return &AuthRepository{DB: db}
}

// FindUser find user in storage and will be return interface or error
func (a *AuthRepository) FindUser(input model.InputAuth) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = ? AND password = ?", input.TableName)

	rows, err := a.DB.Query(query, input.Username, input.Password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	if len(result) <= 0 {
		return nil, fmt.Errorf("invalid username or password")
	}

	// log.Println(*result["fullname"].(*string))

	return result, nil
}

// GoogleUser request to google oauth2 to find user
func (a *AuthRepository) GoogleUser(config *config.ConfigMap, code string) (*http.Response, error) {
	token, err := config.GoogleOauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	return response, nil
}

func scanRows(rows *sql.Rows) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	for rows.Next() {
		columns, _ := rows.ColumnTypes()

		// Scan needs an array of pointers to the values it is setting
		// This creates the object and sets the values correctly
		values := make([]interface{}, len(columns))

		for key, column := range columns {
			var valueType interface{} // for checking type data each column

			switch column.DatabaseTypeName() {
			case "TEXT":
				valueType = new(string)
			case "VARCHAR":
				valueType = new(string)
			default:
				valueType = new(interface{})
			}

			result[column.Name()] = valueType
			values[key] = valueType // assign value of result pointer into values
		}

		err := rows.Scan(values...) // scan values will be affect into result because have same pointer address
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
