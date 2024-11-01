package update_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStudentUpdateWithPersistence_Execute_WhenRepositoryFailsToAddStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCasePersistenceExecute).
		SettingLoggerErrorBehavior("error updating a student", f.fakeError)

	builder.StudentRepositoryMock.
		EXPECT().
		Update(builder.Ctx, gomock.Any()).
		Return(f.fakeError)

	sut := builder.BuildStudentUpdateWithPersistence()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentUpdateWithPersistence_Execute_WhenRepositoryAddsTheStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCasePersistenceExecute).
		SettingLoggerDebugBehavior("student updated")

	builder.StudentRepositoryMock.
		EXPECT().
		Update(builder.Ctx, gomock.Any()).
		Return(nil)

	sut := builder.BuildStudentUpdateWithPersistence()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewOkResult(result.Data), result)
}
