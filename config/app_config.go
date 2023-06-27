package config

import (
	"github.com/rafialariq/digital-bank/utility"
)

type ApiConfig struct {
	ServerPort string
}

type DbConfig struct {
	Host, Port, User, Name, Password, SslMode string
}

type AppConfig struct {
	ApiConfig
	DbConfig
}

func (c *AppConfig) readConfigFile() {
	envFilePath := ".env"
	c.ApiConfig = ApiConfig{
		ServerPort: utility.DotEnv("SERVER_PORT", envFilePath),
	}

	c.DbConfig = DbConfig{
		Host:     utility.DotEnv("DB_HOST", envFilePath),
		Port:     utility.DotEnv("DB_PORT", envFilePath),
		User:     utility.DotEnv("DB_USER", envFilePath),
		Name:     utility.DotEnv("DB_NAME", envFilePath),
		Password: utility.DotEnv("DB_PASSWORD", envFilePath),
		SslMode:  utility.DotEnv("SSL_MODE", envFilePath),
	}

}

func NewConfig() *AppConfig {
	config := AppConfig{}
	config.readConfigFile()
	return &config
}
