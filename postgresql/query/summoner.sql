-- name: InsertSummonerRecord :one
INSERT INTO summoner_records
(
    record_date,
    account_id,
    profile_icon_id,
    revision_date,
    name,
    id,
    puuid,
    summoner_level
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING record_id;

-- name: DeleteSummonerRecord :one
DELETE FROM summoner_records
WHERE record_id = $1
RETURNING record_date, name;

-- name: SelectSummonerRecordsByPuuid :many
SELECT *
FROM summoner_records
WHERE puuid = $1
ORDER BY record_date DESC
LIMIT $2 OFFSET $3;

-- name: SelectSummonerRecordsByName :many
SELECT *
FROM summoner_records
WHERE name = $1
ORDER BY record_date DESC
LIMIT $2 OFFSET $3;

-- name: CountSummonerRecordsByPuuid :one
SELECT COUNT(*)
FROM summoner_records
WHERE puuid = $1;