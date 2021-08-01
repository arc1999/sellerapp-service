package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

//AppConfig interface
type AppConfig struct {
}

// LoadEnv configuration from env file and set it to system environment variable.
func (config AppConfig) LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	keys := viper.AllKeys()
	log.Println(keys)
	for _, k := range keys {
		k = strings.ToUpper(k)
		log.Println(k)
		log.Println(viper.GetString(k))
		envValue, found := os.LookupEnv(k)
		if !found || envValue == "" {
			os.Setenv(k, viper.GetString(k))
		}
	}
}
