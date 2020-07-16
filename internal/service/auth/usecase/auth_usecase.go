package usecase

import (
	"fmt"

	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
)

// AuthUsecase is usecase which has a repository auth in it
type AuthUsecase struct {
	AuthRepo auth.RepositoryAuth
}

// NewAuthUsecase initiate `auth usecase`
func NewAuthUsecase(authRepo auth.RepositoryAuth) auth.UsecaseAuth {
	return &AuthUsecase{AuthRepo: authRepo}
}

// AuthenticationValidation validate input from request before forwarded to repository
func (a *AuthUsecase) AuthenticationValidation(input model.InputAuth) (interface{}, error) {
	if input.TableName == "" {
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

	result, err := a.AuthRepo.FindUser(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AuthUsecase) GoogleAuthentication() (interface{}, error) {
	panic("implement me")
}
