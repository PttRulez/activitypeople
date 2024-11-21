package domain

import "time"

type Food struct {
	ID       int
	Name     string
	Calories int
	Carbs    int
	Fat      int
	Protein  int
	Public   bool
	UserID   int
}

type FoodItem struct {
	Food
	Weight        int
	TotalCalories int
}

type Meal struct {
	Calories int
	Date     time.Time
	Id       int
	Name     string
	Foods    []FoodItem
	UserId   int
}
