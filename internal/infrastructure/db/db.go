package db

import (
	"database/sql"
	"fmt"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		env.GetEnvVar("DB_USER"),
		env.GetEnvVar("DB_PASS"),
		env.GetEnvVar("DB_HOST"),
		env.GetEnvVar("DB_PORT"),
		env.GetEnvVar("DB_NAME"),
	)

	db, err := sql.Open(env.GetEnvVar("DB_DRIVER"), dsn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	return db, err
}
