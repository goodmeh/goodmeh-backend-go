package mapper

import (
	"goodmeh/app/dto/response"
	"goodmeh/app/repository"
)

func ToReviewResponseDto(reviewsWithUsers []repository.GetPlaceReviewsRow, imageUrls [][]string, perPage int) []response.ReviewResponseDto {
	reviewDtos := make([]response.ReviewResponseDto, len(reviewsWithUsers))
	for i, review := range reviewsWithUsers {
		reviewDtos[i] = response.ReviewResponseDto{
			Review:    review.Review,
			User:      review.User,
			ImageUrls: imageUrls[i],
		}
		if review.ReviewReply.ReviewID != "" {
			reviewDtos[i].Reply = &reviewsWithUsers[i].ReviewReply
		}
	}
	return reviewDtos
}
