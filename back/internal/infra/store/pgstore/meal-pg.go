package pgstore

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/pttrulez/activitypeople/internal/domain"
)

func (pg *MealPostgres) Insert(ctx context.Context, m domain.Meal) error {
	const op = "MealPostgres.Insert"

	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	q := pg.sq.Insert("meals").Columns("calories", "date", "name", "user_id").
		Values(m.Calories, m.Date, m.Name, m.UserId).
		Suffix("RETURNING \"id\"")

	row, err := q.RunWith(tx).QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	row.Next()
	var mealId int
	err = row.Scan(&mealId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	row.Close()

	q = pg.sq.Insert("foods_on_meals").Columns("calories", "calories_per_100", "meal_id", "name", "weight")

	for _, f := range m.Foods {
		q = q.Values(&f.Calories, &f.CaloriesPer100, mealId, &f.Name, &f.Weight)
	}

	_, err = q.RunWith(tx).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *MealPostgres) Get(ctx context.Context, f domain.TimeFilters, userID int) (
	[]domain.Meal, error) {
	const op = "MealPostgres.Get"

	tx, err := pg.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Meals
	q := pg.sq.Select("calories", "date", "id", "name").
		From("meals").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.GtOrEq{"date": f.From}).
		Where(sq.LtOrEq{"date": f.Until}).
		OrderBy("date DESC")

	rows, err := q.RunWith(tx).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	meals := make([]domain.Meal, 0)

	for rows.Next() {
		var m domain.Meal
		err := rows.Scan(&m.Calories, &m.Date, &m.Id, &m.Name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		meals = append(meals, m)
	}

	// FoodsInMeals
	mealIDs := make([]int, len(meals))
	for i, m := range meals {
		mealIDs[i] = m.Id
	}

	q = pg.sq.Select("calories", "calories_per_100", "meal_id", "name", "weight").
		From("foods_on_meals").
		Where("meal_id = any(?)", pq.Array(mealIDs))

	rows, err = q.RunWith(tx).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	foods := make([]domain.FoodInMeal, 0)
	for rows.Next() {
		var f domain.FoodInMeal
		err := rows.Scan(&f.Calories, &f.CaloriesPer100, &f.MealId, &f.Name, &f.Weight)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		foods = append(foods, f)
	}

	foodMap := make(map[int][]domain.FoodInMeal)

	for _, f := range foods {
		foodMap[f.MealId] = append(foodMap[f.MealId], f)
	}

	for i, m := range meals {
		meals[i].Foods = foodMap[m.Id]
	}

	return meals, nil
}

type MealPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewMealPostgres(db *sql.DB) *MealPostgres {
	return &MealPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
