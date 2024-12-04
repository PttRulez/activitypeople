package meal

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func (s *MealService) CreateMeal(ctx context.Context, f domain.Meal) error {
	return s.mealRepo.Insert(ctx, f)
}

func NewMealService(mealRepo Repository) *MealService {
	return &MealService{
		mealRepo: mealRepo,
	}
}

type MealService struct {
	mealRepo Repository
}

type Repository interface {
	Insert(ctx context.Context, f domain.Meal) error
}
