package create

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentCreateWithPersistence struct {
	StudentRepository repositories.IStudentRepository
	Tracer            tracer.ITracer
	Logger            logger.ILogger
}

func NewStudentCreateWithPersistence(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentCreateUseCase {
	return &StudentCreateWithPersistence{
		StudentRepository: studentRepository,
		Tracer:            tracer,
		Logger:            logger,
	}
}

func (uc *StudentCreateWithPersistence) Execute(ctx context.Context, studentCreateInput dtos.StudentCreateInput) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentCreateUseCasePersistenceExecute)
	defer span.End()

	student := entities.NewStudent(studentCreateInput.Name)

	uc.Logger.Debug(ctx, "new student", "id", student.ID)

	uc.Tracer.AddAttributes(span, tracer.StudentCreateUseCasePersistenceExecute,
		tracer.Attributes{
			"Entity": student,
		})

	err := uc.StudentRepository.Add(ctx, student)
	if err != nil {
		uc.Logger.Error(ctx, "error adding a new student", err)

		return dtos.NewInternalServerErrorResult()
	}

	uc.Logger.Debug(ctx, "student stored")

	output := &dtos.StudentOutput{
		ID:   student.ID,
		Name: student.Name,
	}

	return dtos.NewCreatedResult(output)
}
