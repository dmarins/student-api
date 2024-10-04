package repositories_test

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
)

var builder *tests.IntegrationTestsBuilder

// TestMain will run before all the tests in the package
func TestMain(m *testing.M) {
	// Setup: Create the builder and initialize the container
	builder = tests.
		NewIntegrationTestsBuilder().
		WithLogger().
		WithTracer()

	// Run the tests
	code := m.Run()

	// Teardown: Stop the container after all tests
	builder.TearDown()

	// Exit with the code returned by the tests
	os.Exit(code)
}

func TestStudentRepository_Add_Success(t *testing.T) {
	student := &entities.Student{
		ID:   "ef9a11c4-9603-4995-b386-37ed45365eb6",
		Name: "nicole",
	}

	sut := builder.BuildStudentRepository()

	err := sut.Add(builder.GetCtx(), student)

	assert.NoError(t, err)
}

func TestStudentRepository_ExistsByName_ShouldBeReturnsTrue(t *testing.T) {
	sut := builder.BuildStudentRepository()

	exists, err := sut.ExistsByName(builder.GetCtx(), "bob")

	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestStudentRepository_ExistsByName_ShouldBeReturnsFalse(t *testing.T) {
	sut := builder.BuildStudentRepository()

	exists, err := sut.ExistsByName(builder.GetCtx(), "sbrubles")

	assert.NoError(t, err)
	assert.False(t, exists)
}
