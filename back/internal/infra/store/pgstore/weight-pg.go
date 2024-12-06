package pgstore

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pttrulez/activitypeople/internal/domain"
)

func (pg *WeightPostgres) Insert(ctx context.Context, w domain.Weight,
	userID int) error {
	const op = "WeightPostgres.Insert"

	q := pg.sq.Insert("weights").
		Columns("date", "weight", "user_id").
		Values(w.Date, w.Weight, userID)

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *WeightPostgres) Get(ctx context.Context, userID int,
	f domain.TimeFilters) ([]domain.Weight, error) {
	const op = "MealPostgres.Get"
	
	q := pg.sq.Select("date", "weight").
		From("weights").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.GtOrEq{"date": f.From}).
		Where(sq.LtOrEq{"date": f.Until}).
		OrderBy("date DESC")

	rows, err := q.RunWith(pg.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	weights := make([]domain.Weight, 0)
	for rows.Next() {
		var w domain.Weight
		err := rows.Scan(&w.Date, &w.Weight)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		weights = append(weights, w)
	}

	return weights, nil
}

type WeightPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewWeightPostgres(db *sql.DB) *WeightPostgres {
	return &WeightPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
