package service

import (
	"context"
	"encoding/json"
	"goodmeh/app/events"
	"goodmeh/app/repository"
	"log"
	"slices"

	"github.com/goodmeh/backend-private/collector"
)

type IReviewService interface {
	GetReviewsImages(reviewIds []string) ([][]string, error)
}

type ReviewService struct {
	ctx      context.Context
	q        *repository.Queries
	eventBus *events.EventBus
}

func (p *ReviewService) WithTx() (
	*ReviewService,
	repository.RollbackFunc,
	repository.CommitFunc,
	error,
) {
	tx, rollback, commit, err := p.q.Begin(p.ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return &ReviewService{p.ctx, tx, p.eventBus}, rollback, commit, nil
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

func (r *ReviewService) InsertReviews(reviewsChan <-chan []collector.ScrapedReview) {
	actualInsertion := func(acc []collector.ScrapedReview) {
		if len(acc) == 0 {
			return
		}
		r, rollback, commit, err := r.WithTx()
		if err != nil {
			log.Printf("failed to begin transaction: %v", err)
			return
		}
		defer rollback(r.ctx)
		// Insert User
		{
			users := make([]repository.InsertUsersParams, len(acc))
			for i, r := range acc {
				users[i] = repository.InsertUsersParams{
					ID:           r.User.ID,
					Name:         r.User.Name,
					PhotoUri:     r.User.PhotoUri,
					ReviewCount:  r.User.ReviewCount,
					PhotoCount:   r.User.PhotoCount,
					RatingCount:  r.User.RatingCount,
					IsLocalGuide: r.User.IsLocalGuide,
				}
			}
			var err error
			r.q.InsertUsers(r.ctx, users).Exec(func(i int, e error) {
				if e != nil {
					err = e
				}
			})
			if err != nil {
				log.Printf("failed to insert users: %v", err)
				return
			}
		}

		// Insert Review
		{
			reviews := make([]repository.InsertReviewsParams, len(acc))
			for i, r := range acc {
				reviews[i] = repository.InsertReviewsParams{
					ID:         r.Review.ID,
					UserID:     r.User.ID,
					Rating:     r.Review.Rating,
					Text:       r.Review.Text,
					CreatedAt:  r.Review.CreatedAt,
					PlaceID:    r.Review.PlaceID,
					PriceRange: r.Review.PriceRange,
				}
			}
			var err error
			r.q.InsertReviews(r.ctx, reviews).Exec(func(i int, e error) {
				if e != nil {
					err = e
				}
			})
			if err != nil {
				log.Printf("failed to insert reviews: %v", err)
				return
			}
		}

		// Insert ReviewReply
		{
			reviewReplies := make([]repository.InsertReviewRepliesParams, 0, len(acc))
			for _, r := range acc {
				if r.Reply == nil {
					continue
				}
				reviewReplies = append(reviewReplies, repository.InsertReviewRepliesParams{
					ReviewID:  r.Review.ID,
					Text:      r.Reply.Text,
					CreatedAt: r.Reply.CreatedAt,
				})
			}
			r.q.InsertReviewReplies(r.ctx, reviewReplies).Exec(func(i int, e error) {
				if e != nil {
					err = e
				}
			})
			if err != nil {
				log.Printf("failed to insert review replies: %v", err)
				return
			}
		}

		// Insert ReviewImage
		{
			reviewImages := make([]repository.InsertReviewImagesParams, 0)
			for _, r := range acc {
				for _, imageUrl := range r.ImageUrls {
					reviewImages = append(reviewImages, repository.InsertReviewImagesParams{
						ReviewID: r.Review.ID,
						ImageUrl: imageUrl,
					})
				}
			}
			r.q.InsertReviewImages(r.ctx, reviewImages).Exec(func(i int, e error) {
				if e != nil {
					err = e
				}
			})
			if err != nil {
				log.Printf("failed to insert review images: %v", err)
				return
			}
		}
		err = commit(r.ctx)
		if err != nil {
			log.Printf("failed to commit transaction: %v", err)
			return
		}
		log.Printf("Inserted %d reviews", len(acc))
	}
	go func() {
		var count uint32
		const MAX_BATCH_SIZE = 5000
		acc := make([]collector.ScrapedReview, 0, MAX_BATCH_SIZE)
		for {
			select {
			case <-r.ctx.Done():
				actualInsertion(acc)
				return
			case reviews, hasMore := <-reviewsChan:
				if !hasMore {
					actualInsertion(acc)
					log.Printf("Finished inserting %d reviews", count)
					r.eventBus.Publish(events.ON_REVIEWS_INSERT_END, nil)
					return
				}
				count += uint32(len(reviews))
				if len(acc)+len(reviews) > MAX_BATCH_SIZE {
					actualInsertion(acc)
					acc = acc[:0]
				}
				acc = append(acc, reviews...)
			}
		}
	}()
}

func NewReviewService(ctx context.Context, q *repository.Queries, eventBus *events.EventBus) *ReviewService {
	r := &ReviewService{ctx, q, eventBus}
	r.eventBus.Subscribe(events.ON_REVIEWS_READY, events.AssertHandler(r.InsertReviews))
	return r
}
