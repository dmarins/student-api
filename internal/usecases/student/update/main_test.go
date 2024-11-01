package update_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
)

type fakeValues struct {
	fakeStudentUpdateInput dtos.StudentUpdateInput
	fakeStudent            entities.Student
	fakeError              error
}

var f *fakeValues

// TestMain will run before all the tests in the package creation_test
func TestMain(m *testing.M) {
	// Setup: Creating fake values
	fakeStudentUpdateInput := dtos.StudentUpdateInput{
		Name: "John Doe",
	}
	fakeStudent := entities.Student{
		ID: "58ecde02-18f6-4896-a716-64abf6724587",
	}

	f = &fakeValues{
		fakeStudentUpdateInput: fakeStudentUpdateInput,
		fakeStudent:            fakeStudent,
		fakeError:              errors.New("fails"),
	}

	// Run the all tests in the package creation_test
	code := m.Run()

	// Exit with the code returned by the tests
	os.Exit(code)
}
