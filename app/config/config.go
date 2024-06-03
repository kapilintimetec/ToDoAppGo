package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DbHost string `mapstructure:"DB_HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbUser string `mapstructure:"DB_USER"`
	DbPass string `mapstructure:"DB_PASSWORD"`
	DbName string `mapstructure:"DB_NAME"`
	DbURL  string
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %s", err))
	}
}
