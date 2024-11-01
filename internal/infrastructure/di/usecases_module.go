package di

import (
	domain_usecases "github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/create"
	"github.com/dmarins/student-api/internal/usecases/student/read"
	"github.com/dmarins/student-api/internal/usecases/student/update"
	"go.uber.org/fx"
)

func useCasesModule() fx.Option {
	return fx.Module("usecases",
		healthCheckUseCase(),
		createStudentUseCase(),
		readingStudentUseCaseModule(),
		updateStudentUseCase(),
	)
}

func healthCheckUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(healthcheck.NewHealthCheck, fx.As(new(domain_usecases.IHealthCheckUseCase))),
	)
}

func createStudentUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(create.NewStudentCreateWithPersistence, fx.ResultTags(`name:"studentCreateWithPersistence"`),
			fx.As(new(domain_usecases.IStudentCreateUseCase))),
		fx.Annotate(create.NewStudentCreateWithValidations, fx.ParamTags(``, ``, ``, `name:"studentCreateWithPersistence"`),
			fx.ResultTags(`name:"studentCreateWithValidations"`), fx.As(new(domain_usecases.IStudentCreateUseCase)),
		),
	)
}

func readingStudentUseCaseModule() fx.Option {
	return fx.Provide(
		fx.Annotate(read.NewStudentReadWithFindById, fx.As(new(domain_usecases.IStudentReadUseCase))),
	)
}

func updateStudentUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(update.NewStudentUpdateWithPersistence, fx.ResultTags(`name:"studentUpdateWithPersistence"`),
			fx.As(new(domain_usecases.IStudentUpdateUseCase))),
		fx.Annotate(update.NewStudentUpdateWithValidations, fx.ParamTags(``, ``, ``, `name:"studentUpdateWithPersistence"`),
			fx.ResultTags(`name:"studentUpdateWithValidations"`), fx.As(new(domain_usecases.IStudentUpdateUseCase)),
		),
	)
}
