package di

import (
	"github.com/dmarins/student-api/internal/adapters/handlers"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"go.uber.org/fx"
)

func provideStudentHandler(
	tracer tracer.ITracer,
	logger logger.ILogger,
	studentCreationUseCase usecases.IStudentCreationUseCase,
	studentReadingUseCase usecases.IStudentReadingUseCase) *handlers.StudentHandler {
	return handlers.NewStudentHandler(tracer, logger, studentCreationUseCase, studentReadingUseCase)
}

func studentHandlerModule() fx.Option {
	return fx.Module("studentHandlers",
		fx.Provide(fx.Annotate(provideStudentHandler, fx.ParamTags(``, ``, `name:"studentCreationWithValidations"`, ``))),
		fx.Invoke(handlers.RegisterStudentRoutes),
	)
}

func healthCheckHandlerModule() fx.Option {
	return fx.Module("healthCheckHandlers",
		fx.Provide(handlers.NewHealthCheckHandler),
		fx.Invoke(handlers.RegisterHealthCheckRoute),
	)
}
