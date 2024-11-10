package delete_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/entities"
)

type fakeValues struct {
	fakeStudent            entities.Student
	fakeError              error
	fakeStudentToBeDeleted string
}

var f *fakeValues

// TestMain will run before all the tests in the package creation_test
func TestMain(m *testing.M) {
	// Setup: Creating fake values
	fakeStudent := entities.Student{
		ID: "58ecde02-18f6-4896-a716-64abf6724587",
	}

	f = &fakeValues{
		fakeStudent:            fakeStudent,
		fakeError:              errors.New("fails"),
		fakeStudentToBeDeleted: "8e99273f-e566-4476-836e-048b1ecd9c4d",
	}

	// Run the all tests in the package creation_test
	code := m.Run()

	// Exit with the code returned by the tests
	os.Exit(code)
}
