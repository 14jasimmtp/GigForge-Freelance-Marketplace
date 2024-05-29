package config

import "github.com/spf13/viper"

type Config struct{
	PORT string `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL"`
	User_SVC string `mapstructure:"USER_SVC"`
}

func LoadConfig() (err error){
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	return 
}