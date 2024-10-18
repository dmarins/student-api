package tests

import (
	"context"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/infrastructure/di"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/gavv/httpexpect/v2"
	"go.uber.org/fx"
)

type E2eTestsBuilder struct {
	Ctx        context.Context
	App        *fx.App
	AppServer  server.IServer
	TestServer *httptest.Server
}

func NewE2eTestsBuilder() *E2eTestsBuilder {
	os.Setenv("APP_ENV", "test")
	env.LoadEnvironmentVariables()

	return &E2eTestsBuilder{
		Ctx: context.Background(),
	}
}

func (b *E2eTestsBuilder) StartCompositionRoot() *E2eTestsBuilder {
	app := di.StartCompositionRoot(fx.Populate(&b.AppServer))

	if err := app.Start(b.Ctx); err != nil {
		log.Fatalf("failed to initialize FX: %v", err)
	}

	b.App = app

	return b
}

func (b *E2eTestsBuilder) StartTestServer() *E2eTestsBuilder {
	b.TestServer = httptest.NewServer(b.AppServer.GetEcho())

	return b
}

func (b *E2eTestsBuilder) Build(t *testing.T) *httpexpect.Expect {
	return httpexpect.Default(t, b.TestServer.URL)
}
