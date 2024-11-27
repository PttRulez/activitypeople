package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/converter"
)

func (c *MealController) CreateMeal(e echo.Context) error {
	user := e.Get("u").(domain.User)

	var req contracts.CreateMealRequest
	err := e.Bind(&req)
	if err != nil {
		return err
	}
	err = ValidateStruct(req)
	if err != nil {
		return err
	}

	meal := converter.FromMealReqToMeal(req)
	meal.UserId = user.Id

	return c.mealService.CreateMeal(e.Request().Context(), meal)
}

func (c *MealController) GetMeals(e echo.Context) error {
	user := e.Get("u").(domain.User)
	fmt.Println("user", user)

	from := e.QueryParam("from")
	until := e.QueryParam("until")
	now := time.Now()

	var f domain.MealFilters

	if from == "" {
		f.From = now.AddDate(0, 0, -10)
	} else {
		t, err := time.Parse(time.DateOnly, from)
		if err != nil {
			return err
		}
		f.From = t
	}

	if until == "" {
		f.Until = now

	} else {
		t, err := time.Parse(time.DateOnly, until)
		if err != nil {
			return err
		}
		f.Until = t
	}

	meals, err := c.mealService.GetMeals(e.Request().Context(), f, user.Id)
	if err != nil {
		return err
	}

	mealsResponse := make([]contracts.MealResponse, len(meals))
	for i, m := range meals {
		mealsResponse[i] = converter.FromMealToMealResponse(m)
	}

	return e.JSON(http.StatusOK, mealsResponse)
}

func NewMealController(mealService MealService) *MealController {
	return &MealController{
		mealService: mealService,
	}
}

type MealService interface {
	CreateMeal(ctx context.Context, f domain.Meal) error
	GetMeals(ctx context.Context, f domain.MealFilters, userId int) ([]domain.Meal, error)
}

type MealController struct {
	mealService MealService
}
