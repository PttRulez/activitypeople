package contracts

import "time"

type CreateWeightRequest struct {
	Date   time.Time `json:"date" validate:"required"`
	Weight float64   `json:"weight" validate:"required,number"`
}

type WeightResponse struct {
	Date   time.Time `json:"date"`
	Weight float64   `json:"weight"`
}

type DiaryDayResponse struct {
	Activities       []ActivityResponse `json:"activities"`
	Calories         int                `json:"calories"`
	CaloriesBurned   int                `json:"caloriesBurned"`
	CaloriesConsumed int                `json:"caloriesConsumed"`
	Date             time.Time          `json:"date"`
	Meals            []MealResponse     `json:"meals"`
	Weight           float64            `json:"weight"`
}
