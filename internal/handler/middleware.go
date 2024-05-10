package handler

import (
	"antiscoof/internal/store"
	"antiscoof/internal/types"
	"context"
	"fmt"
	"net/http"
	"strings"
)

func AddUserToContext(sessionStore store.SessionStore, userStore store.UserStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/public") {
				next.ServeHTTP(w, r)
				return
			}

			userSession, err := sessionStore.GetUserFromSession(r)
			if err != nil {
				fmt.Println("Не получили пользователя из сессии:\n", err)
				next.ServeHTTP(w, r)
				return
			}
			
			user, err := userStore.GetById(r.Context(), userSession.Id)

			authUser := types.AuthenticatedUser{
				Id:       user.Id,
				Email:    user.Email,
				LoggedIn: true,
			}
			ctx := context.WithValue(r.Context(), types.UserContextKey, authUser)
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

		user := getUserFromContext(r)
		if !user.LoggedIn {
			path := r.URL.Path
			htmxRedirect(w, r, "/login?from="+path)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
