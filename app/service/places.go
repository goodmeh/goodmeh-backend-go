package service

import (
	"context"
	"goodmeh/app/repository"
)

type IPlaceService interface {
	GetRandomPlaces() ([]repository.Place, error)
	GetPlaceReviews(placeId string, page, perPage int) ([]repository.GetPlaceReviewsRow, error)
	GetPlaceImages(placeId string, page, perPage int) ([]string, error)
}

type PlaceService struct {
	ctx context.Context
	q   *repository.Queries
}

func NewPlaceService(ctx context.Context, q *repository.Queries) *PlaceService {
	return &PlaceService{ctx, q}
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
