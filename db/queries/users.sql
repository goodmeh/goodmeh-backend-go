-- name: InsertUsers :batchexec
INSERT INTO "user" (
        id,
        name,
        photo_uri,
        review_count,
        photo_count,
        rating_count,
        is_local_guide
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO
UPDATE
SET name = EXCLUDED.name,
    photo_uri = EXCLUDED.photo_uri,
    review_count = EXCLUDED.review_count,
    photo_count = EXCLUDED.photo_count,
    rating_count = EXCLUDED.rating_count,
    is_local_guide = EXCLUDED.is_local_guide;