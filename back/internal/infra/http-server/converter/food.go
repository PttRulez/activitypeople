package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

func FromFoodReqToFood(req contracts.CreateFoodRequest) domain.Food {
	var carbs, fat, protein int

	if req.Carbs == nil {
		carbs = 0
	} else {
		carbs = *req.Carbs
	}
	if req.Fat == nil {
		fat = 0
	} else {
		fat = *req.Fat
	}
	if req.Protein == nil {
		protein = 0
	} else {
		protein = *req.Protein
	}

	return domain.Food{
		Name:           req.Name,
		CreatedByAdmin: false,
		CaloriesPer100: req.CaloriesPer100,
		Carbs:          carbs,
		Fat:            fat,
		Protein:        protein,
	}
}

func FromFoodToFoodresponse(f domain.Food) contracts.FoodResponse {
	return contracts.FoodResponse{
		Name:           f.Name,
		CaloriesPer100: f.CaloriesPer100,
		Carbs:          f.Carbs,
		Fat:            f.Fat,
		Id:             f.ID,
		Protein:        f.Protein,
	}
}
