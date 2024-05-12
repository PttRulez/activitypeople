package model

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

type RegisterUserDto struct {
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
	Role            Role   `json:"role"`
}

type LoginUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
