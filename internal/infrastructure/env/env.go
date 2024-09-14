package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnvVars() error {

	var envFile string

	env := GetEnvVar("APP_ENV")
	if env == "local" {
		envFile = "./.env.local"
	} else {
		envFile = "./.env"
	}

	if err := godotenv.Load(envFile); err != nil {
		return err
	}

	log.Printf("carregando as variáveis de ambiente com o arquivo %s", envFile)

	return nil
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
