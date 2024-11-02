package pgstore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pttrulez/activitypeople/internal/domain"

	_ "github.com/lib/pq"
)

func (pg *UserPostgres) Insert(ctx context.Context, u *domain.User) (int, error) {
	querySting := "INSERT INTO users (email, hashed_password, name, role) VALUES ($1, $2, $3, $4) RETURNING id;"
	row := pg.db.QueryRowContext(ctx, querySting, u.Email, u.HashedPassword, u.Name, u.Role)
	if row.Err() != nil {
		return 0, fmt.Errorf("[UserPostgres.Insert]: %w", row.Err())
	}
	var user domain.User
	err := row.Scan(&user.Id)
	if err != nil {
		return 0, fmt.Errorf("[UserPostgres.Insert]: %w", err)
	}

	return user.Id, nil
}

func (pg *UserPostgres) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	querySting := `SELECT * FROM users WHERE email = $1 LIMIT 1;`
	row := pg.db.QueryRowContext(ctx, querySting, email)
	if row.Err() != nil {
		return nil, fmt.Errorf("[UserPostgres.GetByEmail]: %w", row.Err())
	}

	var u domain.User
	err := row.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		break
	default:
		return nil, fmt.Errorf("[UserPostgres.GetByEmail]: %w", err)
	}

	return &u, nil
}

func (pg *UserPostgres) GetById(ctx context.Context, id int) (*domain.User, error) {
	querySting := `SELECT * FROM users WHERE id = $1 LIMIT 1;`
	row := pg.db.QueryRowContext(ctx, querySting, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("[UserPostgres.GetById]: %w", row.Err())
	}

	var u domain.User
	err := row.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	return &u, nil
}

func (pg *UserPostgres) GetByIdWithStrava(ctx context.Context, id int) (*domain.User, error) {
	querySting := `SELECT u.*, s.access_token, s.refresh_token, s.user_id
    FROM users u
    LEFT JOIN strava_info s ON s.user_id = u.id
    WHERE u.id = $1 LIMIT 1;`
	row := pg.db.QueryRowContext(ctx, querySting, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("[UserPostgres.GetById]: %w", row.Err())
	}

	var strava domain.StravaInfo
	var hasStrava *int
	var u domain.User
	err := row.Scan(&u.Id, &u.Email, &u.HashedPassword, &u.Name, &u.Role,
		&strava.AccessToken, &strava.RefreshToken, &hasStrava)
	if hasStrava == nil {
		u.Strava = nil
	} else {
		u.Strava = &strava
	}
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		break
	default:
		return nil, err
	}

	return &u, nil
}

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func CreateUsersTable(db *sql.DB) {
	querySting := `CREATE TABLE IF NOT EXISTS users
  (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL
  );`
	_, err := db.Exec(querySting)
	if err != nil {
		panic(err)
	}
}

func DropUsersTable(db *sql.DB) {
	querySting := `DROP TABLE IF EXISTS users;`
	_, err := db.Exec(querySting)
	if err != nil {
		panic(err)
	}
}
