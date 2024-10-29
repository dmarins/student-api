package handlers_test

import (
	"net/http"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

func TestStudentHandler_Post_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/students").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Post_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.POST("/studentssss").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Post_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/students").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Post_WithErrorBind(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON("/{}").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Post_WithErrorValidation(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeInvalidInputStudent).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Post_WithStudentAlreadyExists(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStoredInputStudent).
		Expect().
		Status(http.StatusConflict).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewConflictResult().Message)
}

func TestStudentHandler_Post_WhenTheStudentsIsCreated(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeInputStudent).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object()

	response.Value("message").IsEqual("The registration was completed successfully.")
	response.Value("data").Object().Value("id").String().NotEmpty()
	response.Value("data").Object().Value("name").IsEqual(f.fakeInputStudent.Name)
}

func TestStudentHandler_Get_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/students").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Get_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.GET("/studentssss").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Get_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/students").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Get_WhenStudentIsFound(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/students/06b2ec25-3fe0-475e-9077-e77a113f4727").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(f.fakeStudent).Message)
	response.Value("data").Object().Value("id").IsEqual(f.fakeStudent.ID)
	response.Value("data").Object().Value("name").IsEqual(f.fakeStudent.Name)
}

func TestStudentHandler_Get_WhenStudentIsNotFound(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/students/58ecde02-18f6-4896-a716-64abf6724587").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewNotFoundResult().Message)
}
