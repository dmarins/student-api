package delete

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentDeleteWithFindById struct {
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	StudentRepository repositories.IStudentRepository
	Next              usecases.IStudentDeleteUseCase
}

func NewStudentDeleteWithFindById(tracer tracer.ITracer,
	logger logger.ILogger,
	studentRepository repositories.IStudentRepository,
	next usecases.IStudentDeleteUseCase) usecases.IStudentDeleteUseCase {
	return &StudentDeleteWithFindById{
		Tracer:            tracer,
		Logger:            logger,
		StudentRepository: studentRepository,
		Next:              next,
	}
}

func (uc *StudentDeleteWithFindById) Execute(ctx context.Context, studentId string) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentDeleteUseCaseFindByIdExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentDeleteUseCaseFindByIdExecute,
		tracer.Attributes{
			"ID": studentId,
		})

	student, err := uc.StudentRepository.FindById(ctx, studentId)
	if err != nil {
		uc.Logger.Error(ctx, "error finding student by id", err, "id", studentId)

		return dtos.NewInternalServerErrorResult()
	}

	if student == nil {
		uc.Logger.Warn(ctx, "the student was not found according to the id provided", "id", studentId)

		return dtos.NewNotFoundResult()
	}

	return uc.Next.Execute(ctx, studentId)
}
