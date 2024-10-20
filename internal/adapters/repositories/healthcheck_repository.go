package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type HealthCheckRepository struct {
	Postgres db.IDb
	Tracer   tracer.ITracer
}

func NewHealthCheckRepository(tracer tracer.ITracer,
	postgres db.IDb) repositories.IHealthCheckRepository {
	return &HealthCheckRepository{
		Postgres: postgres,
		Tracer:   tracer,
	}
}

func (r *HealthCheckRepository) CheckDbConnection(ctx context.Context) error {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.HealthCheckRepositoryCheckDbConnection)
	defer span.End()

	var result int
	row := r.Postgres.QueryRowContext(ctx, "SELECT 1")

	if err := row.Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no rows returned in health check query: %w", err)
		}

		return err
	}

	return nil
}
