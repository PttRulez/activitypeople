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

func (s *FoodService) Search(ctx context.Context, q string) ([]domain.Food, error) {
	return s.foodRepo.Search(ctx, q)
}

func NewFoodService(foodRepo Repository) *FoodService {
	return &FoodService{
		foodRepo: foodRepo,
	}
}

type FoodService struct {
	foodRepo Repository
}

type Repository interface {
	Insert(ctx context.Context, f domain.Food) error
	Delete(ctx context.Context, foodID int, userID int) error
	Search(ctx context.Context, q string) ([]domain.Food, error)
}
