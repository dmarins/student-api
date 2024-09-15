package server

import (
	"log"
	"net/http"
	"time"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigRequestTimeout() echo.MiddlewareFunc {

	duration, err := time.ParseDuration(env.GetEnvironmentVariable("REQUEST_TIMEOUT"))
	if err != nil {
		log.Fatalf("could not parse REQUEST_TIMEOUT: %s", err)

		duration = time.Second * 30
		log.Printf("using default REQUEST_TIMEOUT of 30s")
	}

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "request timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Fatalf("request timeout, err: %s, code: %v, uri: %s", err, http.StatusGatewayTimeout, c.Request().RequestURI)
		},
		Timeout: duration,
	})
}
