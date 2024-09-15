package di

import (
	"context"

	"github.com/dmarins/student-api/internal/adapters/handlers"
	"github.com/dmarins/student-api/internal/adapters/repositories"
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
			tracer.NewTracer,
			db.NewDatabase,
			server.NewServer,
			repositories.NewStudentRepository,
			create.NewCreateStudentUseCase,
			handlers.NewStudentHandler,
		),
		fx.Invoke(
			handlers.RegisterStudentRoutes,
		),
	)
}
