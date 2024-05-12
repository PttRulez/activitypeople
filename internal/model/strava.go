package model

type StravaInfo struct {
	AccessToken  *string `db:"access_token"`
	Id           int     `db:"id"`
	RefreshToken *string `db:"refresh_token"`
	UserId       int     `db:"user_id"`
}

type UpdateStravaTokens struct {
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	UserId       int    `db:"user_id"`
}
