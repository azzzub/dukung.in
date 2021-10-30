package tools

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while rading config file %s", err)
	}

	var envKey string

	if key == "ENV" {
		envKey = key
	} else {
		envKey = viper.Get("ENV").(string) + "_" + key
	}

	value, ok := viper.Get(envKey).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
