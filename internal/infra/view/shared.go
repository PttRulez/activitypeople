package view

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func GetAuthenticatedUser(ctx context.Context) domain.AuthenticatedUser {
	user, ok := ctx.Value(domain.UserContextKey).(domain.AuthenticatedUser)
	if !ok {
		return domain.AuthenticatedUser{}
	}
	return user
}
