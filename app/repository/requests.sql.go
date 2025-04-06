// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: requests.sql

package repository

import (
	"context"
)

const deleteRequest = `-- name: DeleteRequest :exec
DELETE FROM request
WHERE place_id = $1
    AND status = $2
`

type DeleteRequestParams struct {
	PlaceID string `json:"place_id"`
	Status  int32  `json:"status"`
}

func (q *Queries) DeleteRequest(ctx context.Context, arg DeleteRequestParams) error {
	_, err := q.db.Exec(ctx, deleteRequest, arg.PlaceID, arg.Status)
	return err
}

const getRequest = `-- name: GetRequest :one
SELECT place_id, created_at, status, failed, batch_job_id
FROM request
WHERE place_id = $1
    AND status = $2
`

type GetRequestParams struct {
	PlaceID string `json:"place_id"`
	Status  int32  `json:"status"`
}

func (q *Queries) GetRequest(ctx context.Context, arg GetRequestParams) (Request, error) {
	row := q.db.QueryRow(ctx, getRequest, arg.PlaceID, arg.Status)
	var i Request
	err := row.Scan(
		&i.PlaceID,
		&i.CreatedAt,
		&i.Status,
		&i.Failed,
		&i.BatchJobID,
	)
	return i, err
}

const insertRequestOrSetFailedFalse = `-- name: InsertRequestOrSetFailedFalse :exec
INSERT INTO request (place_id, created_at, status, batch_job_id)
VALUES ($1, NOW(), $2, $3) ON CONFLICT (place_id, status) DO
UPDATE
SET failed = false
`

type InsertRequestOrSetFailedFalseParams struct {
	PlaceID    string  `json:"place_id"`
	Status     int32   `json:"status"`
	BatchJobID *string `json:"batch_job_id"`
}

func (q *Queries) InsertRequestOrSetFailedFalse(ctx context.Context, arg InsertRequestOrSetFailedFalseParams) error {
	_, err := q.db.Exec(ctx, insertRequestOrSetFailedFalse, arg.PlaceID, arg.Status, arg.BatchJobID)
	return err
}

const setRequestFailed = `-- name: SetRequestFailed :exec
UPDATE request
SET failed = $3
WHERE place_id = $1
    AND status = $2
`

type SetRequestFailedParams struct {
	PlaceID string `json:"place_id"`
	Status  int32  `json:"status"`
	Failed  bool   `json:"failed"`
}

func (q *Queries) SetRequestFailed(ctx context.Context, arg SetRequestFailedParams) error {
	_, err := q.db.Exec(ctx, setRequestFailed, arg.PlaceID, arg.Status, arg.Failed)
	return err
}
