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
	"go.uber.org/fx"
)

type (
	IServer interface {
		GetEcho() *echo.Echo
	}

	Server struct {
		echo *echo.Echo
	}
)

func NewServer(lc fx.Lifecycle, logger logger.ILogger) IServer {
	e := setupEchoServer(logger)

	setupLifecycle(lc, e, logger)

	return &Server{
		echo: e,
	}
}

func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}

func setupEchoServer(logger logger.ILogger) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middlewares.RequestContext(logger))
	e.Use(middlewares.Timeout(logger))
	e.Use(middlewares.CORS())

	e.Server.Addr = fmt.Sprintf(
		"%s:%s",
		env.GetEnvironmentVariable("APP_HOST"),
		env.GetEnvironmentVariable("APP_PORT"),
	)

	e.Validator = NewValidator()

	return e
}

func setupLifecycle(lc fx.Lifecycle, e *echo.Echo, logger logger.ILogger) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info(ctx, "starting HTTP server...", "address", e.Server.Addr)

				go e.Server.ListenAndServe()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return gracefulShutdownServer(ctx, e, logger)
			},
		},
	)
}

func gracefulShutdownServer(ctx context.Context, e *echo.Echo, logger logger.ILogger) error {
	duration, err := time.ParseDuration(env.GetEnvironmentVariable("APP_GRACEFUL_SHUTDOWN_TIMEOUT"))
	if err != nil {
		logger.Error(ctx, "could not parse APP_GRACEFUL_SHUTDOWN_TIMEOUT", err)
		duration = time.Second * 5
		logger.Warn(ctx, "using default APP_GRACEFUL_SHUTDOWN_TIMEOUT of 5s")
	}

	shutdownCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	err = logger.Sync()
	if err != nil {
		logger.Error(ctx, "failed to synchronize logs", err)
	}

	err = e.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(shutdownCtx, "failed to gracefully shutdown server", err)
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}
