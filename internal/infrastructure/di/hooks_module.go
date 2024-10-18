package di

import (
	"context"

	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"go.uber.org/fx"
)

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
