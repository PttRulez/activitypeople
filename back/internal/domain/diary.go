package domain

import (
	"fmt"
	"time"
)

type DiaryDay struct {
	Activities       []Activity
	Calories         int
	CaloriesBurned   int
	CaloriesConsumed int
	Date             time.Time
	Meals            []Meal
	Steps            int
	Weight           float64
}

func (d DiaryDay) CalculateCalories(u User) DiaryDay {
	caloriesBurned := 0
	for _, a := range d.Activities {
		caloriesBurned += a.Calories
	}
	fmt.Println("burned", caloriesBurned, d.Steps, d.Date.Format(time.DateOnly))
	caloriesBurned += d.Steps / 100 * u.CaloriesPer100Steps
	fmt.Println("burned", caloriesBurned)

	caloriesConsumed := 0
	for _, m := range d.Meals {
		caloriesConsumed += m.Calories
	}

	calories := caloriesConsumed - caloriesBurned - u.BMR

	d.Calories = calories
	d.CaloriesBurned = caloriesBurned
	d.CaloriesConsumed = caloriesConsumed

	return d
}

type Steps struct {
	Date  time.Time
	Steps int
}

type Weight struct {
	Date   time.Time
	Weight float64
}
