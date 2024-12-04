package activities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
	"github.com/pttrulez/activitypeople/internal/infra/strava"
)

func (s *Service) GetActivities(ctx context.Context, user domain.User,
	filters domain.ActivityFilters) ([]domain.Activity, error) {
	activities, err := s.activityRepo.Get(ctx, user.Id, filters)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (s *Service) HydrateStravaActivity(ctx context.Context, sourceId int,
	u domain.User) error {
	stravaInfo, err := s.stravaRepo.GetByUserId(ctx, u.Id)
	if err != nil {
		return err
	}

	client := s.stravaBase.NewClient(*stravaInfo.AccessToken, *stravaInfo.RefreshToken,
		s.makeStoreTokensFunc(ctx, u.Id))

	activity, err := client.GetActivity(ctx, sourceId)
	if err != nil {
		return fmt.Errorf("get activity from strava:  %w", err)
	}

	calories := int(activity.Calories * 0.89)

	return s.activityRepo.UpdateCalories(ctx, calories, sourceId, u.Id)
}

func (s *Service) SyncStravaActivities(ctx context.Context, user domain.User) error {
	stravaInfo, err := s.stravaRepo.GetByUserId(ctx, user.Id)
	if err != nil {
		return err
	}
	client := s.stravaBase.NewClient(*stravaInfo.AccessToken, *stravaInfo.RefreshToken,
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
	after := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()

	for isPolling {
		// Запрашиваем активности из апи стравы
		stravaActivities, err := client.GetAthleteActivities(ctx, &after)
		if err != nil {
			return err
		}

		activities := make([]domain.Activity, 0, len(stravaActivities))
		for _, activity := range stravaActivities {
			a := FromStravaToActivity(activity)
			activities = append(activities, a)
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
				after = a.StartTimeUnix
			}
		}

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

func (s *Service) OAuthStrava(ctx context.Context, userCode string, userID int) error {
	client := s.stravaBase.NewClient("", "", s.makeStoreTokensFunc(ctx, userID))

	data, err := client.OAuth(userCode)
	if err != nil {
		return err
	}
	_, err = s.stravaRepo.GetByUserId(ctx, userID)
	if errors.Is(err, store.ErrNotFound) {
		//  Вставляем инфу, если авторизуется в страве первый раз
		err = s.stravaRepo.Insert(ctx, data.AccessToken, data.RefreshToken, userID)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	// Апдейтим инфу, если такая уже есть у юзера
	err = s.stravaRepo.UpdateUserStravaInfo(ctx, data.AccessToken, data.RefreshToken,
		userID)
	if err != nil {
		return err
	}

	return nil
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
	UpdateCalories(ctx context.Context, calories, sourceId, userID int) error
}

type StravaRepository interface {
	GetByUserId(ctx context.Context, userId int) (domain.StravaInfo, error)
	Insert(ctx context.Context, accessToken string, refreshToken string, userId int) error
	UpdateUserStravaInfo(ctx context.Context, accessToken string,
		refreshToken string, userId int) error
}
