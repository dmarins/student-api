package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type IStudentReadingUseCase interface {
	Execute(ctx context.Context, studentId string) *dtos.Result
}
