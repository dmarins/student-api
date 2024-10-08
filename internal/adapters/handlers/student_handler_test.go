package handlers_test

import (
	"net/http"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

func TestStudentHandler_Post_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/student").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Post_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.POST("/studentttt").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Post_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/student").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Post_WithErrorBind(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/student").
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

	response := e.POST("/student").
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

	response := e.POST("/student").
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

	response := e.POST("/student").
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
