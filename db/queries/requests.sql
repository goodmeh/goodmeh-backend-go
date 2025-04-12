-- name: GetRequest :one
SELECT *
FROM request
WHERE place_id = $1
    AND status = $2;
-- name: InsertRequestOrSetFailedFalse :exec
INSERT INTO request (place_id, created_at, status, batch_job_id)
VALUES ($1, NOW(), $2, $3) ON CONFLICT (place_id, status) DO
UPDATE
SET failed = false,
    created_at = NOW();
-- name: SetRequestFailed :exec
UPDATE request
SET failed = $3
WHERE place_id = $1
    AND status = $2;
-- name: DeleteRequest :exec
DELETE FROM request
WHERE place_id = $1
    AND status = $2;