package response

import (
	"goodmeh/app/repository"
	"time"
)

type PlacePreviewResponseDto struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Rating          float64    `json:"rating"`
	UserRatingCount int32      `json:"user_rating_count"`
	LastScraped     *time.Time `json:"last_scraped"`
	ImageUrl        *string    `json:"image_url"`
	PrimaryType     *string    `json:"primary_type"`
}

type ReviewResponseDto struct {
	repository.Review
	User      repository.User         `json:"user"`
	Reply     *repository.ReviewReply `json:"reply"`
	ImageUrls []string                `json:"image_urls"`
}

type GetPlaceReviewsResponseDto struct {
	Data    []ReviewResponseDto `json:"data"`
	HasNext bool                `json:"has_next"`
}

type GetPlaceImagesResponseDto struct {
	Data    []string `json:"data"`
	HasNext bool     `json:"has_next"`
}
