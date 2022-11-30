package util

import "github.com/spf13/viper"

type Config struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile("app.env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
