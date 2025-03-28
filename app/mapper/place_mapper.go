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
			LastScraped:     &place.LastScraped,
			ImageUrl:        place.ImageUrl.String,
			PrimaryType:     place.PrimaryType.String,
		}
	}
	return placeDtos
}
