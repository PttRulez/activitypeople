package diaries

import (
	"context"
	"fmt"
	"time"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func (s *DiaryService) GetDiaries(ctx context.Context, u domain.User, from,
	until time.Time) (map[time.Time]domain.DiaryDay, error) {
	activities, err := s.activityRepo.Get(ctx, u.Id, domain.ActivityFilters{
		From:  from,
		Until: until,
	})
	if err != nil {
		return nil, err
	}

	meals, err := s.mealRepo.Get(ctx, domain.MealFilters{
		From:  from,
		Until: until,
	}, u.Id)
	if err != nil {
		return nil, err
	}

	weights, err := s.weightRepo.Get(ctx, u.Id, domain.WeightFilters{
		From:  from,
		Until: until,
	})
	if err != nil {
		return nil, err
	}

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

	fmt.Printf("u.BMR %+v\n", u)
	for date, day := range diariesMap {
		diariesMap[date] = day.CalculateCalories(u.BMR)
	}

	return diariesMap, nil
}

func NewService(activityRepo ActivityRepository, mealRepo MealRepository,
	weightRepo WeightRepository) *DiaryService {
	return &DiaryService{
		activityRepo: activityRepo,
		mealRepo:     mealRepo,
		weightRepo:   weightRepo,
	}
}

type DiaryService struct {
	activityRepo ActivityRepository
	mealRepo     MealRepository
	weightRepo   WeightRepository
}

type ActivityRepository interface {
	Get(ctx context.Context, userID int, filters domain.ActivityFilters) (
		[]domain.Activity, error)
}

type MealRepository interface {
	Get(ctx context.Context, f domain.MealFilters, userID int) (
		[]domain.Meal, error)
}

type WeightRepository interface {
	Get(ctx context.Context, userID int,
		f domain.WeightFilters) ([]domain.Weight, error)
}
