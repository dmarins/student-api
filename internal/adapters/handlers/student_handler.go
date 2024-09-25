package handlers

import (
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	echo "github.com/labstack/echo/v4"
)

type StudentHandler struct {
	CreateStudentUseCase usecases.IStudentCreationUseCase
	Tracer               tracer.ITracer
	Logger               logger.ILogger
}

func NewStudentHandler(
	tracer tracer.ITracer,
	logger logger.ILogger,
	createStudentUseCase usecases.IStudentCreationUseCase) *StudentHandler {
	handler := &StudentHandler{
		CreateStudentUseCase: createStudentUseCase,
		Tracer:               tracer,
		Logger:               logger,
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
	var result *dtos.Result

	if err := c.Bind(&studentInput); err != nil {
		h.Logger.Warn(ctx, "invalid payload, check the data sent", "error", err.Error())

		result = dtos.NewHttpStatusBadRequestResult()
		return echo.NewHTTPError(result.StatusCode, result.Message)
	}

	h.Logger.Debug(ctx, "echo bind ok")

	if err := c.Validate(&studentInput); err != nil {
		h.Logger.Warn(ctx, "invalid field", "error", err.Error())

		result = dtos.NewHttpStatusBadRequestResult()
		return echo.NewHTTPError(result.StatusCode, result.Message)
	}

	h.Logger.Debug(ctx, "echo validate ok")

	output, result := h.CreateStudentUseCase.Execute(ctx, entities.Student{Name: studentInput.Name})
	switch result.StatusCode {
	case http.StatusCreated:
		return c.JSON(result.StatusCode, dtos.NewHttpStatusCreatedResult(output))
	case http.StatusInternalServerError:
		return echo.NewHTTPError(result.StatusCode, result.Errors)
	default:
		return echo.NewHTTPError(result.StatusCode, result)
	}
}
