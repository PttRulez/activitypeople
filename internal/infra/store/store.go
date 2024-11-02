package store

import (
	"context"
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
)

type UserSession struct {
	Id    int
	Email string
}

type SessionStore interface {
	DeleteUserSession(w http.ResponseWriter, r *http.Request) error
	SetUserIntoSession(w http.ResponseWriter, r *http.Request, user UserSession) error
	GetUserFromSession(r *http.Request) (*UserSession, error)
}

type StravaStore interface {
	Insert(ctx context.Context, s *domain.StravaInfo) error
	GetByUserId(ctx context.Context, userId int) (*domain.StravaInfo, error)
	UpdateUserStravaInfo(ctx context.Context, accessToken string,
		refreshToken string, userId int) error
}

type UserStore interface {
	Insert(ctx context.Context, u *domain.User) (int, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetById(ctx context.Context, id int) (*domain.User, error)
	GetByIdWithStrava(ctx context.Context, id int) (*domain.User, error)
}
