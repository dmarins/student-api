package di

import (
	domain_usecases "github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/creation"
	"github.com/dmarins/student-api/internal/usecases/student/reading"
	"go.uber.org/fx"
)

func healthCheckUseCaseModule() fx.Option {
	return fx.Module("healthCheckUseCase",
		fx.Provide(
			fx.Annotate(healthcheck.NewHealthCheck, fx.As(new(domain_usecases.IHealthCheckUseCase))),
		),
	)
}

func createStudentUseCaseModule() fx.Option {
	return fx.Module("createStudentUseCase",
		fx.Provide(
			fx.Annotate(creation.NewStudentCreationWithPersistence, fx.ResultTags(`name:"studentCreationWithPersistence"`),
				fx.As(new(domain_usecases.IStudentCreationUseCase))),
			fx.Annotate(creation.NewStudentCreationWithValidations, fx.ParamTags(``, ``, ``, `name:"studentCreationWithPersistence"`),
				fx.ResultTags(`name:"studentCreationWithValidations"`), fx.As(new(domain_usecases.IStudentCreationUseCase)),
			),
		),
	)
}

func readingStudentUseCaseModule() fx.Option {
	return fx.Module("readingStudentUseCase",
		fx.Provide(
			fx.Annotate(reading.NewStudentReadingWithFindById, fx.As(new(domain_usecases.IStudentReadingUseCase))),
		),
	)
}
