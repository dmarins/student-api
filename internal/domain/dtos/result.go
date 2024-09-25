package dtos

import "net/http"

type Result struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Errors     error  `json:"errors,omitempty"`
}

func newSuccessResult(statusCode int, data any, message string) *Result {
	return &Result{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Errors:     nil,
	}
}

func newWarningResult(statusCode int, message string) *Result {
	return &Result{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
		Errors:     nil,
	}
}

func newErrorResult(statusCode int, errMessage string, errors error) *Result {
	return &Result{
		StatusCode: statusCode,
		Message:    errMessage,
		Data:       nil,
		Errors:     errors,
	}
}

func NewHttpStatusCreatedResult(data any) *Result {
	return newSuccessResult(
		http.StatusCreated,
		data,
		"The registration was completed successfully.",
	)
}

func NewHttpStatusOkResult(data any) *Result {
	return newSuccessResult(
		http.StatusOK,
		data,
		"The operation was performed successfully.",
	)
}

func NewHttpStatusBadRequestResult() *Result {
	return newWarningResult(
		http.StatusBadRequest,
		"The request does not meet the expected format. Please check the data and try again.",
	)
}

func NewHttpStatusNotFoundResult() *Result {
	return newWarningResult(
		http.StatusNotFound,
		"The target was not found.",
	)
}

func NewHttpStatusConflictResult() *Result {
	return newWarningResult(
		http.StatusConflict,
		"The target already exists.",
	)
}

func NewHttpStatusInternalServerErrorResult(err error) *Result {
	return newErrorResult(
		http.StatusInternalServerError,
		"Sorry, something went wrong in our system. Please try again later.",
		err,
	)
}
