package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/pttrulez/activitypeople/internal/domain"
)

func render(r *http.Request, w http.ResponseWriter, component templ.Component) error {
	return component.Render(r.Context(), w)
}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
		}
	}
}

func GetUserFromRequest(r *http.Request) domain.User {
	return GetUserFromContext(r.Context())
}

func GetUserFromContext(ctx context.Context) domain.User {
	user, ok := ctx.Value(userKey).(domain.User)
	if !ok {
		return domain.User{}
	}
	return user
}

func HtmxRedirect(w http.ResponseWriter, r *http.Request, url string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusSeeOther)
		return nil
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
	return nil
}

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
