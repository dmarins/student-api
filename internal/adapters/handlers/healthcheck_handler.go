package handlers

import (
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
	"github.com/dmarins/student-api/internal/infrastructure/server"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	echo "github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
	Tracer             tracer.ITracer
	Logger             logger.ILogger
	HealthCheckUseCase usecases.IHealthCheckUseCase
}

func NewHealthCheckHandler(tracer tracer.ITracer, logger logger.ILogger, healthCheckUseCase usecases.IHealthCheckUseCase) *HealthCheckHandler {
	handler := &HealthCheckHandler{
		Tracer:             tracer,
		Logger:             logger,
		HealthCheckUseCase: healthCheckUseCase,
	}

	return handler
}

func RegisterHealthCheckRoute(s server.IServer, h *HealthCheckHandler) {
	s.GetEcho().GET("/health", h.Get)
}

// HealthCheck godoc
//
//	@Summary		Check if the API is available.
//	@Description	Checks if the API has connectivity to your database.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	dtos.Result
//	@Failure		500	{object}	dtos.Result
//	@Router			/health [get]
func (h *HealthCheckHandler) Get(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.HealthCheckHandlerGet)
	defer span.End()

	h.Tracer.AddAttributes(span, tracer.HealthCheckHandlerGet,
		tracer.Attributes{
			"Cid": ectx.Request().Header.Get(env.ProvideCidHeaderName()),
		})

	return ReturnResult(ectx, h.HealthCheckUseCase.Execute(ctx))
}
