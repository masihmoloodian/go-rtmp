package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServicePort string `mapstructure:"service_port"`
	DBConnection string `mapstructure:"db_connection"`
}

var AppConfig *Config

func LoadAppConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't read app configuraion: %v", err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Can't unmarshal app config: %v", err)
	}
}