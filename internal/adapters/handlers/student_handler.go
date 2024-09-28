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

func (h *StudentHandler) Create(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerCreate)
	defer span.End()

	h.Tracer.AddAttributes(span, tracer.StudentHandlerCreate,
		tracer.Attributes{
			"Tenant": ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
		})

	var studentInput dtos.StudentInput
	if err := ectx.Bind(&studentInput); err != nil {
		h.Logger.Warn(ctx, "invalid payload, check the data sent", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "echo bind ok")

	if err := ectx.Validate(&studentInput); err != nil {
		h.Logger.Warn(ctx, "invalid field", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "echo validate ok")

	student := entities.Student{
		Name: studentInput.Name,
	}

	return ReturnResult(ectx, h.CreateStudentUseCase.Execute(ctx, student))
}
