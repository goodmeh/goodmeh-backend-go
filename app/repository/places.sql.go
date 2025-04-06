// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: places.sql

package repository

import (
	"context"
)

const afterReviewInsert = `-- name: AfterReviewInsert :exec
UPDATE place
SET last_scraped = NOW(),
    recompute_stats = TRUE
WHERE id = $1
`

func (q *Queries) AfterReviewInsert(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, afterReviewInsert, id)
	return err
}

const getPlaceById = `-- name: GetPlaceById :one
SELECT p.id, p.name, p.rating, p.weighted_rating, p.user_rating_count, p.summary, p.last_scraped, p.image_url, p.recompute_stats, p.primary_type, p.business_summary, p.price_range, p.earliest_review_date, p.lat, p.lng
FROM place p
WHERE p.id = $1
`

func (q *Queries) GetPlaceById(ctx context.Context, id string) (Place, error) {
	row := q.db.QueryRow(ctx, getPlaceById, id)
	var i Place
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Rating,
		&i.WeightedRating,
		&i.UserRatingCount,
		&i.Summary,
		&i.LastScraped,
		&i.ImageUrl,
		&i.RecomputeStats,
		&i.PrimaryType,
		&i.BusinessSummary,
		&i.PriceRange,
		&i.EarliestReviewDate,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}

const getPlaceImageUrls = `-- name: GetPlaceImageUrls :many
SELECT review_image.image_url
FROM review_image
    INNER JOIN review ON review_image.review_id = review.review_id
WHERE review.place_id = $1
ORDER BY review.created_at DESC
LIMIT $2 OFFSET $3
`

type GetPlaceImageUrlsParams struct {
	PlaceID string `json:"place_id"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

func (q *Queries) GetPlaceImageUrls(ctx context.Context, arg GetPlaceImageUrlsParams) ([]string, error) {
	rows, err := q.db.Query(ctx, getPlaceImageUrls, arg.PlaceID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var image_url string
		if err := rows.Scan(&image_url); err != nil {
			return nil, err
		}
		items = append(items, image_url)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlaceNames = `-- name: GetPlaceNames :many
SELECT place.id,
    place.name
FROM place
`

type GetPlaceNamesRow struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) GetPlaceNames(ctx context.Context) ([]GetPlaceNamesRow, error) {
	rows, err := q.db.Query(ctx, getPlaceNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPlaceNamesRow
	for rows.Next() {
		var i GetPlaceNamesRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRandomPlaces = `-- name: GetRandomPlaces :many
SELECT id, name, rating, weighted_rating, user_rating_count, summary, last_scraped, image_url, recompute_stats, primary_type, business_summary, price_range, earliest_review_date, lat, lng
FROM place
ORDER BY RANDOM()
LIMIT $1
`

func (q *Queries) GetRandomPlaces(ctx context.Context, limit int32) ([]Place, error) {
	rows, err := q.db.Query(ctx, getRandomPlaces, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Rating,
			&i.WeightedRating,
			&i.UserRatingCount,
			&i.Summary,
			&i.LastScraped,
			&i.ImageUrl,
			&i.RecomputeStats,
			&i.PrimaryType,
			&i.BusinessSummary,
			&i.PriceRange,
			&i.EarliestReviewDate,
			&i.Lat,
			&i.Lng,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPlace = `-- name: InsertPlace :exec
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
    lng = $8
`

type InsertPlaceParams struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	UserRatingCount int32    `json:"user_rating_count"`
	ImageUrl        *string  `json:"image_url"`
	RecomputeStats  bool     `json:"recompute_stats"`
	PrimaryType     *string  `json:"primary_type"`
	Lat             *float64 `json:"lat"`
	Lng             *float64 `json:"lng"`
}

func (q *Queries) InsertPlace(ctx context.Context, arg InsertPlaceParams) error {
	_, err := q.db.Exec(ctx, insertPlace,
		arg.ID,
		arg.Name,
		arg.UserRatingCount,
		arg.ImageUrl,
		arg.RecomputeStats,
		arg.PrimaryType,
		arg.Lat,
		arg.Lng,
	)
	return err
}

const insertPlaceField = `-- name: InsertPlaceField :exec
INSERT INTO place_field (place_id, field_id)
VALUES (
        $1,
        (
            SELECT id
            FROM field
            WHERE name = $2
                and category_id = $3
        )
    ) ON CONFLICT (place_id, field_id) DO NOTHING
`

type InsertPlaceFieldParams struct {
	PlaceID    string `json:"place_id"`
	Name       string `json:"name"`
	CategoryID int32  `json:"category_id"`
}

func (q *Queries) InsertPlaceField(ctx context.Context, arg InsertPlaceFieldParams) error {
	_, err := q.db.Exec(ctx, insertPlaceField, arg.PlaceID, arg.Name, arg.CategoryID)
	return err
}
