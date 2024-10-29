package create_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type fakeValues struct {
	fakeStudentInput dtos.StudentInput
	fakeError        error
}

var f *fakeValues

// TestMain will run before all the tests in the package creation_test
func TestMain(m *testing.M) {
	// Setup: Creating fake values
	fakeStudentInput := dtos.StudentInput{
		Name: "John Doe",
	}

	f = &fakeValues{
		fakeStudentInput: fakeStudentInput,
		fakeError:        errors.New("fails"),
	}

	// Run the all tests in the package creation_test
	code := m.Run()

	// Exit with the code returned by the tests
	os.Exit(code)
}
