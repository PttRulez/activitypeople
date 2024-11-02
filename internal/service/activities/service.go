package activities

import (
	"context"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
)

func (s *Service) GetActivities(ctx context.Context, userAccessToken string) ([]domain.ActivityInfo,
	error) {
	// return s.stravaClient.ObaGetAthleteActivities(ctx, userAccessToken)
	return s.stravaClient.GetAthleteActivities(ctx, userAccessToken)
}

func (s *Service) OAuthStrava(ctx context.Context, userCode string, userID int) error {
	data, err := s.stravaClient.OAuth(userCode)
	if err != nil {
		return err
	}

	err = s.stravaRepo.Insert(ctx, data.AccessToken, data.RefreshToken, userID)
	if err != nil {
		return err
	}

	return nil
}

func NewService(stravaClient *strava.StravaClient, stravaRepo Repository) *Service {
	return &Service{
		stravaClient: stravaClient,
		stravaRepo:   stravaRepo,
	}
}

type Service struct {
	stravaClient *strava.StravaClient
	stravaRepo   Repository
}

type Repository interface {
	Insert(
		ctx context.Context,
		accessToken string,
		refreshToken string,
		userId int,
	) error
}
