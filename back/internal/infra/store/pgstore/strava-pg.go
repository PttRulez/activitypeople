package pgstore

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
)

func (pg *StravaPostgres) GetByUserId(ctx context.Context, userId int) (domain.StravaInfo, error) {

	query := pg.sq.
		Select("access_token", "refresh_token").
		From("strava_info").
		Where(sq.Eq{"user_id": userId})
	row := query.RunWith(pg.db).QueryRowContext(ctx)
	var i domain.StravaInfo
	err := row.Scan(&i.AccessToken, &i.RefreshToken)

	if err == sql.ErrNoRows {
		return domain.StravaInfo{}, store.ErrNotFound
	} else if err != nil {
		return domain.StravaInfo{}, err
	}

	return i, nil
}

func (pg *StravaPostgres) Insert(ctx context.Context, accessToken string,
	refreshToken string, userId int) error {
	const op = "StravaPostgres.Insert"

	queryString := `INSERT INTO strava_info (user_id, access_token, refresh_token) 
		VALUES ($1, $2, $3);`

	_, err := pg.db.ExecContext(ctx, queryString, userId, accessToken, refreshToken)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (pg *StravaPostgres) UpdateUserStravaInfo(ctx context.Context, accessToken string,
	refreshToken string, userId int) error {
	const op = "StravaPostgres.UpdateUserStravaInfo"

	queryString := `UPDATE strava_info SET access_token = $1, refresh_token = $2 WHERE user_id = $3;`
	_, err := pg.db.ExecContext(ctx, queryString, accessToken, refreshToken, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

type StravaPostgres struct {
	db *sql.DB
	sq sq.StatementBuilderType
}

func NewStravaPostgres(db *sql.DB) *StravaPostgres {
	return &StravaPostgres{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
