package read_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
	"github.com/stretchr/testify/assert"
)

func TestStudentSearchWithSearchBy_Execute_WhenCountFails(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentSearchUseCaseSearchByExecute).
		SettingLoggerErrorBehavior("error counting students", f.fakeError)

	pagination := dtos.PaginationRequest{
		Page:      1,
		PageSize:  10,
		SortOrder: "asc",
		SortField: "name",
	}

	filter := dtos.Filter{
		Name: tests.ToPointer("thompson"),
	}

	builder.StudentRepositoryMock.
		EXPECT().
		Count(builder.Ctx, filter).
		Return(0, f.fakeError)

	sut := builder.BuildStudentSearchWithSearchBy()

	result, err := sut.Execute(builder.Ctx, pagination, filter)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestStudentSearchWithSearchBy_Execute_WhenCountReturnsZero(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentSearchUseCaseSearchByExecute)

	pagination := dtos.PaginationRequest{
		Page:      1,
		PageSize:  10,
		SortOrder: "asc",
		SortField: "name",
	}

	filter := dtos.Filter{
		Name: tests.ToPointer("thompson"),
	}

	builder.StudentRepositoryMock.
		EXPECT().
		Count(builder.Ctx, filter).
		Return(0, nil)

	builder.StudentRepositoryMock.
		EXPECT().
		SearchBy(builder.Ctx, pagination, filter).
		Times(0)

	sut := builder.BuildStudentSearchWithSearchBy()

	result, err := sut.Execute(builder.Ctx, pagination, filter)

	assert.Nil(t, err)
	assert.EqualValues(t, result,
		&dtos.PaginationResponse{
			TotalPages:  1,
			CurrentPage: 1,
			PageSize:    10,
			TotalItems:  0,
			Items:       nil,
		})
}

func TestStudentSearchWithSearchBy_Execute_WhenSearchByFails(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentSearchUseCaseSearchByExecute).
		SettingLoggerErrorBehavior("error searching students", f.fakeError)

	pagination := dtos.PaginationRequest{
		Page:      1,
		PageSize:  10,
		SortOrder: "asc",
		SortField: "name",
	}

	filter := dtos.Filter{
		Name: tests.ToPointer("thompson"),
	}

	builder.StudentRepositoryMock.
		EXPECT().
		Count(builder.Ctx, filter).
		Return(2, nil)

	builder.StudentRepositoryMock.
		EXPECT().
		SearchBy(builder.Ctx, pagination, filter).
		Return(nil, f.fakeError)

	sut := builder.BuildStudentSearchWithSearchBy()

	result, err := sut.Execute(builder.Ctx, pagination, filter)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestStudentSearchWithSearchBy_Execute_WhenSearchByReturnsAsExpected(t *testing.T) {
	builder := tests.NewUnitTestsBuilder(t).
		WithValidCtx().
		SettingTracerBehavior(tracer.StudentSearchUseCaseSearchByExecute)

	pagination := dtos.PaginationRequest{
		Page:      1,
		PageSize:  10,
		SortOrder: "asc",
		SortField: "name",
	}

	filter := dtos.Filter{
		Name: tests.ToPointer("thompson"),
	}

	builder.StudentRepositoryMock.
		EXPECT().
		Count(builder.Ctx, filter).
		Return(2, nil)

	builder.StudentRepositoryMock.
		EXPECT().
		SearchBy(builder.Ctx, pagination, filter).
		Return([]*entities.Student{
			{
				ID:   "a",
				Name: "student 1",
			},
			{
				ID:   "b",
				Name: "student 2",
			},
		}, nil)

	sut := builder.BuildStudentSearchWithSearchBy()

	result, err := sut.Execute(builder.Ctx, pagination, filter)

	assert.Nil(t, err)
	assert.EqualValues(t, result,
		&dtos.PaginationResponse{
			TotalPages:  1,
			CurrentPage: 1,
			PageSize:    10,
			TotalItems:  2,
			Items: []*entities.Student{
				{
					ID:   "a",
					Name: "student 1",
				},
				{
					ID:   "b",
					Name: "student 2",
				},
			},
		})
}
