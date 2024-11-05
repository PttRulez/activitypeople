package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/pttrulez/activitypeople/internal/domain"
)

type ctxKey string

var userKey ctxKey = "user"

func AddUserToContextMiddleware(sessionStore SessionStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/public") {
				next.ServeHTTP(w, r)
				return
			}

			user, err := sessionStore.GetUserFromSession(r)
			if err != nil {
				ctx := context.WithValue(r.Context(), userKey, domain.User{})
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			ctx := context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func OnlyAuthenticatedMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}

		user, ok := r.Context().Value(userKey).(domain.User)
		if !ok || user.Id == 0 {
			HtmxRedirect(w, r, "/login?from="+r.URL.Path)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// func OnlyAuthenticated(sessionStore SessionStore) func(next http.Handler) http.Handler {
// 	return func (next http.Handler) http.Handler {
// 		fn := func(w http.ResponseWriter, r *http.Request) {
// 			if strings.Contains(r.URL.Path, "/public") {
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			_, err := sessionStore.GetUserFromSession(r)
// 			if err != nil {
// 				handler.HtmxRedirect(w, r, "/login?from="+r.URL.Path)
// 			}
// 			next.ServeHTTP(w, r)
// 		}
// 		return http.HandlerFunc(fn)
// 	}
// }
