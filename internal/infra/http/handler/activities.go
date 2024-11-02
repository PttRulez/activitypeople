package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/view/activities"
)

func (c *ActivitiesController) GetActivitiesPage(w http.ResponseWriter, r *http.Request) error {
	user := GetUserFromRequest(r)

	// // создаём страва клиент
	// stravaApi := strava.NewStravaClient(c.cfg.Strava.ClientID, c.cfg.Strava.ClientSecret)

	// // TODO опеределиться как ходим в страву - сами или через стороннюю либу
	// go func() {
	// 	stravaApi.ObaGetAthleteActivities()
	// }()

	// activitiesInfo, err := stravaApi.GetAthleteActivities()

	activitiesInfo, err := c.activitiesService.GetActivities(r.Context(), *user.Strava.AccessToken)
	if err != nil {
		fmt.Println("GetActivitiesPage after stravaApi.GetAthleteActivities(): ", err)
		return activities.Index([]domain.ActivityInfo{}).Render(r.Context(), w)
	}

	return activities.Index(activitiesInfo).Render(r.Context(), w)
}

type ActivitiesController struct {
	activitiesService AcitivitiesService
}

func NewActivitiesController(
	activitiesService AcitivitiesService,
) *ActivitiesController {
	return &ActivitiesController{
		activitiesService: activitiesService,
	}
}

type AcitivitiesService interface {
	GetActivities(ctx context.Context, userAccessToken string) ([]domain.ActivityInfo, error)
}
