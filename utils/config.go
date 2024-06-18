package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVE"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	ServerAddres string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}