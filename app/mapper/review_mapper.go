package mapper

import (
	"encoding/json"
	"goodmeh/app/dto/response"
	"goodmeh/app/repository"
)

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
