package tests

import (
	"context"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/mocks"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/dmarins/student-api/internal/usecases/student/creation"
	"go.uber.org/mock/gomock"
)

var requestContextEnvVarName = contextKey(env.ProvideRequestContextName())

type (
	contextKey string

	UnitTestsBuilder struct {
		Ctx                   context.Context
		Ctrl                  *gomock.Controller
		TracerMock            *mocks.MockITracer
		SpanMock              *mocks.MockISpan
		LoggerMock            *mocks.MockILogger
		StudentRepositoryMock *mocks.MockIStudentRepository
		Next                  *mocks.MockIStudentCreationUseCase
	}
)

func NewUnitTestsBuilder(t *testing.T) *UnitTestsBuilder {
	ctrl := gomock.NewController(t)
	loggerMock := mocks.NewMockILogger(ctrl)
	tracerMock := mocks.NewMockITracer(ctrl)
	spanMock := mocks.NewMockISpan(ctrl)
	studentRepositoryMock := mocks.NewMockIStudentRepository(ctrl)
	next := mocks.NewMockIStudentCreationUseCase(ctrl)

	return &UnitTestsBuilder{
		Ctrl:                  ctrl,
		LoggerMock:            loggerMock,
		TracerMock:            tracerMock,
		SpanMock:              spanMock,
		StudentRepositoryMock: studentRepositoryMock,
		Next:                  next,
	}
}

func (b *UnitTestsBuilder) WithValidCtx() *UnitTestsBuilder {
	requestContext := dtos.RequestContext{
		TenantId: "x-tenant",
		Cid:      "cid",
	}

	b.Ctx = context.WithValue(
		context.Background(),
		requestContextEnvVarName,
		requestContext,
	)

	return b
}

func (b *UnitTestsBuilder) SettingTracerBehavior(spanName string, times int, attributes tracer.Attributes) *UnitTestsBuilder {
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
		AddAttributes(b.SpanMock, spanName, attributes).
		Times(times)

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

func (b *UnitTestsBuilder) BuildStudentCreationWithValidations() usecases.IStudentCreationUseCase {
	return creation.NewStudentCreationWithValidations(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock, b.Next)
}

func (b *UnitTestsBuilder) BuildStudentCreationWithPersistence() usecases.IStudentCreationUseCase {
	return creation.NewStudentCreationWithPersistence(b.TracerMock, b.LoggerMock, b.StudentRepositoryMock)
}
