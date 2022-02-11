package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is a default configuration struc for this app
type Config struct {
	Version  string
	DB_User  string
	DB_Pass  string
	DB_Host  string
	DB_Port  string
	DB_Name  string
	Web_Port string
	Env      string
}

// InitConfig initiates this application default configuration
func InitConfig() (*Config, error) {

	config := &Config{
		Version:  viper.GetString("VERSION"),
		DB_User:  viper.GetString("DB_USER"),
		DB_Pass:  viper.GetString("DB_PASS"),
		DB_Host:  viper.GetString("DB_HOST"),
		DB_Port:  viper.GetString("DB_PORT"),
		DB_Name:  viper.GetString("DB_NAME"),
		Web_Port: viper.GetString("WEB_PORT"),
		Env:      viper.GetString("ENV"),
	}

	if len(config.Version) == 0 {
		return nil, fmt.Errorf("VERSION is not set")
	}

	if len(config.DB_User) == 0 {
		return nil, fmt.Errorf("DB_USER is not set")
	}

	if len(config.DB_Pass) == 0 {
		return nil, fmt.Errorf("DB_PASS is not set")
	}

	if len(config.DB_Host) == 0 {
		return nil, fmt.Errorf("DB_HOST is not set")
	}

	if len(config.DB_Port) == 0 {
		return nil, fmt.Errorf("DB_PORT is not set")
	}

	if len(config.DB_Name) == 0 {
		return nil, fmt.Errorf("DB_NAME is not set")
	}

	if len(config.Web_Port) == 0 {
		return nil, fmt.Errorf("WEB_PORT is not set")
	}

	if len(config.Env) == 0 {
		return nil, fmt.Errorf("ENV is not set")
	}

	return config, nil
}
