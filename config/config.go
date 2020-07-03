package config

// ConfigMap is configuration option for authentication
type ConfigMap struct {
	DB *Database
}

// NewConfigMap return a configmap with a default value in table name (`users`)
func NewConfigMap() *ConfigMap {
	return &ConfigMap{
		DB: &Database{
			DBHost:     "",
			DBPassword: "",
			DBUsername: "",
			DBName:     "",
			TableName:  "users",
		},
	}
}
