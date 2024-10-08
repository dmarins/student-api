package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IHealthCheckUseCase interface {
	Execute(ctx context.Context) *dtos.Result
}
