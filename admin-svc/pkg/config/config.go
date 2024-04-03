package config

import "github.com/spf13/viper"

type Config struct{
	PORT string `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL"`
}

func LoadConfig()(config *Config,err error){
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}