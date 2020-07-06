package model

import "github.com/confus1on/authez/config"

// InputAuth is input config for authentication
type InputAuth struct {
	DB       config.Database `json:"database"`
	Username string          `json:"username"`
	Password string          `json:"password"`
}
