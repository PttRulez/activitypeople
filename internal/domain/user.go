package domain

type Role string

const (
	Admin Role = "ADMIN"
	Scoof Role = "SCOOF"
)

const UserContextKey = "user"

type User struct {
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Id             int    `db:"id"`
	Name           string `db:"name"`
	Password       string ``
	Role           Role   `db:"role"`
	Strava         *StravaInfo
}

type AuthenticatedUser struct {
	Id       int
	Email    string
	Name     string
	LoggedIn bool
	Strava   *StravaInfo
}

type StravaInfo struct {
	AccessToken  *string `db:"access_token"`
	Id           int     `db:"id"`
	RefreshToken *string `db:"refresh_token"`
	UserId       int     `db:"user_id"`
}
