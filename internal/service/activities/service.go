package activities

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
)

func (s *Service) GetActivities(ctx context.Context, userAccessToken string,
	refreshToken string, userID int) ([]domain.Activity,
	error) {
	client := s.stravaBase.NewClient(userAccessToken, refreshToken, s.makeStoreTokensFunc(
		ctx, userID))
	client.ObaGetAthleteActivities(ctx)
	return client.GetAthleteActivities(ctx)
}

func (s *Service) OAuthStrava(ctx context.Context, userCode string, userID int) error {
	client := s.stravaBase.NewClient("", "", s.makeStoreTokensFunc(ctx, userID))

	data, err := client.OAuth(userCode)
	if err != nil {
		return err
	}

	err = s.stravaRepo.Insert(ctx, data.AccessToken, data.RefreshToken, userID)
	if err != nil {
		return err
	}

	return nil
}

func NewService(stravaBase *strava.Base, stravaRepo Repository) *Service {
	return &Service{
		stravaBase: stravaBase,
		stravaRepo: stravaRepo,
	}
}

type Service struct {
	stravaBase *strava.Base
	stravaRepo Repository
}

type Repository interface {
	Insert(ctx context.Context, accessToken string, refreshToken string, userId int) error
	UpdateUserStravaInfo(ctx context.Context, accessToken string,
		refreshToken string, userId int) error
}
