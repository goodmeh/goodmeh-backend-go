package controller

import (
	"goodmeh/app/mapper"
	"goodmeh/app/service"
	"net/http"

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

func (p *PlacesController) Init(r *gin.RouterGroup) {
	g := r.Group("/v1/places")
	g.GET("/:id", p.GetPlace)
	g.GET("/discover", p.GetRandomPlaces)
}

func NewPlacesController(placeService service.IPlaceService) *PlacesController {
	return &PlacesController{placeService}
}
