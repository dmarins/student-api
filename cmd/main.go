package main

import (
	"context"
	"log"

	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

func main() {
	// Env vars
	err := env.InitEnvVars()
	if err != nil {
		log.Fatalf("failed to initialize env vars: %v", err)
	}

	// Log
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	// Tracing
	tracer, err := tracer.InitTracer()
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := tracer.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown tracer: %v", err)
		}
	}()

	// Database
	database, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// DI
	app := NewDependencyInjectionContainer(database, logger, tracer)

	// Http Server
	app.HttpServer.Start()
}
