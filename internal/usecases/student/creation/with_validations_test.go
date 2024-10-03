package creation_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/mocks"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/dmarins/student-api/internal/usecases/student/creation"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type studentCreationWithValidationsTestBuilder struct {
	ctrl                  *gomock.Controller
	tracerMock            *mocks.MockITracer
	spanMock              *mocks.MockISpan
	loggerMock            *mocks.MockILogger
	studentRepositoryMock *mocks.MockIStudentRepository
	nextStepOfUseCaseMock *mocks.MockIStudentCreationUseCase
	ctx                   context.Context
	fakeStudent           entities.Student
	fakeError             error
}

func newstudentCreationWithValidationsTestBuilder(t *testing.T) *studentCreationWithValidationsTestBuilder {
	ctx, ctrl, tracerMock, loggerMock := tests.SetupTest(t)
	spanMock := mocks.NewMockISpan(ctrl)
	studentRepositoryMock := mocks.NewMockIStudentRepository(ctrl)
	nextStepOfUseCaseMock := mocks.NewMockIStudentCreationUseCase(ctrl)

	return &studentCreationWithValidationsTestBuilder{
		ctrl:                  ctrl,
		tracerMock:            tracerMock,
		loggerMock:            loggerMock,
		spanMock:              spanMock,
		studentRepositoryMock: studentRepositoryMock,
		nextStepOfUseCaseMock: nextStepOfUseCaseMock,
		ctx:                   ctx,
		fakeStudent:           entities.Student{Name: "John Doe"},
		fakeError:             errors.New("fail"),
	}
}

func (b *studentCreationWithValidationsTestBuilder) withTracerMock() *studentCreationWithValidationsTestBuilder {
	b.tracerMock.
		EXPECT().
		NewSpanContext(b.ctx, tracer.StudentCreationUseCaseValidationsExecute).
		Return(b.spanMock, b.ctx).
		Times(1)

	b.spanMock.
		EXPECT().
		End().
		Times(1)

	b.tracerMock.
		EXPECT().
		AddAttributes(b.spanMock, tracer.StudentCreationUseCaseValidationsExecute,
			tracer.Attributes{
				"Entity": b.fakeStudent,
			}).
		Times(1)

	return b
}

func (b *studentCreationWithValidationsTestBuilder) whereExistsByNameFails() *studentCreationWithValidationsTestBuilder {
	b.studentRepositoryMock.
		EXPECT().
		ExistsByName(b.ctx, b.fakeStudent.Name).
		Return(false, b.fakeError)

	return b
}

func (b *studentCreationWithValidationsTestBuilder) whereExistsByNameReturnsTrue() *studentCreationWithValidationsTestBuilder {
	b.studentRepositoryMock.
		EXPECT().
		ExistsByName(b.ctx, b.fakeStudent.Name).
		Return(true, nil)

	return b
}

func (b *studentCreationWithValidationsTestBuilder) whereExistsByNameReturnsFalse() *studentCreationWithValidationsTestBuilder {
	b.studentRepositoryMock.
		EXPECT().
		ExistsByName(b.ctx, b.fakeStudent.Name).
		Return(false, nil)

	return b
}

func (b *studentCreationWithValidationsTestBuilder) withLoggerError() *studentCreationWithValidationsTestBuilder {
	b.loggerMock.
		EXPECT().
		Error(b.ctx, "error checking if student exists", b.fakeError, "name", b.fakeStudent.Name)

	return b
}

func (b *studentCreationWithValidationsTestBuilder) withLoggerWarn() *studentCreationWithValidationsTestBuilder {
	b.loggerMock.
		EXPECT().
		Warn(b.ctx, "there is already a student with the same name", "name", b.fakeStudent.Name)

	return b
}

func (b *studentCreationWithValidationsTestBuilder) withNextStepFails() *studentCreationWithValidationsTestBuilder {
	b.nextStepOfUseCaseMock.
		EXPECT().
		Execute(b.ctx, b.fakeStudent).
		Return(dtos.NewInternalServerErrorResult())

	return b
}

func (b *studentCreationWithValidationsTestBuilder) withNextStepReturnSuccess() *studentCreationWithValidationsTestBuilder {
	b.nextStepOfUseCaseMock.
		EXPECT().
		Execute(b.ctx, b.fakeStudent).
		Return(dtos.NewCreatedResult(b.fakeStudent))

	return b
}

func (b *studentCreationWithValidationsTestBuilder) build() usecases.IStudentCreationUseCase {
	return creation.NewStudentCreationWithValidations(b.tracerMock, b.loggerMock, b.studentRepositoryMock, b.nextStepOfUseCaseMock)
}

func TestStudentCreationWithValidations_Execute_WhenExistsByNameFails(t *testing.T) {
	builder := newstudentCreationWithValidationsTestBuilder(t).
		withTracerMock().
		whereExistsByNameFails().
		withLoggerError()

	sut := builder.build()

	result := sut.Execute(builder.ctx, builder.fakeStudent)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenStudentAlreadyExists(t *testing.T) {
	builder := newstudentCreationWithValidationsTestBuilder(t).
		withTracerMock().
		whereExistsByNameReturnsTrue().
		withLoggerWarn()

	sut := builder.build()

	result := sut.Execute(builder.ctx, builder.fakeStudent)

	assert.EqualValues(t, dtos.NewConflictResult(), result)
}

func TestStudentCreationWithValidations_Execute_WhenNextStepFails(t *testing.T) {
	builder := newstudentCreationWithValidationsTestBuilder(t).
		withTracerMock().
		whereExistsByNameReturnsFalse().
		withNextStepFails()

	sut := builder.build()

	result := sut.Execute(builder.ctx, builder.fakeStudent)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithValidations_Execute_Success(t *testing.T) {
	builder := newstudentCreationWithValidationsTestBuilder(t).
		withTracerMock().
		whereExistsByNameReturnsFalse().
		withNextStepReturnSuccess()

	sut := builder.build()

	result := sut.Execute(builder.ctx, builder.fakeStudent)

	assert.EqualValues(t, dtos.NewCreatedResult(builder.fakeStudent), result)
}
