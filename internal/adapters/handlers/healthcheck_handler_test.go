package handlers_test

import (
	"net/http"
	"testing"
)

func TestHealthCheckHandler_Get_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/health").
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

func TestHealthCheckHandler_Get_WhenReturnsHealthy(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/health").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual("healthy")
}
