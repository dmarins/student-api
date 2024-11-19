package create_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

type fakeValues struct {
	fakeStudentCreateInput dtos.StudentCreateInput
	fakeError              error
}

var f *fakeValues

// TestMain will run before all the tests in the package creation_test
func TestMain(m *testing.M) {
	// Setup: Creating fake values
	fakeStudentCreateInput := dtos.StudentCreateInput{
		Name: "John Doe",
	}

	f = &fakeValues{
		fakeStudentCreateInput: fakeStudentCreateInput,
		fakeError:              errors.New("fails"),
	}

	// Run the all tests in the package creation_test
	code := m.Run()

	// Exit with the code returned by the tests
	os.Exit(code)
}
