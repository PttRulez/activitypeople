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

func (u *User) JSON() UserInfo {
	stravaLinked := false
	if u.Strava.AccessToken != nil {
		stravaLinked = true
	}

	return UserInfo{
		Email:        u.Email,
		Name:         u.Name,
		Role:         u.Role,
		StravaLinked: stravaLinked,
	}
}

type UserInfo struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Role         Role   `json:"role"`
	StravaLinked bool   `json:"stravaLinked"`
}
