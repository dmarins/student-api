package main

import (
	"database/sql"

	"github.com/dmarins/student-api/internal/adapters/handlers"
	"github.com/dmarins/student-api/internal/adapters/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/usecases/create"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type DependencyInjectionContainer struct {
	HttpServer     server.IHttpServer
	StudentHandler *handlers.StudentHandler
}

func NewDependencyInjectionContainer(db *sql.DB, logger *zap.Logger, tracerProvider trace.TracerProvider) *DependencyInjectionContainer {
	httpServer := server.NewEchoHttpServer()
	studentHandler := handlers.NewStudentHandler(httpServer, NewCreateStudentUseCase(db, logger, tracerProvider))

	return &DependencyInjectionContainer{
		HttpServer:     httpServer,
		StudentHandler: studentHandler,
	}
}

func NewCreateStudentUseCase(db *sql.DB, logger *zap.Logger, tracerProvider trace.TracerProvider) usecases.ICreateStudentUseCase {
	studentRepo := repositories.NewStudentRepository(db)
	createStudentUseCase := create.NewCreateStudentUseCase(studentRepo)

	return createStudentUseCase
}
