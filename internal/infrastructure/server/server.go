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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type (
	IServer interface {
		GetEcho() *echo.Echo
		// Start(ctx context.Context) error
		// GracefulShutdown(ctx context.Context) error
	}

	Server struct {
		echo *echo.Echo
	}
)

func NewServer(lc fx.Lifecycle) IServer {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(ConfigRequestContext())
	e.Use(ConfigRequestTimeout())
	e.Use(ConfigCORS())

	e.Server.Addr = fmt.Sprintf("%s:%s",
		env.GetEnvironmentVariable("APP_HOST"),
		env.GetEnvironmentVariable("APP_PORT"),
	)

	e.Validator = ConfigCustomValidator()

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info(ctx, "starting HTTP server...")
				go e.Server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				duration, err := time.ParseDuration(env.GetEnvironmentVariable("APP_GRACEFUL_SHUTDOWN_TIMEOUT"))
				if err != nil {
					logger.Warn(ctx, "could not parse APP_GRACEFUL_SHUTDOWN_TIMEOUT", err)

					duration = time.Second * 5
					logger.Warn(ctx, "using default APP_GRACEFUL_SHUTDOWN_TIMEOUT of 5s", err)
				}

				if err := logger.Sync(); err != nil {
					logger.Error(ctx, "failed to synchronize logger", err)
				}

				shutdownCtx, cancel := context.WithTimeout(ctx, duration)
				defer cancel()

				err = e.Shutdown(shutdownCtx)
				if err != nil {
					logger.Error(shutdownCtx, "failed to gracefully shutdown server", err)
					return err
				}

				quit := make(chan os.Signal, 1)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit

				return nil
			},
		},
	)

	return &Server{
		echo: e,
	}
}

func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}
