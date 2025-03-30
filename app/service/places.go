package service

import (
	"context"
	"database/sql"
	"errors"
	"goodmeh/app/events"
	"goodmeh/app/repository"
	"log"
	"time"

	"github.com/goodmeh/backend-private/collector"
	googlereviews "github.com/goodmeh/backend-private/collector/google_reviews"
)

type IPlaceService interface {
	GetRandomPlaces() ([]repository.Place, error)
	GetPlaceReviews(placeId string, page, perPage int) ([]repository.GetPlaceReviewsRow, error)
	GetPlaceNames() (map[string]string, error)
	GetPlaceImages(placeId string, page, perPage int) ([]string, error)
	RequestPlace(placeId string) (*repository.Place, error)
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

func (p *PlaceService) InsertPlace(place collector.ScrapedPlace) error {
	return p.q.InsertPlace(p.ctx, repository.InsertPlaceParams{
		ID:              place.Place.ID,
		Name:            place.Place.Name,
		UserRatingCount: place.Place.UserRatingCount,
		ImageUrl:        place.Place.ImageUrl,
		RecomputeStats:  true,
		PrimaryType:     place.Place.PrimaryType,
		Lat:             place.Place.Lat,
		Lng:             place.Place.Lng,
	})
}

func (p *PlaceService) InsertPlaceFields(place collector.ScrapedPlace) error {
	p.eventBus.Publish(events.INSERT_NEW_FIELDS, place.Fields)
	categories, err := p.q.GetFieldCategories(p.ctx)
	if err != nil {
		return err
	}
	categoryMap := make(map[string]int32)
	for _, category := range categories {
		categoryMap[category.Name] = category.ID
	}
	for _, field := range place.Fields {
		err = p.q.InsertPlaceField(p.ctx, repository.InsertPlaceFieldParams{
			PlaceID:    place.Place.ID,
			CategoryID: categoryMap[field[0]],
			Name:       field[1],
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PlaceService) ScrapePlace(placeId string, laterThan *time.Time) {
	log.Printf("Scraping reviews for place %s", placeId)
	c := googlereviews.NewGoogleReviewsCollector(placeId, nil)
	reviewsChan, placeChan, err := c.Collect(p.ctx)
	if err != nil {
		return
	}
	var place collector.ScrapedPlace
	select {
	case place = <-placeChan:
	case <-p.ctx.Done():
		return
	}
	log.Printf("Scraped place %s", placeId)
	p.eventBus.Publish(events.ON_PLACE_SCRAPE, place.Place)
	log.Printf("Inserting place %s", placeId)
	err = p.InsertPlace(place)
	if err != nil {
		log.Printf("failed to insert place: %v", err)
		return
	}
	log.Printf("Inserting place fields %s", placeId)
	err = p.InsertPlaceFields(place)
	if err != nil {
		log.Printf("failed to insert place fields: %v", err)
		return
	}
	p.eventBus.Publish(events.ON_REVIEWS_READY, reviewsChan)
}

func (p *PlaceService) RequestPlace(placeId string) (*repository.Place, error) {
	const REVIEW_SCRAPE_INTERVAL_DAYS = 7
	place, err := p.q.GetPlaceById(p.ctx, placeId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	placePointer := &place
	if errors.Is(err, sql.ErrNoRows) {
		placePointer = nil
	}
	if place.LastScraped != nil &&
		time.Since(*place.LastScraped) < REVIEW_SCRAPE_INTERVAL_DAYS*24*time.Hour {
		return placePointer, nil
	}
	go p.ScrapePlace(placeId, place.LastScraped)
	return placePointer, nil
}
