package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IStudentDeleteUseCase interface {
	Execute(ctx context.Context, studentId string) *dtos.Result
}
