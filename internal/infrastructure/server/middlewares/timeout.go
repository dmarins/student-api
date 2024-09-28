package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Timeout(logger logger.ILogger) echo.MiddlewareFunc {
	duration, err := time.ParseDuration(env.ProvideRequestTimeoutInSeconds())
	if err != nil {
		logger.Error(context.TODO(), "could not parse REQUEST_TIMEOUT", err)

		duration = time.Second * 30
		logger.Warn(context.TODO(), "using default REQUEST_TIMEOUT of 30s")
	}

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "request timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			logger.Error(context.TODO(),
				"request timeout", err,
				"uri", c.Request().RequestURI,
				"code", strconv.Itoa(http.StatusGatewayTimeout),
			)
		},
		Timeout: duration,
	})
}
