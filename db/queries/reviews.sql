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
-- name: GetReviewImageUrls :many
SELECT review_image.review_id,
    JSON_AGG(review_image.image_url) AS image_urls
FROM review_image
WHERE review_id = ANY(@review_ids::text [])
GROUP BY review_image.review_id
ORDER BY review_image.review_id;
-- name: InsertReviews :batchexec
INSERT INTO review (
        id,
        user_id,
        rating,
        text,
        created_at,
        place_id,
        price_range
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO
UPDATE
SET user_id = EXCLUDED.user_id,
    rating = EXCLUDED.rating,
    text = EXCLUDED.text,
    created_at = EXCLUDED.created_at,
    weight = EXCLUDED.weight,
    place_id = EXCLUDED.place_id,
    price_range = EXCLUDED.price_range,
    summary = EXCLUDED.summary,
    business_summary = EXCLUDED.business_summary;
-- name: InsertReviewReplies :batchexec
INSERT INTO review_reply (
        review_id,
        text,
        created_at
    )
VALUES ($1, $2, $3) ON CONFLICT (review_id) DO
UPDATE
SET text = EXCLUDED.text,
    created_at = EXCLUDED.created_at;
-- name: InsertReviewImages :batchexec
INSERT INTO review_image (review_id, image_url)
VALUES ($1, $2) ON CONFLICT (review_id, image_url) DO NOTHING;
-- name: GetReviewsWithEnoughText :many
SELECT text
FROM review
WHERE place_id = $1
    AND text != ''
    AND LENGTH(text) > 50;