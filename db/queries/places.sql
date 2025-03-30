-- name: GetRandomPlaces :many
SELECT *
FROM place
ORDER BY RANDOM()
LIMIT $1;
-- name: GetPlaceNames :many
SELECT place.id,
    place.name
FROM place;
-- name: GetPlaceById :one
SELECT p.*
FROM place p
WHERE p.id = $1;
-- name: InsertPlaceField :exec
INSERT INTO place_field (place_id, field_id)
VALUES (
        $1,
        (
            SELECT id
            FROM field
            WHERE name = $2
                and category_id = $3
        )
    ) ON CONFLICT (place_id, field_id) DO NOTHING;
-- name: GetPlaceImageUrls :many
SELECT review_image.image_url
FROM review_image
    INNER JOIN review ON review_image.review_id = review.review_id
WHERE review.place_id = $1
ORDER BY review.created_at DESC
LIMIT $2 OFFSET $3;
-- name: InsertPlace :exec
INSERT INTO place (
        id,
        name,
        user_rating_count,
        image_url,
        recompute_stats,
        primary_type,
        lat,
        lng
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    ) ON CONFLICT (id) DO
UPDATE
SET name = $2,
    user_rating_count = $3,
    image_url = $4,
    recompute_stats = $5,
    primary_type = $6,
    lat = $7,
    lng = $8;