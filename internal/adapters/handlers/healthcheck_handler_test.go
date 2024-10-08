package handlers_test

import (
	"net/http"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

func TestHealthCheckHandler_Get_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/health").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestHealthCheckHandler_Get_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.GET("/healthhhh").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestHealthCheckHandler_Get_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/health").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestHealthCheckHandler_Get_WhenReturnsHealthy(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/health").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual("healthy")
}
