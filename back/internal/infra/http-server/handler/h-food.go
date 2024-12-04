package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/converter"
)

func (c *FoodController) CreateFood(e echo.Context) error {
	user := e.Get("u").(domain.User)

	var req contracts.CreateFoodRequest
	err := e.Bind(&req)
	if err != nil {
		return err
	}
	err = ValidateStruct(req)
	if err != nil {
		return err
	}

	food := converter.FromFoodReqToFood(req)
	if user.Role == domain.Admin {
		food.CreatedByAdmin = true
	}

	err = c.foodService.CreateFood(e.Request().Context(), food, user.Id)
	if err != nil {
		return err
	}

	return e.String(http.StatusCreated, "Food created successfully")
}

func (c *FoodController) DeleteFood(e echo.Context) error {
	user := e.Get("u").(domain.User)

	idStr := e.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("not valid Id format")
	}

	err = c.foodService.DeleteFood(e.Request().Context(), id, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *FoodController) Search(e echo.Context) error {
	q := e.QueryParam("q")

	foods, err := c.foodService.Search(e.Request().Context(), q)
	if err != nil {
		return err
	}

	res := make([]contracts.FoodResponse, len(foods))
	for i, f := range foods {
		res[i] = converter.FromFoodToFoodresponse(f)
	}

	return e.JSON(http.StatusOK, res)
}

func (c *FoodController) CreateWeight(e echo.Context) error {
	user := e.Get("u").(domain.User)

	var req contracts.CreateWeightRequest
	err := e.Bind(&req)
	if err != nil {
		return err
	}
	err = ValidateStruct(req)
	if err != nil {
		return err
	}

	weight := converter.FromWeightReqToWeight(req)

	err = c.foodService.CreateWeight(e.Request().Context(), weight, user.Id)
	if err != nil {
		return err
	}

	return e.String(http.StatusCreated, "Weight created successfully")
}

func NewFoodController(activitiesService AcitivitiesService, foodService FoodService) *FoodController {
	return &FoodController{
		activitiesService: activitiesService,
		foodService:       foodService,
	}
}

type FoodController struct {
	activitiesService AcitivitiesService
	foodService       FoodService
	mealService       MealService
}
