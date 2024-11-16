package read

import (
	"context"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type StudentSearchWithSearchBy struct {
	Tracer            tracer.ITracer
	Logger            logger.ILogger
	StudentRepository repositories.IStudentRepository
}

func NewStudentSearchWithSearchBy(tracer tracer.ITracer, logger logger.ILogger, studentRepository repositories.IStudentRepository) usecases.IStudentSearchUseCase {
	return &StudentSearchWithSearchBy{
		Tracer:            tracer,
		Logger:            logger,
		StudentRepository: studentRepository,
	}
}

func (uc *StudentSearchWithSearchBy) Execute(ctx context.Context, pagination dtos.PaginationRequest, filter dtos.Filter) *dtos.Result {
	span, ctx := uc.Tracer.NewSpanContext(ctx, tracer.StudentSearchUseCaseSearchByExecute)
	defer span.End()

	uc.Tracer.AddAttributes(span, tracer.StudentSearchUseCaseSearchByExecute,
		tracer.Attributes{
			"Pagination": pagination,
			"Filter":     filter,
		})

	count, err := uc.StudentRepository.Count(ctx, filter)
	if err != nil {
		uc.Logger.Error(ctx, "error counting students", err)

		return dtos.NewInternalServerErrorResult()
	}

	if count <= 0 {
		paginationResponse := dtos.NewPaginationResponse(0, pagination.Page, pagination.PageSize, nil)
		return dtos.NewOkResult(paginationResponse)
	}

	students, err := uc.StudentRepository.SearchBy(ctx, pagination, filter)
	if err != nil {
		uc.Logger.Error(ctx, "error searching students", err)

		return dtos.NewInternalServerErrorResult()
	}

	paginationResponse := dtos.NewPaginationResponse(count, pagination.Page, pagination.PageSize, students)
	return dtos.NewOkResult(paginationResponse)
}
