package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/dmarins/student-api/internal/domain/dtos"
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

	err := r.Postgres.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM students WHERE name = $1)", name).Scan(&exists)
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

func (r *StudentRepository) Update(ctx context.Context, student *entities.Student) error {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositoryUpdate)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositoryUpdate,
		tracer.Attributes{
			"Entity": student,
		})

	_, err := r.Postgres.ExecContext(ctx, "UPDATE students SET name = $1 WHERE id = $2", student.Name, student.ID)

	return err
}

func (r *StudentRepository) Delete(ctx context.Context, studentId string) error {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositoryDelete)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositoryDelete,
		tracer.Attributes{
			"ID": studentId,
		})

	_, err := r.Postgres.ExecContext(ctx, "DELETE FROM students WHERE id = $1", studentId)

	return err
}

func (r *StudentRepository) Count(ctx context.Context, filter dtos.Filter) (int, error) {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositoryCount)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositoryCount,
		tracer.Attributes{
			"Filter": filter,
		})

	query := "SELECT COUNT(*) FROM students"
	args := make([]interface{}, 0)

	if filter.Name != nil {
		query += " WHERE name LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", *filter.Name))
	}

	var count int

	err := r.Postgres.QueryRowContext(ctx, query, args[:]...).Scan(&count)
	return count, err
}

func (r *StudentRepository) SearchBy(ctx context.Context, pagination dtos.PaginationRequest, filter dtos.Filter) ([]*entities.Student, error) {
	span, ctx := r.Tracer.NewSpanContext(ctx, tracer.StudentRepositorySearchBy)
	defer span.End()

	r.Tracer.AddAttributes(span, tracer.StudentRepositorySearchBy,
		tracer.Attributes{
			"Pagination": pagination,
			"Filter":     filter,
		})

	query := "SELECT id, name FROM students"
	args := make([]interface{}, 0)

	if filter.Name != nil {
		query += " WHERE name LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", *filter.Name))
	}

	if strings.EqualFold(pagination.SortField, dtos.FILTER_NAME) {
		query += " ORDER BY name "
	} else {
		query += " ORDER BY id "
	}

	if pagination.IsASC() {
		query += dtos.ORDER_ASC
	} else {
		query += dtos.ORDER_DESC
	}

	query += " LIMIT $2 OFFSET $3"
	args = append(args, pagination.PageSize, pagination.Offset())

	rows, err := r.Postgres.QueryContext(ctx, query, args[:]...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []*entities.Student
	for rows.Next() {
		var student entities.Student

		err := rows.Scan(&student.ID, &student.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		students = append(students, &student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
