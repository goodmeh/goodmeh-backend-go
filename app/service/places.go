package service

import (
	"context"
	"goodmeh/app/repository"
)

type IPlaceService interface {
	GetRandomPlaces() ([]repository.Place, error)
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
