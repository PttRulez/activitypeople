package domain

type Role string

const (
	Admin       Role = "ADMIN"
	RegularUser Role = "REGULAR"
)

type User struct {
	Email          string
	Id             int
	HashedPassword string
	Name           string
	Password       string
	Role           Role
	Strava         StravaInfo
	UserSettings
}

type UserSettings struct {
	BMR                 int
	CaloriesPer100Steps float64
}

type StravaInfo struct {
	AccessToken  *string `db:"access_token"`
	RefreshToken *string `db:"refresh_token"`
}

func (u *User) JSON() UserInfo {
	stravaLinked := false
	if u.Strava.AccessToken != nil {
		stravaLinked = true
	}

	return UserInfo{
		CaloriesPer100Steps: u.CaloriesPer100Steps,
		BMR:                 u.BMR,
		Email:               u.Email,
		Name:                u.Name,
		Role:                u.Role,
		StravaLinked:        stravaLinked,
	}
}

type UserInfo struct {
	CaloriesPer100Steps float64 `json:"caloriesPer100Steps"`
	BMR                 int     `json:"bmr"`
	Email               string  `json:"email"`
	Name                string  `json:"name"`
	Role                Role    `json:"role"`
	StravaLinked        bool    `json:"stravaLinked"`
}
