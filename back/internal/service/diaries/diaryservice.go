package diaries

import (
	"context"
	"fmt"
	"time"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func (s *DiaryService) GetDiaries(ctx context.Context, userID int, from,
	until time.Time) (map[time.Time]domain.DiaryDay, error) {
	activities, err := s.activityRepo.Get(ctx, userID, domain.ActivityFilters{
		From:  from,
		Until: until,
	})
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	meals, err := s.mealRepo.Get(ctx, domain.TimeFilters{
		From:  from,
		Until: until,
	}, userID)
	if err != nil {
		return nil, err
	}

	weights, err := s.weightRepo.Get(ctx, userID, domain.TimeFilters{
		From:  from,
		Until: until,
	})
	if err != nil {
		return nil, err
	}

	steps, err := s.stepsRepo.Get(ctx, userID, domain.TimeFilters{
		From:  from,
		Until: until,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("STEPS: %+v\n", steps)

	diariesMap := make(map[time.Time]domain.DiaryDay, 0)

	for _, a := range activities {
		if d, ok := diariesMap[a.Date]; ok {
			d.Activities = append(d.Activities, a)
		} else {
			diariesMap[a.Date] = domain.DiaryDay{
				Activities: []domain.Activity{a},
				Date:       a.Date,
			}
		}
	}
	for _, m := range meals {
		if d, ok := diariesMap[m.Date]; ok {
			d.Meals = append(d.Meals, m)
			diariesMap[m.Date] = d
		} else {
			diariesMap[m.Date] = domain.DiaryDay{
				Date:  m.Date,
				Meals: []domain.Meal{m},
			}
		}
	}
	for _, w := range weights {
		if d, ok := diariesMap[w.Date]; ok {
			d.Weight = w.Weight
			diariesMap[w.Date] = d
		} else {
			diariesMap[w.Date] = domain.DiaryDay{Date: w.Date, Weight: w.Weight}
		}
	}
	for _, s := range steps {
		fmt.Println(steps)
		if d, ok := diariesMap[s.Date]; ok {
			d.Steps = s.Steps
			diariesMap[s.Date] = d
		} else {
			diariesMap[s.Date] = domain.DiaryDay{Date: s.Date, Steps: s.Steps}
		}
	}

	for date, day := range diariesMap {
		diariesMap[date] = day.CalculateCalories(user)
	}

	return diariesMap, nil
}

func NewService(
	activityRepo ActivityRepository,
	mealRepo MealRepository,
	stepsRepo StepsRepository,
	userRepo UserRepository,
	weightRepo WeightRepository,
) *DiaryService {
	return &DiaryService{
		activityRepo: activityRepo,
		mealRepo:     mealRepo,
		stepsRepo:    stepsRepo,
		userRepo:     userRepo,
		weightRepo:   weightRepo,
	}
}

type DiaryService struct {
	activityRepo ActivityRepository
	mealRepo     MealRepository
	stepsRepo    StepsRepository
	userRepo     UserRepository
	weightRepo   WeightRepository
}

type ActivityRepository interface {
	Get(ctx context.Context, userID int, filters domain.ActivityFilters) (
		[]domain.Activity, error)
}

type MealRepository interface {
	Get(ctx context.Context, f domain.TimeFilters, userID int) (
		[]domain.Meal, error)
}

type StepsRepository interface {
	Get(ctx context.Context, userID int,
		f domain.TimeFilters) ([]domain.Steps, error)
}
type UserRepository interface {
	GetById(ctx context.Context, userID int) (domain.User,
		error)
}
type WeightRepository interface {
	Get(ctx context.Context, userID int,
		f domain.TimeFilters) ([]domain.Weight, error)
}
