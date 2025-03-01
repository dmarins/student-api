package create_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentCreateWithNameCheck_Execute_WhenTheRepositoryFailsToCheckIfTheStudentExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreateUseCaseValidationsExecute).
		SettingLoggerErrorBehavior("error checking if student exists", f.fakeError, "name", f.fakeStudentCreateInput.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentCreateInput.Name).
		Return(false, f.fakeError)

	sut := builder.BuildStudentCreateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentCreateInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreateWithNameCheck_Execute_WhenTheStudentAlreadyExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreateUseCaseValidationsExecute).
		SettingLoggerWarnBehavior("there is already a student with the same name", "name", f.fakeStudentCreateInput.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentCreateInput.Name).
		Return(true, nil)

	sut := builder.BuildStudentCreateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentCreateInput)

	assert.EqualValues(t, dtos.NewConflictResult(), result)
}

func TestStudentCreateWithNameCheck_Execute_WhenAnErrorIsReturnedByTheNextDecorator(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreateUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentCreateInput.Name).
		Return(false, nil)

	builder.StudentCreateUseCaseMock.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentCreateInput).
		Return(dtos.NewInternalServerErrorResult())

	sut := builder.BuildStudentCreateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentCreateInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreateWithNameCheck_Execute_WhenTheStudentDoesNotExist(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreateUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentCreateInput.Name).
		Return(false, nil)

	builder.StudentCreateUseCaseMock.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentCreateInput).
		Return(dtos.NewCreatedResult(f.fakeStudentCreateInput))

	sut := builder.BuildStudentCreateWithNameCheck()

	result := sut.Execute(builder.Ctx, f.fakeStudentCreateInput)

	assert.EqualValues(t, dtos.NewCreatedResult(result.Data), result)
}
