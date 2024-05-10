package session

import (
	"antiscoof/internal/store"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func (s GorillaCookiesSessionsStore) SetUserIntoSession(w http.ResponseWriter, r *http.Request,
	user store.UserSession) error {
	userSession, err := s.store.Get(r, s.userSessionKey)
	if err != nil {
		return fmt.Errorf("SetUserIntoSession: %w", err)
	}

	// кука живет один час
	userSession.Options.MaxAge = 3600

	// заполняем куку
	userSession.Values["id"] = user.Id
	userSession.Values["email"] = user.Email

	return userSession.Save(r, w)
}

func (s GorillaCookiesSessionsStore) GetUserFromSession(r *http.Request) (*store.UserSession, error) {
	userSession, err := s.store.Get(r, s.userSessionKey)
	if err != nil {
		return nil, fmt.Errorf("GetUserFromSession: %w", err)
	}

	var user store.UserSession
	if userSession.Values["id"] == nil {
		return nil, fmt.Errorf("userSessionCookie user id is not set")
	}

	user = store.UserSession{
		Id:    userSession.Values["id"].(int),
		Email: userSession.Values["email"].(string),
	}

	return &user, nil
}

func (s GorillaCookiesSessionsStore) DeleteUserSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.store.Get(r, s.userSessionKey)
	session.Options.MaxAge = -1
	return session.Save(r, w)
}

type GorillaCookiesSessionsStore struct {
	userSessionKey string
	store          *sessions.CookieStore
}

func NewGorillaCookiesSessionsStore(secret []byte, userSessionKey string) *GorillaCookiesSessionsStore {
	store := sessions.NewCookieStore(secret)

	return &GorillaCookiesSessionsStore{
		store:          store,
		userSessionKey: userSessionKey,
	}
}
