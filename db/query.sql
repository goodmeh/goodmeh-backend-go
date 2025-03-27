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
-- name: GetPlaceImageUrls :many
SELECT review_image.image_url
FROM review_image
    INNER JOIN review ON review_image.review_id = review.review_id
WHERE review.place_id = $1
ORDER BY review.created_at DESC
LIMIT $2 OFFSET $3;