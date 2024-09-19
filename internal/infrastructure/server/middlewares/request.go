package middlewares

import (
	"context"
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestContext(logger logger.ILogger) echo.MiddlewareFunc {
	headerCidKey := env.GetEnvironmentVariable("HEADER_CID")
	headerTenantKey := env.GetEnvironmentVariable("HEADER_TENANT")
	requestContextKey := env.GetEnvironmentVariable("REQUEST_CONTEXT")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			cid := c.Request().Header.Get(headerCidKey)
			if cid == "" {
				cid = uuid.New().String()
				c.Request().Header.Set(headerCidKey, cid)
			}

			tenant := c.Request().Header.Get(headerTenantKey)
			if tenant == "" {
				logger.Warn(ctx, "the x-tenant header was not provided",
					dtos.Field{
						Key: "cid", Value: cid,
					})
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "the x-tenant header was not provided",
				})
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
