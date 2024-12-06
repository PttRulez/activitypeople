package pgstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"

	_ "github.com/lib/pq"
)

func (pg *UserPostgres) Insert(ctx context.Context, email, hashedPassword, name string) (
	domain.User, error) {
	const op = "UserPostgres.Insert"

	q := pg.sq.Insert("users").
		Columns("email", "hashed_password", "name", "role").
		Values(email, hashedPassword, name, domain.RegularUser).
		Suffix("RETURNING id, email, name, role")

	var user domain.User
	err := q.RunWith(pg.db).ScanContext(ctx, &user.Id, &user.Email, &user.Name, &user.Role)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (pg *UserPostgres) GetById(ctx context.Context, userID int) (domain.User,
	error) {
	const op = "UserPostgres.GetByEmail"

	q := pg.sq.Select("u.bmr", "u.calories_per_100_steps", "u.id", "u.email",
		"u.hashed_password", "u.name", "u.role", "s.access_token", "s.refresh_token").
		From("users u").
		LeftJoin("strava_info s ON s.user_id = u.id").
		Where(sq.Eq{"u.id": userID})

	var u domain.User
	err := q.RunWith(pg.db).ScanContext(ctx, &u.BMR, &u.CaloriesPer100Steps, &u.Id, &u.Email,
		&u.HashedPassword, &u.Name, &u.Role, &u.Strava.AccessToken, &u.Strava.RefreshToken)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, store.ErrNotFound
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, nil
}

func (pg *UserPostgres) GetByEmail(ctx context.Context, email string) (domain.User,
	error) {
	const op = "UserPostgres.GetByEmail"

	q := pg.sq.Select("u.bmr", "u.calories_per_100_steps", "u.id", "u.email", "u.hashed_password",
		"u.name", "u.role", "s.access_token", "s.refresh_token").
		From("users u").
		LeftJoin("strava_info s ON s.user_id = u.id").
		Where(sq.Eq{"u.email": email})

	var u domain.User
	err := q.RunWith(pg.db).ScanContext(ctx, &u.BMR, &u.CaloriesPer100Steps, &u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role,
		&u.Strava.AccessToken, &u.Strava.RefreshToken)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, store.ErrNotFound
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, nil
}

func (pg *UserPostgres) SaveSettings(ctx context.Context, f domain.UserSettings, userId int) error {
	const op = "UserPostgres.SaveSettings"

	q := pg.sq.Update("users").
		Set("bmr", f.BMR).
		Set("calories_per_100_steps", f.CaloriesPer100Steps).
		Where(sq.Eq{"id": userId})

	_, err := q.RunWith(pg.db).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

type UserPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
