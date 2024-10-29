package repositories_test

import (
	"testing"

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
