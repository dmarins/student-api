package env

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once
)

func LoadEnvironmentVariables() {
	var envFile string

	once.Do(func() {
		env := ProvideAppEnv()
		if env == "local" {
			envFile = "./.env.local"
		} else {
			envFile = "./.env"
		}

		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("failed to load environments variables, env: %s, env_file: %s", env, envFile)
		}
	})
}

func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func ProvideAppEnv() string {
	return GetEnvironmentVariable("APP_ENV")
}
