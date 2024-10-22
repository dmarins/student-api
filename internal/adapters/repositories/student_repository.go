package repositories

import (
	"context"
	"database/sql"

	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentRepository struct {
	Postgres db.IDb
	Tracer   tracer.ITracer
}

func NewStudentRepository(tracer tracer.ITracer,
	postgres db.IDb) repositories.IStudentRepository {
	return &StudentRepository{
		Postgres: postgres,
		Tracer:   tracer,
	}
}

func (r *StudentRepository) Add(ctx context.Context, student *entities.Student) error {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositoryAdd)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositoryAdd,
		tracer.Attributes{
			"Entity": student,
		})

	_, err := r.Postgres.ExecContext(ctx, "INSERT INTO students (id, name) VALUES ($1, $2)", student.ID, student.Name)

	return err
}

func (r *StudentRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositoryExistsByName)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositoryExistsByName,
		tracer.Attributes{
			"Name": name,
		})

	var exists bool

	err := r.Postgres.
		QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM students WHERE name = $1)", name).
		Scan(&exists)

	return exists, err
}

func (r *StudentRepository) FindById(ctx context.Context, studentId string) (*entities.Student, error) {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositoryFindById)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositoryFindById,
		tracer.Attributes{
			"ID": studentId,
		})

	row := r.Postgres.QueryRowContext(ctx, "SELECT Id, Name FROM students WHERE Id = $1", studentId)

	var student entities.Student
	err := row.Scan(&student.ID, &student.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &student, nil
}
