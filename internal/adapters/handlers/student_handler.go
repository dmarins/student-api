package handlers

import (
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	CreateStudentUseCase usecases.ICreateStudentUseCase
}

func NewStudentHandler(createStudentUseCase usecases.ICreateStudentUseCase) *StudentHandler {
	return &StudentHandler{
		CreateStudentUseCase: createStudentUseCase,
	}
}

func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var studentInput dtos.StudentInput
	if err := c.Bind(&studentInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// if err := c.Validate(&studentInput); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	result, err := h.CreateStudentUseCase.Execute(c.Request().Context(), entities.Student{Name: studentInput.Name})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
