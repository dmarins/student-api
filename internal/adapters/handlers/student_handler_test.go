package handlers_test

import (
	"net/http"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
)

func TestStudentHandler_Create_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Create_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.POST("/v1/studentssss").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Create_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/v1/students").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Create_WithErrorBind(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON("/{}").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Create_WithErrorValidation(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentCreateInputInvalid).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Create_WithStudentAlreadyExists(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentCreateInputStored).
		Expect().
		Status(http.StatusConflict).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewConflictResult().Message)
}

func TestStudentHandler_Create_WhenTheStudentsIsCreated(t *testing.T) {
	e := builder.Build(t)

	response := e.POST("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentCreateInput).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object()

	response.Value("message").IsEqual("The registration was completed successfully.")
	response.Value("data").Object().Value("id").String().NotEmpty()
	response.Value("data").Object().Value("name").IsEqual(f.fakeStudentCreateInput.Name)
}

func TestStudentHandler_Read_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Read_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.GET("/v1/studentssss").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Read_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Read_WhenStudentIsFound(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students/06b2ec25-3fe0-475e-9077-e77a113f4727").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(f.fakeStudent).Message)
	response.Value("data").Object().Value("id").IsEqual(f.fakeStudent.ID)
	response.Value("data").Object().Value("name").IsEqual(f.fakeStudent.Name)
}

func TestStudentHandler_Read_WhenStudentIsNotFound(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students/58ecde02-18f6-4896-a716-64abf6724587").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewNotFoundResult().Message)
}

func TestStudentHandler_Update_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/v1/students/8e99273f-e566-4476-836e-048b1ecd9c4d").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Update_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.PUT("/v1/studentssss/dbf54856-9a98-4672-9c90-e9da71a1f893").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Update_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.PUT("/v1/students/dbf54856-9a98-4672-9c90-e9da71a1f893").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Update_WithErrorBind(t *testing.T) {
	e := builder.Build(t)

	response := e.PUT("/v1/students/dbf54856-9a98-4672-9c90-e9da71a1f893").
		WithHeader("x-tenant", "sbrubles").
		WithJSON("/{}").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Update_WithErrorValidation(t *testing.T) {
	e := builder.Build(t)

	response := e.PUT("/v1/students/dbf54856-9a98-4672-9c90-e9da71a1f893").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentUpdateInputInvalid).
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Update_WithStudentAlreadyExists(t *testing.T) {
	e := builder.Build(t)

	response := e.PUT("/v1/students/06b2ec25-3fe0-475e-9077-e77a113f4727").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentUpdateInputStored).
		Expect().
		Status(http.StatusConflict).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewConflictResult().Message)
}

func TestStudentHandler_Update_WhenTheStudentsIsUpdated(t *testing.T) {
	e := builder.Build(t)

	response := e.PUT("/v1/students/e6e84c46-6ddf-4d9a-b27a-ddb74b4d63bb").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentUpdateInput).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual("The operation was performed successfully.")
	response.Value("data").Object().Value("id").String().NotEmpty()
	response.Value("data").Object().Value("name").IsEqual(f.fakeStudentUpdateInput.Name)
}

func TestStudentHandler_Delete_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/v1/students/8e99273f-e566-4476-836e-048b1ecd9c4d").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Delete_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.DELETE("/v1/studentssss/8e99273f-e566-4476-836e-048b1ecd9c4d").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Delete_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.DELETE("/v1/students/8e99273f-e566-4476-836e-048b1ecd9c4d").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Delete_WithErrorValidation(t *testing.T) {
	e := builder.Build(t)

	response := e.DELETE("/v1/students/1").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Delete_WhenTheStudentsIsDeleted(t *testing.T) {
	e := builder.Build(t)

	e.DELETE("/v1/students/8e99273f-e566-4476-836e-048b1ecd9c4d").
		WithHeader("x-tenant", "sbrubles").
		WithJSON(f.fakeStudentToBeDeleted).
		Expect().
		Status(http.StatusNoContent)
}

func TestStudentHandler_Search_WithWrongMethod(t *testing.T) {
	e := builder.Build(t)

	e.PATCH("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusMethodNotAllowed)
}

func TestStudentHandler_Search_WithWrongPath(t *testing.T) {
	e := builder.Build(t)

	e.GET("/v1/studentssss").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusNotFound)
}

func TestStudentHandler_Search_WhenTenantIsNotSent(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Search_NoPaginationAndNoFilters(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPageOnly(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPageSizeOnly(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("pageSize", 10).
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPageAndPageSize(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPaginationErrorBind(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", "abc").
		WithQuery("pageSize", "def").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusBadRequest).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewBadRequestResult().Message)
}

func TestStudentHandler_Search_IncreasingPagination(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithQuery("pageSize", 3).
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(3)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(3)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPageAndPageSizeEqualToZero(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 0).
		WithQuery("pageSize", 0).
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPageAndPageSizeNegatives(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", -1).
		WithQuery("pageSize", -1).
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPaginationAndSortOrderOnly(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithQuery("sortOrder", "desc").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPaginationAndSortFieldOnly(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithQuery("sortField", "name").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPaginationAndOrdination(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithQuery("sortOrder", "desc").
		WithQuery("sortField", "name").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(7)
}

func TestStudentHandler_Search_WithPaginationAndOrdinationAndFilterByName(t *testing.T) {
	e := builder.Build(t)

	response := e.GET("/v1/students").
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithQuery("sortOrder", "desc").
		WithQuery("sortField", "name").
		WithQuery("name", "thompson").
		WithHeader("x-tenant", "sbrubles").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	response.Value("message").IsEqual(dtos.NewOkResult(nil).Message)
	response.Value("data").Object().Value("total_pages").IsEqual(1)
	response.Value("data").Object().Value("current_page").IsEqual(1)
	response.Value("data").Object().Value("page_size").IsEqual(10)
	response.Value("data").Object().Value("total_items").IsEqual(2)
}
