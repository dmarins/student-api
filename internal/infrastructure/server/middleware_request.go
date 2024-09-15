package server

import (
	"context"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RequestContext struct {
	TenantId string
	Cid      string
}

func ConfigRequestContext() echo.MiddlewareFunc {
	return newRequestContext
}

func newRequestContext(next echo.HandlerFunc) echo.HandlerFunc {
	headerCid := env.GetEnvironmentVariable("HEADER_CID")

	return func(c echo.Context) error {
		cid := c.Request().Header.Get(headerCid)
		if cid == "" {
			cid = uuid.New().String()
			c.Request().Header.Set(headerCid, cid)
		}

		tenant := c.Request().Header.Get(env.GetEnvironmentVariable("HEADER_TENANT"))

		rctx := RequestContext{
			TenantId: tenant,
			Cid:      cid,
		}

		ctx := context.WithValue(c.Request().Context(), env.GetEnvironmentVariable("REQUEST_CONTEXT"), rctx)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
