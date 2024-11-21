package domain

type Role string

const (
	Admin Role = "ADMIN"
	Scoof Role = "SCOOF"
)

type User struct {
	Email          string
	Id             int
	HashedPassword string
	Name           string
	Password       string
	Role           Role
	Strava         StravaInfo
}

type StravaInfo struct {
	AccessToken  *string `db:"access_token"`
	RefreshToken *string `db:"refresh_token"`
}
