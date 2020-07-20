package auth

import (
	"github.com/confus1on/authez/config"
	"github.com/confus1on/authez/internal/model"
)

// UsecaseAuth is abstract will be use in usecase package
type UsecaseAuth interface {
	AuthenticationValidation(input model.InputAuth) (interface{}, error)
	GoogleAuthentication(config *config.ConfigMap, state, code string) (interface{}, error)
}
