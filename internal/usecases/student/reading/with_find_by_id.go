package reading

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentReadingWithFindById struct {
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	StudentRepository repositories.IStudentRepository
}

func NewStudentReadingWithFindById(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentReadingUseCase {
	return &StudentReadingWithFindById{
		Tracer:            tracer,
		Logger:            logger,
		StudentRepository: studentRepository,
	}
}

func (uc *StudentReadingWithFindById) Execute(ctx context.Context, studentId string) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentReadingUseCaseFindByIdExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentReadingUseCaseFindByIdExecute,
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

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return dtos.NewOkResult(output)
}
