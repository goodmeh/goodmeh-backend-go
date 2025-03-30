-- name: GetRandomPlaces :many
SELECT *
FROM place
ORDER BY RANDOM()
LIMIT $1;
-- name: GetPlaceReviews :many
SELECT sqlc.embed(r),
    sqlc.embed(u),
    sqlc.embed(rr)
FROM review r
    INNER JOIN "user" u ON r.user_id = u.id
    LEFT JOIN review_reply rr ON r.id = rr.review_id
WHERE place_id = $1
ORDER BY r.created_at DESC
LIMIT $2 OFFSET $3;
-- name: GetPlaceImageUrls :many
SELECT review_image.image_url
FROM review_image
    INNER JOIN review ON review_image.review_id = review.review_id
WHERE review.place_id = $1
ORDER BY review.created_at DESC
LIMIT $2 OFFSET $3;
-- name: GetReviewImageUrls :many
SELECT review_image.review_id,
    JSON_AGG(review_image.image_url) AS image_urls
FROM review_image
WHERE review_id = ANY(@review_ids::text [])
GROUP BY review_image.review_id
ORDER BY review_image.review_id;
-- name: GetFieldCategories :many
SELECT *
FROM field_category;
-- name: InsertField :exec
INSERT INTO field (name, category_id)
VALUES ($1, $2) ON CONFLICT (name, category_id) DO NOTHING;
-- name: GetPlaceNames :many
SELECT place.id,
    place.name
FROM place;