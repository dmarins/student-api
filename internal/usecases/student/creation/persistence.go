package creation

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/google/uuid"
)

type StudentCreationUseCasePersistence struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
}

func NewStudentCreationUseCasePersistence(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentCreationUseCase {
	return &StudentCreationUseCasePersistence{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
	}
}

func (uc *StudentCreationUseCasePersistence) Execute(ctx context.Context, student entities.Student) (*dtos.StudentOutput, *dtos.Result) {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentCreationUseCasePersistenceExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentCreationUseCasePersistenceExecute,
		tracer.Attributes{
			"Entity": student,
		})

	student.ID = uuid.New().String()

	uc.Logger.Debug(ctx, "new student", "id", student.ID, "name", student.Name)

	err := uc.StudentRepository.Add(ctx, &student)
	if err != nil {
		uc.Logger.Error(ctx, "error adding a new student", err)
		return nil, dtos.NewHttpStatusInternalServerErrorResult(err)
	}

	uc.Logger.Debug(ctx, "stored")

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return output, dtos.NewHttpStatusCreatedResult(output)
}
