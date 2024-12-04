package httpserver

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	var validateErr validator.ValidationErrors
	if errors.As(err, &validateErr) {
		c.JSON(http.StatusUnprocessableEntity, validationErrsToResponse(validateErr))
		return
	}

	c.String(http.StatusInternalServerError, err.Error())
}

func validationErrsToResponse(errs validator.ValidationErrors) map[string]string {
	mappedErrors := map[string]string{}
	for _, err := range errs {
		fieldName := err.Field()

		switch err.ActualTag() {
		case "required":
			mappedErrors[fieldName] += fmt.Sprintf("Поле %s обязательно для заполнения", err.Field())
		case "email":
			mappedErrors[fieldName] += fmt.Sprintf("Поле %s должно быть валидным email'ом", err.Field())
		default:
			mappedErrors[fieldName] += fmt.Sprintf("Неверно заполнено поле %s", err.Field())
		}
	}

	return mappedErrors
}
