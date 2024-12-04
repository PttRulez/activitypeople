package handler

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
)

type AcitivitiesService interface {
	GetActivities(ctx context.Context, user domain.User,
		filters domain.ActivityFilters) ([]domain.Activity, error)
	HydrateStravaActivity(ctx context.Context, sourceId int, user domain.User) error
	OAuthStrava(ctx context.Context, userCode string, userID int) error
	SyncStravaActivities(ctx context.Context, user domain.User) error
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	Register(ctx context.Context, email, password, name string) (domain.User, error)
}

type FoodService interface {
	CreateFood(ctx context.Context, f domain.Food, userID int) error
	CreateWeight(ctx context.Context, w domain.Weight, userID int) error
	DeleteFood(ctx context.Context, foodID int, userID int) error
	Search(ctx context.Context, q string) ([]domain.Food, error)
}

type MealService interface {
	CreateMeal(ctx context.Context, f domain.Meal) error
	GetMeals(ctx context.Context, f domain.MealFilters, userId int) ([]domain.Meal, error)
}
