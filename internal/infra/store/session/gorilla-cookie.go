package session

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/pttrulez/activitypeople/internal/domain"
)

func (s GorillaCookiesSessionsStore) SetUserIntoSession(w http.ResponseWriter, r *http.Request,
	user domain.User) error {
	userSession, err := s.store.Get(r, userKey)
	if err != nil {
		return fmt.Errorf("SetUserIntoSession: %w", err)
	}

	// кука живет один час
	userSession.Options.MaxAge = 3600

	// заполняем куку
	userSession.Values[idKey] = user.Id
	userSession.Values[emailKey] = user.Email
	userSession.Values[stravaAccessKey] = user.Strava.AccessToken
	userSession.Values[stravaRefreshKey] = user.Strava.RefreshToken

	return userSession.Save(r, w)
}

func (s GorillaCookiesSessionsStore) GetUserFromSession(r *http.Request) (domain.User, error) {
	userSession, err := s.store.Get(r, userKey)
	if err != nil {
		return domain.User{}, fmt.Errorf("GetUserFromSession: %w", err)
	}

	if userSession.Values[idKey] == nil {
		return domain.User{}, fmt.Errorf("userSessionCookie user id is not set")
	}

	stravaAccessToken, _ := userSession.Values[stravaAccessKey].(string)
	stravaRefreshToken, _ := userSession.Values[stravaRefreshKey].(string)
	user := domain.User{
		Id:    userSession.Values[idKey].(int),
		Email: userSession.Values[emailKey].(string),
		Strava: domain.StravaInfo{
			AccessToken:  &stravaAccessToken,
			RefreshToken: &stravaRefreshToken,
		},
	}

	return user, nil
}

func (s GorillaCookiesSessionsStore) DeleteUserSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.store.Get(r, userKey)
	session.Options.MaxAge = -1
	return session.Save(r, w)
}

type GorillaCookiesSessionsStore struct {
	store *sessions.CookieStore
}

func NewGorillaCookiesSessionsStore(secret []byte) *GorillaCookiesSessionsStore {
	store := sessions.NewCookieStore(secret)

	return &GorillaCookiesSessionsStore{
		store: store,
	}
}

var (
	userKey          = "User"
	idKey            = "Id"
	emailKey         = "Email"
	stravaAccessKey  = "StravaAccess"
	stravaRefreshKey = "StravaRefresh"
)
