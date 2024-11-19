package create

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentCreateWithNameCheck struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	Next              usecases.IStudentCreateUseCase
}

func NewStudentCreateWithNameCheck(tracer tracer.ITracer,
	logger logger.ILogger,
	studentRepository repositories.IStudentRepository,
	next usecases.IStudentCreateUseCase) usecases.IStudentCreateUseCase {
	return &StudentCreateWithNameCheck{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
		Next:              next,
	}
}

func (uc *StudentCreateWithNameCheck) Execute(ctx context.Context, studentCreateInput dtos.StudentCreateInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentCreateUseCaseValidationsExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentCreateUseCaseValidationsExecute,
		tracer.Attributes{
			"Payload": studentCreateInput,
		})

	exists, err := uc.StudentRepository.ExistsByName(ctx, studentCreateInput.Name)
	if err != nil {
		uc.Logger.Error(ctx, "error checking if student exists", err, "name", studentCreateInput.Name)

		return dtos.NewInternalServerErrorResult()
	}

	if exists {
		uc.Logger.Warn(ctx, "there is already a student with the same name", "name", studentCreateInput.Name)

		return dtos.NewConflictResult()
	}

	return uc.Next.Execute(ctx, studentCreateInput)
}
