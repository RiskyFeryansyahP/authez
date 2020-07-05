package model

import "github.com/confus1on/authez/config"

type InputAuth struct {
	DB       config.Database `json:"database"`
	Username string          `json:"username"`
	Password string          `json:"password"`
}
