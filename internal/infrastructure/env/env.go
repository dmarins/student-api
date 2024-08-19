package env

import (
	"github.com/spf13/viper"
)

func InitEnvVars() error {

	viper.SetConfigFile("./.env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
