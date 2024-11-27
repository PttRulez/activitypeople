package domain

import "time"

type Food struct {
	Calories       int
	Carbs          int
	CreatedByAdmin bool
	Fat            int
	ID             int
	Name           string
	Protein        int
	UserID         int
}

type FoodInMeal struct {
	Calories int
	Name     string
	Id       int
	Weight   int
	MealId   int
}

type Meal struct {
	Calories int
	Date     time.Time
	Id       int
	Name     string
	Foods    []FoodInMeal
	UserId   int
}

type MealFilters struct {
	From  time.Time
	Until time.Time
}
