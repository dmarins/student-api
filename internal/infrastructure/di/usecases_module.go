package di

import (
	domain_usecases "github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/creation"
	"github.com/dmarins/student-api/internal/usecases/student/reading"
	"go.uber.org/fx"
)

func useCasesModule() fx.Option {
	return fx.Module("usecases",
		healthCheckUseCase(),
		createStudentUseCase(),
		readingStudentUseCaseModule(),
	)
}

func healthCheckUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(healthcheck.NewHealthCheck, fx.As(new(domain_usecases.IHealthCheckUseCase))),
	)
}

func createStudentUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(creation.NewStudentCreationWithPersistence, fx.ResultTags(`name:"studentCreationWithPersistence"`),
			fx.As(new(domain_usecases.IStudentCreationUseCase))),
		fx.Annotate(creation.NewStudentCreationWithValidations, fx.ParamTags(``, ``, ``, `name:"studentCreationWithPersistence"`),
			fx.ResultTags(`name:"studentCreationWithValidations"`), fx.As(new(domain_usecases.IStudentCreationUseCase)),
		),
	)
}

func readingStudentUseCaseModule() fx.Option {
	return fx.Provide(
		fx.Annotate(reading.NewStudentReadingWithFindById, fx.As(new(domain_usecases.IStudentReadingUseCase))),
	)
}
