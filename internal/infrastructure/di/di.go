package di

import (
	"context"

	"github.com/dmarins/student-api/internal/adapters/handlers"
	"github.com/dmarins/student-api/internal/adapters/repositories"
	domain_repositories "github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/dmarins/student-api/internal/usecases/student/create"
	"go.uber.org/fx"
)

func StartCompositionRoot(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(
			func() context.Context { return ctx },
		),
		infrastructureModule(),
		createStudentUseCaseModule(),
		studentHandlerModule(),
	)
}

func infrastructureModule() fx.Option {
	return fx.Module("infrastructure",
		fx.Provide(
			fx.Annotate(tracer.NewTracer, fx.As(new(tracer.ITracer))),
			fx.Annotate(db.NewDatabase, fx.As(new(db.IDb))),
			fx.Annotate(server.NewServer, fx.As(new(server.IServer))),
		),
	)
}

func createStudentUseCaseModule() fx.Option {
	return fx.Module("createStudentUseCase",
		fx.Provide(
			fx.Annotate(repositories.NewStudentRepository, fx.As(new(domain_repositories.IStudentRepository))),
			fx.Annotate(create.NewCreateStudentUseCase, fx.As(new(usecases.ICreateStudentUseCase))),
		),
	)
}

func studentHandlerModule() fx.Option {
	return fx.Module("studentHandlers",
		fx.Provide(handlers.NewStudentHandler),
		fx.Invoke(handlers.RegisterStudentRoutes),
	)
}
