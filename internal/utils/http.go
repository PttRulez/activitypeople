package utils

import (
	"antiscoof/internal/model"
	"net/http"
)

func HtmxRedirect(w http.ResponseWriter, r *http.Request, url string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusSeeOther)
		return nil
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
	return nil
}

func GetUserFromContext(r *http.Request) model.AuthenticatedUser {
	user, ok := r.Context().Value(model.UserContextKey).(model.AuthenticatedUser)
	if !ok {
		return model.AuthenticatedUser{}
	}
	return user
}
