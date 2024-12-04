package food

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func (s *FoodService) CreateFood(ctx context.Context, f domain.Food, userID int) error {
	f.UserID = userID
	return s.foodRepo.Insert(ctx, f)
}

func (s *FoodService) DeleteFood(ctx context.Context, foodID int, userID int) error {
	return s.foodRepo.Delete(ctx, foodID, userID)
}

func (s *FoodService) CreateMeal(ctx context.Context, f domain.Meal) error {
	return s.mealRepo.Insert(ctx, f)
}

func (s *FoodService) CreateWeight(ctx context.Context, w domain.Weight,
	userID int) error {
	return s.weightRepo.Insert(ctx, w, userID)
}

func (s *FoodService) GetMeals(ctx context.Context, f domain.MealFilters, userId int) (
	[]domain.Meal, error) {
	return s.mealRepo.Get(ctx, f, userId)
}

func (s *FoodService) Search(ctx context.Context, q string) ([]domain.Food, error) {
	return s.foodRepo.Search(ctx, q)
}

func NewFoodService(foodRepo FoodRepository, mealRepo MealRepository, weightRepo WeightRepository) *FoodService {
	return &FoodService{
		foodRepo:   foodRepo,
		mealRepo:   mealRepo,
		weightRepo: weightRepo,
	}
}

type FoodService struct {
	foodRepo   FoodRepository
	mealRepo   MealRepository
	weightRepo WeightRepository
}

type FoodRepository interface {
	Insert(ctx context.Context, f domain.Food) error
	Delete(ctx context.Context, foodID int, userID int) error
	Search(ctx context.Context, q string) ([]domain.Food, error)
}

type MealRepository interface {
	Insert(ctx context.Context, f domain.Meal) error
	Get(ctx context.Context, f domain.MealFilters, userID int) ([]domain.Meal, error)
}

type WeightRepository interface {
	Insert(ctx context.Context, w domain.Weight, userID int) error
}
