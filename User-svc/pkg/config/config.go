package config

import "github.com/spf13/viper"

type Config struct{
	DB_URL string `mapstructure:"DB_URL"`
	PORT string `mapstructure:"PORT"`
	ATokenSecret string `mapstructure:"ATokenSecret"`
	OtpEmail string `mapstructure:"OTP_Email"`
	Email_Password string `mapstructure:"Email_Password"`
	DB_Path string `mapstructure:"DB_Path"`
}

func LoadConfig() (err error){
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	
	err = viper.ReadInConfig()

	if err != nil{
		return
	}
	return
}