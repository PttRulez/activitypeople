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
	// The activity's distance, in meters
	Distance    float64 `json:"distance"`
	ElapsedTime int     `json:"elapsed_time"`
	Id          int     `json:"id"`

	// The activity's moving time, in seconds
	MovingTime int `json:"moving_time"`

	// The name of the activity
	Name           string    `json:"name"`
	PhotoCount     int       `json:"photo_count"`
	StartDate      time.Time `json:"start_date"`
	StartDateLocal time.Time `json:"start_date_local"`
	SportType      SportType `json:"type"`
	WorkoutType    int       `json:"workout_type"`
}
