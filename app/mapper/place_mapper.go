package mapper

import (
	"goodmeh/app/dto/response"
	"goodmeh/app/repository"
)

func ToPlacePreviewResponseDtos(places []repository.Place) []response.PlacePreviewResponseDto {
	placeDtos := make([]response.PlacePreviewResponseDto, len(places))
	for i, place := range places {
		placeDtos[i] = response.PlacePreviewResponseDto{
			ID:              place.ID,
			Name:            place.Name,
			Rating:          place.Rating,
			UserRatingCount: place.UserRatingCount,
			LastScraped:     place.LastScraped,
			ImageUrl:        place.ImageUrl,
			PrimaryType:     place.PrimaryType,
		}
	}
	return placeDtos
}

func ToPlaceResponseDto(place repository.Place) response.PlaceResponseDto {
	return response.PlaceResponseDto{
		ID:                 place.ID,
		Name:               place.Name,
		Rating:             place.Rating,
		WeightedRating:     place.WeightedRating,
		UserRatingCount:    place.UserRatingCount,
		Summary:            place.Summary,
		LastScraped:        place.LastScraped,
		ImageUrl:           place.ImageUrl,
		PrimaryType:        place.PrimaryType,
		BusinessSummary:    place.BusinessSummary,
		PriceRange:         place.PriceRange,
		EarliestReviewDate: place.EarliestReviewDate,
		Lat:                place.Lat,
		Lng:                place.Lng,
	}
}
