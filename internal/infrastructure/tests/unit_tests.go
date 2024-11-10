package tests

import (
	"context"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/mocks"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/usecases/healthcheck"
	"github.com/dmarins/student-api/internal/usecases/student/create"
	"github.com/dmarins/student-api/internal/usecases/student/delete"
	"github.com/dmarins/student-api/internal/usecases/student/read"
	"github.com/dmarins/student-api/internal/usecases/student/update"
	"go.uber.org/mock/gomock"
)

type UnitTestsBuilder struct {
	Ctx                       context.Context
	Ctrl                      *gomock.Controller
	TracerMock                *mocks.MockITracer
	SpanMock                  *mocks.MockISpan
	LoggerMock                *mocks.MockILogger
	StudentRepositoryMock     *mocks.MockIStudentRepository
	HealthCheckRepositoryMock *mocks.MockIHealthCheckRepository
	StudentCreateUseCaseMock  *mocks.MockIStudentCreateUseCase
	StudentUpdateUseCaseMock  *mocks.MockIStudentUpdateUseCase
	StudentDeleteUseCaseMock  *mocks.MockIStudentDeleteUseCase
}

func NewUnitTestsBuilder(t *testing.T) *UnitTestsBuilder {
	ctrl := gomock.NewController(t)
	loggerMock := mocks.NewMockILogger(ctrl)
	tracerMock := mocks.NewMockITracer(ctrl)
	spanMock := mocks.NewMockISpan(ctrl)
	studentRepositoryMock := mocks.NewMockIStudentRepository(ctrl)
	healthCheckRepositoryMock := mocks.NewMockIHealthCheckRepository(ctrl)
	studentCreateUseCaseMock := mocks.NewMockIStudentCreateUseCase(ctrl)
	studentUpdateUseCaseMock := mocks.NewMockIStudentUpdateUseCase(ctrl)
	studentDeleteUseCaseMock := mocks.NewMockIStudentDeleteUseCase(ctrl)

	return &UnitTestsBuilder{
		Ctrl:                      ctrl,
		LoggerMock:                loggerMock,
		TracerMock:                tracerMock,
		SpanMock:                  spanMock,
		StudentRepositoryMock:     studentRepositoryMock,
		HealthCheckRepositoryMock: healthCheckRepositoryMock,
		StudentCreateUseCaseMock:  studentCreateUseCaseMock,
		StudentUpdateUseCaseMock:  studentUpdateUseCaseMock,
		StudentDeleteUseCaseMock:  studentDeleteUseCaseMock,
	}
}

func (b *UnitTestsBuilder) WithValidCtx() *UnitTestsBuilder {
	requestContext := dtos.RequestContext{
		TenantId: "x-tenant",
		Cid:      "cid",
	}

	b.Ctx = context.WithValue(
		context.Background(),
		env.ProvideRequestContextName(),
		requestContext,
	)

	return b
}

func (b *UnitTestsBuilder) SettingTracerBehavior(spanName string) *UnitTestsBuilder {
	b.TracerMock.
		EXPECT().
		NewSpanContext(b.Ctx, spanName).
		Return(b.SpanMock, b.Ctx).
		Times(1)

	b.SpanMock.
		EXPECT().
		End().
		Times(1)

	b.TracerMock.
		EXPECT().
		AddAttributes(b.SpanMock, spanName, gomock.Any()).
		AnyTimes()

	return b
}

func (b *UnitTestsBuilder) SettingLoggerDebugBehavior(debugMessage string, fields ...any) *UnitTestsBuilder {
	b.LoggerMock.
		EXPECT().
		Debug(b.Ctx, debugMessage, fields...)

	return b
}

func (b *UnitTestsBuilder) SettingLoggerWarnBehavior(debugMessage string, fields ...any) *UnitTestsBuilder {
	b.LoggerMock.
		EXPECT().
		Warn(b.Ctx, debugMessage, fields...)

	return b
}

func (b *UnitTestsBuilder) SettingLoggerErrorBehavior(debugMessage string, err error, fields ...any) *UnitTestsBuilder {
	b.LoggerMock.
		EXPECT().
		Error(b.Ctx, debugMessage, err, fields...)

	return b
}

func (b *UnitTestsBuilder) BuildHealthCheckUseCase() usecases.IHealthCheckUseCase {
	return healthcheck.NewHealthCheck(b.TracerMock, b.LoggerMock, b.HealthCheckRepositoryMock)
}

func (b *UnitTestsBuilder) BuildStudentCreateWithNameCheck() usecases.IStudentCreateUseCase {
	return create.NewStudentCreateWithNameCheck(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock, b.StudentCreateUseCaseMock)
}

func (b *UnitTestsBuilder) BuildStudentCreateWithPersistence() usecases.IStudentCreateUseCase {
	return create.NewStudentCreateWithPersistence(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}

func (b *UnitTestsBuilder) BuildStudentReadWithFindByID() usecases.IStudentReadUseCase {
	return read.NewStudentReadWithFindById(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}

func (b *UnitTestsBuilder) BuildStudentUpdateWithNameCheck() usecases.IStudentUpdateUseCase {
	return update.NewStudentUpdateWithNameCheck(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock, b.StudentUpdateUseCaseMock)
}

func (b *UnitTestsBuilder) BuildStudentUpdateWithFindById() usecases.IStudentUpdateUseCase {
	return update.NewStudentUpdateWithFindById(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock, b.StudentUpdateUseCaseMock)
}

func (b *UnitTestsBuilder) BuildStudentUpdateWithPersistence() usecases.IStudentUpdateUseCase {
	return update.NewStudentUpdateWithPersistence(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}

func (b *UnitTestsBuilder) BuildStudentDeleteWithFindById() usecases.IStudentDeleteUseCase {
	return delete.NewStudentDeleteWithFindById(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock, b.StudentDeleteUseCaseMock)
}

func (b *UnitTestsBuilder) BuildStudentDeleteWithPersistence() usecases.IStudentDeleteUseCase {
	return delete.NewStudentDeleteWithPersistence(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}
