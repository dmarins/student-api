package repositories_test

import (
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
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
			SortOrder: "asc",
			SortField: "name",
		},
		dtos.Filter{
			Name: ToPointer("a"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 3)
}

func TestStudentRepository_SearchBy_WhenRepositoryReturnsThreePageOfData(t *testing.T) {
	sut := builder.BuildStudentRepository()

	results, err := sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      1,
			PageSize:  1,
			SortOrder: "asc",
			SortField: "name",
		},
		dtos.Filter{
			Name: ToPointer("a"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, results[0].Name, "alice")

	results, err = sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      2,
			PageSize:  1,
			SortOrder: "asc",
			SortField: "name",
		},
		dtos.Filter{
			Name: ToPointer("a"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, results[0].Name, "ashley")

	results, err = sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      3,
			PageSize:  1,
			SortOrder: "asc",
			SortField: "name",
		},
		dtos.Filter{
			Name: ToPointer("a"),
		},
	)

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, results[0].Name, "megan")
}

func TestStudentRepository_SearchBy_WhenTheQueryFails(t *testing.T) {
	sut := failedBuilder.BuildStudentRepository()

	results, err := sut.SearchBy(builder.Ctx,
		dtos.PaginationRequest{
			Page:      1,
			PageSize:  10,
			SortOrder: "asc",
			SortField: "name",
		},
		dtos.Filter{
			Name: ToPointer("a"),
		},
	)

	assert.Error(t, err)
	assert.Nil(t, results)
}
