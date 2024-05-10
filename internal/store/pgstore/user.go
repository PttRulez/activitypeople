package pgstore

import (
	"antiscoof/internal/model"
	"context"
	"database/sql"
	"fmt"
)

func (pg *UserPostgres) Insert(ctx context.Context, u *model.User) (int, error) {
	querySting := "INSERT INTO users (email, hashed_password, name, role) VALUES ($1, $2, $3, $4);"
	res, err := pg.db.ExecContext(ctx, querySting, u.Email, u.HashedPassword, u.Name, u.Role)
	if err != nil {
		return 0, fmt.Errorf("[UserPostgres.Insert]: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("[UserPostgres.Insert]: %w", err)
	}

	return int(id), nil
}

func (pg *UserPostgres) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	querySting := `SELECT * FROM users WHERE email = $1 LIMIT 1;`
	row := pg.db.QueryRowContext(ctx, querySting, email)
	if row.Err() != nil {
		return nil, fmt.Errorf("[UserPostgres.GetByEmail]: %w", row.Err())
	}

	var u model.User
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

func (pg *UserPostgres) GetById(ctx context.Context, id int) (*model.User, error) {
	querySting := `SELECT * FROM users WHERE id = $1 LIMIT 1;`
	row := pg.db.QueryRowContext(ctx, querySting, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("[UserPostgres.GetById]: %w", row.Err())
	}

	var u model.User
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
