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
	"github.com/dmarins/student-api/internal/usecases/student/read"
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
	Next                      *mocks.MockIStudentCreateUseCase
}

func NewUnitTestsBuilder(t *testing.T) *UnitTestsBuilder {
	ctrl := gomock.NewController(t)
	loggerMock := mocks.NewMockILogger(ctrl)
	tracerMock := mocks.NewMockITracer(ctrl)
	spanMock := mocks.NewMockISpan(ctrl)
	studentRepositoryMock := mocks.NewMockIStudentRepository(ctrl)
	healthCheckRepositoryMock := mocks.NewMockIHealthCheckRepository(ctrl)
	next := mocks.NewMockIStudentCreateUseCase(ctrl)

	return &UnitTestsBuilder{
		Ctrl:                      ctrl,
		LoggerMock:                loggerMock,
		TracerMock:                tracerMock,
		SpanMock:                  spanMock,
		StudentRepositoryMock:     studentRepositoryMock,
		HealthCheckRepositoryMock: healthCheckRepositoryMock,
		Next:                      next,
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

func (b *UnitTestsBuilder) BuildStudentCreateWithValidations() usecases.IStudentCreateUseCase {
	return create.NewStudentCreateWithValidations(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock, b.Next)
}

func (b *UnitTestsBuilder) BuildStudentCreateWithPersistence() usecases.IStudentCreateUseCase {
	return create.NewStudentCreateWithPersistence(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}

func (b *UnitTestsBuilder) BuildStudentReadWithFindByID() usecases.IStudentReadUseCase {
	return read.NewStudentReadWithFindById(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}
