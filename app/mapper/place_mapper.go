package mapper

import (
	"encoding/json"
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
			LastScraped:     &place.LastScraped.Time,
			ImageUrl:        place.ImageUrl.String,
			PrimaryType:     place.PrimaryType.String,
		}
	}
	return placeDtos
}

func ToReviewResponseDto(reviewsWithUsers []repository.GetPlaceReviewsRow, imageUrls [][]string, perPage int) []response.ReviewResponseDto {
	reviewDtos := make([]response.ReviewResponseDto, len(reviewsWithUsers))
	for i, review := range reviewsWithUsers {
		reviewDtos[i] = response.ReviewResponseDto{
			GetPlaceReviewsRow: review,
			ImageUrls:          imageUrls[i],
		}
		json.Unmarshal(review.User, &reviewDtos[i].User)
		if review.Reply != nil {
			json.Unmarshal(review.Reply, &reviewDtos[i].Reply)
		}
	}
	return reviewDtos
}
