package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/converter"
)

func (c *ActivitiesController) OAuthStrava(e echo.Context) error {
	code := e.QueryParam("code")
	user := e.Get("u").(domain.User)

	return c.activitiesService.OAuthStrava(e.Request().Context(), code, user.Id)
}

func (c *ActivitiesController) GetActivities(e echo.Context) error {
	user := e.Get("u").(domain.User)
	now := time.Now()
	year := now.Year()
	month := now.Month()

	if e.QueryParam("year") != "" {
		yearNumber, _ := strconv.Atoi(e.QueryParam("year"))
		if yearNumber <= year && 2000 < yearNumber {
			year = yearNumber
		}
	}

	if e.QueryParam("month") != "" {
		monthNumber, err := strconv.Atoi(e.QueryParam("month"))
		if err == nil && monthNumber >= 1 || monthNumber <= 12 {
			month = time.Month(monthNumber)
		}
	}

	daysInMonth := 32 - time.Date(year, month, 32, 0, 0, 0, 0, time.UTC).Day()

	activitiesList, err := c.activitiesService.GetActivities(e.Request().Context(), user,
		domain.ActivityFilters{
			From:  time.Date(year, month, 1, 0, 0, 0, 0, time.UTC),
			Until: time.Date(year, month, daysInMonth, 0, 0, 0, 0, time.UTC),
		})
	if err != nil {
		return err
	}

	activitiesResponse := make([]contracts.ActivityResponse, 0, len(activitiesList))
	for _, a := range activitiesList {
		activitiesResponse = append(activitiesResponse,
			converter.FromActivityToActivityResponse(a))
	}
	return e.JSON(http.StatusOK, activitiesResponse)
}

func (c *ActivitiesController) SyncStrava(e echo.Context) error {
	user := e.Get("u").(domain.User)
	return c.activitiesService.SyncActivities(e.Request().Context(), user)
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
	GetActivities(ctx context.Context, user domain.User,
		filters domain.ActivityFilters) ([]domain.Activity, error)
	OAuthStrava(ctx context.Context, userCode string, userID int) error
	SyncActivities(ctx context.Context, user domain.User) error
}
