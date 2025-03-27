// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Field struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	CategoryID int32  `json:"category_id"`
}

type FieldCategory struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Place struct {
	ID                 string           `json:"id"`
	Name               string           `json:"name"`
	Rating             float64          `json:"rating"`
	WeightedRating     float64          `json:"weighted_rating"`
	UserRatingCount    int32            `json:"user_rating_count"`
	Summary            pgtype.Text      `json:"summary"`
	LastScraped        pgtype.Timestamp `json:"last_scraped"`
	ImageUrl           pgtype.Text      `json:"image_url"`
	RecomputeStats     bool             `json:"recompute_stats"`
	PrimaryType        pgtype.Text      `json:"primary_type"`
	BusinessSummary    pgtype.Text      `json:"business_summary"`
	PriceRange         pgtype.Int4      `json:"price_range"`
	EarliestReviewDate pgtype.Timestamp `json:"earliest_review_date"`
	Lat                pgtype.Float8    `json:"lat"`
	Lng                pgtype.Float8    `json:"lng"`
}

type PlaceField struct {
	PlaceID string `json:"place_id"`
	FieldID int32  `json:"field_id"`
}

type PlaceKeyword struct {
	PlaceID string `json:"place_id"`
	Keyword string `json:"keyword"`
	Count   int32  `json:"count"`
}

type Review struct {
	ID              string           `json:"id"`
	UserID          string           `json:"user_id"`
	Rating          int32            `json:"rating"`
	Text            string           `json:"text"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	Weight          int32            `json:"weight"`
	PlaceID         string           `json:"place_id"`
	PriceRange      pgtype.Int4      `json:"price_range"`
	Summary         pgtype.Text      `json:"summary"`
	BusinessSummary pgtype.Text      `json:"business_summary"`
}

type ReviewImage struct {
	ReviewID string `json:"review_id"`
	ImageUrl string `json:"image_url"`
}

type ReviewReply struct {
	ReviewID  string           `json:"review_id"`
	Text      string           `json:"text"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type User struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	PhotoUri     pgtype.Text `json:"photo_uri"`
	ReviewCount  int32       `json:"review_count"`
	PhotoCount   int32       `json:"photo_count"`
	RatingCount  int32       `json:"rating_count"`
	IsLocalGuide bool        `json:"is_local_guide"`
	Score        int32       `json:"score"`
}
