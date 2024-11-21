package handler

// import (
// 	"context"
// 	"fmt"
// 	"math"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/pttrulez/activitypeople/internal/domain"
// 	activitiesview "github.com/pttrulez/activitypeople/internal/infra/view/pages/activities"
// )

// func (c *ActivitiesController) GetActivitiesPage(w http.ResponseWriter, r *http.Request) error {
// 	user := GetUserFromRequest(r)
// 	now := time.Now()
// 	year := now.Year()
// 	month := now.Month()

// 	if r.URL.Query().Get("year") != "" {
// 		yearNumber, _ := strconv.Atoi(r.URL.Query().Get("year"))
// 		if yearNumber <= year && 2000 < yearNumber {
// 			year = yearNumber
// 		}
// 	}

// 	if r.URL.Query().Get("month") != "" {
// 		monthNumber, err := strconv.Atoi(r.URL.Query().Get("month"))
// 		if err == nil && monthNumber >= 1 || monthNumber <= 12 {
// 			month = time.Month(monthNumber)
// 		}
// 	}

// 	t := time.Date(year, month, 32, 0, 0, 0, 0, time.UTC)
// 	daysInMonth := 32 - t.Day()
// 	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
// 	numberOfRows := math.Ceil((float64(daysInMonth) + float64(firstDay.Weekday()) - 1) / 7)
// 	activitiesDate := time.Date(year, month, 1, 1, 1, 1, 1, time.Local)

// 	activitiesList, err := c.activitiesService.GetActivities(r.Context(), user,
// 		domain.ActivityFilters{
// 			From:  time.Date(year, month, 1, 0, 0, 0, 0, time.UTC),
// 			Until: time.Date(year, month, daysInMonth, 0, 0, 0, 0, time.UTC),
// 		})
// 	if err != nil {
// 		return activitiesview.Index([]domain.DiaryDay{},
// 			createPagination(activitiesDate), user).Render(r.Context(), w)
// 	}

// 	days := make([]domain.DiaryDay, int(numberOfRows)*7)
// 	for _, a := range activitiesList {
// 		index := int(firstDay.Weekday()) + int(a.Date.Day()) - 2
// 		days[index].Activity = a
// 	}

// 	return activitiesview.Index(days, createPagination(activitiesDate), user).
// 		Render(r.Context(), w)
// }

// func createPagination(d time.Time) activitiesview.Pagination {
// 	prevDate := d.AddDate(0, -1, 1)
// 	prevButtonLink := fmt.Sprintf("/activities?month=%d&year=%d", int(prevDate.Month()), prevDate.Year())
// 	nextButtonLink := ""

// 	now := time.Now()
// 	if !(d.Year() == now.Year() && d.Month() == now.Month()) {
// 		nextDate := d.AddDate(0, 1, 1)
// 		nextButtonLink = fmt.Sprintf("/activities?month=%d&year=%d", int(nextDate.Month()), nextDate.Year())
// 	}

// 	return activitiesview.Pagination{
// 		NextButtonLink: nextButtonLink,
// 		PrevButtonLink: prevButtonLink,
// 		MonthName:      d.Month().String(),
// 	}
// }

// type ActivitiesController struct {
// 	Handler
// 	activitiesService AcitivitiesService
// }

// func NewActivitiesController(
// 	activitiesService AcitivitiesService,
// ) *ActivitiesController {
// 	return &ActivitiesController{
// 		activitiesService: activitiesService,
// 	}
// }

// type AcitivitiesService interface {
// 	GetActivities(ctx context.Context, user domain.User, filters domain.ActivityFilters) (
// 		[]domain.Activity, error)
// }
