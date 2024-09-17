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
	IDb interface {
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	}

	Db struct {
		postgres *sql.DB
	}
)

func NewDatabase(ctx context.Context) IDb {
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

	return &Db{
		postgres: db,
	}
}

func (d *Db) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.postgres.ExecContext(ctx, query, args...)
}

func (d *Db) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return d.postgres.QueryRowContext(ctx, query, args...)
}
