package handlers

import (
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	CreateStudentUseCase usecases.ICreateStudentUseCase
	Tracer               tracer.ITracer
}

func NewStudentHandler(
	tracer tracer.ITracer,
	createStudentUseCase usecases.ICreateStudentUseCase) *StudentHandler {
	handler := &StudentHandler{
		CreateStudentUseCase: createStudentUseCase,
		Tracer:               tracer,
	}

	return handler
}

func RegisterStudentRoutes(s server.IServer, h *StudentHandler) {
	s.GetEcho().POST("/student", h.Create)
}

func (h *StudentHandler) Create(c echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(c.Request(), tracer.StudentHandlerCreate)
	defer span.End()

	h.Tracer.AddAttributes(span, tracer.StudentHandlerCreate,
		tracer.Attributes{
			"Tenant": c.Request().Header.Get(env.GetEnvironmentVariable("HEADER_TENANT")),
		})

	var studentInput dtos.StudentInput
	if err := c.Bind(&studentInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&studentInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	h.Tracer.AddAttributes(span, tracer.StudentHandlerCreate,
		tracer.Attributes{
			"Payload": studentInput,
		})

	result, err := h.CreateStudentUseCase.Execute(ctx, entities.Student{Name: studentInput.Name})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
