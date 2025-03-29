package service

import (
	"context"
	"goodmeh/app/events"
	"goodmeh/app/repository"
)

type IPlaceService interface {
	GetRandomPlaces() ([]repository.Place, error)
	GetPlaceReviews(placeId string, page, perPage int) ([]repository.GetPlaceReviewsRow, error)
	GetPlaceNames() (map[string]string, error)
	GetPlaceImages(placeId string, page, perPage int) ([]string, error)
}

type PlaceService struct {
	ctx      context.Context
	q        *repository.Queries
	eventBus *events.EventBus
}

func NewPlaceService(ctx context.Context, q *repository.Queries, eventBus *events.EventBus) *PlaceService {
	return &PlaceService{ctx, q, eventBus}
}

func (p *PlaceService) GetRandomPlaces() ([]repository.Place, error) {
	return p.q.GetRandomPlaces(p.ctx, 10)
}

func (p *PlaceService) GetPlaceReviews(placeId string, page, perPage int) ([]repository.GetPlaceReviewsRow, error) {
	limit := int32(perPage)
	offset := int32(page * perPage)
	return p.q.GetPlaceReviews(p.ctx, repository.GetPlaceReviewsParams{
		PlaceID: placeId,
		Limit:   limit,
		Offset:  offset,
	})
}

func (p *PlaceService) GetPlaceImages(placeId string, page, perPage int) ([]string, error) {
	limit := int32(perPage)
	offset := int32(page * perPage)
	return p.q.GetPlaceImageUrls(p.ctx, repository.GetPlaceImageUrlsParams{
		PlaceID: placeId,
		Limit:   limit,
		Offset:  offset,
	})
}

func (p *PlaceService) GetPlaceNames() (map[string]string, error) {
	placeNames, err := p.q.GetPlaceNames(p.ctx)
	if err != nil {
		return nil, err
	}
	placeNamesMap := make(map[string]string)
	for _, placeName := range placeNames {
		placeNamesMap[placeName.ID] = placeName.Name
	}
	return placeNamesMap, nil
}
