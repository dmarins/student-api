package repositories

import (
	"context"
	"database/sql"

	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) repositories.IStudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) Save(ctx context.Context, student *entities.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (id, name) VALUES ($1, $2)", student.ID, student.Name)

	return err
}

func (r *StudentRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool

	err := r.db.
		QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM students WHERE name = $1)", name).
		Scan(&exists)

	return exists, err
}
