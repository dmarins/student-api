package healthcheck_test

import (
	"errors"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

var fakeError error = errors.New("fails")

func TestHealthCheckUseCase_Execute_WhenRepositoryFailsToCheckDbConnection(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.HealthCheckExecute, 0, nil).
		SettingLoggerErrorBehavior("error checking db connection", fakeError)

	builder.HealthCheckRepositoryMock.
		EXPECT().
		CheckDbConnection(builder.Ctx).
		Return(fakeError)

	sut := builder.BuildHealthCheckUseCase()

	result := sut.Execute(builder.Ctx)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestHealthCheckUseCase_Execute_WhenRepositoryCheckDbConnection(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.HealthCheckExecute, 0, nil)

	builder.HealthCheckRepositoryMock.
		EXPECT().
		CheckDbConnection(builder.Ctx).
		Return(nil)

	sut := builder.BuildHealthCheckUseCase()

	result := sut.Execute(builder.Ctx)

	expectedResult := dtos.NewOkResult(nil)
	expectedResult.Message = "healthy"

	assert.EqualValues(t, expectedResult, result)
}
