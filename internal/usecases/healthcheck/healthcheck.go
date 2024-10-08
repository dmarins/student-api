package healthcheck

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type HealthCheck struct {
	HealthCheckRepository repositories.IHealthCheckRepository
	Tracer                tracer.ITracer
	Logger                logger.ILogger
}

func NewHealthCheck(tracer tracer.ITracer, logger logger.ILogger, healthCheckRepository repositories.IHealthCheckRepository) usecases.IHealthCheckUseCase {
	return &HealthCheck{
		HealthCheckRepository: healthCheckRepository,
		Tracer:                tracer,
		Logger:                logger,
	}
}

func (uc *HealthCheck) Execute(ctx context.Context) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.HealthCheckExecute)
	defer span.End()

	err := uc.HealthCheckRepository.CheckDbConnection(ctx)
	if err != nil {
		uc.Logger.Error(ctx, "error checking db connection", err)

		return dtos.NewInternalServerErrorResult()
	}

	result := dtos.NewOkResult(nil)
	result.Message = "healthy"

	return result
}
