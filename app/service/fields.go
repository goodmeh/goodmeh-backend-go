package service

import (
	"context"
	"goodmeh/app/events"
	"goodmeh/app/repository"
	"log"
)

type IFieldService interface {
	InsertFields(fields [][2]string)
}

type FieldService struct {
	ctx      context.Context
	q        *repository.Queries
	eventBus *events.EventBus
}

func NewFieldService(ctx context.Context, q *repository.Queries, eventBus *events.EventBus) *FieldService {
	f := &FieldService{ctx, q, eventBus}
	f.eventBus.Subscribe(events.INSERT_NEW_FIELDS, events.AssertHandler(f.InsertFields))
	return f
}

func (f *FieldService) InsertFields(fields [][2]string) {
	categories, err := f.q.GetFieldCategories(f.ctx)
	if err != nil {
		log.Printf("Error getting field categories: %v", err)
		return
	}
	categoryMap := make(map[string]int32)
	for _, category := range categories {
		categoryMap[category.Name] = category.ID
	}
	for _, group := range fields {
		categoryName := group[0]
		fieldName := group[1]
		err = f.q.InsertField(f.ctx, repository.InsertFieldParams{
			Name:       fieldName,
			CategoryID: categoryMap[categoryName],
		})
		if err != nil {
			log.Printf("Error inserting field %s: %v", fieldName, err)
			return
		}
	}
}
