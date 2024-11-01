package update

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentUpdateWithNameCheck struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	Next              usecases.IStudentUpdateUseCase
}

func NewStudentUpdateWithNameCheck(tracer tracer.ITracer,
	logger logger.ILogger,
	studentRepository repositories.IStudentRepository,
	next usecases.IStudentUpdateUseCase) usecases.IStudentUpdateUseCase {
	return &StudentUpdateWithNameCheck{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
		Next:              next,
	}
}

func (uc *StudentUpdateWithNameCheck) Execute(ctx context.Context, studentUpdateInput dtos.StudentUpdateInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentUpdateUseCaseValidationsExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentUpdateUseCaseValidationsExecute,
		tracer.Attributes{
			"Payload": studentUpdateInput,
		})

	exists, err := uc.StudentRepository.ExistsByName(ctx, studentUpdateInput.Name)
	if err != nil {
		uc.Logger.Error(ctx, "error checking if student exists", err, "name", studentUpdateInput.Name)

		return dtos.NewInternalServerErrorResult()
	}

	if exists {
		uc.Logger.Warn(ctx, "there is already a student with the same name", "name", studentUpdateInput.Name)

		return dtos.NewConflictResult()
	}

	return uc.Next.Execute(ctx, studentUpdateInput)
}
