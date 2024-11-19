package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IStudentReadUseCase interface {
	Execute(ctx context.Context, studentId string) *dtos.Result
}

type IStudentSearchUseCase interface {
	Execute(ctx context.Context, pagination dtos.PaginationInput, filter dtos.Filter) *dtos.Result
}
