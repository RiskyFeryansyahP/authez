package model

// InputAuth is input config for authentication
type InputAuth struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	TableName string `json:"table_name"`
}
