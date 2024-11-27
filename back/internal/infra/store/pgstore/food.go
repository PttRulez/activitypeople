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

	q := pg.sq.Delete("foods").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"id": foodID})

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *FoodPostgres) Insert(ctx context.Context, f domain.Food) error {
	const op = "FoodPostgres.Insert"
	q := pg.sq.Insert("foods").
		Columns("calories", "carbs", "created_by_admin", "fat", "name", "protein", "user_id").
		Values(f.Calories, f.Carbs, f.CreatedByAdmin, f.Fat, f.Name, f.Protein, f.UserID)

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *FoodPostgres) Search(ctx context.Context, q string) ([]domain.Food, error) {
	const op = "FoodPostgres.Search"

	query := pg.sq.Select("id", "carbs", "calories", "fat", "name", "created_by_admin",
		"protein", "user_id").
		From("foods").
		Where(sq.ILike{"name": fmt.Sprintf("%%%v%%", q)})

	rows, err := query.RunWith(pg.db).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	fmt.Println(query.ToSql())
	foods := make([]domain.Food, 0)
	for rows.Next() {
		var f domain.Food
		err := rows.Scan(&f.ID, &f.Carbs, &f.Calories, &f.Fat, &f.Name, &f.CreatedByAdmin,
			&f.Protein, &f.UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		foods = append(foods, f)
	}

	return foods, nil
}

type FoodPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewFoodPostgres(db *sql.DB) *FoodPostgres {
	return &FoodPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
