package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IStudentCreateUseCase interface {
	Execute(ctx context.Context, studentInput dtos.StudentInput) *dtos.Result
}
