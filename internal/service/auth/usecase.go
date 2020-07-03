package auth

import (
	"github.com/confus1on/authez/internal/model"
)

type UsecaseAuth interface {
	AuthenticationValidation(input model.InputAuth, typeConnection string) (interface{}, error)
}