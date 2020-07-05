package auth

import (
	"github.com/confus1on/authez/internal/model"
)

type RepositoryAuth interface {
	FindUser(input model.InputAuth, typeConnection string) (interface{}, error)
}