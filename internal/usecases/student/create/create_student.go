package create

import (
	"context"
	"errors"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/google/uuid"
)

type CreateStudentUseCase struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
}

func NewCreateStudentUseCase(tracer tracer.ITracer,
	logger logger.ILogger,
	studentRepository repositories.IStudentRepository) usecases.ICreateStudentUseCase {
	return &CreateStudentUseCase{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
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
		uc.Logger.Error(ctx, "error checking if student exists", err, "name", student.Name)
		return nil, err
	}
	if exists {
		uc.Logger.Warn(ctx, "there is already a student with the same name", "name", student.Name)
		return nil, errors.New("student already exists")
	}

	student.ID = uuid.New().String()

	uc.Logger.Debug(ctx, "new student", "id", student.ID, "name", student.Name)

	err = uc.StudentRepository.Add(ctx, &student)
	if err != nil {
		uc.Logger.Error(ctx, "error adding a new student", err)
		return nil, err
	}

	uc.Logger.Debug(ctx, "stored")

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return output, nil
}
