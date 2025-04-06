-- name: GetRequest :one
SELECT *
FROM request
WHERE place_id = $1
    AND status = $2;
-- name: InsertRequest :exec
INSERT INTO request (place_id, created_at, status, batch_job_id)
VALUES ($1, NOW(), $2, $3);
-- name: SetRequestFailed :exec
UPDATE request
SET failed = TRUE
WHERE place_id = $1
    AND status = $2;
-- name: DeleteRequest :exec
DELETE FROM request
WHERE place_id = $1
    AND status = $2;