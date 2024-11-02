package handler

import (
	"context"
	"fmt"
	"net/http"
)

func (c *StravaController) StravaOAuthCallback(w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")
	if code == "" {
		return HtmxRedirect(w, r, "/")
	}

	user := GetUserFromRequest(r)
	err := c.stravaService.OAuthStrava(r.Context(), code, user.Id)
	if err != nil {
		fmt.Printf("StravaOAuthCallback error: %s\n", err)
		return HtmxRedirect(w, r, "/")
	}

	return HtmxRedirect(w, r, "/")
}

type StravaController struct {
	stravaService StravaService
}

type StravaService interface {
	OAuthStrava(ctx context.Context, userCode string, userID int) error
}

func NewStravaController(
	stravaService StravaService,
) *StravaController {
	return &StravaController{
		stravaService: stravaService,
	}
}
