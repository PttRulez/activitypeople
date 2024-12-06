package user

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func (s *UserService) SaveSettings(ctx context.Context, f domain.UserSettings,
	userId int) error {
	return s.userRepo.SaveSettings(ctx, f, userId)
}

func NewService(userRepo Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type UserService struct {
	userRepo Repository
}

type Repository interface {
	SaveSettings(ctx context.Context, s domain.UserSettings, userId int) error
}
