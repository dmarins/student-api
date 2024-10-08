package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

func Timeout(logger logger.ILogger) echo.MiddlewareFunc {
	ctx := context.Background()

	duration, err := time.ParseDuration(env.ProvideRequestTimeoutInSeconds())
	if err != nil {
		logger.Error(ctx, "could not parse REQUEST_TIMEOUT", err)

		duration = time.Second * 30
		logger.Warn(ctx, "using default REQUEST_TIMEOUT of 30s")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			headerCidKey := env.ProvideCidHeaderName()
			cid := c.Request().Header.Get(headerCidKey)

			ctx, cancel := context.WithTimeout(c.Request().Context(), duration)
			defer cancel()

			c.SetRequest(c.Request().WithContext(ctx))

			done := make(chan error, 1)

			go func() {
				done <- next(c)
			}()

			select {
			case err := <-done:
				return err
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					logger.Error(ctx, "request timeout", err, "cid", cid)
					return c.JSON(http.StatusGatewayTimeout, dtos.NewGatewayTimeoutErrorResult())
				}

				return nil
			}
		}
	}
}
