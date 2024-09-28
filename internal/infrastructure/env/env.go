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

func getEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func ProvideAppEnv() string {
	return getEnvironmentVariable("APP_ENV")
}

func ProvideAppName() string {
	return getEnvironmentVariable("APP_NAME")
}

func ProvideAppHost() string {
	return getEnvironmentVariable("APP_HOST")
}

func ProvideAppPort() string {
	return getEnvironmentVariable("APP_PORT")
}

func ProvideAppGracefulShutdownTimeoutInSeconds() string {
	return getEnvironmentVariable("APP_GRACEFUL_SHUTDOWN_TIMEOUT")
}

func ProvideTenantHeaderName() string {
	return getEnvironmentVariable("HEADER_TENANT")
}

func ProvideCidHeaderName() string {
	return getEnvironmentVariable("HEADER_CID")
}

func ProvideRequestContextName() string {
	return getEnvironmentVariable("REQUEST_CONTEXT")
}

func ProvideRequestTimeoutInSeconds() string {
	return getEnvironmentVariable("REQUEST_TIMEOUT")
}

func ProvideDbHost() string {
	return getEnvironmentVariable("POSTGRES_HOST")
}

func ProvideDbPort() string {
	return getEnvironmentVariable("POSTGRES_PORT")
}

func ProvideDbUser() string {
	return getEnvironmentVariable("POSTGRES_USER")
}

func ProvideDbPassword() string {
	return getEnvironmentVariable("POSTGRES_PASSWORD")
}

func ProvideDbName() string {
	return getEnvironmentVariable("POSTGRES_DB")
}
