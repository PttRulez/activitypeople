package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
)

func AddUserToContext(sessionStore store.SessionStore, userStore store.UserStore,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/public") {
				next.ServeHTTP(w, r)
				return
			}

			userSession, err := sessionStore.GetUserFromSession(r)
			if err != nil {
				fmt.Printf("Не получили пользователя из сессии:\n %v", err)
				next.ServeHTTP(w, r)
				return
			}

			user, err := userStore.GetByIdWithStrava(r.Context(), userSession.Id)
			if err != nil || user == nil {
				fmt.Println("Не получили пользователя из базы данных:\n", err)

				next.ServeHTTP(w, r)
				return
			}

			authUser := domain.AuthenticatedUser{
				Id:       user.Id,
				Email:    user.Email,
				LoggedIn: true,
				Name:     user.Name,
				Strava:   user.Strava,
			}
			ctx := context.WithValue(r.Context(), domain.UserContextKey, authUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func OnlyAuthenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}

		user := GetUserFromRequest(r)
		if !user.LoggedIn {
			path := r.URL.Path
			HtmxRedirect(w, r, "/login?from="+path)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
