package handler

import (
	"fmt"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store/pgstore"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
)

func (c *StravaController) StravaOAuthCallback(w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")
	if code == "" {
		return HtmxRedirect(w, r, "/")
	}
	data, err := c.stravaApi.OAuth(code)
	if err != nil {
		fmt.Printf("StravaOAuthCallback error: %s\n", err)
		return HtmxRedirect(w, r, "/")
	}
	user := GetUserFromRequest(r)
	err = c.stravaRepo.Insert(r.Context(), &domain.StravaInfo{
		AccessToken:  &data.AccessToken,
		RefreshToken: &data.RefreshToken,
		UserId:       user.Id,
	})
	if err != nil {
		fmt.Printf("StravaOAuthCallback error: %s\n", err)
		return HtmxRedirect(w, r, "/")
	}

	return HtmxRedirect(w, r, "/")
}

type StravaController struct {
	stravaApi  *strava.StravaApi
	stravaRepo *pgstore.StravaPostgres
}

func NewStravaController(
	stravaApi *strava.StravaApi,
	stravaRepo *pgstore.StravaPostgres,
) *StravaController {
	return &StravaController{
		stravaApi:  stravaApi,
		stravaRepo: stravaRepo,
	}
}
