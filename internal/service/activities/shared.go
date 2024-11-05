package activities

import "context"

func (s *Service) makeStoreTokensFunc(ctx context.Context, userID int) func(
	accessToken, refreshToken string) error {
	return func(accessToken, refreshToken string) error {
		return s.stravaRepo.UpdateUserStravaInfo(ctx, accessToken, refreshToken, userID)
	}
}
