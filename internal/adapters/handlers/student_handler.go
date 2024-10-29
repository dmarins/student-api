package handlers

import (
	"net/http"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/dmarins/student-api/internal/infrastructure/uuid"
	echo "github.com/labstack/echo/v4"
)

type StudentHandler struct {
	Tracer               tracer.ITracer
	Logger               logger.ILogger
	StudentCreateUseCase usecases.IStudentCreateUseCase
	StudentReadUseCase   usecases.IStudentReadUseCase
}

func NewStudentHandler(
	tracer tracer.ITracer,
	logger logger.ILogger,
	studentCreateUseCase usecases.IStudentCreateUseCase,
	studentReadingUseCase usecases.IStudentReadUseCase) *StudentHandler {
	handler := &StudentHandler{
		Tracer:               tracer,
		Logger:               logger,
		StudentCreateUseCase: studentCreateUseCase,
		StudentReadUseCase:   studentReadingUseCase,
	}

	return handler
}

func RegisterStudentRoutes(s server.IServer, h *StudentHandler) {
	routesGroup := s.GetEcho().Group("/students")
	routesGroup.POST("", h.Post)
	routesGroup.GET("/:id", h.Get)
}

func (h *StudentHandler) Post(ectx echo.Context) error {
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

	return ReturnResult(ectx, h.StudentCreateUseCase.Execute(ctx, studentInput))
}

func (h *StudentHandler) Get(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerGet)
	defer span.End()

	studentId := ectx.Param("id")

	h.Tracer.AddAttributes(span, tracer.StudentHandlerGet,
		tracer.Attributes{
			"Tenant":    ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
			"StudentId": studentId,
		})

	if ok := uuid.IsValid(studentId); !ok {
		h.Logger.Warn(ctx, "id format is invalid", "id", studentId)

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "validate ok")

	return ReturnResult(ectx, h.StudentReadUseCase.Execute(ctx, studentId))
}
