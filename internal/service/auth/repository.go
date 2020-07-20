package auth

import (
	"net/http"

	"github.com/confus1on/authez/config"
	"github.com/confus1on/authez/internal/model"
)

// RepositoryAuth is absract will be use in repository package
type RepositoryAuth interface {
	FindUser(input model.InputAuth) (interface{}, error)
	GoogleUser(config *config.ConfigMap, code string) (*http.Response, error)
}
