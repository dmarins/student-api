package repositories_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestStudentRepository_Add_WhenRepositoryAddsTheStudent(t *testing.T) {
	sut := builder.BuildStudentRepository()

	err := sut.Add(builder.Ctx, f.fakeNewStudent)

	assert.NoError(t, err)
}

func TestStudentRepository_Add_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	err := sut.Add(builder.Ctx, f.fakeNewStudent)

	assert.Error(t, err)
}

func TestStudentRepository_ExistsByName_WhenTheStudentAlreadyExists(t *testing.T) {
	sut := builder.BuildStudentRepository()

	exists, err := sut.ExistsByName(builder.Ctx, f.fakeStoredStudent.Name)

	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestStudentRepository_ExistsByName_WhenTheStudentDoesNotExist(t *testing.T) {
	sut := builder.BuildStudentRepository()

	exists, err := sut.ExistsByName(builder.Ctx, f.fakeNewStudent.Name+"1")

	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestStudentRepository_ExistsByName_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	exists, err := sut.ExistsByName(builder.Ctx, f.fakeStoredStudent.ID)

	assert.Error(t, err)
	assert.False(t, exists)
}

func TestStudentRepository_FindById_WhenTheStudentIsNotFound(t *testing.T) {
	sut := builder.BuildStudentRepository()

	student, err := sut.FindById(builder.Ctx, f.fakeNotFoundStudent.ID)

	assert.NoError(t, err)
	assert.Nil(t, student)
}

func TestStudentRepository_FindById_WhenTheStudentIsFound(t *testing.T) {
	sut := builder.BuildStudentRepository()

	student, err := sut.FindById(builder.Ctx, f.fakeStoredStudent.ID)

	assert.NoError(t, err)
	assert.EqualValues(t, student, f.fakeStoredStudent)
}

func TestStudentRepository_FindById_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	student, err := sut.FindById(builder.Ctx, f.fakeStoredStudent.ID)

	assert.Error(t, err)
	assert.Nil(t, student)
}

func TestStudentRepository_Update_WhenRepositoryUpdatesTheStudent(t *testing.T) {
	sut := builder.BuildStudentRepository()

	err := sut.Update(builder.Ctx, f.fakeNewStudent)

	assert.NoError(t, err)
}

func TestStudentRepository_Update_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	err := sut.Update(builder.Ctx, f.fakeNewStudent)

	assert.Error(t, err)
}

func TestStudentRepository_Delete_WhenRepositoryRemovesTheStudent(t *testing.T) {
	sut := builder.BuildStudentRepository()

	err := sut.Delete(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.NoError(t, err)
}

func TestStudentRepository_Delete_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	err := sut.Delete(builder.Ctx, f.fakeStudentToBeDeleted)

	assert.Error(t, err)
}

func TestStudentRepository_SearchBy_WhenRepositoryReturnsOnePageOfData(t *testing.T) {
	sut := builder.BuildStudentRepository()

	results, err := sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      1,
			PageSize:  10,
			SortOrder: tests.ToPointer("asc"),
			SortField: tests.ToPointer("name"),
		},
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, results[0].Name, "michael thompson")
	assert.Equal(t, results[1].Name, "will thompson")

	results, err = sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      2,
			PageSize:  10,
			SortOrder: tests.ToPointer("asc"),
			SortField: tests.ToPointer("name"),
		},
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 0)
}

func TestStudentRepository_SearchBy_WhenRepositoryReturnsTwoPageOfData(t *testing.T) {
	sut := builder.BuildStudentRepository()

	results, err := sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      1,
			PageSize:  1,
			SortOrder: tests.ToPointer("asc"),
			SortField: tests.ToPointer("name"),
		},
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, results[0].Name, "michael thompson")

	results, err = sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      2,
			PageSize:  1,
			SortOrder: tests.ToPointer("asc"),
			SortField: tests.ToPointer("name"),
		},
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, results[0].Name, "will thompson")
}

func TestStudentRepository_SearchBy_ToSortOrderDescAndSortFieldId(t *testing.T) {
	sut := builder.BuildStudentRepository()

	results, err := sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      1,
			PageSize:  10,
			SortOrder: tests.ToPointer("desc"),
			SortField: tests.ToPointer("id"),
		},
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, results[0].Name, "will thompson")
	assert.Equal(t, results[1].Name, "michael thompson")
}

func TestStudentRepository_SearchBy_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	results, err := sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      1,
			PageSize:  10,
			SortOrder: tests.ToPointer("asc"),
			SortField: tests.ToPointer("name"),
		},
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.Error(t, err)
	assert.Nil(t, results)
}

func TestStudentRepository_Count_WhenRepositoryReturnsCount(t *testing.T) {
	sut := builder.BuildStudentRepository()

	count, err := sut.Count(builder.Ctx,
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, count, 2)
}

func TestStudentRepository_Count_WhenRepositoryReturnsZero(t *testing.T) {
	sut := builder.BuildStudentRepository()

	count, err := sut.Count(builder.Ctx,
		dtos.Filter{
			Name: tests.ToPointer("sbrubles"),
		},
	)

	assert.NoError(t, err)
	assert.Zero(t, count)
}

func TestStudentRepository_Count_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	count, err := sut.Count(builder.Ctx,
		dtos.Filter{
			Name: tests.ToPointer("thompson"),
		},
	)

	assert.Error(t, err)
	assert.Zero(t, count)
}
