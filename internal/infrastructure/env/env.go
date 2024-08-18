package env

import (
	"github.com/spf13/viper"
)

func InitEnvVars() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AddConfigPath("./cmd")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
