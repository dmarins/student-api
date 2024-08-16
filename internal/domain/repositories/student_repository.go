package repositories

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/entities"
)

type IStudentRepository interface {
	Save(ctx context.Context, student *entities.Student) error
	ExistsByName(ctx context.Context, name string) (bool, error)
}
