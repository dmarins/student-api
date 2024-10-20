package creation

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentCreationWithPersistence struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
}

func NewStudentCreationWithPersistence(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentCreationUseCase {
	return &StudentCreationWithPersistence{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
	}
}

func (uc *StudentCreationWithPersistence) Execute(ctx context.Context, studentInput dtos.StudentInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentCreationUseCasePersistenceExecute)
	defer span.End()

	student := entities.NewStudent(studentInput.Name)

	uc.Logger.Debug(ctx, "new student", "id", student.ID)

	uc.Tracer.AddAttributes(span, tracer.StudentCreationUseCasePersistenceExecute,
		tracer.Attributes{
			"Entity": student,
		})

	err := uc.StudentRepository.Add(ctx, student)
	if err != nil {
		uc.Logger.Error(ctx, "error adding a new student", err)

		return dtos.NewInternalServerErrorResult()
	}

	uc.Logger.Debug(ctx, "student stored")

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return dtos.NewCreatedResult(output)
}
