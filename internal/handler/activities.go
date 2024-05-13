package handler

import (
	"antiscoof/internal/config"
	stravaclient "antiscoof/internal/service/strava-client"
	stravamodels "antiscoof/internal/service/strava-client/stravamodels"
	"antiscoof/internal/store"
	"antiscoof/internal/view/activities"
	"fmt"
	"net/http"
)

func (c *ActivitiesController) GetActivitiesPage(w http.ResponseWriter, r *http.Request) error {
	stravaApi := stravaclient.NewStravaApiFromRequest(r, c.cfg.Strava.ClientID, c.cfg.Strava.ClientSecret, c.stravaStore)
	go func() {
		stravaApi.ObaGetAthleteActivities()
	}()
	activitiesInfo, err := stravaApi.GetAthleteActivities()
	if err != nil {
		fmt.Println("GetActivitiesPage after stravaApi.GetAthleteActivities(): ", err)
		return activities.Index([]stravamodels.ActivityInfo{}).Render(r.Context(), w)
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
