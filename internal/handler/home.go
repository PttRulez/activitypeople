package handler

import (
	"antiscoof/internal/utils"
	"antiscoof/internal/view/home"
	"net/http"
)

func (c *HomeController) HandlerHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := utils.GetUserFromContext(r)

	if !user.LoggedIn {
		utils.HtmxRedirect(w, r, "/login")
	}
	return render(r, w, home.Index(c.StravaOAuthLink))
}

type HomeController struct {
	StravaOAuthLink string
}

func NewHomeController(StravaOAuthLink string) *HomeController {
	return &HomeController{
		StravaOAuthLink: StravaOAuthLink,
	}
}
