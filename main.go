package main

import (
	log "github.com/mgutz/logxi/v1"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Couldn't read the config file")
	}
}

func main() {
}
