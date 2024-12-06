package pgstore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
)

func (pg *StepsPostgres) Get(ctx context.Context, userID int, f domain.TimeFilters) (
	[]domain.Steps, error) {
	const op = "StepsPostgres.Get"

	q := pg.sq.Select("date", "steps").
		From("steps").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.GtOrEq{"date": f.From}).
		Where(sq.LtOrEq{"date": f.Until}).
		OrderBy("date DESC")

	rows, err := q.RunWith(pg.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	steps := make([]domain.Steps, 0)
	for rows.Next() {
		fmt.Println("letsgo")
		var w domain.Steps
		err = rows.Scan(&w.Date, &w.Steps)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		steps = append(steps, w)
	}

	return steps, nil
}

func (pg *StepsPostgres) GetByDate(ctx context.Context, date time.Time, userID int) (
	domain.Steps, error) {
	const op = "StepsPostgres.GetByDate"

	q := pg.sq.Select("date", "steps").
		From("steps").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"date": date})

	row := q.RunWith(pg.db).QueryRowContext(ctx)
	var s domain.Steps
	err := row.Scan(&s.Date, &s.Steps)
	if err == sql.ErrNoRows {
		return s, store.ErrNotFound
	} else if err != nil {
		return s, fmt.Errorf("%s: %w", op, err)
	}

	return s, nil
}

func (pg *StepsPostgres) Insert(ctx context.Context, s domain.Steps, userID int) error {
	const op = "StepPostgres.Insert"

	q := pg.sq.Insert("steps").
		Columns("date", "steps", "user_id").
		Values(s.Date, s.Steps, userID)

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *StepsPostgres) Update(ctx context.Context, s domain.Steps, userID int) error {
	const op = "StepPostgres.Insert"

	q := pg.sq.Update("steps").
		Set("steps", s.Steps).
		Where(sq.Eq{"date": s.Date}).
		Where(sq.Eq{"user_id": userID})

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

type StepsPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewStepsPostgres(db *sql.DB) *StepsPostgres {
	return &StepsPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
