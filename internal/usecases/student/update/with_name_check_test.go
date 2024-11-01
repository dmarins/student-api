package update_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentUpdateWithNameCheck_Execute_WhenTheRepositoryFailsToCheckIfTheStudentExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute).
		SettingLoggerErrorBehavior("error checking if student exists", f.fakeError, "name", f.fakeStudentUpdateInput.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentUpdateInput.Name).
		Return(false, f.fakeError)

	sut := builder.BuildStudentUpdateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentUpdateWithNameCheck_Execute_WhenTheStudentAlreadyExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute).
		SettingLoggerWarnBehavior("there is already a student with the same name", "name", f.fakeStudentUpdateInput.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentUpdateInput.Name).
		Return(true, nil)

	sut := builder.BuildStudentUpdateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewConflictResult(), result)
}

func TestStudentUpdateWithNameCheck_Execute_WhenAnErrorIsReturnedByTheNextDecorator(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentUpdateInput.Name).
		Return(false, nil)

	builder.StudentUpdateUseCaseMock.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentUpdateInput).
		Return(dtos.NewInternalServerErrorResult())

	sut := builder.BuildStudentUpdateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentUpdateWithNameCheck_Execute_WhenTheStudentDoesNotExist(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentUpdateUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentUpdateInput.Name).
		Return(false, nil)

	builder.StudentUpdateUseCaseMock.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentUpdateInput).
		Return(dtos.NewCreatedResult(f.fakeStudentUpdateInput))

	sut := builder.BuildStudentUpdateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentUpdateInput)

	assert.EqualValues(t, dtos.NewCreatedResult(result.Data), result)
}
