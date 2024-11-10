package repositories

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/entities"
)

type IStudentRepository interface {
	Add(ctx context.Context, student *entities.Student) error
	ExistsByName(ctx context.Context, name string) (bool, error)
	FindById(ctx context.Context, studentId string) (*entities.Student, error)
	Update(ctx context.Context, student *entities.Student) error
	Delete(ctx context.Context, studentId string) error
}
