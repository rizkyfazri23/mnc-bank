package config

import "github.com/rizkyfazri23/mnc/utils"

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host, Port, User, Password, Name, SslMode string
}

type AppConfig struct {
	ApiConfig
	DbConfig
}

func (c *AppConfig) readConfigFile() {
	api := utils.DotEnv("SERVER_PORT")
	c.DbConfig = DbConfig{
		Host:     utils.DotEnv("DB_HOST"),
		Port:     utils.DotEnv("DB_PORT"),
		User:     utils.DotEnv("DB_USER"),
		Password: utils.DotEnv("DB_PASSWORD"),
		Name:     utils.DotEnv("DB_NAME"),
		SslMode:  utils.DotEnv("SSL_MODE"),
	}
	c.ApiConfig = ApiConfig{ApiPort: api}
}

func NewConfig() AppConfig {
	cfg := AppConfig{}
	cfg.readConfigFile()
	return cfg
}
