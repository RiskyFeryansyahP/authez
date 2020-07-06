package auth

import (
	"github.com/confus1on/authez/internal/model"
)

// RepositoryAuth is absract will be use in repository package
type RepositoryAuth interface {
	FindUser(input model.InputAuth, typeConnection string) (interface{}, error)
}
