package di

import (
	"github.com/dmarins/student-api/internal/adapters/repositories"
	domain_repositories "github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/creation"
	"go.uber.org/fx"
)

func healthCheckUseCaseModule() fx.Option {
	return fx.Module("healthCheckUseCase",
		fx.Provide(
			fx.Annotate(repositories.NewHealthCheckRepository, fx.As(new(domain_repositories.IHealthCheckRepository))),
			fx.Annotate(healthcheck.NewHealthCheck, fx.As(new(usecases.IHealthCheckUseCase))),
		),
	)
}

func createStudentUseCaseModule() fx.Option {
	return fx.Module("createStudentUseCase",
		fx.Provide(
			fx.Annotate(repositories.NewStudentRepository, fx.As(new(domain_repositories.IStudentRepository))),
			fx.Annotate(creation.NewStudentCreationWithPersistence,
				fx.ResultTags(`name:"studentCreationWithPersistence"`), fx.As(new(usecases.IStudentCreationUseCase))),
			fx.Annotate(creation.NewStudentCreationWithValidations, fx.ParamTags(``, ``, ``, `name:"studentCreationWithPersistence"`),
				fx.ResultTags(`name:"studentCreationWithValidations"`), fx.As(new(usecases.IStudentCreationUseCase)),
			),
		),
	)
}
