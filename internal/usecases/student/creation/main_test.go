package creation_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/infrastructure/tracer"
)

type fakeValues struct {
	fakeStudent          entities.Student
	fakeError            error
	fakeTracerAttributes tracer.Attributes
}

var f *fakeValues

// TestMain will run before all the tests in the package creation_test
func TestMain(m *testing.M) {
	// Setup: Creating fake values
	fakeStudent := entities.Student{
		Name: "John Doe",
	}

	f = &fakeValues{
		fakeStudent: fakeStudent,
		fakeError:   errors.New("fails"),
		fakeTracerAttributes: tracer.Attributes{
			"Entity": fakeStudent,
		},
	}

	// Run the all tests in the package creation_test
	code := m.Run()

	// Exit with the code returned by the tests
	os.Exit(code)
}
