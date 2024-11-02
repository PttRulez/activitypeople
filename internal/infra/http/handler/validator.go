package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(target any) (bool, map[string]string) {
	m := make(map[string]string)
	validate := validator.New()
	if err := validate.Struct(target); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			m[e.Field()] = getValErrMessage(e)
		}
	}
	return len(m) == 0, m
}

func getValErrMessage(err validator.FieldError) string {
	switch err.ActualTag() {
	case "required":
		return fmt.Sprintf("Поле %s обязательно для заполнения", err.Field())
	case "email":
		return fmt.Sprintf("Поле %s должно быть валидным email'ом", err.Field())
	default:
		return err.Error()
		// return fmt.Sprintf("Поле %s должно неверно заполнено", err.Field())
	}
}
