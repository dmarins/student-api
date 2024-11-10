package delete_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStudentDeleteWithPersistence_Execute_WhenRepositoryFailsToDeleteStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentDeleteUseCasePersistenceExecute).
		SettingLoggerErrorBehavior("error deleting a student", f.fakeError)

	builder.StudentRepositoryMock.
		EXPECT().
		Delete(builder.Ctx, gomock.Any()).
		Return(f.fakeError)

	sut := builder.BuildStudentDeleteWithPersistence()

	result := sut.Execute(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentDeleteWithPersistence_Execute_WhenRepositoryRemovesTheStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentDeleteUseCasePersistenceExecute).
		SettingLoggerDebugBehavior("student deleted")

	builder.StudentRepositoryMock.
		EXPECT().
		Delete(builder.Ctx, gomock.Any()).
		Return(nil)

	sut := builder.BuildStudentDeleteWithPersistence()

	result := sut.Execute(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.EqualValues(t, dtos.NewNoCotentResult(), result)
}
