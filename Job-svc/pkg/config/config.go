package config

import "github.com/spf13/viper"

type Config struct{
	PORT string `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL"`
	User_SVC string `mapstructure:"USER_SVC"`
}

func LoadConfig() (err error){
	viper.AddConfigPath("./")
	viper.SetConfigName("job-svc_config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.BindEnv("PORT", "PORT")
	viper.BindEnv("DB_URL", "DB_URL")
	viper.BindEnv("USER_SVC", "USER_SVC")

	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	return 
}