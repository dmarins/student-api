package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IStudentCreationUseCase interface {
	Execute(ctx context.Context, studentInput dtos.StudentInput) *dtos.Result
}
