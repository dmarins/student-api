package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.PATCH, echo.OPTIONS},
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
		})
}
