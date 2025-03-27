-- name: GetRandomPlaces :many
SELECT *
FROM place
ORDER BY RANDOM()
LIMIT $1;
-- name: GetPlaceReviews :many
SELECT *
FROM review
WHERE place_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;