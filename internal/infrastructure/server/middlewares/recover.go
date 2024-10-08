package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

func Recover(logger logger.ILogger) echo.MiddlewareFunc {
	headerCidKey := env.ProvideCidHeaderName()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cid := c.Request().Header.Get(headerCidKey)
			ctx := c.Request().Context()

			defer func() {
				if r := recover(); r != nil {
					var err error
					switch r := r.(type) {
					case string:
						err = fmt.Errorf(r)
					case error:
						err = r
					default:
						err = fmt.Errorf("unknown panic: %v", r)
					}

					logger.Error(ctx, "panic recovered", err, "cid", cid)
					c.JSON(http.StatusInternalServerError, dtos.NewInternalServerErrorResult())
				}
			}()

			return next(c)
		}
	}
}
