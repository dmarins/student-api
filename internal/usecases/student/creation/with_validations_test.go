package creation_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentCreationWithValidations_Execute_WhenTheRepositoryFailsToCheckIfTheStudentExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute).
		SettingLoggerErrorBehavior("error checking if student exists", f.fakeError, "name", f.fakeStudentInput.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentInput.Name).
		Return(false, f.fakeError)

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudentInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenTheStudentAlreadyExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute).
		SettingLoggerWarnBehavior("there is already a student with the same name", "name", f.fakeStudentInput.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentInput.Name).
		Return(true, nil)

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudentInput)

	assert.EqualValues(t, dtos.NewConflictResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenAnErrorIsReturnedByTheNextDecorator(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentInput.Name).
		Return(false, nil)

	builder.Next.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentInput).
		Return(dtos.NewInternalServerErrorResult())

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudentInput)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenTheStudentDoesNotExist(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudentInput.Name).
		Return(false, nil)

	builder.Next.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudentInput).
		Return(dtos.NewCreatedResult(f.fakeStudentInput))

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudentInput)

	assert.EqualValues(t, dtos.NewCreatedResult(result.Data), result)
}
