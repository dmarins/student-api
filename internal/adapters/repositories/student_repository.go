package repositories

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/infrastructure/db"
)

type StudentRepository struct {
	postgres db.IDb
}

func NewStudentRepository(postgres db.IDb) repositories.IStudentRepository {
	return &StudentRepository{
		postgres: postgres,
	}
}

func (r *StudentRepository) Save(ctx context.Context, student *entities.Student) error {
	_, err := r.postgres.ExecContext(ctx, "INSERT INTO students (id, name) VALUES ($1, $2)", student.ID, student.Name)

	return err
}

func (r *StudentRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool

	err := r.postgres.
		QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM students WHERE name = $1)", name).
		Scan(&exists)

	return exists, err
}
