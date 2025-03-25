package response

import "time"

type PlacePreviewResponseDto struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Rating          float64    `json:"rating"`
	UserRatingCount int32      `json:"user_rating_count"`
	LastScraped     *time.Time `json:"last_scraped"`
	ImageUrl        string     `json:"image_url"`
	PrimaryType     string     `json:"primary_type"`
}
