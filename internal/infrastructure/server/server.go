package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	IServer interface {
		GetEcho() *echo.Echo
		ListenAndServe(ctx context.Context, logger logger.ILogger)
		GracefulShutdownServer(ctx context.Context, logger logger.ILogger) error
	}

	Server struct {
		echo *echo.Echo
	}
)

func NewServer(logger logger.ILogger) IServer {
	e := echo.New()

	e.Use(middlewares.CORS())
	e.Use(middleware.Logger())
	e.Use(middlewares.RequestContext(logger))
	e.Use(middlewares.Timeout(logger))
	e.Use(middlewares.Recover(logger))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Server.Addr = fmt.Sprintf("%s:%s", env.ProvideAppHost(), env.ProvideAppPort())

	e.Validator = NewValidator()

	return &Server{
		echo: e,
	}
}

func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}

func (s *Server) ListenAndServe(ctx context.Context, logger logger.ILogger) {
	logger.Info(ctx, "HTTP server started", "address", s.echo.Server.Addr)

	s.echo.Server.ListenAndServe()
}

func (s *Server) GracefulShutdownServer(ctx context.Context, logger logger.ILogger) error {
	if env.ProvideAppEnv() == "test" {
		return nil
	}

	duration, err := time.ParseDuration(env.ProvideAppGracefulShutdownTimeoutInSeconds())
	if err != nil {
		logger.Error(ctx, "could not parse APP_GRACEFUL_SHUTDOWN_TIMEOUT", err)

		duration = time.Second * 5
		logger.Warn(ctx, "using default APP_GRACEFUL_SHUTDOWN_TIMEOUT of 5s")
	}

	shutdownCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	err = s.echo.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(shutdownCtx, "failed to gracefull server shutdown", err)
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info(ctx, "Graceful server shutdown completed successfully")

	return nil
}
