package update

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentUpdateWithPersistence struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
}

func NewStudentUpdateWithPersistence(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentUpdateUseCase {
	return &StudentUpdateWithPersistence{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
	}
}

func (uc *StudentUpdateWithPersistence) Execute(ctx context.Context, studentUpdateInput dtos.StudentUpdateInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentUpdateUseCasePersistenceExecute)
	defer span.End()

	student := &entities.Student{
		ID:   studentUpdateInput.ID,
		Name: studentUpdateInput.Name,
	}

	uc.Tracer.AddAttributes(span, tracer.StudentUpdateUseCasePersistenceExecute,
		tracer.Attributes{
			"Entity": student,
		})

	err := uc.StudentRepository.Update(ctx, student)
	if err != nil {
		uc.Logger.Error(ctx, "error updating a student", err)

		return dtos.NewInternalServerErrorResult()
	}

	uc.Logger.Debug(ctx, "student updated")

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return dtos.NewOkResult(output)
}
