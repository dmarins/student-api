package delete

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentDeleteWithPersistence struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
}

func NewStudentDeleteWithPersistence(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentDeleteUseCase {
	return &StudentDeleteWithPersistence{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
	}
}

func (uc *StudentDeleteWithPersistence) Execute(ctx context.Context, studentId string) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentDeleteUseCasePersistenceExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentDeleteUseCasePersistenceExecute,
		tracer.Attributes{
			"ID": studentId,
		})

	err := uc.StudentRepository.Delete(ctx, studentId)
	if err != nil {
		uc.Logger.Error(ctx, "error deleting a student", err)

		return dtos.NewInternalServerErrorResult()
	}

	uc.Logger.Debug(ctx, "student deleted")

	return dtos.NewNoCotentResult()
}
