package middlewares

import (
	"context"
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/uuid"
	"github.com/labstack/echo/v4"
)

var agnosticRoute string = "/health"
var tenant string
var cid string

func RequestContext(logger logger.ILogger) echo.MiddlewareFunc {
	headerCidKey := env.ProvideCidHeaderName()
	headerTenantKey := env.ProvideTenantHeaderName()
	requestContextKey := env.ProvideRequestContextName()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			path := c.Request().URL.Path

			if path != agnosticRoute {
				cid := c.Request().Header.Get(headerCidKey)
				if cid == "" {
					cid = uuid.NewId()
					c.Request().Header.Set(headerCidKey, cid)
				}

				tenant = c.Request().Header.Get(headerTenantKey)
				if tenant == "" {
					logger.Warn(ctx, "the x-tenant header was not provided", "cid", cid)
					return c.JSON(http.StatusBadRequest, dtos.NewBadRequestResult())
				}
			}

			rctx := dtos.RequestContext{
				TenantId: tenant,
				Cid:      cid,
			}

			ctxChanged := context.WithValue(ctx, requestContextKey, rctx)
			c.SetRequest(c.Request().WithContext(ctxChanged))

			return next(c)
		}
	}
}
