package tests

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"time"

	"github.com/dmarins/student-api/internal/adapters/repositories"
	domain_repositories "github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationTestsBuilder struct {
	ctx         context.Context
	pgContainer *postgres.PostgresContainer
	dbConn      *sql.DB
	postgresDb  db.IDb
	logger      logger.ILogger
	tracer      tracer.ITracer
}

func NewIntegrationTestsBuilder() *IntegrationTestsBuilder {
	ctx := context.Background()

	pgContainer, err := postgres.Run(
		ctx,
		"docker.io/postgres:16.4-alpine3.20",
		postgres.WithDatabase("students"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		postgres.WithInitScripts(filepath.Join("../../../migrations", "000001_create_students_table.up.sql")),
		postgres.WithInitScripts(filepath.Join("../../../migrations", "000002_insert_students.up.sql")),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		log.Fatalf("Failed to start postgres container: %s", err)
	}

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable", "application_name=student-api-integration-tests")
	if err != nil {
		log.Fatalf("Failed to get postgres container port: %s", err)
	}

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to postgres container: %s", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping postgres container: %s", err)
	}

	postgresDb := db.NewIntegrationTestDatabase(dbConn)

	return &IntegrationTestsBuilder{
		ctx:         ctx,
		pgContainer: pgContainer,
		dbConn:      dbConn,
		postgresDb:  postgresDb,
	}
}

func (b *IntegrationTestsBuilder) WithLogger() *IntegrationTestsBuilder {
	b.logger = logger.NewLogger()

	return b
}

func (b *IntegrationTestsBuilder) WithTracer() *IntegrationTestsBuilder {
	b.tracer = tracer.NewTracer(b.ctx, b.logger, env.ProvideAppName(), env.ProvideAppEnv())

	return b
}

func (b *IntegrationTestsBuilder) BuildStudentRepository() domain_repositories.IStudentRepository {
	return repositories.NewStudentRepository(b.tracer, b.postgresDb)
}

func (b *IntegrationTestsBuilder) TearDown() {
	b.dbConn.Close()
	b.pgContainer.Terminate(b.ctx)
}

func (b *IntegrationTestsBuilder) GetCtx() context.Context {
	return b.ctx
}
