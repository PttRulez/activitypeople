package stravaclient

import "time"

type StravaOAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type StravaOAuthRefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int    `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type StravaActivityInfo struct {
	Distance    float64   `json:"distance"`
	ElapsedTime int       `json:"elapsed_time"`
	Id          int       `json:"id"`
	MovingTime  int       `json:"moving_time"`
	Name        string    `json:"name"`
	SportType   string    `json:"sport_type"`
	StartDate   time.Time `json:"start_date"`
	Type        string    `json:"type"`
	WorkoutType int       `json:"workout_type"`
}
