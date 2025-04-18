// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: reviews.sql

package repository

import (
	"context"
)

const getPlaceReviews = `-- name: GetPlaceReviews :many
SELECT r.id, r.user_id, r.rating, r.text, r.created_at, r.weight, r.place_id, r.price_range, r.summary, r.business_summary,
    u.id, u.name, u.photo_uri, u.review_count, u.photo_count, u.rating_count, u.is_local_guide, u.long_review_count, u.score,
    rr.review_id, rr.text, rr.created_at
FROM review r
    INNER JOIN "user" u ON r.user_id = u.id
    LEFT JOIN review_reply rr ON r.id = rr.review_id
WHERE place_id = $1
ORDER BY r.created_at DESC
LIMIT $2 OFFSET $3
`

type GetPlaceReviewsParams struct {
	PlaceID string `json:"place_id"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetPlaceReviewsRow struct {
	Review      Review      `json:"review"`
	User        User        `json:"user"`
	ReviewReply ReviewReply `json:"review_reply"`
}

func (q *Queries) GetPlaceReviews(ctx context.Context, arg GetPlaceReviewsParams) ([]GetPlaceReviewsRow, error) {
	rows, err := q.db.Query(ctx, getPlaceReviews, arg.PlaceID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPlaceReviewsRow
	for rows.Next() {
		var i GetPlaceReviewsRow
		if err := rows.Scan(
			&i.Review.ID,
			&i.Review.UserID,
			&i.Review.Rating,
			&i.Review.Text,
			&i.Review.CreatedAt,
			&i.Review.Weight,
			&i.Review.PlaceID,
			&i.Review.PriceRange,
			&i.Review.Summary,
			&i.Review.BusinessSummary,
			&i.User.ID,
			&i.User.Name,
			&i.User.PhotoUri,
			&i.User.ReviewCount,
			&i.User.PhotoCount,
			&i.User.RatingCount,
			&i.User.IsLocalGuide,
			&i.User.LongReviewCount,
			&i.User.Score,
			&i.ReviewReply.ReviewID,
			&i.ReviewReply.Text,
			&i.ReviewReply.CreatedAt,
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

const getReviewImageUrls = `-- name: GetReviewImageUrls :many
SELECT review_image.review_id,
    JSON_AGG(review_image.image_url) AS image_urls
FROM review_image
WHERE review_id = ANY($1::text [])
GROUP BY review_image.review_id
ORDER BY review_image.review_id
`

type GetReviewImageUrlsRow struct {
	ReviewID  string `json:"review_id"`
	ImageUrls []byte `json:"image_urls"`
}

func (q *Queries) GetReviewImageUrls(ctx context.Context, reviewIds []string) ([]GetReviewImageUrlsRow, error) {
	rows, err := q.db.Query(ctx, getReviewImageUrls, reviewIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReviewImageUrlsRow
	for rows.Next() {
		var i GetReviewImageUrlsRow
		if err := rows.Scan(&i.ReviewID, &i.ImageUrls); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReviewsWithEnoughText = `-- name: GetReviewsWithEnoughText :many
SELECT text
FROM review
WHERE place_id = $1
    AND text != ''
    AND LENGTH(text) > 50
`

func (q *Queries) GetReviewsWithEnoughText(ctx context.Context, placeID string) ([]string, error) {
	rows, err := q.db.Query(ctx, getReviewsWithEnoughText, placeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
		items = append(items, text)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
