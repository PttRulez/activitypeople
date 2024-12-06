package domain

import "time"

type Food struct {
	CaloriesPer100 int
	Carbs          int
	CreatedByAdmin bool
	Fat            int
	ID             int
	Name           string
	Protein        int
	UserID         int
}

type FoodInMeal struct {
	Calories       int
	CaloriesPer100 int
	Name           string
	Weight         int
	MealId         int
}

type Meal struct {
	Calories int
	Date     time.Time
	Id       int
	Name     string
	Foods    []FoodInMeal
	UserId   int
}
