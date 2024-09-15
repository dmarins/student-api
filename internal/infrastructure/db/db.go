package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	_ "github.com/lib/pq"
)

type (
	IDb interface{}

	Db struct{}
)

func NewDatabase(ctx context.Context) *sql.DB {
	dsn := fmt.
		Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.GetEnvironmentVariable("POSTGRES_HOST"),
			env.GetEnvironmentVariable("POSTGRES_PORT"),
			env.GetEnvironmentVariable("POSTGRES_USER"),
			env.GetEnvironmentVariable("POSTGRES_PASSWORD"),
			env.GetEnvironmentVariable("POSTGRES_DB"),
		)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal(ctx, "failed to open database", err)
		return nil
	}

	err = db.PingContext(ctx)
	if err != nil {
		logger.Fatal(ctx, "failed to verify connection to database", err)
		return nil
	}

	return db
}
