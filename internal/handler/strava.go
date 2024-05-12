package handler

import (
	"antiscoof/internal/model"
	stravaclient "antiscoof/internal/service/strava-client"
	"antiscoof/internal/store/pgstore"
	"antiscoof/internal/utils"
	"fmt"
	"net/http"
)

func (c *StravaController) StravaOAuthCallback(w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")
	if code == "" {
		return utils.HtmxRedirect(w, r, "/")
	}
	data, err := c.stravaApi.OAuth(code)
	if err != nil {
		fmt.Printf("StravaOAuthCallback error: %s\n", err)
		return utils.HtmxRedirect(w, r, "/")
	}
	user := utils.GetUserFromContext(r)
	err = c.stravaRepo.Insert(r.Context(), &model.StravaInfo{
		AccessToken:  &data.AccessToken,
		RefreshToken: &data.RefreshToken,
		UserId:       user.Id,
	})
	if err != nil {
		fmt.Printf("StravaOAuthCallback error: %s\n", err)
		return utils.HtmxRedirect(w, r, "/")
	}

	return utils.HtmxRedirect(w, r, "/")
}

type StravaController struct {
	stravaApi  *stravaclient.StravaApi
	stravaRepo *pgstore.StravaPostgres
}

func NewStravaController(
	stravaApi *stravaclient.StravaApi,
	stravaRepo *pgstore.StravaPostgres,
) *StravaController {
	return &StravaController{
		stravaApi:  stravaApi,
		stravaRepo: stravaRepo,
	}
}
