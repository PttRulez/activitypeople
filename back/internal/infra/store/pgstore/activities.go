package pgstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pttrulez/activitypeople/internal/domain"
)

func (pg *ActivitiesPostgres) Get(ctx context.Context, userID int,
	filters domain.ActivityFilters) ([]domain.Activity, error) {

	queryString := `SELECT calories, description, distance, date, elevate, heartrate, name,
		pace, pace_string, source, source_id, sport_type, total_time
    FROM activities WHERE user_id = $1`

	argCount := 2
	argArr := make([]interface{}, 1, argCount)
	argArr[0] = userID

	if filters.Source != "" {
		queryString += fmt.Sprintf(" AND source = $%d", argCount)
		argArr = append(argArr, filters.Source)
		argCount++
	}
	if !filters.From.IsZero() {
		queryString += fmt.Sprintf(" AND date >= $%d", argCount)
		argArr = append(argArr, filters.From)
		argCount++
	}
	if !filters.Until.IsZero() {
		queryString += fmt.Sprintf(" AND date <= $%d", argCount)
		argArr = append(argArr, filters.Until)
		argCount++
	}
	queryString += " ORDER BY date DESC;"

	rows, err := pg.db.QueryContext(ctx, queryString, argArr...)
	if err != nil {
		return nil, fmt.Errorf("[ActivitiesPostgres.Get]: %w", err)
	}
	defer rows.Close()
	var activities []domain.Activity
	for rows.Next() {
		var activity domain.Activity
		err = rows.Scan(
			&activity.Calories,
			&activity.Description,
			&activity.Distance,
			&activity.Date,
			&activity.Elevate,
			&activity.Heartrate,
			&activity.Name,
			&activity.Pace,
			&activity.PaceString,
			&activity.Source,
			&activity.SourceId,
			&activity.SportType,
			&activity.TotalTime,
		)
		if err != nil {
			return nil, fmt.Errorf("[ActivitiesPostgres.Get]: %w", err)
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

func (pg *ActivitiesPostgres) Insert(ctx context.Context, a domain.Activity) error {
	queryString := `INSERT INTO activities (calories, description, distance, date, elevate, heartrate, name,
		pace, pace_string, source, source_id, sport_type, total_time, user_id) VALUES`

	_, err := pg.db.ExecContext(ctx, queryString,
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
		a.TotalTime,
		a.UserId)
	if err != nil {
		return fmt.Errorf("[ActivitiesPostgres.InsertBulk]: %w", err)
	}

	return nil
}

func (pg *ActivitiesPostgres) InsertBulk(ctx context.Context, activities []domain.Activity) error {
	queryString := `INSERT INTO activities (calories, description, distance, date, elevate, 
		heartrate, name, pace, pace_string, source, source_id, sport_type, total_time, user_id)
		VALUES`

	args := make([]interface{}, 0, len(activities)*8)
	for i, a := range activities {
		i = i * 14
		queryString += fmt.Sprintf("($%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v, $%v),",
			i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9, i+10, i+11, i+12, i+13, i+14)
		args = append(args,
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
			a.TotalTime,
			a.UserId)
	}
	queryString = strings.TrimSuffix(queryString, ",")
	queryString = queryString + ";"

	_, err := pg.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return fmt.Errorf("[ActivitiesPostgres.InsertBulk]: %w", err)
	}

	return nil
}

type ActivitiesPostgres struct {
	db *sql.DB
}

func NewActivitiesPostgres(db *sql.DB) *ActivitiesPostgres {
	return &ActivitiesPostgres{db: db}
}
