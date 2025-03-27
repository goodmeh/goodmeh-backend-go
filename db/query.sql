-- name: GetRandomPlaces :many
SELECT * FROM place ORDER BY RANDOM() LIMIT $1;