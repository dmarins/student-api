package di

import (
	"context"

	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"go.uber.org/fx"
)

func registerHooks() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, tracer tracer.ITracer, logger logger.ILogger, db db.IDb, server server.IServer) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.ListenAndServe(ctx, logger)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Sync(ctx)
				tracer.Shutdown(ctx, logger)
				db.Close(ctx, logger)
				return server.GracefulShutdownServer(ctx, logger)
			},
		})
	})
}
