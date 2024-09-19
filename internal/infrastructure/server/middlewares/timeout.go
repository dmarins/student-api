package middlewares

import (
	"net/http"
	"time"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Timeout(logger logger.ILogger) echo.MiddlewareFunc {
	duration, err := time.ParseDuration(env.GetEnvironmentVariable("REQUEST_TIMEOUT"))
	if err != nil {
		logger.Error(nil, "could not parse REQUEST_TIMEOUT, using default of 30s", err)
		duration = time.Second * 30
		logger.Warn(nil, "using default REQUEST_TIMEOUT of 30s")
	}

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "request timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			logger.Error(nil, "request timeout", err,
				dtos.Field{
					Key: "uri", Value: c.Request().RequestURI,
				},
				dtos.Field{
					Key: "code", Value: http.StatusGatewayTimeout,
				},
			)
		},
		Timeout: duration,
	})
}
