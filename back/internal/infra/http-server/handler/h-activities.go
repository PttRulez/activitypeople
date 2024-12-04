package handler

import (
	"fmt"
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

func (c *ActivitiesController) HydrateStravaActivity(e echo.Context) error {
	user := e.Get("u").(domain.User)

	sourceIdStr := e.Param("sourceId")
	sourceId, err := strconv.Atoi(sourceIdStr)
	if err != nil {
		return fmt.Errorf("wrong sourceId: %w", err)
	}

	return c.activitiesService.HydrateStravaActivity(e.Request().Context(),
		sourceId, user)
}

func (c *ActivitiesController) GetActivities(e echo.Context) error {
	user := e.Get("u").(domain.User)
	until := time.Now()
	from := until.AddDate(0, 0, -10)
	var err error

	if e.QueryParam("from") != "" {
		from, err = time.Parse(("2006-01-02"), e.QueryParam("from"))
		if err != nil {
			return e.String(http.StatusUnprocessableEntity, "wrong from format")
		}
	}

	if e.QueryParam("until") != "" {
		until, err = time.Parse(("2006-01-02"), e.QueryParam("until"))
		if err != nil {
			return e.String(http.StatusUnprocessableEntity, "wrong from until")
		}
	}

	activitiesList, err := c.activitiesService.GetActivities(e.Request().Context(), user,
		domain.ActivityFilters{
			From:  from,
			Until: until,
		})
	if err != nil {
		return err
	}

	r := make([]contracts.ActivityDayResponse, 0)
	curDate := activitiesList[0].Date
	curIndex := -1

	for _, a := range activitiesList {
		if a.Date.Equal(curDate) && curIndex != -1 {
			r[curIndex].Activities = append(r[curIndex].Activities,
				converter.FromActivityToActivityResponse(a))
		} else {
			curDate = a.Date
			r = append(r, contracts.ActivityDayResponse{
				Date: curDate.Format(time.DateOnly),
				Activities: []contracts.ActivityResponse{
					converter.FromActivityToActivityResponse(a),
				},
			})
			curIndex++
		}
	}
	return e.JSON(http.StatusOK, r)
}

func (c *ActivitiesController) SyncStrava(e echo.Context) error {
	user := e.Get("u").(domain.User)
	return c.activitiesService.SyncStravaActivities(e.Request().Context(), user)
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
