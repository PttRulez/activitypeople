package pgstore

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pttrulez/activitypeople/internal/domain"
)

func (pg *FoodPostgres) Delete(ctx context.Context, foodID int, userID int) error {
	const op = "FoodPostgres.Delete"

	q := sq.Delete("foods").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"id": foodID})

	_, err := q.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *FoodPostgres) Insert(ctx context.Context, a domain.Food) error {
	const op = "FoodPostgres.Insert"
	q := sq.Insert("foods").
		Columns("carbs", "calories", "fat", "name", "public", "protein", "user_id").
		Values(a.Carbs, a.Calories, a.Fat, a.Name, a.Public, a.Protein, a.UserID)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *FoodPostgres) Search(ctx context.Context, q string) ([]domain.Food, error) {
	const op = "FoodPostgres.Search"

	query := sq.Select("carbs", "calories", "fat", "name", "public", "protein", "user_id").
		From("foods").
		Where(sq.Like{"name": q})

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var foods []domain.Food
	for rows.Next() {
		var f domain.Food
		err := rows.Scan(&f.Carbs, &f.Calories, &f.Fat, &f.Name, &f.Public, &f.Protein, &f.UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		foods = append(foods, f)
	}

	return foods, nil
}

type FoodPostgres struct {
	db *sql.DB
}

func NewFoodPostgres(db *sql.DB) *FoodPostgres {
	return &FoodPostgres{db: db}
}
