package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read ENV file")
	}

	return config
}
