package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

func FromDiaryToDiaryResponse(d domain.DiaryDay) contracts.DiaryDayResponse {
	activities := make([]contracts.ActivityResponse, len(d.Activities))
	for i, a := range d.Activities {
		activities[i] = FromActivityToActivityResponse(a)
	}

	meals := make([]contracts.MealResponse, len(d.Meals))
	for i, a := range d.Meals {
		meals[i] = FromMealToMealResponse(a)
	}

	return contracts.DiaryDayResponse{
		Activities:       activities,
		Calories:         d.Calories,
		Date:             d.Date,
		CaloriesBurned:   d.CaloriesBurned,
		CaloriesConsumed: d.CaloriesConsumed,
		Meals:            meals,
		Steps:            d.Steps,
		Weight:           d.Weight,
	}
}

func FromWeightReqToWeight(req contracts.CreateWeightRequest) domain.Weight {
	return domain.Weight{
		Date:   req.Date,
		Weight: req.Weight,
	}
}

func FromWeightToWeightResponse(w domain.Weight) contracts.WeightResponse {
	return contracts.WeightResponse{
		Date:   w.Date,
		Weight: w.Weight,
	}
}

func FromStepsReqToSteps(req contracts.CreateStepsRequest) domain.Steps {
	return domain.Steps{
		Date:  req.Date,
		Steps: req.Steps,
	}
}

func FromStepsToStepsResponse(w domain.Steps) contracts.StepsResponse {
	return contracts.StepsResponse{
		Date:  w.Date,
		Steps: w.Steps,
	}
}
