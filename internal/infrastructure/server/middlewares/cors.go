package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	allowedMethods := []string{
		echo.POST,
		echo.GET,
		echo.PUT,
		echo.PATCH,
		echo.OPTIONS,
	}

	allowedHeaders := []string{
		echo.HeaderOrigin,
		echo.HeaderContentType,
		echo.HeaderAccept,
		echo.HeaderAuthorization,
	}

	allowedOrigins := []string{"*"}

	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods:     allowedMethods,
			AllowOrigins:     allowedOrigins,
			AllowHeaders:     allowedHeaders,
			AllowCredentials: true,
		})
}
