package handlers_test

import (
	"os"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/infrastructure/tests"
)

type fakeValues struct {
	fakeInputStudent        *dtos.StudentInput
	fakeInvalidInputStudent *dtos.StudentInput
	fakeStoredInputStudent  *dtos.StudentInput
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
		fakeInputStudent: &dtos.StudentInput{
			Name: "john doe",
		},
		fakeInvalidInputStudent: &dtos.StudentInput{
			Name: "",
		},
		fakeStoredInputStudent: &dtos.StudentInput{
			Name: "alice",
		},
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
