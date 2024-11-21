package session

import "github.com/pttrulez/activitypeople/internal/domain"

type User struct {
	Id               int
	Email            string
	Name             string
	Role             domain.Role
	StravaAcessToken string
}
