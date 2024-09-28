package usecases

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
)

type IStudentCreationUseCase interface {
	Execute(ctx context.Context, student entities.Student) *dtos.Result
}
