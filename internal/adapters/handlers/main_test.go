package handlers_test

import (
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
)

type fakeValues struct {
	fakeStudentCreateInput        *dtos.StudentCreateInput
	fakeStudentUpdateInput        *dtos.StudentUpdateInput
	fakeStudentCreateInputInvalid *dtos.StudentCreateInput
	fakeStudentUpdateInputInvalid *dtos.StudentUpdateInput
	fakeStudentCreateInputStored  *dtos.StudentCreateInput
	fakeStudentUpdateInputStored  *dtos.StudentUpdateInput
	fakeStudent                   *entities.Student
	fakeStudentToBeDeleted        string
}

var f *fakeValues
var builder *tests.E2eTestsBuilder

// TestMain will run before all the tests in the package handlers_test
func TestMain(m *testing.M) {
	// Setup: Create the builder and initialize the container
	builder = tests.
		NewE2eTestsBuilder().
		StartCompositionRoot().
		StartTestServer()

	f = &fakeValues{
		fakeStudentCreateInput: &dtos.StudentCreateInput{
			Name: "john doe",
		},
		fakeStudentCreateInputInvalid: &dtos.StudentCreateInput{
			Name: "",
		},
		fakeStudentCreateInputStored: &dtos.StudentCreateInput{
			Name: "alice",
		},
		fakeStudent: &entities.Student{
			ID:   "06b2ec25-3fe0-475e-9077-e77a113f4727",
			Name: "alice",
		},
		fakeStudentUpdateInput: &dtos.StudentUpdateInput{
			Name: "ashley updated",
		},
		fakeStudentUpdateInputInvalid: &dtos.StudentUpdateInput{
			Name: "",
		},
		fakeStudentUpdateInputStored: &dtos.StudentUpdateInput{
			Name: "alice",
		},
		fakeStudentToBeDeleted: "8e99273f-e566-4476-836e-048b1ecd9c4d",
	}

	// Run all tests in handlers_test
	code := m.Run()

	// Teardown: Stop all after all tests
	func() {
		builder.App.Stop(builder.Ctx)
		builder.TestServer.Close()
	}()

	// Exit with the code returned by the tests
	os.Exit(code)
}
