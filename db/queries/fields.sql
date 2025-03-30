-- name: GetFieldCategories :many
SELECT *
FROM field_category;
-- name: InsertField :exec
INSERT INTO field (name, category_id)
VALUES ($1, $2) ON CONFLICT (name, category_id) DO NOTHING;