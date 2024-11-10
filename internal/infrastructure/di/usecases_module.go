package di

import (
	domain_usecases "github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/create"
	"github.com/dmarins/student-api/internal/usecases/student/delete"
	"github.com/dmarins/student-api/internal/usecases/student/read"
	"github.com/dmarins/student-api/internal/usecases/student/update"
	"go.uber.org/fx"
)

func useCasesModule() fx.Option {
	return fx.Module("usecases",
		healthCheckUseCase(),
		createStudentUseCase(),
		readStudentUseCaseModule(),
		updateStudentUseCase(),
		deleteStudentUseCase(),
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
		fx.Annotate(create.NewStudentCreateWithNameCheck, fx.ParamTags(``, ``, ``, `name:"studentCreateWithPersistence"`),
			fx.ResultTags(`name:"studentCreateWithNameCheck"`), fx.As(new(domain_usecases.IStudentCreateUseCase)),
		),
	)
}

func readStudentUseCaseModule() fx.Option {
	return fx.Provide(
		fx.Annotate(read.NewStudentReadWithFindById, fx.As(new(domain_usecases.IStudentReadUseCase))),
	)
}

func updateStudentUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(update.NewStudentUpdateWithPersistence, fx.ResultTags(`name:"studentUpdateWithPersistence"`),
			fx.As(new(domain_usecases.IStudentUpdateUseCase))),
		fx.Annotate(update.NewStudentUpdateWithNameCheck, fx.ParamTags(``, ``, ``, `name:"studentUpdateWithPersistence"`),
			fx.ResultTags(`name:"studentUpdateWithNameCheck"`), fx.As(new(domain_usecases.IStudentUpdateUseCase))),
		fx.Annotate(update.NewStudentUpdateWithFindById, fx.ParamTags(``, ``, ``, `name:"studentUpdateWithNameCheck"`),
			fx.ResultTags(`name:"studentUpdateWithFindById"`), fx.As(new(domain_usecases.IStudentUpdateUseCase)),
		),
	)
}

func deleteStudentUseCase() fx.Option {
	return fx.Provide(
		fx.Annotate(delete.NewStudentDeleteWithPersistence, fx.ResultTags(`name:"studentDeleteWithPersistence"`),
			fx.As(new(domain_usecases.IStudentDeleteUseCase))),
		fx.Annotate(delete.NewStudentDeleteWithFindById, fx.ParamTags(``, ``, ``, `name:"studentDeleteWithPersistence"`),
			fx.ResultTags(`name:"studentDeleteWithFindById"`), fx.As(new(domain_usecases.IStudentDeleteUseCase)),
		),
	)
}
