package handler

import (
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
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

func GetUserFromRequest(r *http.Request) domain.AuthenticatedUser {
	user, ok := r.Context().Value(domain.UserContextKey).(domain.AuthenticatedUser)
	if !ok {
		return domain.AuthenticatedUser{}
	}
	return user
}
