package repository

import (
	"database/sql"
	"fmt"
	"github.com/confus1on/authez/config"
	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
)

type AuthRepository struct {
	Config *config.ConfigMap
}

func NewAuthRepository() auth.RepositoryAuth {
	cfg := config.NewConfigMap()

	return &AuthRepository{Config: cfg}
}

func (a AuthRepository) FindUser(input model.InputAuth, typeConnection string) (interface{}, error) {
	newDatabase := config.NewDatabase(input.DB)

	var query string
	
	// change default value database with new database input value
	a.Config.DB = newDatabase

	db, err := a.Config.DB.SetConnectionDB(typeConnection)
	if err != nil {
		return nil, err
	}

	switch typeConnection {
	case "mysql":
		query = fmt.Sprintf("SELECT * FROM %s WHERE username = ? AND password = ?", input.DB.TableName)
		break
	case "postgresql":
		query = fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password = $2", input.DB.TableName)
		break
	}

	rows, err := db.Query(query, input.Username, input.Password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func scanRows(rows *sql.Rows) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}

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

		err = rows.Scan(values...) // scan values will be affect into result because have same pointer address
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
