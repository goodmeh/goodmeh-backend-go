package service

import (
	"context"
	"goodmeh/app/repository"
)

type IFieldService interface {
	InsertFields(fieldNames, categoryNames []string) error
}

type FieldService struct {
	ctx context.Context
	q   *repository.Queries
}

func NewFieldService(ctx context.Context, q *repository.Queries) *FieldService {
	return &FieldService{ctx, q}
}

func (f *FieldService) InsertFields(fieldNames, categoryNames []string) error {
	categories, err := f.q.GetFieldCategories(f.ctx)
	if err != nil {
		return err
	}
	categoryMap := make(map[string]int32)
	for _, category := range categories {
		categoryMap[category.Name] = category.ID
	}
	for i, fieldName := range fieldNames {
		err = f.q.InsertField(f.ctx, repository.InsertFieldParams{
			Name:       fieldName,
			CategoryID: categoryMap[categoryNames[i]],
		})
		if err != nil {
			return err
		}
	}
	return nil
}
