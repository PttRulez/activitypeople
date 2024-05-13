package models

import "time"

type ActivityInfo struct {
	// The activity's distance, in meters
	Distance    float64 `json:"distance"`
	ElapsedTime int     `json:"elapsed_time"`
	Id          int     `json:"id"`

	// The activity's moving time, in seconds
	MovingTime int `json:"moving_time"`

	// The name of the activity
	Name        string    `json:"name"`
	PhotoCount  int       `json:"photo_count"`
	StartDate   time.Time `json:"start_date"`
	SportType   SportType `json:"type"`
	WorkoutType int       `json:"workout_type"`
}
