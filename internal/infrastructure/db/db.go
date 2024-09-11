package db

import (
	"database/sql"
	"fmt"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	_ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.GetEnvVar("POSTGRES_HOST"),
		env.GetEnvVar("POSTGRES_PORT"),
		env.GetEnvVar("POSTGRES_USER"),
		env.GetEnvVar("POSTGRES_PASSWORD"),
		env.GetEnvVar("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}
