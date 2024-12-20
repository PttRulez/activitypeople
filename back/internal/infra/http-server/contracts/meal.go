package contracts

import "time"

type CreateMealRequest struct {
	Calories *int         `json:"calories" validate:"required,number"`
	Date     time.Time    `json:"date" validate:"required"`
	Name     string       `json:"name" validate:"required"`
	Foods    []FoodInMeal `json:"foods" validate:"required,dive"`
}

type FoodInMeal struct {
	Calories       *int   `json:"calories" validate:"required"`
	CaloriesPer100 *int   `json:"caloriesPer100" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Weight         *int   `json:"weight"  validate:"required"`
}

type FoodInMealResponse struct {
	Calories int    `json:"calories"`
	Name     string `json:"name"`
	Weight   int    `json:"weight"`
}

type MealResponse struct {
	Calories int                  `json:"calories"`
	Date     time.Time            `json:"date"`
	Id       int                  `json:"id"`
	Name     string               `json:"name"`
	Foods    []FoodInMealResponse `json:"foods"`
}
