package service

type IPlaceService interface{}

type PlaceService struct{}

func NewPlaceService() *PlaceService {
	return &PlaceService{}
}
