package store

import (
	"antiscoof/internal/model"
	"context"
	"net/http"
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
	Insert(ctx context.Context, s *model.StravaInfo) error
	GetByUserId(ctx context.Context, userId int) (*model.StravaInfo, error)
	UpdateUserStravaInfo(ctx context.Context, s *model.UpdateStravaTokens) error
}

type UserStore interface {
	Insert(ctx context.Context, u *model.User) (int, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByIdWithStrava(ctx context.Context, id int) (*model.User, error)
}
