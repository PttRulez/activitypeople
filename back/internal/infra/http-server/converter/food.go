package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http_server/contracts"
)

func FromFoodReqToFood(req contracts.CreateFoodRequest) domain.Food {
	return domain.Food{
		Name:     req.Name,
		Calories: req.Calories,
		Carbs:    req.Carbs,
		Fat:      req.Fat,
		Protein:  req.Protein,
		Public:   req.Public,
	}
}
