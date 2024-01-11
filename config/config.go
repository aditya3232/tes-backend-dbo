package config

import (
	"log"

	"github.com/spf13/viper"
)

var CONFIG = load()

type envConfigs struct {
	DEBUG int `mapstructure:"DEBUG"`

	APP_PORT string `mapstructure:"APP_PORT"`
	APP_HOST string `mapstructure:"APP_HOST"`

	DB_HOST    string `mapstructure:"DB_HOST"`
	DB_PORT    string `mapstructure:"DB_PORT"`
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PASS    string `mapstructure:"DB_PASS"`
	DB_NAME    string `mapstructure:"DB_NAME"`
	DB_CHARSET string `mapstructure:"DB_CHARSET"`
	DB_LOC     string `mapstructure:"DB_LOC"`
}

func load() (config *envConfigs) {
	viper.AddConfigPath("../tes-backend-dbo/config")

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
