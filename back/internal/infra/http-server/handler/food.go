package handler

import (
	"context"
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

func NewFoodController(foodService FoodService) *FoodController {
	return &FoodController{
		foodService: foodService,
	}
}

type FoodController struct {
	foodService FoodService
}

type FoodService interface {
	CreateFood(ctx context.Context, f domain.Food, userID int) error
	DeleteFood(ctx context.Context, foodID int, userID int) error
	Search(ctx context.Context, q string) ([]domain.Food, error)
}
