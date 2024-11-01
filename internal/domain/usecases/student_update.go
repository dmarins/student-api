package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IStudentUpdateUseCase interface {
	Execute(ctx context.Context, studentUpdateInput dtos.StudentUpdateInput) *dtos.Result
}
