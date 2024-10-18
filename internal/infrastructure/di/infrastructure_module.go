package di

import (
	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"go.uber.org/fx"
)

func infrastructureModule() fx.Option {
	return fx.Module("infrastructure",
		fx.Provide(
			fx.Annotate(logger.NewLogger, fx.As(new(logger.ILogger))),
			fx.Annotate(tracer.NewTracer, fx.As(new(tracer.ITracer))),
			fx.Annotate(db.NewDatabase, fx.As(new(db.IDb))),
			fx.Annotate(server.NewServer, fx.As(new(server.IServer))),
		),
	)
}
