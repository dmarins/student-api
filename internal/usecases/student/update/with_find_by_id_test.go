package update_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentUpdateWithFindById_Execute_WhenTheRepositoryFailsToFindTheStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute).
		SettingLoggerErrorBehavior("error finding student by id", f.fakeError, "id", f.fakeStudentUpdateInput.ID)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudentUpdateInput.ID).
		Return(nil, f.fakeError)

	sut := builder.BuildStudentUpdateWithFindById()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentUpdateWithFindById_Execute_WhenTheStudentIsNotFound(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute).
		SettingLoggerWarnBehavior("the student was not found according to the id provided", "id", f.fakeStudentUpdateInput.ID)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudentUpdateInput.ID).
		Return(nil, nil)

	sut := builder.BuildStudentUpdateWithFindById()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewNotFoundResult(), result)
}

func TestStudentUpdateWithFindById_Execute_WhenTheStudentIsFound(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudentUpdateInput.ID).
		Return(&f.fakeStudent, nil)

	builder.StudentUpdateUseCaseMock.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentUpdateInput).
		Return(dtos.NewOkResult(f.fakeStudentUpdateInput))

	sut := builder.BuildStudentUpdateWithFindById()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewOkResult(result.Data), result)
}
