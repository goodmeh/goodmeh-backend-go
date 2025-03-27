package controller

import (
	"goodmeh/app/dto/response"
	"goodmeh/app/mapper"
	"goodmeh/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPlacesController interface {
	Init(*gin.RouterGroup)
}

type PlacesController struct {
	placeService service.IPlaceService
}

func (p *PlacesController) GetPlace(c *gin.Context) {
	// TODO: implement GetPlace
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Not implemented",
	})
}

func (p *PlacesController) GetRandomPlaces(c *gin.Context) {
	placesModel, err := p.placeService.GetRandomPlaces()
	places := mapper.ToPlacePreviewResponseDtos(placesModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, places)
}

func (p *PlacesController) GetPlaceReviews(c *gin.Context) {
	id := c.Param("id")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page",
		})
		return
	}
	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid per_page",
		})
		return
	}
	reviewsModel, err := p.placeService.GetPlaceReviews(id, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.GetPlaceReviewsResponseDto{
		Data:    reviewsModel,
		HasNext: len(reviewsModel) == perPage,
	})

}

func (p *PlacesController) Init(r *gin.RouterGroup) {
	g := r.Group("/v1/places")
	g.GET("/:id", p.GetPlace)
	g.GET("/:id/reviews", p.GetPlaceReviews)
	g.GET("/discover", p.GetRandomPlaces)
}

func NewPlacesController(placeService service.IPlaceService) *PlacesController {
	return &PlacesController{placeService}
}
