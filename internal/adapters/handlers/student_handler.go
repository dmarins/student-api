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
	StudentUpdateUseCase usecases.IStudentUpdateUseCase
	StudentDeleteUseCase usecases.IStudentDeleteUseCase
	StudentSearchUseCase usecases.IStudentSearchUseCase
}

func NewStudentHandler(
	tracer tracer.ITracer,
	logger logger.ILogger,
	studentCreateUseCase usecases.IStudentCreateUseCase,
	studentReadingUseCase usecases.IStudentReadUseCase,
	studentUpdateUseCase usecases.IStudentUpdateUseCase,
	studentDeleteUseCase usecases.IStudentDeleteUseCase,
	studentSearchUseCase usecases.IStudentSearchUseCase) *StudentHandler {
	handler := &StudentHandler{
		Tracer:               tracer,
		Logger:               logger,
		StudentCreateUseCase: studentCreateUseCase,
		StudentReadUseCase:   studentReadingUseCase,
		StudentUpdateUseCase: studentUpdateUseCase,
		StudentDeleteUseCase: studentDeleteUseCase,
		StudentSearchUseCase: studentSearchUseCase,
	}

	return handler
}

func RegisterStudentRoutes(s server.IServer, h *StudentHandler) {
	routesGroup := s.GetEcho().Group("/students")

	routesGroup.POST("", h.Create)
	routesGroup.GET("/:id", h.Read)
	routesGroup.PUT("/:id", h.Update)
	routesGroup.DELETE("/:id", h.Delete)
	routesGroup.GET("", h.Search)
}

// StudentCreate godoc
//
//	@Summary		Allows you to create a student.
//	@Description	Allows you to create a student after validating duplicate names.
//	@Tags			students
//	@Param			x-tenant	header	string					true	"To identify the tenant"
//	@Param			x-cid		header	string					false	"To identify the request"
//	@Param			payload		body	dtos.StudentCreateInput	true	"To create a student"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	dtos.Result
//	@Failure		400	{object}	dtos.Result
//	@Failure		409	{object}	dtos.Result
//	@Failure		500	{object}	dtos.Result
//	@Router			/students [post]
func (h *StudentHandler) Create(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerCreate)
	defer span.End()

	h.Tracer.AddAttributes(span, tracer.StudentHandlerCreate,
		tracer.Attributes{
			"Tenant": ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
		})

	var studentCreateInput dtos.StudentCreateInput
	if err := ectx.Bind(&studentCreateInput); err != nil {
		h.Logger.Warn(ctx, "invalid payload, check the data sent", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "echo bind ok")

	if err := ectx.Validate(&studentCreateInput); err != nil {
		h.Logger.Warn(ctx, "invalid field", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "echo validate ok")

	return ReturnResult(ectx, h.StudentCreateUseCase.Execute(ctx, studentCreateInput))
}

// StudentRead godoc
//
//	@Summary		Allows you to get the details of a student.
//	@Description	Allows you to get the details of a student by ID.
//	@Tags			students
//	@Param			x-tenant	header	string	true	"To identify the tenant"
//	@Param			x-cid		header	string	false	"To identify the request"
//	@Param			id			path	string	true	"Student identifier"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dtos.Result
//	@Failure		400	{object}	dtos.Result
//	@Failure		404	{object}	dtos.Result
//	@Failure		500	{object}	dtos.Result
//	@Router			/students/{id} [get]
func (h *StudentHandler) Read(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerRead)
	defer span.End()

	studentId := ectx.Param("id")

	h.Tracer.AddAttributes(span, tracer.StudentHandlerRead,
		tracer.Attributes{
			"Tenant":    ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
			"StudentId": studentId,
		})

	ok := uuid.IsValid(studentId)
	if !ok {
		h.Logger.Warn(ctx, "identifier format is invalid", "id", studentId)

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "identifier format ok")

	return ReturnResult(ectx, h.StudentReadUseCase.Execute(ctx, studentId))
}

// StudentUpdate godoc
//
//	@Summary		Allows you to update a student data.
//	@Description	Allows you to update a student's data after finding them and validating if there are duplicate names.
//	@Tags			students
//	@Param			x-tenant	header	string					true	"To identify the tenant"
//	@Param			x-cid		header	string					false	"To identify the request"
//	@Param			id			path	string					true	"Student identifier"
//	@Param			payload		body	dtos.StudentUpdateInput	true	"To update a student"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dtos.Result
//	@Failure		400	{object}	dtos.Result
//	@Failure		404	{object}	dtos.Result
//	@Failure		409	{object}	dtos.Result
//	@Failure		500	{object}	dtos.Result
//	@Router			/students/{id} [put]
func (h *StudentHandler) Update(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerUpdate)
	defer span.End()

	studentId := ectx.Param("id")

	h.Tracer.AddAttributes(span, tracer.StudentHandlerUpdate,
		tracer.Attributes{
			"Tenant":    ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
			"StudentId": studentId,
		})

	ok := uuid.IsValid(studentId)
	if !ok {
		h.Logger.Warn(ctx, "identifier format is invalid", "id", studentId)

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "identifier format ok")

	var studentUpdateInput dtos.StudentUpdateInput
	if err := ectx.Bind(&studentUpdateInput); err != nil {
		h.Logger.Warn(ctx, "invalid payload, check the data sent", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "echo bind ok")

	studentUpdateInput.ID = studentId
	if err := ectx.Validate(&studentUpdateInput); err != nil {
		h.Logger.Warn(ctx, "invalid field", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "echo validate ok")

	return ReturnResult(ectx, h.StudentUpdateUseCase.Execute(ctx, studentUpdateInput))
}

func (h *StudentHandler) Delete(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerDelete)
	defer span.End()

	studentId := ectx.Param("id")

	h.Tracer.AddAttributes(span, tracer.StudentHandlerDelete,
		tracer.Attributes{
			"Tenant":    ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
			"StudentId": studentId,
		})

	ok := uuid.IsValid(studentId)
	if !ok {
		h.Logger.Warn(ctx, "identifier format is invalid", "id", studentId)

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Logger.Debug(ctx, "identifier format ok")

	return ReturnResult(ectx, h.StudentDeleteUseCase.Execute(ctx, studentId))
}

// StudentSearch godoc
//
//	@Summary		Allows you to search the students.
//	@Description	Allows you to search for students by controlling pagination, sorting and filtering results.
//	@Tags			students
//	@Param			x-tenant	header	string					true	"To identify the tenant"
//	@Param			x-cid		header	string					false	"To identify the request"
//	@Param			pagination	query	dtos.PaginationInput	true	"Pagination and sorting"
//	@Param			filter		query	dtos.Filter				false	"Filters"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dtos.Result
//	@Failure		400	{object}	dtos.Result
//	@Failure		500	{object}	dtos.Result
//	@Router			/students [get]
func (h *StudentHandler) Search(ectx echo.Context) error {
	span, ctx := h.Tracer.NewRootSpan(ectx.Request(), tracer.StudentHandlerSearch)
	defer span.End()

	h.Tracer.AddAttributes(span, tracer.StudentHandlerSearch,
		tracer.Attributes{
			"Tenant": ectx.Request().Header.Get(env.ProvideTenantHeaderName()),
		})

	var paginationRequest dtos.PaginationInput
	if err := ectx.Bind(&paginationRequest); err != nil {
		h.Logger.Warn(ctx, "invalid pagination, check the data sent", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	paginationRequest = *dtos.NewPaginationInput(
		paginationRequest.Page,
		paginationRequest.PageSize,
		paginationRequest.SortOrder,
		paginationRequest.SortField,
	)

	h.Tracer.AddAttributes(span, tracer.StudentHandlerSearch,
		tracer.Attributes{
			"Pagination": paginationRequest,
		})

	h.Logger.Debug(ctx, "echo bind pagination ok")

	var filter dtos.Filter
	if err := ectx.Bind(&filter); err != nil {
		h.Logger.Warn(ctx, "invalid filter, check the data sent", "error", err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, dtos.NewBadRequestResult().Message)
	}

	h.Tracer.AddAttributes(span, tracer.StudentHandlerSearch,
		tracer.Attributes{
			"Filter": filter,
		})

	h.Logger.Debug(ctx, "echo bind filter ok")

	return ReturnResult(ectx, h.StudentSearchUseCase.Execute(ctx, paginationRequest, filter))
}
