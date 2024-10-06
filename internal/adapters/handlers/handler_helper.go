package handlers

import (
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/labstack/echo/v4"
)

func ReturnResult(ectx echo.Context, result *dtos.Result) error {
	switch result.Code {
	// case http.StatusOK:
	// 	return ectx.JSON(http.StatusOK, result)
	case http.StatusCreated:
		return ectx.JSON(http.StatusCreated, result)
	// case http.StatusBadRequest:
	// 	return echo.NewHTTPError(http.StatusBadRequest, result.Message)
	// case http.StatusNotFound:
	// 	return echo.NewHTTPError(http.StatusNotFound, result.Message)
	case http.StatusConflict:
		return echo.NewHTTPError(http.StatusConflict, result.Message)
	case http.StatusInternalServerError:
		return echo.NewHTTPError(http.StatusInternalServerError, result.Message)
	}

	return nil
}
