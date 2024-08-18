package main

import (
	"context"
	"log"

	"github.com/dmarins/student-api/internal/infrastructure/db"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	// Env vars
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./cmd")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("No .env file found: %v", err)
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
	e := echo.New()
	e.POST("/student", app.StudentHandler.CreateStudent)
	e.Logger.Fatal(e.Start(":8080"))
}
