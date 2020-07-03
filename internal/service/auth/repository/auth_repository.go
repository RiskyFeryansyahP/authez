package repository

import (
	"fmt"
	"github.com/confus1on/authez/config"
	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
	"log"
	"reflect"
)

type AuthRepository struct {
	Config *config.ConfigMap
}

func NewAuthRepository() auth.RepositoryAuth {
	cfg := config.NewConfigMap()

	return &AuthRepository{Config: cfg}
}

func (a AuthRepository) FindUser(input model.InputAuth, typeConnection string) (interface{}, error) {
	result := map[string]interface{}{}

	newDatabase := config.NewDatabase(input.DB)
	
	// change default value database with new database input value
	a.Config.DB = newDatabase

	db, err := a.Config.DB.SetConnectionDB(typeConnection)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1 AND password = $2", input.DB.TableName)

	rows, err := db.Query(query, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		log.Println("1")
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}

		// Scan needs an array of pointers to the values it is setting
		// This creates the object and sets the values correctly
		values := make([]interface{}, len(columns))

		for key, column := range columns {
			result[column.Name()] = reflect.New(column.ScanType()).Interface()
			values[key] = result[column.Name()] // assign value of result pointer into values
		}

		err = rows.Scan(values...) // scan values will be affect into result because have same pointer address
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
