// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: summoners.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const listAll = `-- name: ListAll :one
SELECT id, puuid, account_id, summoner_id, level, profile_icon_id, name, last_revision, time_stamp FROM summoners
WHERE id = $1 LIMIT 1
`

func (q *Queries) ListAll(ctx context.Context, id pgtype.UUID) (Summoner, error) {
	row := q.db.QueryRow(ctx, listAll, id)
	var i Summoner
	err := row.Scan(
		&i.ID,
		&i.Puuid,
		&i.AccountID,
		&i.SummonerID,
		&i.Level,
		&i.ProfileIconID,
		&i.Name,
		&i.LastRevision,
		&i.TimeStamp,
	)
	return i, err
}
