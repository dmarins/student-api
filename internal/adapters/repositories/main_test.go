package repositories_test

import (
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
)

type fakeValues struct {
	fakeNewStudent    *entities.Student
	fakeStoredStudent *entities.Student
}

var f *fakeValues
var builder *tests.IntegrationTestsBuilder

// TestMain will run before all the tests in the package repositories_test
func TestMain(m *testing.M) {
	// Setup: Create the builder and initialize the container
	builder = tests.
		NewIntegrationTestsBuilder().
		WithLogger().
		WithTracer()

	f = &fakeValues{
		fakeNewStudent: &entities.Student{
			ID:   "ef9a11c4-9603-4995-b386-37ed45365eb6",
			Name: "john doe",
		},
		fakeStoredStudent: &entities.Student{
			ID:   "06b2ec25-3fe0-475e-9077-e77a113f4727",
			Name: "alice",
		},
	}

	// Run all tests in repositories_test
	code := m.Run()

	// Teardown: Stop the container after all tests
	func() {
		builder.DbConn.Close()
		builder.PgContainer.Terminate(builder.Ctx)
	}()

	// Exit with the code returned by the tests
	os.Exit(code)
}