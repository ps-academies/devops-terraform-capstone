package settings

import (
"github.com/spf13/viper"
)

type IPostgresConfiguration interface {
	Host() string
	Port() int
	Username() string
	Password() string
	DB() string
}

const (
	postgresKey         = "postgres."
	postgresHostKey     = postgresKey + "host"
	postgresPortKey     = postgresKey + "port"
	postgresUsernameKey = postgresKey + "username"
	postgresPasswordKey = postgresKey + "password"
	postgresDBKey       = postgresKey + "db"
)

type PostgresConfiguration struct{}

func NewPostgresConfiguration() *PostgresConfiguration {
	viper.SetDefault(postgresHostKey, "localhost")
	viper.SetDefault(postgresPortKey, 5432)
	viper.SetDefault(postgresUsernameKey, "postgres")
	viper.SetDefault(postgresPasswordKey, "postgres")
	viper.SetDefault(postgresDBKey, "todos")
	return &PostgresConfiguration{}
}

func (rds PostgresConfiguration) Host() string {
	return viper.GetString(postgresHostKey)
}

func (rds PostgresConfiguration) Port() int {
	return viper.GetInt(postgresPortKey)
}

func (rds PostgresConfiguration) Username() string {
	return viper.GetString(postgresUsernameKey)
}

func (rds PostgresConfiguration) Password() string {
	return viper.GetString(postgresPasswordKey)
}

func (rds PostgresConfiguration) DB() string {
	return viper.GetString(postgresDBKey)
}

