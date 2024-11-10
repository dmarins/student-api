package delete_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentDeleteWithFindById_Execute_WhenTheRepositoryFailsToFindTheStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentDeleteUseCaseFindByIdExecute).
		SettingLoggerErrorBehavior("error finding student by id", f.fakeError, "id", f.fakeStudentToBeDeleted)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudentToBeDeleted).
		Return(nil, f.fakeError)

	sut := builder.BuildStudentDeleteWithFindById()

	result := sut.Execute(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentDeleteWithFindById_Execute_WhenTheStudentIsNotFound(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentDeleteUseCaseFindByIdExecute).
		SettingLoggerWarnBehavior("the student was not found according to the id provided", "id", f.fakeStudentToBeDeleted)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudentToBeDeleted).
		Return(nil, nil)

	sut := builder.BuildStudentDeleteWithFindById()

	result := sut.Execute(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.EqualValues(t, dtos.NewNotFoundResult(), result)
}

func TestStudentDeleteWithFindById_Execute_WhenTheStudentIsFound(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentDeleteUseCaseFindByIdExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudentToBeDeleted).
		Return(&f.fakeStudent, nil)

	builder.StudentDeleteUseCaseMock.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentToBeDeleted).
		Return(dtos.NewNoCotentResult())

	sut := builder.BuildStudentDeleteWithFindById()

	result := sut.Execute(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.EqualValues(t, dtos.NewNoCotentResult(), result)
}
