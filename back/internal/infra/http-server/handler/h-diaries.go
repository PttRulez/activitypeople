package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/converter"
)

func (c *DiaryController) GetDiaries(e echo.Context) error {
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

	diaries, err := c.diaryService.GetDiaries(e.Request().Context(), user.Id, from, until)
	if err != nil {
		return err
	}

	drs := make(map[string]contracts.DiaryDayResponse)
	for _, d := range diaries {
		drs[d.Date.Format("2006-01-02")] = converter.FromDiaryToDiaryResponse(d)
	}

	var res contracts.DiariesResponse
	res.BMR = user.BMR
	res.Diaries = drs

	return e.JSON(http.StatusOK, res)
}

func NewDiaryController(diaryService DiaryService) *DiaryController {
	return &DiaryController{diaryService: diaryService}
}

type DiaryController struct {
	diaryService DiaryService
}
