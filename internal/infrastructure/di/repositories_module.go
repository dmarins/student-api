package di

import (
	"github.com/dmarins/student-api/internal/adapters/repositories"
	domain_repositories "github.com/dmarins/student-api/internal/domain/repositories"
	"go.uber.org/fx"
)

func repositoriesModule() fx.Option {
	return fx.Module("repositories",
		fx.Provide(
			fx.Annotate(repositories.NewHealthCheckRepository, fx.As(new(domain_repositories.IHealthCheckRepository))),
			fx.Annotate(repositories.NewStudentRepository, fx.As(new(domain_repositories.IStudentRepository))),
		),
	)
}
