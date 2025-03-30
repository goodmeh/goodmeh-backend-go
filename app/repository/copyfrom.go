// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: copyfrom.go

package repository

import (
	"context"
)

// iteratorForInsertReviewImages implements pgx.CopyFromSource.
type iteratorForInsertReviewImages struct {
	rows                 []InsertReviewImagesParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertReviewImages) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertReviewImages) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ReviewID,
		r.rows[0].ImageUrl,
	}, nil
}

func (r iteratorForInsertReviewImages) Err() error {
	return nil
}

func (q *Queries) InsertReviewImages(ctx context.Context, arg []InsertReviewImagesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"review_image"}, []string{"review_id", "image_url"}, &iteratorForInsertReviewImages{rows: arg})
}

// iteratorForInsertReviewReplies implements pgx.CopyFromSource.
type iteratorForInsertReviewReplies struct {
	rows                 []InsertReviewRepliesParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertReviewReplies) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertReviewReplies) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ReviewID,
		r.rows[0].Text,
		r.rows[0].CreatedAt,
	}, nil
}

func (r iteratorForInsertReviewReplies) Err() error {
	return nil
}

func (q *Queries) InsertReviewReplies(ctx context.Context, arg []InsertReviewRepliesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"review_reply"}, []string{"review_id", "text", "created_at"}, &iteratorForInsertReviewReplies{rows: arg})
}

// iteratorForInsertReviews implements pgx.CopyFromSource.
type iteratorForInsertReviews struct {
	rows                 []InsertReviewsParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertReviews) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertReviews) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].UserID,
		r.rows[0].Rating,
		r.rows[0].Text,
		r.rows[0].CreatedAt,
		r.rows[0].PlaceID,
		r.rows[0].PriceRange,
	}, nil
}

func (r iteratorForInsertReviews) Err() error {
	return nil
}

func (q *Queries) InsertReviews(ctx context.Context, arg []InsertReviewsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"review"}, []string{"id", "user_id", "rating", "text", "created_at", "place_id", "price_range"}, &iteratorForInsertReviews{rows: arg})
}

// iteratorForInsertUsers implements pgx.CopyFromSource.
type iteratorForInsertUsers struct {
	rows                 []InsertUsersParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertUsers) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertUsers) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ID,
		r.rows[0].Name,
		r.rows[0].PhotoUri,
		r.rows[0].ReviewCount,
		r.rows[0].PhotoCount,
		r.rows[0].RatingCount,
		r.rows[0].IsLocalGuide,
		r.rows[0].Score,
	}, nil
}

func (r iteratorForInsertUsers) Err() error {
	return nil
}

func (q *Queries) InsertUsers(ctx context.Context, arg []InsertUsersParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"user"}, []string{"id", "name", "photo_uri", "review_count", "photo_count", "rating_count", "is_local_guide", "score"}, &iteratorForInsertUsers{rows: arg})
}
