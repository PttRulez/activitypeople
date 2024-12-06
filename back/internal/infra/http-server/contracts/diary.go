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

type CreateStepsRequest struct {
	Date  time.Time `json:"date" validate:"required"`
	Steps int       `json:"steps" validate:"required,number"`
}

type StepsResponse struct {
	Date  time.Time `json:"date"`
	Steps int       `json:"steps"`
}

type DiaryDayResponse struct {
	Activities       []ActivityResponse `json:"activities"`
	Calories         int                `json:"calories"`
	CaloriesBurned   int                `json:"caloriesBurned"`
	CaloriesConsumed int                `json:"caloriesConsumed"`
	Date             time.Time          `json:"date"`
	Meals            []MealResponse     `json:"meals"`
	Steps            int                `json:"steps"`
	Weight           float64            `json:"weight"`
}

type DiariesResponse struct {
	BMR                 int                         `json:"bmr"`
	CaloriesPer100Steps int                         `json:"caloriesPer100Steps"`
	Diaries             map[string]DiaryDayResponse `json:"diaries"`
}
