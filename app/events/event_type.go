package events

import "github.com/goodmeh/backend-private/collector"

const (
	ON_PLACE_SCRAPE = iota
	ON_REVIEWS_READY
	ON_REVIEWS_INSERT_END
	INSERT_NEW_FIELDS
)

type OnReviewsReadyParams struct {
	PlaceId     string
	ReviewsChan <-chan []collector.ScrapedReview
}
