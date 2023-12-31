// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: summoners.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countSummonerRecords = `-- name: CountSummonerRecords :one
SELECT count(*)
FROM summoner_records
`

func (q *Queries) CountSummonerRecords(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countSummonerRecords)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countSummonerRecordsByPuuid = `-- name: CountSummonerRecordsByPuuid :one
SELECT count(*)
FROM summoner_records
WHERE puuid = $1
`

func (q *Queries) CountSummonerRecordsByPuuid(ctx context.Context, puuid string) (int64, error) {
	row := q.db.QueryRow(ctx, countSummonerRecordsByPuuid, puuid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteSummoner = `-- name: DeleteSummoner :exec
DELETE FROM summoner_records WHERE record_id = $1
`

func (q *Queries) DeleteSummoner(ctx context.Context, recordID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteSummoner, recordID)
	return err
}

const insertSummoner = `-- name: InsertSummoner :one
INSERT INTO summoner_records (puuid, account_id, id, summoner_level, profile_icon_id, name, revision_date, record_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING record_id
`

type InsertSummonerParams struct {
	Puuid         string
	AccountID     string
	ID            string
	SummonerLevel int64
	ProfileIconID int32
	Name          string
	RevisionDate  int64
	RecordDate    pgtype.Timestamp
}

func (q *Queries) InsertSummoner(ctx context.Context, arg InsertSummonerParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, insertSummoner,
		arg.Puuid,
		arg.AccountID,
		arg.ID,
		arg.SummonerLevel,
		arg.ProfileIconID,
		arg.Name,
		arg.RevisionDate,
		arg.RecordDate,
	)
	var record_id pgtype.UUID
	err := row.Scan(&record_id)
	return record_id, err
}

const selectSummonerRecentByPuuid = `-- name: SelectSummonerRecentByPuuid :one
SELECT record_id, record_date, account_id, profile_icon_id, revision_date, name, id, puuid, summoner_level
FROM summoner_records
WHERE puuid = $1
ORDER BY record_date DESC
LIMIT 1
`

func (q *Queries) SelectSummonerRecentByPuuid(ctx context.Context, puuid string) (SummonerRecord, error) {
	row := q.db.QueryRow(ctx, selectSummonerRecentByPuuid, puuid)
	var i SummonerRecord
	err := row.Scan(
		&i.RecordID,
		&i.RecordDate,
		&i.AccountID,
		&i.ProfileIconID,
		&i.RevisionDate,
		&i.Name,
		&i.ID,
		&i.Puuid,
		&i.SummonerLevel,
	)
	return i, err
}

const selectSummonerRecords = `-- name: SelectSummonerRecords :many
SELECT record_id, record_date, account_id, profile_icon_id, revision_date, name, id, puuid, summoner_level
FROM summoner_records
ORDER BY record_date DESC
LIMIT $1
OFFSET $2
`

type SelectSummonerRecordsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) SelectSummonerRecords(ctx context.Context, arg SelectSummonerRecordsParams) ([]SummonerRecord, error) {
	rows, err := q.db.Query(ctx, selectSummonerRecords, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SummonerRecord
	for rows.Next() {
		var i SummonerRecord
		if err := rows.Scan(
			&i.RecordID,
			&i.RecordDate,
			&i.AccountID,
			&i.ProfileIconID,
			&i.RevisionDate,
			&i.Name,
			&i.ID,
			&i.Puuid,
			&i.SummonerLevel,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectSummonerRecordsByPuuid = `-- name: SelectSummonerRecordsByPuuid :many
SELECT record_id, record_date, account_id, profile_icon_id, revision_date, name, id, puuid, summoner_level
FROM summoner_records
WHERE puuid = $1
ORDER BY record_date DESC
LIMIT $2
OFFSET $3
`

type SelectSummonerRecordsByPuuidParams struct {
	Puuid  string
	Limit  int32
	Offset int32
}

func (q *Queries) SelectSummonerRecordsByPuuid(ctx context.Context, arg SelectSummonerRecordsByPuuidParams) ([]SummonerRecord, error) {
	rows, err := q.db.Query(ctx, selectSummonerRecordsByPuuid, arg.Puuid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SummonerRecord
	for rows.Next() {
		var i SummonerRecord
		if err := rows.Scan(
			&i.RecordID,
			&i.RecordDate,
			&i.AccountID,
			&i.ProfileIconID,
			&i.RevisionDate,
			&i.Name,
			&i.ID,
			&i.Puuid,
			&i.SummonerLevel,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectSummonerRecordsNoIds = `-- name: SelectSummonerRecordsNoIds :many
SELECT record_date, revision_date, name, summoner_level
FROM summoner_records
ORDER BY record_date DESC
LIMIT $1
OFFSET $2
`

type SelectSummonerRecordsNoIdsParams struct {
	Limit  int32
	Offset int32
}

type SelectSummonerRecordsNoIdsRow struct {
	RecordDate    pgtype.Timestamp
	RevisionDate  int64
	Name          string
	SummonerLevel int64
}

func (q *Queries) SelectSummonerRecordsNoIds(ctx context.Context, arg SelectSummonerRecordsNoIdsParams) ([]SelectSummonerRecordsNoIdsRow, error) {
	rows, err := q.db.Query(ctx, selectSummonerRecordsNoIds, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectSummonerRecordsNoIdsRow
	for rows.Next() {
		var i SelectSummonerRecordsNoIdsRow
		if err := rows.Scan(
			&i.RecordDate,
			&i.RevisionDate,
			&i.Name,
			&i.SummonerLevel,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
