-- name: InsertUsers :copyfrom
INSERT INTO "user" (
        id,
        name,
        photo_uri,
        review_count,
        photo_count,
        rating_count,
        is_local_guide,
        score
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);