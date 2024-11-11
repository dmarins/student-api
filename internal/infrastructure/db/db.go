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
		Close(ctx context.Context, logger logger.ILogger)
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	}

	Db struct {
		postgres *sql.DB
	}
)

func NewDatabase(ctx context.Context, logger logger.ILogger) IDb {
	host := env.ProvideDbHost()
	port := env.ProvideDbPort()

	dsn := fmt.
		Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			env.ProvideDbUser(),
			env.ProvideDbPassword(),
			env.ProvideDbName(),
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

	logger.Info(ctx, "db connected", "address", fmt.Sprintf("%s:%s", host, port))

	return &Db{
		postgres: db,
	}
}

func NewIntegrationTestDatabase(db *sql.DB) IDb {
	return &Db{
		postgres: db,
	}
}

func (d *Db) Close(ctx context.Context, logger logger.ILogger) {
	if d.postgres != nil {
		err := d.postgres.Close()
		if err != nil {
			logger.Error(ctx, "failed to close database connection", err)
			return
		}
	}

	logger.Info(ctx, "database connection closed successfully")
}

func (d *Db) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.postgres.ExecContext(ctx, query, args...)
}

func (d *Db) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return d.postgres.QueryRowContext(ctx, query, args...)
}

func (d *Db) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.postgres.QueryContext(ctx, query, args...)
}
