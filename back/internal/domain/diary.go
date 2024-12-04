package domain

import "time"

type DiaryDay struct {
	Activities       []Activity
	Calories         int
	CaloriesBurned   int
	CaloriesConsumed int
	Date             time.Time
	Meals            []Meal
	Weight           float64
}

type Weight struct {
	Date   time.Time
	Weight float64
}

type WeightFilters struct {
	From  time.Time
	Until time.Time
}

func (d DiaryDay) CalculateCalories(bmr int) DiaryDay {
	caloriesBurned := 0
	for _, a := range d.Activities {
		caloriesBurned += a.Calories
	}

	caloriesConsumed := 0
	for _, m := range d.Meals {
		caloriesConsumed += m.Calories
	}

	calories := caloriesConsumed - caloriesBurned - bmr

	d.Calories = calories
	d.CaloriesBurned = caloriesBurned
	d.CaloriesConsumed = caloriesConsumed

	return d
}
