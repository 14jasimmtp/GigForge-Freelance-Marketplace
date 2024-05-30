package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT         string `mapstructure:"PORT"`
	AUTH_SVC     string `mapstructure:"AUTH_SVC"`
	ATokenSecret string `mapstructure:"ATokenSecret"`
	CHAT_SVC     string `mapstructure:"CHAT_SVC"`
}

func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("api-gateway_config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	

	return
}
