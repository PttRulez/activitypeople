package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

func FromFoodReqToFood(req contracts.CreateFoodRequest) domain.Food {
	var carbs, fat, protein int

	if req.Carbs == nil {
		carbs = 0
	}
	if req.Fat == nil {
		fat = 0
	}
	if req.Protein == nil {
		protein = 0
	}

	return domain.Food{
		Name:           req.Name,
		CreatedByAdmin: false,
		Calories:       req.Calories,
		Carbs:          carbs,
		Fat:            fat,
		Protein:        protein,
	}
}

func FromFoodToFoodresponse(f domain.Food) contracts.FoodResponse {
	return contracts.FoodResponse{
		Name:     f.Name,
		Calories: f.Calories,
		Carbs:    f.Carbs,
		Fat:      f.Fat,
		Id:       f.ID,
		Protein:  f.Protein,
	}
}
