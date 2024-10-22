package di

import (
	"github.com/dmarins/student-api/internal/adapters/handlers"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"go.uber.org/fx"
)

func handlersModule() fx.Option {
	return fx.Module("handlers",
		healthCheckHandler(),
		studentHandler(),
		fx.Invoke(
			handlers.RegisterHealthCheckRoute,
			handlers.RegisterStudentRoutes,
		),
	)
}

func healthCheckHandler() fx.Option {
	return fx.Provide(
		handlers.NewHealthCheckHandler,
	)
}

func provideStudentHandler(
	tracer tracer.ITracer,
	logger logger.ILogger,
	studentCreationUseCase usecases.IStudentCreationUseCase,
	studentReadingUseCase usecases.IStudentReadingUseCase) *handlers.StudentHandler {
	return handlers.NewStudentHandler(tracer, logger, studentCreationUseCase, studentReadingUseCase)
}

func studentHandler() fx.Option {
	return fx.Provide(
		fx.Annotate(provideStudentHandler, fx.ParamTags(``, ``, `name:"studentCreationWithValidations"`, ``)),
	)
}
