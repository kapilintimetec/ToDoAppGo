package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

var config *Config

func GetConfig() (*Config, error) {
	if config == nil {
		err := initConfig()
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config = &Config{
		DbHost:     viper.GetString("DB_HOST"),
		DbPort:     viper.GetString("DB_PORT"),
		DbUser:     viper.GetString("DB_USER"),
		DbPassword: viper.GetString("DB_PASSWORD"),
		DbName:     viper.GetString("DB_NAME"),
	}

	return nil
}
