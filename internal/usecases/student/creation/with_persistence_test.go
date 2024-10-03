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

type studentCreationWithPersistenceTestBuilder struct {
	ctrl                  *gomock.Controller
	tracerMock            *mocks.MockITracer
	spanMock              *mocks.MockISpan
	loggerMock            *mocks.MockILogger
	studentRepositoryMock *mocks.MockIStudentRepository
	nextStepOfUseCaseMock *mocks.MockIStudentCreationUseCase
	ctx                   context.Context
	fakeStudent           entities.Student
	fakeError             error
	fakeOutputStudent     *dtos.StudentOutput
}

func newstudentCreationWithPersistenceTestBuilder(t *testing.T) *studentCreationWithPersistenceTestBuilder {
	ctx, ctrl, tracerMock, loggerMock := tests.SetupTest(t)
	spanMock := mocks.NewMockISpan(ctrl)
	studentRepositoryMock := mocks.NewMockIStudentRepository(ctrl)

	return &studentCreationWithPersistenceTestBuilder{
		ctrl:                  ctrl,
		tracerMock:            tracerMock,
		loggerMock:            loggerMock,
		spanMock:              spanMock,
		studentRepositoryMock: studentRepositoryMock,
		ctx:                   ctx,
		fakeStudent:           entities.Student{Name: "John Doe"},
		fakeError:             errors.New("fail"),
		fakeOutputStudent:     &dtos.StudentOutput{Name: "John Doe"},
	}
}

func (b *studentCreationWithPersistenceTestBuilder) withTracerMock() *studentCreationWithPersistenceTestBuilder {
	b.tracerMock.
		EXPECT().
		NewSpanContext(b.ctx, tracer.StudentCreationUseCasePersistenceExecute).
		Return(b.spanMock, b.ctx).
		Times(1)

	b.spanMock.
		EXPECT().
		End().
		Times(1)

	b.tracerMock.
		EXPECT().
		AddAttributes(b.spanMock, tracer.StudentCreationUseCasePersistenceExecute,
			tracer.Attributes{
				"Entity": b.fakeStudent,
			}).
		Times(1)

	return b
}

func (b *studentCreationWithPersistenceTestBuilder) withLoggerDebug(debugMessage string, hasArgs bool) *studentCreationWithPersistenceTestBuilder {
	if hasArgs {
		b.loggerMock.
			EXPECT().
			Debug(b.ctx, debugMessage, "id", gomock.Any())

		return b
	}

	b.loggerMock.
		EXPECT().
		Debug(b.ctx, debugMessage)

	return b
}

func (b *studentCreationWithPersistenceTestBuilder) whereAddFails() *studentCreationWithPersistenceTestBuilder {
	b.studentRepositoryMock.
		EXPECT().
		Add(b.ctx, gomock.Any()).
		Return(b.fakeError)

	return b
}

func (b *studentCreationWithPersistenceTestBuilder) whereAddNoFails() *studentCreationWithPersistenceTestBuilder {
	b.studentRepositoryMock.
		EXPECT().
		Add(b.ctx, gomock.Any()).
		Return(nil)

	return b
}

func (b *studentCreationWithPersistenceTestBuilder) withLoggerError() *studentCreationWithPersistenceTestBuilder {
	b.loggerMock.
		EXPECT().
		Error(b.ctx, "error adding a new student", b.fakeError)

	return b
}

func (b *studentCreationWithPersistenceTestBuilder) build() usecases.IStudentCreationUseCase {
	return creation.NewStudentCreationWithPersistence(b.tracerMock, b.loggerMock, b.studentRepositoryMock)
}

func TestStudentCreationWithPersistence_Execute_WhenAddFails(t *testing.T) {
	builder := newstudentCreationWithPersistenceTestBuilder(t).
		withTracerMock().
		withLoggerDebug("new student", true).
		whereAddFails().
		withLoggerError()

	sut := builder.build()

	result := sut.Execute(builder.ctx, builder.fakeStudent)

	assert.EqualValues(t, dtos.NewInternalServerErrorResult(), result)
}

func TestStudentCreationWithPersistence_Execute_Success(t *testing.T) {
	builder := newstudentCreationWithPersistenceTestBuilder(t).
		withTracerMock().
		withLoggerDebug("new student", true).
		whereAddNoFails().
		withLoggerDebug("student stored", false)

	sut := builder.build()

	result := sut.Execute(builder.ctx, builder.fakeStudent)

	assert.EqualValues(t, dtos.NewCreatedResult(result.Data), result)
}
