package repositories

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
)

type IStudentRepository interface {
	Add(ctx context.Context, student *entities.Student) error
	ExistsByName(ctx context.Context, name string) (bool, error)
	FindById(ctx context.Context, studentId string) (*entities.Student, error)
	Update(ctx context.Context, student *entities.Student) error
	Delete(ctx context.Context, studentId string) error
	SearchBy(ctx context.Context, pagination dtos.PaginationInput, filter dtos.Filter) ([]*entities.Student, error)
	Count(ctx context.Context, filter dtos.Filter) (int, error)
}
