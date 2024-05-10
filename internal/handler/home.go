package handler

import (
	"antiscoof/internal/view/home"
	"net/http"
)

func HandlerHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := getUserFromContext(r)
	if !user.LoggedIn {
		htmxRedirect(w, r, "/login")
	}
	return render(r, w, home.Index())
}
