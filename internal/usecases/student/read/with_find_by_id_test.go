package read_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentReadWithFindById_Execute_WhenTheRepositoryFailsToFindTheStudent(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentReadUseCaseFindByIdExecute).
		SettingLoggerErrorBehavior("error finding student by id", f.fakeError, "id", f.fakeStudent.ID)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudent.ID).
		Return(nil, f.fakeError)

	sut := builder.BuildStudentReadWithFindByID()

	result := sut.Execute(builder.Ctx, f.fakeStudent.ID)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentReadWithFindById_Execute_WhenTheStudentIsNotFound(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentReadUseCaseFindByIdExecute).
		SettingLoggerWarnBehavior("the student was not found according to the id provided", "id", f.fakeStudent.ID)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudent.ID).
		Return(nil, nil)

	sut := builder.BuildStudentReadWithFindByID()

	result := sut.Execute(builder.Ctx, f.fakeStudent.ID)

	assert.EqualValues(t, dtos.NewNotFoundResult(), result)
}

func TestStudentReadWithFindById_Execute_WhenTheStudentIsFound(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentReadUseCaseFindByIdExecute)

	builder.StudentRepositoryMock.
		EXPECT().
		FindById(builder.Ctx, f.fakeStudent.ID).
		Return(&f.fakeStudent, nil)

	sut := builder.BuildStudentReadWithFindByID()

	result := sut.Execute(builder.Ctx, f.fakeStudent.ID)

	assert.EqualValues(t, dtos.NewOkResult(result.Data), result)
}
