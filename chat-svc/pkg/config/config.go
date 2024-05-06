package config

import "github.com/spf13/viper"

type Config struct{
	PORT string `mapstructure:"PORT"`
}

func LoadConfig() (config *Config,err error){
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err =viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	
	err = viper.Unmarshal(&config)
	return

}