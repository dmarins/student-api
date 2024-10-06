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
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute, 1, f.fakeTracerAttributes).
		SettingLoggerErrorBehavior("error checking if student exists", f.fakeError, "name", f.fakeStudent.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudent.Name).
		Return(false, f.fakeError)

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudent)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenTheStudentAlreadyExists(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute, 1, f.fakeTracerAttributes).
		SettingLoggerWarnBehavior("there is already a student with the same name", "name", f.fakeStudent.Name)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudent.Name).
		Return(true, nil)

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudent)

	assert.EqualValues(t, dtos.NewConflictResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenAnErrorIsReturnedByTheNextDecorator(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute, 1, f.fakeTracerAttributes)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudent.Name).
		Return(false, nil)

	builder.Next.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudent).
		Return(dtos.NewInternalServerErrorResult())

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudent)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenTheNextDecoratorReturnsSuccess(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentCreationUseCaseValidationsExecute, 1, f.fakeTracerAttributes)

	builder.StudentRepositoryMock.
		EXPECT().
		ExistsByName(builder.Ctx, f.fakeStudent.Name).
		Return(false, nil)

	builder.Next.
		EXPECT().
		Execute(builder.Ctx, f.fakeStudent).
		Return(dtos.NewCreatedResult(f.fakeStudent))

	sut := builder.BuildStudentCreationWithValidations()

	result := sut.Execute(builder.Ctx, f.fakeStudent)

	assert.EqualValues(t, dtos.NewCreatedResult(f.fakeStudent), result)
}
