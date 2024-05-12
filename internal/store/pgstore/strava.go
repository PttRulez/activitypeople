package pgstore

import (
	"antiscoof/internal/model"
	"context"
	"database/sql"
	"fmt"
)

func (pg *StravaPostgres) Insert(ctx context.Context, s *model.StravaInfo) error {
	queryString := `INSERT INTO strava_info (user_id, access_token, refresh_token) 
		VALUES ($1, $2, $3);`

	_, err := pg.db.ExecContext(ctx, queryString, s.UserId, s.AccessToken, s.RefreshToken)
	if err != nil {
		return fmt.Errorf("[StravaPostgres.Insert]: %w", err)
	}

	return nil
}

func (pg *StravaPostgres) UpdateUserStravaInfo(ctx context.Context, s *model.UpdateStravaTokens) error {
	queryString := `UPDATE strava_info SET access_token = $1, refresh_token = $2 WHERE user_id = $3;`
	fmt.Printf("UpdateUserStravaInfo: %v, %v, %v\n", s.AccessToken, s.RefreshToken, s.UserId)
	_, err := pg.db.ExecContext(ctx, queryString, s.AccessToken, s.RefreshToken, s.UserId)
	fmt.Printf("UpdateUserStravaInfo err: %v\n", err)
	if err != nil {
		return fmt.Errorf("[StravaPostgres.UpdateUserStravaInfo]: %w", err)
	}

	return nil
}

func (pg *StravaPostgres) GetByUserId(ctx context.Context, userId int) (*model.StravaInfo,
	error) {
	queryString := `SELECT * FROM strava_info WHERE user_id = $1 LIMIT 1;`
	row := pg.db.QueryRowContext(ctx, queryString, userId)

	var s model.StravaInfo
	err := row.Scan(&s.Id, &s.UserId, &s.AccessToken, &s.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("[StravaPostgres.GetByUserId]: %w", err)
	}

	return &s, nil
}

type StravaPostgres struct {
	db *sql.DB
}

func NewStravaPostgres(db *sql.DB) *StravaPostgres {
	return &StravaPostgres{db: db}
}

func CreateStravaInfoTable(db *sql.DB) {
	querySting := `CREATE TABLE IF NOT EXISTS strava_info
  (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL,
    access_token VARCHAR(255),
    refresh_token VARCHAR(255)
  );`
	_, err := db.Exec(querySting)
	if err != nil {
		panic(err)
	}
}

func DropStravaInfoTable(db *sql.DB) {
	querySting := `DROP TABLE IF EXISTS strava_info;`
	_, err := db.Exec(querySting)
	if err != nil {
		panic(err)
	}
}
