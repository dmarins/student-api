package create

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentCreateWithValidations struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	Next              usecases.IStudentCreateUseCase
}

func NewStudentCreateWithValidations(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository, next usecases.IStudentCreateUseCase) usecases.IStudentCreateUseCase {
	return &StudentCreateWithValidations{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
		Next:              next,
	}
}

func (uc *StudentCreateWithValidations) Execute(ctx context.Context, studentInput dtos.StudentInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentCreateUseCaseValidationsExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentCreateUseCaseValidationsExecute,
		tracer.Attributes{
			"Payload": studentInput,
		})

	exists, err := uc.StudentRepository.ExistsByName(ctx, studentInput.Name)
	if err != nil {
		uc.Logger.Error(ctx, "error checking if student exists", err, "name", studentInput.Name)

		return dtos.NewInternalServerErrorResult()
	}

	if exists {
		uc.Logger.Warn(ctx, "there is already a student with the same name", "name", studentInput.Name)

		return dtos.NewConflictResult()
	}

	return uc.Next.Execute(ctx, studentInput)
}