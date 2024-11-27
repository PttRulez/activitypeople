package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

func FromMealReqToMeal(req contracts.CreateMealRequest) domain.Meal {
	foods := make([]domain.FoodInMeal, len(req.Foods))
	for i, f := range req.Foods {
		foods[i] = domain.FoodInMeal{
			Id:       f.Id,
			Name:     f.Name,
			Weight:   f.Weight,
			Calories: f.Calories,
		}
	}

	return domain.Meal{
		Date:     req.Date.Time,
		Calories: req.Calories,
		Name:     req.Name,
		Foods:    foods,
	}
}

func FromMealToMealResponse(m domain.Meal) contracts.MealResponse {
	foods := make([]contracts.FoodInMealResponse, len(m.Foods))
	for i, f := range m.Foods {
		foods[i] = contracts.FoodInMealResponse{
			Calories: f.Calories,
			Name:     f.Name,
			Weight:   f.Weight,
		}
	}

	return contracts.MealResponse{
		Calories: m.Calories,
		Date:     m.Date,
		Id:       m.Id,
		Name:     m.Name,
		Foods:    foods,
	}
}
