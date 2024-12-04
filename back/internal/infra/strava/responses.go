package strava

import "time"

type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OAuthRefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type ActivityResponse struct {
	Calories float64 `json:"calories"`
}

type AthleteActivityResponse struct {
	AverageHeartrate float64 `json:"average_heartrate"`

	// The activity's distance, in meters
	Distance float64 `json:"distance"`

	// The activity's total time, in seconds
	ElapsedTime int `json:"elapsed_time"`

	HasHeartrate bool `json:"has_heartrate"`

	Id int64 `json:"id"`

	MaxHeartrate float64 `json:"max_heartrate"`

	// The activity's moving time, in seconds
	MovingTime float64 `json:"moving_time"`

	// The name of the activity
	Name string `json:"name"`

	PhotoCount int `json:"photo_count"`

	StartDate time.Time `json:"start_date"`

	StartDateLocal time.Time `json:"start_date_local"`

	SportType SportType `json:"type"`

	TotalElevationGain float64 `json:"total_elevation_gain"`

	WorkoutType int `json:"workout_type"`
}
