package pgstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"

	_ "github.com/lib/pq"
)

func (pg *UserPostgres) Insert(ctx context.Context, email, hashedPassword, name string) (
	domain.User, error) {
	querySting := `INSERT INTO users (email, hashed_password, name, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, name, role;`

	row := pg.db.QueryRowContext(ctx, querySting, email, hashedPassword, name, domain.Scoof)

	var user domain.User
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Role)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (pg *UserPostgres) GetByEmail(ctx context.Context, email string) (domain.User,
	error) {
	querySting := `SELECT u.id, u.email, u.hashed_password, u.name, u.role, s.access_token,
		s.refresh_token
    FROM users u
    LEFT JOIN strava_info s ON s.user_id = u.id
    WHERE u.email = $1 LIMIT 1;`

	row := pg.db.QueryRowContext(ctx, querySting, email)

	var u domain.User
	err := row.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role,
		&u.Strava.AccessToken, &u.Strava.RefreshToken)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, store.ErrNotFound
	}
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
