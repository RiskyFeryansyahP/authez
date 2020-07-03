package usecase

import (
	"fmt"
	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
)

type AuthUsecase struct {
	AuthRepo auth.RepositoryAuth
}

func NewAuthUsecase(authRepo auth.RepositoryAuth) auth.UsecaseAuth {
	return &AuthUsecase{AuthRepo: authRepo}
}

func (a *AuthUsecase) AuthenticationValidation(input model.InputAuth, typeConnection string) (interface{}, error) {
	if input.DB.DBHost == "" {
		err := fmt.Errorf("database host cant be empty")
		return nil, err
	}

	if input.DB.DBUsername == "" {
		err := fmt.Errorf("database username cant be empty")
		return nil, err
	}

	if input.DB.DBPassword == "" {
		err := fmt.Errorf("database password cant be empty")
		return nil, err
	}

	if input.DB.DBName == "" {
		err := fmt.Errorf("database name cant be empty")
		return nil, err
	}

	if input.DB.TableName == "" {
		err := fmt.Errorf("table name cant be empty")
		return nil, err
	}

	if input.Username == "" {
		err := fmt.Errorf("username cant be empty")
		return nil, err
	}

	if input.Password == "" {
		err := fmt.Errorf("password cant be empty")
		return nil, err
	}

	result, err := a.AuthRepo.FindUser(input, typeConnection)
	if err != nil {
		return nil, err
	}

	return result, nil
}