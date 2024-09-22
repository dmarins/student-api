package create

import (
	"context"
	"errors"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/google/uuid"
)

type CreateStudentUseCase struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
}

func NewCreateStudentUseCase(tracer tracer.ITracer,
	studentRepository repositories.IStudentRepository) usecases.ICreateStudentUseCase {
	return &CreateStudentUseCase{
		StudentRepository: studentRepository,
		Tracer:            tracer,
	}
}

func (uc *CreateStudentUseCase) Execute(ctx context.Context, student entities.Student) (*dtos.StudentOutput, error) {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.CreateStudentUseCaseExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.CreateStudentUseCaseExecute,
		tracer.Attributes{
			"Entity": student,
		})

	exists, err := uc.StudentRepository.ExistsByName(ctx, student.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("student already exists")
	}

	student.ID = uuid.New().String()

	err = uc.StudentRepository.Add(ctx, &student)
	if err != nil {
		return nil, err
	}

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return output, nil
}
