package update

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentUpdateWithValidations struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	Next              usecases.IStudentUpdateUseCase
}

func NewStudentUpdateWithValidations(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository, next usecases.IStudentUpdateUseCase) usecases.IStudentUpdateUseCase {
	return &StudentUpdateWithValidations{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
		Next:              next,
	}
}

func (uc *StudentUpdateWithValidations) Execute(ctx context.Context, studentUpdateInput dtos.StudentUpdateInput) *dtos.Result {
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
