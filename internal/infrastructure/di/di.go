package di

import (
	"context"

	"go.uber.org/fx"
)

func StartCompositionRoot(options ...fx.Option) *fx.App {
	baseOptions := []fx.Option{
		fx.Provide(
			func() context.Context { return context.Background() },
		),
		registerHooks(),
		infrastructureModule(),
		repositoriesModule(),
		healthCheckUseCaseModule(),
		createStudentUseCaseModule(),
		readingStudentUseCaseModule(),
		healthCheckHandlerModule(),
		studentHandlerModule(),
	}

	allOptions := append(baseOptions, options...)
	return fx.New(allOptions...)
}
