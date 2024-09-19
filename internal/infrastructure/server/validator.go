package server

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewValidator() *CustomValidator {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("json"), ",")[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return &CustomValidator{validate: validate}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validate.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		for _, ve := range validationErrors {
			errorMessages = append(errorMessages, ve.Field()+" is invalid: "+ve.Tag())
		}

		return echo.NewHTTPError(http.StatusBadRequest, strings.Join(errorMessages, ", "))
	}

	return nil
}
