package handler

import (
	"context"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/view/activities"
)

func (c *ActivitiesController) GetActivitiesPage(w http.ResponseWriter, r *http.Request) error {
	user := GetUserFromRequest(r)

	activitiesInfo, err := c.activitiesService.GetActivities(r.Context(),
		*user.Strava.AccessToken, *user.Strava.RefreshToken, user.Id)
	if err != nil {
		return activities.Index([]domain.Activity{}, user).Render(r.Context(), w)
	}

	return activities.Index(activitiesInfo, user).Render(r.Context(), w)
}

type ActivitiesController struct {
	Handler
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
	GetActivities(ctx context.Context, userAccessToken, refreshToken string, userID int) ([]domain.Activity, error)
}
