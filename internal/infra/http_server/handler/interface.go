package handler

import (
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
)

type SessionStore interface {
	SetUserIntoSession(w http.ResponseWriter, r *http.Request, user domain.User) error

	GetUserFromSession(r *http.Request) (domain.User, error)

	DeleteUserSession(w http.ResponseWriter, r *http.Request) error
}

type Handler struct {
	session SessionStore
}
