package handler

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/pttrulez/activitypeople/internal/domain"
// 	"github.com/pttrulez/activitypeople/internal/infra/http_server/contracts"
// 	"github.com/pttrulez/activitypeople/internal/infra/http_server/converter"
// 	"github.com/pttrulez/activitypeople/internal/logger"
// )

// func (c *FoodController) CreateFood(w http.ResponseWriter, r *http.Request) error {
// 	user := GetUserFromRequest(r)

// 	var err error
// 	var req contracts.CreateFoodRequest
// 	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		logger.Debug()
// 		w.WriteHeader(http.StatusCreated)
// 		return err
// 	}

// 	if err = c.validator.Struct(req); err != nil {
// 		var validateErr validator.ValidationErrors
// 		errors.As(err, &validateErr)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return validateErr
// 	}

// 	err = c.foodService.CreateFood(r.Context(), converter.FromFoodReqToFood(req), user.Id)
// 	if err != nil {
// 		return err
// 	}
// 	w.WriteHeader(http.StatusBadRequest)
// 	return nil
// }

// func (c *FoodController) DeleteFood(w http.ResponseWriter, r *http.Request) error {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		return errors.New("id is required")
// 	}

// 	user := GetUserFromRequest(r)

// 	err = c.foodService.DeleteFood(r.Context(), id, user.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *FoodController) Search(w http.ResponseWriter, r *http.Request) error {
// 	q := r.URL.Query().Get("q")

// 	foods, err := c.foodService.Search(r.Context(), q)
// 	if err != nil {
// 		return err
// 	}

// 	return writeJSON(w, http.StatusOK, foods)
// }

// func NewFoodController(foodService FoodService, validator *validator.Validate) *FoodController {
// 	return &FoodController{
// 		foodService: foodService,
// 		validator:   validator,
// 	}
// }

// type FoodController struct {
// 	foodService FoodService
// 	validator   *validator.Validate
// }

// type FoodService interface {
// 	CreateFood(ctx context.Context, f domain.Food, userID int) error
// 	DeleteFood(ctx context.Context, foodID int, userID int) error
// 	Search(ctx context.Context, q string) ([]domain.Food, error)
// }
