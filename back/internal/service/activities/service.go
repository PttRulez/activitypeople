package activities

import (
	"context"
	"fmt"
	"time"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
)

func (s *Service) GetActivities(ctx context.Context, user domain.User, filters domain.ActivityFilters) ([]domain.Activity, error) {
	activities, err := s.activityRepo.Get(ctx, user.Id, filters)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (s *Service) SyncActivities(ctx context.Context, user domain.User) error {
	client := s.stravaBase.NewClient(*user.Strava.AccessToken, *user.Strava.RefreshToken,
		s.makeStoreTokensFunc(ctx, user.Id))
	// Собираем все страва активности юзера и делаем сет, чтобы не доабвлять дубли
	existingActivities, err := s.activityRepo.Get(ctx, user.Id, domain.ActivityFilters{
		Source: domain.Strava,
	})
	if err != nil {
		return err
	}

	setOfStravaIds := make(map[int64]struct{})
	for _, a := range existingActivities {
		setOfStravaIds[a.SourceId] = struct{}{}
	}

	var activitiesToInsert []domain.Activity

	isPolling := true
	after := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	for isPolling {
		fmt.Println("polling...")
		// Запрашиваем активности из апи стравы
		activities, err := client.GetAthleteActivities(ctx, &after)
		if err != nil {
			return err
		}

		if len(activities) == 0 {
			isPolling = false
		}

		for i, a := range activities {
			if _, ok := setOfStravaIds[a.SourceId]; !ok {
				a.UserId = user.Id
				activitiesToInsert = append(activitiesToInsert, a)
			}

			// Выставляем after датой последнего обработанного активити
			if i == len(activities)-1 {
				after = a.Date
			}
		}

		fmt.Println("sleeping...")
		time.Sleep(time.Second * 1)
	}

	if len(activitiesToInsert) > 0 {
		err = s.activityRepo.InsertBulk(ctx, activitiesToInsert)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) OAuthStrava(ctx context.Context, userCode string, userID int) (string, string, error) {
	client := s.stravaBase.NewClient("", "", s.makeStoreTokensFunc(ctx, userID))

	data, err := client.OAuth(userCode)
	if err != nil {
		return "", "", err
	}

	err = s.stravaRepo.Insert(ctx, data.AccessToken, data.RefreshToken, userID)
	if err != nil {
		return "", "", err
	}

	return data.AccessToken, data.RefreshToken, nil
}

func NewService(activityRepo ActivitiesRepository, stravaBase *strava.Base,
	stravaRepo StravaRepository) *Service {
	return &Service{
		activityRepo: activityRepo,
		stravaBase:   stravaBase,
		stravaRepo:   stravaRepo,
	}
}

type Service struct {
	activityRepo ActivitiesRepository
	stravaBase   *strava.Base
	stravaRepo   StravaRepository
}

type ActivitiesRepository interface {
	Get(ctx context.Context, userID int, filters domain.ActivityFilters) (
		[]domain.Activity, error)
	InsertBulk(ctx context.Context, activities []domain.Activity) error
}

type StravaRepository interface {
	Insert(ctx context.Context, accessToken string, refreshToken string, userId int) error
	UpdateUserStravaInfo(ctx context.Context, accessToken string,
		refreshToken string, userId int) error
}