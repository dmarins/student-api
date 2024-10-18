package di

import (
	"context"

	"github.com/dmarins/student-api/internal/adapters/handlers"
	"github.com/dmarins/student-api/internal/adapters/repositories"
	domain_repositories "github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/creation"
	"go.uber.org/fx"
)

func StartCompositionRoot(ctx context.Context, options ...fx.Option) *fx.App {
	baseOptions := []fx.Option{
		fx.Provide(
			func() context.Context { return ctx },
			fx.Annotate(
				func() string { return env.ProvideAppName() },
				fx.ResultTags(`name:"appName"`),
			),
			fx.Annotate(
				func() string { return env.ProvideAppEnv() },
				fx.ResultTags(`name:"appEnv"`),
			),
		),
		infrastructureModule(),
		healthCheckUseCaseModule(),
		createStudentUseCaseModule(),
		healthCheckHandlerModule(),
		studentHandlerModule(),
		registerHooks(),
	}

	allOptions := append(baseOptions, options...)

	return fx.New(allOptions...)
}

func registerHooks() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, server server.IServer, tracer tracer.ITracer, logger logger.ILogger) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.ListenAndServe(ctx, logger)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Sync(ctx)
				tracer.Shutdown(ctx, logger)
				return server.GracefulShutdownServer(ctx, logger)
			},
		})
	})
}

func provideTracer(ctx context.Context, logger logger.ILogger, appName string, env string) tracer.ITracer {
	return tracer.NewTracer(ctx, logger, appName, env)
}

func infrastructureModule() fx.Option {
	return fx.Module("infrastructure",
		fx.Provide(
			fx.Annotate(logger.NewLogger, fx.As(new(logger.ILogger))),
			fx.Annotate(provideTracer, fx.ParamTags(``, ``, `name:"appName"`, `name:"appEnv"`)),
			fx.Annotate(db.NewDatabase, fx.As(new(db.IDb))),
			fx.Annotate(server.NewServer, fx.As(new(server.IServer))),
		),
	)
}

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
				fx.ResultTags(`name:"studentCreationWithPersistence"`),
				fx.As(new(usecases.IStudentCreationUseCase)),
			),
			fx.Annotate(creation.NewStudentCreationWithValidations,
				fx.ParamTags(``, ``, ``, `name:"studentCreationWithPersistence"`),
				fx.ResultTags(`name:"studentCreationWithValidations"`),
				fx.As(new(usecases.IStudentCreationUseCase)),
			),
		),
	)
}

func provideStudentHandler(tracer tracer.ITracer, logger logger.ILogger, studentCreationUseCase usecases.IStudentCreationUseCase) *handlers.StudentHandler {
	return handlers.NewStudentHandler(tracer, logger, studentCreationUseCase)
}

func studentHandlerModule() fx.Option {
	return fx.Module("studentHandlers",
		fx.Provide(fx.Annotate(provideStudentHandler, fx.ParamTags(``, ``, `name:"studentCreationWithValidations"`))),
		fx.Invoke(handlers.RegisterStudentRoutes),
	)
}

func healthCheckHandlerModule() fx.Option {
	return fx.Module("healthCheckHandlers",
		fx.Provide(handlers.NewHealthCheckHandler),
		fx.Invoke(handlers.RegisterHealthCheckRoute),
	)
}
