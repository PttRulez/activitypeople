package handler

import (
	"fmt"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/config"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
	"github.com/pttrulez/activitypeople/internal/infra/view/activities"
)

func (c *ActivitiesController) GetActivitiesPage(w http.ResponseWriter, r *http.Request) error {
	user := GetUserFromRequest(r)

	params := strava.CreateStravaApiParams{
		AppClientId:     c.cfg.Strava.ClientID,
		AppClientSecret: c.cfg.Strava.ClientSecret,
		StoreTokensFunc: func(accessToken string, refreshToken string) error {
			// TODO вызывать нужно метод серивиса, а не репы
			return c.stravaStore.UpdateUserStravaInfo(r.Context(), accessToken, refreshToken, user.Id)
		},
		UserAccessToken:  *user.Strava.AccessToken,
		UserRefreshToken: *user.Strava.RefreshToken,
	}

	// создаём старава клиента
	stravaApi := strava.NewStravaApi(params)

	// TODO опеределиться как ходим в страву - сами или через стороннюю либу
	go func() {
		stravaApi.ObaGetAthleteActivities()
	}()

	activitiesInfo, err := stravaApi.GetAthleteActivities()
	if err != nil {
		fmt.Println("GetActivitiesPage after stravaApi.GetAthleteActivities(): ", err)
		return activities.Index([]domain.ActivityInfo{}).Render(r.Context(), w)
	}
	return activities.Index(activitiesInfo).Render(r.Context(), w)
}

type ActivitiesController struct {
	cfg         *config.Config
	stravaStore store.StravaStore
}

func NewActivitiesController(cfg *config.Config, stravaStore store.StravaStore) *ActivitiesController {
	return &ActivitiesController{
		cfg:         cfg,
		stravaStore: stravaStore,
	}
}
