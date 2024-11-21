package handler

import (
	"net/http"

	"github.com/pttrulez/activitypeople/internal/domain"
)

type SessionStore interface {
	ClearUserSession(w http.ResponseWriter, r *http.Request) error

	GetUserFromSession(r *http.Request) (domain.User, error)

	SetUserIntoSession(w http.ResponseWriter, r *http.Request, user domain.User) error
}

type Handler struct {
	session SessionStore
}
