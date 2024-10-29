package create_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStudentCreateWithPersistence_Execute_WhenRepositoryFailsToAddStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingLoggerDebugBehavior("new student", "id", gomock.Any()).
		SettingTracerBehavior(tracer.StudentCreateUseCasePersistenceExecute).
		SettingLoggerErrorBehavior("error adding a new student", f.fakeError)

	builder.StudentRepositoryMock.
		EXPECT().
		Add(builder.Ctx, gomock.Any()).
		Return(f.fakeError)

	sut := builder.BuildStudentCreateWithPersistence()

	result := sut.Execute(builder.Ctx, f.fakeStudentInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreateWithPersistence_Execute_WhenRepositoryAddsTheStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreateUseCasePersistenceExecute).
		SettingLoggerDebugBehavior("new student", "id", gomock.Any()).
		SettingLoggerDebugBehavior("student stored")

	builder.StudentRepositoryMock.
		EXPECT().
		Add(builder.Ctx, gomock.Any()).
		Return(nil)

	sut := builder.BuildStudentCreateWithPersistence()

	result := sut.Execute(builder.Ctx, f.fakeStudentInput)

	assert.EqualValues(t, dtos.NewCreatedResult(result.Data), result)
}
