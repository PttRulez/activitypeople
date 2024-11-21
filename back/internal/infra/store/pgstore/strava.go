package pgstore

import (
	"context"
	"database/sql"
	"fmt"
)

func (pg *StravaPostgres) Insert(ctx context.Context, accessToken string,
	refreshToken string, userId int) error {
	queryString := `INSERT INTO strava_info (user_id, access_token, refresh_token) 
		VALUES ($1, $2, $3);`

	_, err := pg.db.ExecContext(ctx, queryString, userId, accessToken, refreshToken)
	if err != nil {
		return fmt.Errorf("[StravaPostgres.Insert]: %w", err)
	}

	return nil
}

func (pg *StravaPostgres) UpdateUserStravaInfo(ctx context.Context, accessToken string,
	refreshToken string, userId int) error {
	queryString := `UPDATE strava_info SET access_token = $1, refresh_token = $2 WHERE user_id = $3;`
	_, err := pg.db.ExecContext(ctx, queryString, accessToken, refreshToken, userId)
	if err != nil {
		return fmt.Errorf("[StravaPostgres.UpdateUserStravaInfo]: %w", err)
	}

	return nil
}

type StravaPostgres struct {
	db *sql.DB
}

func NewStravaPostgres(db *sql.DB) *StravaPostgres {
	return &StravaPostgres{db: db}
}
