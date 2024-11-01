package update

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentUpdateWithFindById struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	Next              usecases.IStudentUpdateUseCase
}

func NewStudentUpdateWithFindById(tracer tracer.ITracer,
	logger logger.ILogger,
	studentRepository repositories.IStudentRepository,
	next usecases.IStudentUpdateUseCase) usecases.IStudentUpdateUseCase {
	return &StudentUpdateWithFindById{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
		Next:              next,
	}
}

func (uc *StudentUpdateWithFindById) Execute(ctx context.Context, studentUpdateInput dtos.StudentUpdateInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentUpdateUseCaseValidationsExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentUpdateUseCaseValidationsExecute,
		tracer.Attributes{
			"Payload": studentUpdateInput,
		})

	student, err := uc.StudentRepository.FindById(ctx, studentUpdateInput.ID)
	if err != nil {
		uc.Logger.Error(ctx, "error finding student by id", err, "id", studentUpdateInput.ID)

		return dtos.NewInternalServerErrorResult()
	}

	if student == nil {
		uc.Logger.Warn(ctx, "the student was not found according to the id provided", "id", studentUpdateInput.ID)

		return dtos.NewNotFoundResult()
	}

	return uc.Next.Execute(ctx, studentUpdateInput)
}
