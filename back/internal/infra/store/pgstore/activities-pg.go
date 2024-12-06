package pgstore

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
)

func (pg *ActivitiesPostgres) GetLatestStravaUnix(ctx context.Context, userID int) (int64, error) {
	const op = "ActivitiesPostgres.GetLatestStravaUnix"

	q := pg.sq.Select("max(start_time_unix)").
		From("activities").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"source": domain.Strava})

	row := q.RunWith(pg.db).QueryRowContext(ctx)
	var i int64
	err := row.Scan(&i)
	if err == sql.ErrNoRows {
		return i, store.ErrNotFound
	} else if err != nil {
		return i, fmt.Errorf("%s: %w", op, err)
	}

	return i, nil
}

func (pg *ActivitiesPostgres) Get(ctx context.Context, userID int,
	filters domain.ActivityFilters) ([]domain.Activity, error) {
	const op = "ActivitiesPostgres.Get"

	q := pg.sq.Select("calories", "description", "distance", "date", "elevate",
		"heartrate", "id", "name", "pace", "pace_string", "source", "source_id", "sport_type",
		"start_time_unix", "total_time").
		From("activities").
		Where(sq.Eq{"user_id": userID})

	if filters.Source != "" {
		q = q.Where(sq.Eq{"source": filters.Source})
	}

	if !filters.From.IsZero() {
		q = q.Where(sq.GtOrEq{"date": filters.From})
	}

	if !filters.Until.IsZero() {
		q = q.Where(sq.LtOrEq{"date": filters.Until})
	}

	q = q.OrderBy("date DESC")

	rows, err := q.RunWith(pg.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var activities []domain.Activity
	for rows.Next() {
		var a domain.Activity
		err = rows.Scan(
			&a.Calories,
			&a.Description,
			&a.Distance,
			&a.Date,
			&a.Elevate,
			&a.Heartrate,
			&a.Id,
			&a.Name,
			&a.Pace,
			&a.PaceString,
			&a.Source,
			&a.SourceId,
			&a.SportType,
			&a.StartTimeUnix,
			&a.TotalTime,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		activities = append(activities, a)
	}

	return activities, nil
}

func (pg *ActivitiesPostgres) Insert(ctx context.Context, a domain.Activity) error {
	const op = "ActivitiesPostgres.Insert"

	q := pg.sq.Insert("activities").
		Columns("calories", "description", "distance", "date", "elevate", "heartrate",
			"name", "pace", "pace_string", "source", "source_id", "sport_type", "start_time_unix",
			"total_time", "user_id").
		Values(a.Calories, a.Description, a.Distance, a.Date, a.Elevate, a.Heartrate,
			a.Name, a.Pace, a.PaceString, a.Source, a.SourceId, a.SportType, a.StartTimeUnix,
			a.TotalTime, a.UserId)

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *ActivitiesPostgres) UpdateCalories(ctx context.Context,
	calories, sourceId, userID int) error {
	const op = "ActivitiesPostgres.UpdateCalories"

	q := pg.sq.Update("activities").
		Set("calories", calories).
		Where(sq.Eq{"source_id": sourceId}).
		Where(sq.Eq{"user_id": userID})

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *ActivitiesPostgres) InsertBulk(ctx context.Context, activities []domain.Activity) error {
	const op = "ActivitiesPostgres.InsertBulk"

	q := pg.sq.Insert("activities").
		Columns("calories", "description", "distance", "date", "elevate", "heartrate",
			"name", "pace", "pace_string", "source", "source_id", "sport_type",
			"start_time_unix", "total_time", "user_id")

	for _, a := range activities {
		q = q.Values(
			a.Calories,
			a.Description,
			a.Distance,
			a.Date.Format("2006-01-02"),
			a.Elevate,
			a.Heartrate,
			a.Name,
			a.Pace,
			a.PaceString,
			a.Source,
			a.SourceId,
			a.SportType,
			a.StartTimeUnix,
			a.TotalTime,
			a.UserId)
	}

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

type ActivitiesPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewActivitiesPostgres(db *sql.DB) *ActivitiesPostgres {
	return &ActivitiesPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
