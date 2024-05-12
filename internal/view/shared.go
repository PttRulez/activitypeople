package view

import (
	"antiscoof/internal/model"
	"context"
)

func GetAuthenticatedUser(ctx context.Context) model.AuthenticatedUser {
	user, ok := ctx.Value(model.UserContextKey).(model.AuthenticatedUser)
	if !ok {
		return model.AuthenticatedUser{}
	}
	return user
}
