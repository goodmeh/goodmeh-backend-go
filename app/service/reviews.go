package service

import (
	"context"
	"encoding/json"
	"goodmeh/app/repository"
	"slices"
)

type IReviewService interface {
	GetReviewsImages(reviewIds []string) ([][]string, error)
}

type ReviewService struct {
	ctx context.Context
	q   *repository.Queries
}

func (r *ReviewService) GetReviewsImages(reviewIds []string) ([][]string, error) {
	slices.Sort(reviewIds)
	rows, err := r.q.GetReviewImageUrls(r.ctx, reviewIds)
	reviewIdIndex := 0
	images := make([][]string, len(reviewIds))
	// len(rows) <= len(reviewIds) and both are sorted according to reviewId
	for _, row := range rows {
		for reviewIds[reviewIdIndex] != row.ReviewID {
			reviewIdIndex++
		}
		err := json.Unmarshal(row.ImageUrls, &images[reviewIdIndex])
		if err != nil {
			return nil, err
		}
	}
	return images, err
}

func NewReviewService(ctx context.Context, q *repository.Queries) *ReviewService {
	return &ReviewService{ctx, q}
}
