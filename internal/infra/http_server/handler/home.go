package handler

import (
	"net/http"

	"github.com/pttrulez/activitypeople/internal/infra/view/home"
)

func (c *HomeController) HandlerHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, home.Index(c.StravaOAuthLink, GetUserFromRequest(r)))
}

type HomeController struct {
	StravaOAuthLink string
}

func NewHomeController(StravaOAuthLink string) *HomeController {
	return &HomeController{
		StravaOAuthLink: StravaOAuthLink,
	}
}
